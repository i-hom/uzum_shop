package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/pkg/shopV1"
)

func (s *ShopAPI) GetCart(ctx context.Context, req *emptypb.Empty) (*shopV1.Cart_GetResponse, error) {
	cart, err := s.service.GetCart(ctx)
	if err != nil {
		return &shopV1.Cart_GetResponse{}, err
	}

	var responseItems []*shopV1.Product

	for _, item := range *cart {
		responseItems = append(responseItems, &shopV1.Product{
			ProductId:   item.ID.String(),
			Name:        item.Name,
			Price:       float32(item.Price),
			Description: item.Description,
			Quantity:    uint32(item.Quantity)})
	}

	return &shopV1.Cart_GetResponse{Products: responseItems}, nil
}
