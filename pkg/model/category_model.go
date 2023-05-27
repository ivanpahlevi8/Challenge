package model

type CategoryModel struct {
	Category string
	AllItems []ItemModel
}

// create get set method for total price
func (category *CategoryModel) GetCategory() string {
	return category.Category
}

func (category *CategoryModel) SetCategory(categ string) {
	category.Category = categ
}

// create get set method for all items
func (checkout *CategoryModel) GetAllItemsCategory() []ItemModel {
	return checkout.AllItems
}

func (checkout *CategoryModel) SetAllItemsCategory(allItems []ItemModel) {
	checkout.AllItems = allItems
}
