package repo

import (
	"INTERFACE/02-practical/models"
	"fmt"
)

type RealDbRepo struct{}

func NewRealDbRepo() *RealDbRepo {
	return &RealDbRepo{}
}

func (r *RealDbRepo) SaveItem(item models.ItemDbModel) {
	fmt.Println("Item saved to db: %s, Price: %.2f\n", item.Name, item.Price)
}
