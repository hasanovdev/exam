package controller

import (
	"exam/pkg/file"
	"fmt"
)

func (c *Controller) makeMap10() map[string]int {
	shopCarts, _ := file.ReadShopCart("./data/shop_cart.json")

	myMap := make(map[string]int)
	for _, cart := range shopCarts {
		if cart.Status {
			myMap[cart.UserId] += cart.Count
		}
	}
	return myMap
}

func (c *Controller) Task_10() {
	res := c.makeMap10()

	max := 0
	name := ""
	for usrId, count := range res {
		if count > max {
			max = count
			name = c.GetUsrName()[usrId]
		}
	}

	fmt.Printf("\nName: %s\nCount: %d\n", name, max)

}
