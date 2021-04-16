package router

import (
    "github.com/gin-gonic/gin"
    jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

    "go-admin/app/research/models"
    "go-admin/app/research/service/dto"
    "go-admin/common/actions"
    "go-admin/common/middleware"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerTabDisquisitionRouter)
}

// 需认证的路由代码
func registerTabDisquisitionRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/tabdisquisition").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.TabDisquisition{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.TabDisquisitionSearch), func() interface{} {
            list := make([]models.TabDisquisition, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.TabDisquisitionById), nil))
        r.POST("", actions.CreateAction(new(dto.TabDisquisitionControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.TabDisquisitionControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.TabDisquisitionById)))
    }
}
