package model

type Product struct {
	ID          int    `gorm:"column:id;primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Price       int    `gorm:"column:price" json:"price"`
}
