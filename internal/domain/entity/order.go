package entity

import "errors"

// Order representa a entidade de negócio principal do sistema.
// Contém a lógica e as regras de negócio intrínsecas a uma ordem,
// sendo independente de qualquer tecnologia externa.
type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// NewOrder é a função construtora para criar uma nova instância de Order.
// Ela garante que uma ordem só possa ser criada em um estado válido,
// chamando o método de validação.
func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

// Validate aplica as regras de negócio para garantir a integridade da entidade.
// Retorna um erro se qualquer uma das regras for violada.
func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id é obrigatório")
	}
	if o.Price <= 0 {
		return errors.New("preço deve ser maior que zero")
	}
	if o.Tax < 0 {
		return errors.New("imposto não pode ser negativo")
	}
	return nil
}

// CalculateFinalPrice executa a lógica de negócio para calcular o preço final da ordem.
func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
