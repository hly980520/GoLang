package api

import "github.com/gin-gonic/gin"

type Api interface {
	QueryPage(*gin.Context)
	Details(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Add(*gin.Context)
}
