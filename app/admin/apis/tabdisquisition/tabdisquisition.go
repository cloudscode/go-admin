package tabdisquisition

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
)

type TabDisquisition struct {
	apis.Api
}

// GetTabDisquisitionList 获取科研成果列表
// @Summary 获取科研成果列表
// @Description 获取科研成果列表
// @Tags 科研成果
// @Param title query string false "论文题目"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TabDisquisition}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tabdisquisition [get]
// @Security Bearer
func (e TabDisquisition) GetTabDisquisitionList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.TabDisquisitionSearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	println("***************************** %S",p.DataScope)
	list := make([]models.TabDisquisition, 0)
	var count int64
	serviceTabDisquisition := service.TabDisquisition{}
	serviceTabDisquisition.Log = log
	serviceTabDisquisition.Orm = db
	err = serviceTabDisquisition.GetTabDisquisitionPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get TabDisquisition Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetTabDisquisition 获取科研成果
// @Summary 获取科研成果
// @Description 获取科研成果
// @Tags 科研成果
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.TabDisquisition} "{"code": 200, "data": [...]}"
// @Router /api/v1/tabdisquisition/{id} [get]
// @Security Bearer
func (e TabDisquisition) GetTabDisquisition(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TabDisquisitionById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.TabDisquisition

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceTabDisquisition := service.TabDisquisition{}
	serviceTabDisquisition.Log = log
	serviceTabDisquisition.Orm = db
	err = serviceTabDisquisition.GetTabDisquisition(control, p, &object)
	if err != nil {
		log.Errorf("Get TabDisquisition error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertTabDisquisition 创建科研成果
// @Summary 创建科研成果
// @Description 创建科研成果
// @Tags 科研成果
// @Accept application/json
// @Product application/json
// @Param data body dto.TabDisquisitionControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tabdisquisition [post]
// @Security Bearer
func (e TabDisquisition) InsertTabDisquisition(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TabDisquisitionControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate TabDisquisition model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceTabDisquisition := service.TabDisquisition{}
	serviceTabDisquisition.Orm = db
	serviceTabDisquisition.Log = log
	err = serviceTabDisquisition.InsertTabDisquisition(object.(*models.TabDisquisition))
	if err != nil {
		log.Errorf("Insert TabDisquisition error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateTabDisquisition 修改科研成果
// @Summary 修改科研成果
// @Description 修改科研成果
// @Tags 科研成果
// @Accept application/json
// @Product application/json
// @Param data body dto.TabDisquisitionControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tabdisquisition/{id} [put]
// @Security Bearer
func (e TabDisquisition) UpdateTabDisquisition(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TabDisquisitionControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate TabDisquisition model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceTabDisquisition := service.TabDisquisition{}
	serviceTabDisquisition.Orm = db
	serviceTabDisquisition.Log = log
	err = serviceTabDisquisition.UpdateTabDisquisition(object.(*models.TabDisquisition), p)
	if err != nil {
		log.Errorf("Update TabDisquisition error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteTabDisquisition 删除科研成果
// @Summary 删除科研成果
// @Description 删除科研成果
// @Tags 科研成果
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tabdisquisition [delete]
// @Security Bearer
func (e TabDisquisition) DeleteTabDisquisition(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TabDisquisitionById)

	//删除操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	// 设置编辑人
	control.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceTabDisquisition := service.TabDisquisition{}
	serviceTabDisquisition.Orm = db
	serviceTabDisquisition.Log = log
	err = serviceTabDisquisition.RemoveTabDisquisition(control, p)
	if err != nil {
		log.Errorf("Remove TabDisquisition error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}