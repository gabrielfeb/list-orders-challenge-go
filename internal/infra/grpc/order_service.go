package service

import (
	"context"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/grpc/pb"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createUC usecase.CreateOrderUseCase, listUC usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createUC,
		ListOrdersUseCase:  listUC,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.CreateOrderInputDTO{
		Price: in.Price,
		Tax:   in.Tax,
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:         output.ID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, o := range output {
		orders = append(orders, &pb.Order{
			Id:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}
	return &pb.ListOrdersResponse{Orders: orders}, nil
}
