package services

import (
	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/go-micro-grpc/server/Services"
)

func GetOrderByID(c *gin.Context, id int) (res Services.ProdListResponse, err error) {
	prodS, ok := c.Keys["prodGRPCService"].(Services.ProdService)
	if !ok {
		return
	}
	req := Services.ProdRequest{
		Size: int32(id),
	}
	r, err := prodS.GetProdsList(c, &req)
	res = *r
	return
}
