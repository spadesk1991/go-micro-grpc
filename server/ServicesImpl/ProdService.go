package ServicesImpl

import (
	"context"
	"strconv"

	services "github.com/spadesk1991/go-micro-grpc/server/Services"
)

type ProdService struct {
}

func (ps *ProdService) GetProdsList(ctx context.Context, in *services.ProdRequest, res *services.ProdListResponse) (err error) {
	for i := 0; i < int(in.Size); i++ {
		res.Data = append(res.Data, &services.ProdModel{
			ProdID:   1000 + int32(i),
			ProdName: "prod_name" + strconv.Itoa(i),
		})
	}
	return
}
