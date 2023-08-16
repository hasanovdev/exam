package controller

import (
	"exam/pkg/file"
	"fmt"

	"github.com/fatih/color"
)

type itemSum struct {
	Price int
	Count int
	Total int
	Time  string
}

func (c *Controller) makeMap3() map[string]map[string]itemSum {
	shopCarts, _ := file.ReadShopCart("./data/shop_cart.json")

	myMap := make(map[string]map[string]itemSum)

	price, count := 0, 0
	for _, cart := range shopCarts {
		if _, ok := myMap[cart.UserId]; !ok {
			price, count = 0, 0

			price += c.GetProduct()[cart.ProductId].Price
			count += cart.Count

			valueMap := make(map[string]itemSum)

			valueMap[cart.ProductId] = itemSum{
				Price: price,
				Count: count,
				Total: price * count,
				Time:  cart.Time,
			}

			myMap[cart.UserId] = valueMap

		} else {
			if _, ok := myMap[cart.UserId][cart.ProductId]; !ok {
				price, count = 0, 0

				price += c.GetProduct()[cart.ProductId].Price
				count += cart.Count

				myMap[cart.UserId][cart.ProductId] = itemSum{
					Price: price,
					Count: count,
					Total: price * count,
					Time:  cart.Time,
				}
			} else {

				price, count = myMap[cart.UserId][cart.ProductId].Price, myMap[cart.UserId][cart.ProductId].Count

				price += myMap[cart.UserId][cart.ProductId].Price
				count += myMap[cart.UserId][cart.ProductId].Count

				myMap[cart.UserId][cart.ProductId] = itemSum{
					Price: price,
					Count: count,
					Total: price * count,
					Time:  cart.Time,
				}
			}
		}
	}

	return myMap
}

func (c *Controller) Task_3() {
	res := c.makeMap3()

	i := 1
	for userId, valuMap := range res {
		color.Green("%s\n", c.GetUsrName()[userId])
		for productId, Summ := range valuMap {
			fmt.Printf("%d. Name:%s, Price:%d, Count:%d, Total:%d, Time:%s\n", i, c.GetProduct()[productId].Name, Summ.Price, Summ.Count, Summ.Total, Summ.Time)
			i++
		}
	}
}
