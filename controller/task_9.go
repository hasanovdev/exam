package controller

import (
	"exam/pkg/file"
	"fmt"
)

func (c *Controller) makeMap9() map[string]int {
	shopCarts, _ := file.ReadShopCart("./data/shop_cart.json")

	myMap := make(map[string]int)
	for _, cart := range shopCarts {
		if cart.Status {
			myMap[c.GetProduct()[cart.ProductId].CategoryID] += cart.Count
		}
	}

	return myMap
}

func (c *Controller) Task_9() {
	shopCartsMap := c.makeMap9()

	for catId, count := range shopCartsMap {
		fmt.Printf("Name: %s, Count: %d\n", c.GetCatName()[catId], count)
	}
}
