package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("model model module init function")
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
