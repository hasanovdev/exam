package controller

import (
	"exam/pkg/file"
	"fmt"
	"sort"
)

type productCount struct {
	productId string
	count     int
}

func (c *Controller) makeMap8() map[string]productCount {
	shopCarts, _ := file.ReadShopCart("./data/shop_cart.json")
	myMap := make(map[string]productCount)

	for _, cart := range shopCarts {
		if cart.Status {
			time := cart.Time[:10]
			if _, ok := myMap[time]; !ok {
				myMap[time] = productCount{
					productId: cart.ProductId,
					count:     cart.Count,
				}
			} else {
				if cart.ProductId == myMap[time].productId {
					c := cart.Count + myMap[time].count
					myMap[time] = productCount{
						productId: cart.ProductId,
						count:     c,
					}
				} else {
					if cart.Count > myMap[time].count {
						myMap[time] = productCount{
							productId: cart.ProductId,
							count:     cart.Count,
						}
					}
				}
			}
		}
	}
	return myMap
}

func (c *Controller) makeSlice8() []string {
	shopCartMap := c.makeMap8()
	mySlice := make([]string, 0, len(shopCartMap))
	for time := range shopCartMap {
		mySlice = append(mySlice, time)
	}

	sort.SliceStable(mySlice, func(i, j int) bool {
		return shopCartMap[mySlice[i]].count > shopCartMap[mySlice[j]].count
	})

	return mySlice
}

func (c *Controller) Task_8() {
	res := c.makeSlice8()
	myMap := c.makeMap8()

	i := 1
	for _, date := range res {
		fmt.Printf("%d. Name: %s, Time: %s, Count: %d\n", i, c.GetProduct()[myMap[date].productId].Name, date, myMap[date].count)
		i++
	}
}
