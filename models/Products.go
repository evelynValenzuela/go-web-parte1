package models

type Product struct {
	Id 			int		`json:"id"`
	Name 		string	`json:"name"`
	Quantity 	int		`json:"quantity"`
	CodeValue 	string	`json:"code_value"`
	IsPublised 	bool	`json:"is_pusblised"`
	Expiration 	string	`json:"expiration"`
	Price 		float64	`json:"price"`
}