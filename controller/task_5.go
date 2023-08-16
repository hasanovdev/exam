package controller

import (
	"exam/pkg/file"
	"fmt"
)

func (c *Controller) makeMap5() map[string]int {
	shopCarts, _ := file.ReadShopCart("./data/shop_cart.json")

	myMap := make(map[string]int)

	for _, cart := range shopCarts {
		if cart.Status {
			myMap[cart.ProductId] += cart.Count
		}
	}

	return myMap
}

func (c *Controller) Task_5() {
	shopCartsMap := c.makeMap5()

	for prId, count := range shopCartsMap {
		fmt.Printf("Name: %s, Count: %d\n", c.GetProduct()[prId].Name, count)
	}
}
