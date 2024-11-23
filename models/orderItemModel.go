package models

type OrderItem struct {
	ID        string   `gorm:"type:text;primaryKey" json:"id"`
	OrderID   string   `gorm:"type:text;notNull" json:"orderId"`
	Order     Order    `gorm:"foreignKey:OrderID;references:ID;constraint:onDelete:Cascade"`
	ProductID string   `gorm:"type:text;notNull" json:"productId"`
	Product   Product  `gorm:"foreignKey:ProductID;references:ID;constraint:onDelete:Cascade" json:"product"`
	Quantity  int      `gorm:"type:integer;notNull" json:"quantity"`
}

type OrderItemResponse struct {
	ID			string				`json:"id"`
	OrderID		string				`json:"orderId"`
	ProductID	string				`json:"productId"`
	Quantity	int					`json:"quantity"`
	Product		ProductResponse		`json:"product"`
}