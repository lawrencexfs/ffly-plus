package router

import (
	"github.com/colinrs/ffly-plus/controller"
	"github.com/colinrs/ffly-plus/router/api"
	"github.com/colinrs/ffly-plus/router/middleware"

	//nolint: golint
	_ "github.com/colinrs/ffly-plus/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server ...
type Server struct {
	GinEngine *gin.Engine
}

// Resp 返回基本状态定义
type Resp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// 返回状态码
var (
	SUCCESS    = Resp{Code: 0, Message: "成功", Detail: "Welcome use mds server"}
	SUCCESS_V  = Resp{Code: 1, Detail: "Welcome visitor here"}
	FAILURE    = Resp{Code: 10001, Message: "失败", Detail: "内部调用失败"}
	ReqArgsErr = Resp{Code: 9001, Message: "请求参数错误"}
)

// 通用回复
func HelloRespV(c *gin.Context) {
	c.JSON(200, SUCCESS_V)
}

// InitRouter ...
func InitRouter() *Server {
	server := new(Server)
	gin.SetMode(gin.DebugMode)
	server.GinEngine = gin.Default()
	server.GinEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GinEngine.Use(middleware.AcclogSetUp())
	server.GinEngine.Use(middleware.SentinelMiddleware())

	server.GinEngine.GET("/", HelloRespV)

	registerBaseAPI(server)
	apiGroupV1 := server.GinEngine.Group("/api/v1")
	api.RegisterAPIV1(apiGroupV1)

	return server
}

// registerBaseAPI ...
func registerBaseAPI(server *Server) {
	server.GinEngine.GET("/health", controller.Health)
	server.GinEngine.GET("/version", controller.Version)
}
