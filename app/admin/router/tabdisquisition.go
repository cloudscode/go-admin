package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/common/actions"

	"go-admin/app/admin/apis/tabdisquisition"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTabDisquisitionRouter)
}

// registerTabDisquisitionRouter
func registerTabDisquisitionRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &tabdisquisition.TabDisquisition{}
	r := v1.Group("/tabdisquisition").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("",actions.PermissionAction(), api.GetTabDisquisitionList)
		r.GET("/:id", api.GetTabDisquisition)
		r.POST("", api.InsertTabDisquisition)
		r.PUT("/:id", api.UpdateTabDisquisition)
		r.DELETE("", api.DeleteTabDisquisition)
	}
}