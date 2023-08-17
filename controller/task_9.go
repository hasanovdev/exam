package controller

import (
	"exam/pkg/file"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.AppendHeader(table.Row{"no", "name", "count"})

	i := 1
	for catId, count := range shopCartsMap {
		t.AppendRow(table.Row{i, c.GetCatName()[catId], count})
		i++
	}

	fmt.Println(t.Render())
}
