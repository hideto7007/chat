package lib

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

    // docsのディレクトリを指定
	_ "chat/docs" // ←追記
)

// Register はSwaggerのエンドポイントをルーターに登録します
func Register(r *gin.Engine) {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}