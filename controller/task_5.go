package controller

import (
	"exam/pkg/file"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.AppendHeader(table.Row{"No", "Name", "Count"})
	i := 1
	for prId, count := range shopCartsMap {
		t.AppendRow(table.Row{i, c.GetProduct()[prId].Name, count})
		// fmt.Printf("Name: %s, Count: %d\n", c.GetProduct()[prId].Name, count)
		i++
	}

	fmt.Println(t.Render())
}
