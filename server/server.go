package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	address string
	log *zap.Logger
	gin *gin.Context
}