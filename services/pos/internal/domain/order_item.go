package domain

type OrderItem struct {
	ID       string
	OrderID  string
	Name     string
	Price    int64 // cents
	Quantity int
}
