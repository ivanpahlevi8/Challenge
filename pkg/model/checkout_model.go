package model

type CheckoutModel struct {
	TotalPrice float64
	AllItems   []string
}

// create get set method for total price
func (checkout *CheckoutModel) GetTotalPrice() float64 {
	return checkout.TotalPrice
}

func (checkout *CheckoutModel) SetTotalPrice(totalPrice float64) {
	checkout.TotalPrice = totalPrice
}

// create get set method for all items
func (checkout *CheckoutModel) GetAllItems() []string {
	return checkout.AllItems
}

func (checkout *CheckoutModel) SetAllItems(allItems []string) {
	checkout.AllItems = allItems
}
