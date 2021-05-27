package service

import (
	"errors"

	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/app/admin/service/dto"
	"go-admin/common/service"
)

type TabDisquisition struct {
	service.Service
}

// GetTabDisquisitionPage 获取TabDisquisition列表
func (e *TabDisquisition) GetTabDisquisitionPage(c *dto.TabDisquisitionSearch, p *actions.DataPermission, list *[]models.TabDisquisition, count *int64) error {
	var err error
	var data models.TabDisquisition

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetTabDisquisitionPage error:%s", err)
		return err
	}
	return nil
}

// GetTabDisquisition 获取TabDisquisition对象
func (e *TabDisquisition) GetTabDisquisition(d *dto.TabDisquisitionById, p *actions.DataPermission, model *models.TabDisquisition) error {
	var data models.TabDisquisition

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTabDisquisition error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertTabDisquisition 创建TabDisquisition对象
func (e *TabDisquisition) InsertTabDisquisition(c *models.TabDisquisition) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertTabDisquisition error:%s", err)
		return err
	}
	return nil
}

// UpdateTabDisquisition 修改TabDisquisition对象
func (e *TabDisquisition) UpdateTabDisquisition(c *models.TabDisquisition, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateTabDisquisition error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveTabDisquisition 删除TabDisquisition
func (e *TabDisquisition) RemoveTabDisquisition(d *dto.TabDisquisitionById, p *actions.DataPermission) error {
	var data models.TabDisquisition

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTabDisquisition error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}