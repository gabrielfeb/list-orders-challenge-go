package graph

import (
	"context"
	"fmt" // Adicione este import

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/graph/model"
)

// ESTA STRUCT PRECISA TER OS CAMPOS ABAIXO.
// É aqui que a injeção de dependência acontece.
type Resolver struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

// CreateOrder é o resolver para a mutation createOrder.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error) {
	dto := dto.OrderInputDTO{
		Price: input.Preco,
		Tax:   input.Imposto,
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         fmt.Sprintf("%.0f", output.ID),
		Preco:      output.Price,
		Imposto:    output.Tax,
		PrecoFinal: output.FinalPrice,
	}, nil
}

// Orders é o resolver para a query orders.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for _, o := range output {
		orders = append(orders, &model.Order{
			ID:         fmt.Sprintf("%.0f", o.ID),
			Preco:      o.Price,
			Imposto:    o.Tax,
			PrecoFinal: o.FinalPrice,
		})
	}
	return orders, nil
}

// Mutation returns a generated resolver.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns a generated resolver.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
