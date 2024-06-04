package entity

// GroceryInfo 商品的实体类
type GroceryInfo struct {
	//CommonInfo
	Id          int64
	Name        string
	Description string
	Price       int64
	ImageURL    string
	Category    string
}
