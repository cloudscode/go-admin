package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type TabDisquisition struct {
    models.Model
    
    DisquisitionId string `json:"disquisitionId" gorm:"type:varchar(32);comment:DisquisitionId"` 
    Kind string `json:"kind" gorm:"type:varchar(16);comment:论文类型"` 
    Title string `json:"title" gorm:"type:varchar(256);comment:论文题目"` 
    Publications string `json:"publications" gorm:"type:varchar(256);comment:刊物名称"` 
    Levels string `json:"levels" gorm:"type:varchar(16);comment:刊物类别"` 
    VolumeNo string `json:"volumeNo" gorm:"type:varchar(64);comment:卷号"` 
    Period string `json:"period" gorm:"type:varchar(64);comment:期号"` 
    PublishingDate time.Time `json:"publishingDate" gorm:"type:datetime(3);comment:发表/出版时间"` 
    WordLength int64 `json:"wordLength" gorm:"type:double;comment:字数"` 
    KnowledgeClass string `json:"knowledgeClass" gorm:"type:varchar(16);comment:学科门类"` 
    FirstKnowledge string `json:"firstKnowledge" gorm:"type:varchar(16);comment:一级学科"` 
    Source string `json:"source" gorm:"type:varchar(16);comment:成果来源"` 
    PublishingRange string `json:"publishingRange" gorm:"type:varchar(16);comment:发表范围"` 
    Issn string `json:"issn" gorm:"type:varchar(64);comment:ISSN号"` 
    Cn string `json:"cn" gorm:"type:varchar(64);comment:CN号"` 
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