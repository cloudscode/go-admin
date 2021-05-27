package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type TabDisquisition struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(255);comment:论文题目"` 
    models.ModelTime
    models.ControlBy
}

func (TabDisquisition) TableName() string {
    return "tab_disquisition"
}

func (e *TabDisquisition) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TabDisquisition) GetId() interface{} {
	return e.Id
}