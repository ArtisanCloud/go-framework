package models

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/database"
	"github.com/ArtisanCloud/go-libs/object"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"math"
	"reflect"
	"sync"
)

type Model interface {
	GetTableName(needFull bool) string
	GetUUID() string
}

const STATUS_DRAFT = "Draft"
const STATUS_CANCELED = "Canceled"

const APPROVAL_STATUS_DRAFT = "Draft"
const APPROVAL_STATUS_PENDING = "Pending"
const APPROVAL_STATUS_APPROVED = "Approved"
const APPROVAL_STATUS_REJECTED = "Rejected"

var ArrayModelFields *object.HashMap = &object.HashMap{}


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
