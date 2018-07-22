package models

type Product struct {
	ID uint `gorm:"primary_key"`
	Code string
	Price uint
}
type Product2 struct {
	ID uint `gorm:"primary_key"`
	Code string
	Price uint

}
func (Product2) TableName() string {
	return "xproduct"
}

func (p Product) TableName() string {
	if p.Code == "L1212" {
		return "x_product1"
	} else {
		return "x_product2"
	}
}