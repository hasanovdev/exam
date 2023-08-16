package controller

import (
	"exam/models"
	"fmt"
)

func (c *Controller) makeMap4() map[string]int {
	myMap := make(map[string]int)

	r, _ := c.Strg.Order().GetList(&models.OrderGetListRequest{Offset: 1, Limit: 1000})
	orders := r.Orders

	for _, order := range orders {
		if _, ok := myMap[order.UserId]; !ok {
			myMap[order.UserId] = order.Sum
		} else {
			myMap[order.UserId] += order.Sum
		}
	}

	return myMap
}

func (c *Controller) Task_4() {
	res := c.makeMap4()

	for userId, totalSum := range res {
		fmt.Printf("Name: %s, Total Buy Price: %d\n", c.GetUsrName()[userId], totalSum)
	}
}
