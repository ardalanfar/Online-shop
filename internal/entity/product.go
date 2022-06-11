package entity

//Entity Product

type Product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Number   string `json:"number"`
	Describe string `json:"describe"`
}
