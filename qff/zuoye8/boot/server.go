package boot

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main/app/global"
	"net/http"
)

func ServerSetup() {
	config := global.Config.Server

	gin.SetMode(config.Mode)
	server := &http.Server{
		Addr:           config.Addr(),
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Info("initialize server success", zap.String("port", config.Addr()))
	global.Logger.Error(server.ListenAndServe().Error())
}
