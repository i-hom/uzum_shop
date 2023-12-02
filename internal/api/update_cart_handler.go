package api

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/pkg/shopV1"
)

func (s *ShopAPI) UpdateCart(ctx context.Context, req *shopV1.Cart_UpdateRequest) (*emptypb.Empty, error) {
	err := s.service.UpdateProductInCart(ctx, uuid.MustParse(req.ProductId), int32(req.Quantity))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
