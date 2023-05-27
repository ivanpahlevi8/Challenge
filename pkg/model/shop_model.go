package model

// create shopping char list model
type ShopModel struct {
	Id       string   `json:"id"`
	AllItems []string `json:"all_items"`
}

// create method for get and set id
func (shop *ShopModel) GetId() string {
	return shop.Id
}

func (shop *ShopModel) SetId(id string) {
	shop.Id = id
}

// createg method for get and set allItems
func (shop *ShopModel) GetAllItems() []string {
	return shop.AllItems
}

func (shop *ShopModel) SetAllItems(allItems []string) {
	shop.AllItems = allItems
}

// create method for get total item
func (shop *ShopModel) GetTotalItem() int32 {
	numItem := int32(len(shop.AllItems))

	return numItem
}

// // create method to get first data
// func (shop *ShopModel) GetFirstData() ItemModel {
// 	getItem := shop.AllItems[0]

// 	return getItem
// }

// // create method to get last data
// func (shop *ShopModel) GetLastData() ItemModel {
// 	getItem := shop.AllItems[shop.GetTotalItem()-1]

// 	return getItem
// }

// // create method to get data at current index
// func (shop *ShopModel) GetAtIndex(index int) ItemModel {
// 	getItem := shop.AllItems[index]

// 	return getItem
// }
