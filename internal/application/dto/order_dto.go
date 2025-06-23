package dto

type OrderInputDTO struct{ Price, Tax float64 }
type OrderOutputDTO struct{ ID, Price, Tax, FinalPrice float64 }
