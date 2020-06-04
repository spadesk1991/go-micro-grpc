package ServicesImpl

import (
	"context"

	services "github.com/spadesk1991/go-micro-grpc/Services"
)

type ProdService struct {
}

func (ps *ProdService) GetProdsList(ctx context.Context, in *services.ProdRequest, res *services.ProdListResponse) (err error) {
	return
}
