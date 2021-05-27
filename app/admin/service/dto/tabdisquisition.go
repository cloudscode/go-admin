package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TabDisquisitionSearch struct {
	dto.Pagination     `search:"-"`
    Title string `form:"title"  search:"type:contains;column:title;table:tab_disquisition" comment:"论文题目"`
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
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type TabDisquisitionControl struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Title string `json:"title" comment:"论文题目"`
}

func (s *TabDisquisitionControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *TabDisquisitionControl) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Errorf("ShouldBindUri error: %s", err.Error())
        return errors.New("数据绑定出错")
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Errorf("ShouldBind error: %s", err.Error())
        err = errors.New("数据绑定出错")
    }
    return err
}

func (s *TabDisquisitionControl) GenerateM() (common.ActiveRecord, error) {
	return &models.TabDisquisition{
        Model:     common.Model{ Id: s.Id },
        Title:  s.Title,
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
		log.Errorf("ShouldBindUri error: %s", err.Error())
		return errors.New("数据绑定出错")
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Errorf("ShouldBind error: %s", err.Error())
		err = errors.New("数据绑定出错")
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