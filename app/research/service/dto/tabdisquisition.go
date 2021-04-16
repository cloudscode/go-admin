package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/research/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TabDisquisitionSearch struct {
	dto.Pagination     `search:"-"`
    
}

func (m *TabDisquisitionSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *TabDisquisitionSearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *TabDisquisitionSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Warnf("ShouldBind error: %s", err.Error())
    }
    return err
}

type TabDisquisitionControl struct {
    
    Id int `uri:"id" comment:""` // 
    DisquisitionId string `json:"disquisitionId" comment:""`
    
    Kind string `json:"kind" comment:"论文类型"`
    
    Title string `json:"title" comment:"论文题目"`
    
    Publications string `json:"publications" comment:"刊物名称"`
    
    Levels string `json:"levels" comment:"刊物类别"`
    
    VolumeNo string `json:"volumeNo" comment:"卷号"`
    
    Period string `json:"period" comment:"期号"`
    
    PublishingDate time.Time `json:"publishingDate" comment:"发表/出版时间"`
    
    WordLength int64 `json:"wordLength" comment:"字数"`
    
    KnowledgeClass string `json:"knowledgeClass" comment:"学科门类"`
    
    FirstKnowledge string `json:"firstKnowledge" comment:"一级学科"`
    
    Source string `json:"source" comment:"成果来源"`
    
    PublishingRange string `json:"publishingRange" comment:"发表范围"`
    
    Issn string `json:"issn" comment:"ISSN号"`
    
    Cn string `json:"cn" comment:"CN号"`
    
}

func (s *TabDisquisitionControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *TabDisquisitionControl) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Warnf("ShouldBindUri error: %s", err.Error())
        return err
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Warnf("ShouldBind error: %s", err.Error())
    }
    return err
}

func (s *TabDisquisitionControl) GenerateM() (common.ActiveRecord, error) {
	return &models.TabDisquisition{
	
        Model:     common.Model{ Id: s.Id },
        DisquisitionId:  s.DisquisitionId,
        Kind:  s.Kind,
        Title:  s.Title,
        Publications:  s.Publications,
        Levels:  s.Levels,
        VolumeNo:  s.VolumeNo,
        Period:  s.Period,
        PublishingDate:  s.PublishingDate,
        WordLength:  s.WordLength,
        KnowledgeClass:  s.KnowledgeClass,
        FirstKnowledge:  s.FirstKnowledge,
        Source:  s.Source,
        PublishingRange:  s.PublishingRange,
        Issn:  s.Issn,
        Cn:  s.Cn,
	}, nil
}

func (s *TabDisquisitionControl) GetId() interface{} {
	return s.Id
}

type TabDisquisitionById struct {
	dto.ObjectById
}

func (s *TabDisquisitionById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *TabDisquisitionById) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Warnf("ShouldBindUri error: %s", err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Warnf("ShouldBind error: %s", err.Error())
	}
	return err
}

func (s *TabDisquisitionById) SetUpdateBy(id int) {

}

func (s *TabDisquisitionById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *TabDisquisitionById) GenerateM() (common.ActiveRecord, error) {
	return &models.TabDisquisition{}, nil
}
