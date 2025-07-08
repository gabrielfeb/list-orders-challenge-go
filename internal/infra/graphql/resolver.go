package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"list-orders-challenge-go/internal/usecase"
)

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewResolver(createUC usecase.CreateOrderUseCase, listUC usecase.ListOrdersUseCase) *Resolver {
	return &Resolver{
		CreateOrderUseCase: createUC,
		ListOrdersUseCase:  listUC,
	}
}

func (r *Resolver) CreateOrder(params graphql.ResolveParams) (interface{}, error) {
	input, _ := params.Args["input"].(map[string]interface{})

	dto := usecase.CreateOrderInputDTO{
		Price: input["price"].(float64),
		Tax:   input["tax"].(float64),
	}

	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *Resolver) ListOrders(params graphql.ResolveParams) (interface{}, error) {
	output, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *Resolver) Handler() *handler.Handler {
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.String},
			"price":       &graphql.Field{Type: graphql.Float},
			"tax":         &graphql.Field{Type: graphql.Float},
			"final_price": &graphql.Field{Type: graphql.Float},
		},
	})

	queryFields := graphql.Fields{
		"listOrders": &graphql.Field{
			Type:    graphql.NewList(orderType),
			Resolve: r.ListOrders,
		},
	}

	mutationFields := graphql.Fields{
		"createOrder": &graphql.Field{
			Type: orderType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name: "OrderInput",
							Fields: graphql.InputObjectConfigFieldMap{
								"price": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
								"tax":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
							},
						},
					)),
				},
			},
			Resolve: r.CreateOrder,
		},
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields})
	rootMutation := graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
