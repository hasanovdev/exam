package controller

import (
	"exam/pkg/file"
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	day := color.YellowString("time")
	t.AppendHeader(table.Row{"no", "name", day, "count"})

	i := 1
	for _, date := range res {
		t.AppendRow(table.Row{i, c.GetProduct()[myMap[date].productId].Name, color.YellowString(date), myMap[date].count})
		i++
	}

	fmt.Println(t.Render())
}
