package dto

type OrderInputDTO struct {
	Price float64 `json:"preco"`
	Tax   float64 `json:"imposto"`
}

type OrderOutputDTO struct {
	ID         int     `json:"id"`
	Price      float64 `json:"preco"`
	Tax        float64 `json:"imposto"`
	FinalPrice float64 `json:"preco_final"`
}
