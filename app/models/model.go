package models

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/database"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"math"
	"reflect"
	"sync"
	"time"
)


type Model interface {
	GetTableName(needFull bool) string
	GetID() interface{}
	GetPrimaryKey() string
	GetForeignKey() string
}

type MyModel struct {
	ID   int    `gorm:"autoIncrement:true;unique; column:id; ->;<-:create" json:"-"`
	UUID string `gorm:"primaryKey;autoIncrement:false;unique; column:uuid; ->;<-:create " json:"uuid" sql:"index"`

	CreatedAt time.Time `gorm:"column:created_at; ->;<-:create " json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type MyRelationship struct {
	ID        int       `gorm:"autoIncrement:true;unique; column:id; ->;<-:create" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at; ->;<-:create " json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}


const UNIQUE_ID = "uuid"

const MODEL_STATUS_DRAFT int8 = 0
const MODEL_STATUS_ACTIVE int8 = 1
const MODEL_STATUS_CANCELED int8 = 2
const MODEL_STATUS_PENDING int8 = 3
const MODEL_STATUS_INACTIVE int8 = 4

const APPROVAL_STATUS_DRAFT int8 = 0
const APPROVAL_STATUS_PENDING int8 = 1
const APPROVAL_STATUS_APPROVED int8 = 3
const APPROVAL_STATUS_REJECTED int8 = 4

var ArrayModelFields *object.HashMap = &object.HashMap{}

func NewMyModel() *MyModel {
	now := time.Now()
	return &MyModel{
		UUID:      uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewMyRelationship() *MyRelationship {
	now := time.Now()
	return &MyRelationship{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (mdl *MyModel) GetID() interface{} {
	return mdl.UUID
}

func (mdl *MyModel) GetPrimaryKey() string {
	return "uuid"
}
func (mdl *MyModel) GetForeignKey() string {
	return "model_uuid"
}

func (mdl *MyRelationship) GetID() interface{} {
	return ""
}

func (mdl *MyRelationship) GetPrimaryKey() string {
	return "id"
}
func (mdl *MyRelationship) GetForeignKey() string {
	return "model_id"
}


/**
 * Scope Where Conditions
 */
func WhereUUID(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//return db.Where("integration_id__c=@value", sql.Named("value", uuid))
		return db.Where("integration_id__c=?", uuid)
	}
}

func GetFirst(db *gorm.DB, conditions *gorm.DB, model interface{}, preloads []string) (err error) {

	if conditions != nil {
		db = db.Where(conditions)
	}

	// add preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			if preload != "" {
				db.Preload(preload)
			}
		}
	}

	result := db.First(model)

	return result.Error
}

func GetList(db *gorm.DB, conditions *gorm.DB,
	models interface{}, preloads []string,
	page int, pageSize int) (paginator *database.Pagination, err error) {

	// add pagination
	paginator = database.NewPagination(page, pageSize, "")
	var totalRows int64
	db.Model(models).Count(&totalRows)
	paginator.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(paginator.Limit)))
	paginator.TotalPages = totalPages

	db = db.Scopes(
		Paginate(page, pageSize),
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	// add preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			if preload != "" {
				db.Preload(preload)
			}
		}
	}

	// chunk datas
	result := db.Find(models)
	err = result.Error
	if err != nil {
		return paginator, err
	}

	paginator.Data = models

	return paginator, nil
}

/**
 * Association Relationship
 */
func AssociationRelationship(db *gorm.DB, conditions *gorm.DB, mdl interface{}, relationship string, withClauseAssociations bool) *gorm.Association {

	tx := db.Model(mdl)

	if withClauseAssociations {
		tx.Preload(clause.Associations)
	}

	if conditions != nil {
		tx = tx.Where(conditions)
	}

	return tx.Association(relationship)
}


func ClearAssociation(db *gorm.DB, object Model, foreignKey string, pivot Model) error {
	result := db.Exec("DELETE FROM "+pivot.GetTableName(true)+" WHERE "+foreignKey+"=?", object.GetID())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetModelFields(model interface{}) (fields []string) {

	// check if it has been loaded
	modelType := reflect.TypeOf(model)
	modelName := modelType.Name()
	if (*ArrayModelFields)[modelName] != nil {
		return (*ArrayModelFields)[modelName].([]string)
	}

	fmt.Printf("parse object ~%s~ model fields \n", modelName)
	gormSchema, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		println(err)
		return fields
	}

	fields = []string{}
	for _, field := range gormSchema.Fields {
		if field.DBName != "" && !field.PrimaryKey && !field.Unique && field.Updatable {
			fields = append(fields, field.DBName)
		}
	}
	(*ArrayModelFields)[modelName] = fields
	fmt.Printf("parsed object ~%s~ model fields and fields count is %d \n\n", modelName, len(fields))

	return fields
}
