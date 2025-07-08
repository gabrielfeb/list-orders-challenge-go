package service

import (
	"context"
	"strconv"

	"list-orders-challenge-go/internal/infra/grpc/pb"
	"list-orders-challenge-go/internal/usecase"
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

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	price, err := strconv.ParseFloat(in.Price, 64)
	if err != nil {
		return nil, err
	}

	tax, err := strconv.ParseFloat(in.Tax, 64)
	if err != nil {
		return nil, err
	}

	dto := usecase.CreateOrderInputDTO{
		Price: price,
		Tax:   tax,
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:         output.ID,
			Price:      float32(output.Price),
			Tax:        float32(output.Tax),
			FinalPrice: float32(output.FinalPrice),
		},
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
			Price:      float32(o.Price),
			Tax:        float32(o.Tax),
			FinalPrice: float32(o.FinalPrice),
		})

	}
	return &pb.ListOrdersResponse{Orders: orders}, nil
}
