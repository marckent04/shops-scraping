package database

import (
	"gorm.io/gorm"
	"shops-scraping/shared"
)

type CartLine struct {
	gorm.Model
	shared.Article
}

func GetCartLines() (result []CartLine, err error) {
	err = db.Model(CartLine{}).Find(&result).Error
	return
}

func ClearCart() error {
	return db.Delete(&CartLine{}, "id IS NOT NULL").Error
}

func CreateCartLine(data shared.Article) (CartLine, error) {
	line := CartLine{
		Article: data,
	}
	err := db.Create(&line).Error
	return line, err
}

func DeleteCartLine(id int) error {
	line := CartLine{}
	return db.Delete(&line, id).Error
}
