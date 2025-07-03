package graphql

import (
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func NewGraphQLHandler(listOrdersUC *usecase.ListOrdersUseCase) *Resolver {
	return &Resolver{ListOrdersUC: listOrdersUC}
}

func (r *Resolver) Handle() http.Handler {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"orders": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
						Name: "Order",
						Fields: graphql.Fields{
							"id":          &graphql.Field{Type: graphql.String},
							"price":       &graphql.Field{Type: graphql.Float},
							"tax":         &graphql.Field{Type: graphql.Float},
							"final_price": &graphql.Field{Type: graphql.Float},
						},
					})),
					Resolve: r.resolveOrders,
				},
			},
		}),
	})

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func (r *Resolver) resolveOrders(p graphql.ResolveParams) (interface{}, error) {
	return r.ListOrdersUC.Execute(p.Context)
}
