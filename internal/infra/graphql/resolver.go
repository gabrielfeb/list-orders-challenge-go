package graphql

import (
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Resolver struct {
	ListOrdersUseCase usecase.ListOrdersUseCase
}

func NewResolver(listUC usecase.ListOrdersUseCase) *Resolver {
	return &Resolver{
		ListOrdersUseCase: listUC,
	}
}

func (r *Resolver) ListOrdersHandler() http.Handler {
	orderType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Order",
			Fields: graphql.Fields{
				"id":          &graphql.Field{Type: graphql.String},
				"price":       &graphql.Field{Type: graphql.Float},
				"tax":         &graphql.Field{Type: graphql.Float},
				"final_price": &graphql.Field{Type: graphql.Float},
			},
		},
	)

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"listOrders": &graphql.Field{
					Type: graphql.NewList(orderType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return r.ListOrdersUseCase.Execute()
					},
				},
			},
		},
	)

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
