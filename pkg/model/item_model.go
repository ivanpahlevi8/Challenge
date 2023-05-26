package model

// create model for item shopping
type ItemModel struct {
	Id           string  `json:"id"`
	ItemName     string  `json:"item_name"`
	ItemCatagory string  `json:"item_category"`
	ItemPrice    float64 `json:"item_price"`
	ItemQuantity int32   `json:"item_quantity"`
}

// create get set method for id
func (item *ItemModel) GetId() string {
	return item.Id
}

func (item *ItemModel) SetId(id string) {
	item.Id = id
}

// create get set method for item name
func (item *ItemModel) GetItemName() string {
	return item.ItemName
}

func (item *ItemModel) SetItemName(itemName string) {
	item.ItemName = itemName
}

// create get methid for item category
func (item *ItemModel) GetItemCategory() string {
	return item.ItemName
}

func (item *ItemModel) SetItemCategory(itemName string) {
	item.ItemName = itemName
}

// create get set method for item price
func (item *ItemModel) GetItemPrice() float64 {
	return item.ItemPrice
}

func (item *ItemModel) SetItemPrice(itemPrice float64) {
	item.ItemPrice = itemPrice
}

// create get set method for item quantity
func (item *ItemModel) GetItemQuantity() int32 {
	return item.ItemQuantity
}

func (item *ItemModel) SetItemQuantity(itemQuantity int32) {
	item.ItemQuantity = itemQuantity
}
