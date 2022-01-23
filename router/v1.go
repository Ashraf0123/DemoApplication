package router

import (
	"github.com/labstack/echo/v4"
	v1 "tempApi/api/v1"
)

func V1Routes(e *echo.Group) {
	e.GET("", v1.CacheController().Get)
	e.POST("", v1.CacheController().Create)
	e.PUT("", v1.CacheController().Update)
	e.DELETE("", v1.CacheController().Delete)
}
