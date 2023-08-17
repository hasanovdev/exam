package controller

import (
	"exam/models"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.AppendHeader(table.Row{"no", "name", "total buy price"})

	i := 1
	for userId, totalSum := range res {
		t.AppendRow(table.Row{i, c.GetUsrName()[userId], totalSum})
		// fmt.Printf("Name: %s, Total Buy Price: %d\n", c.GetUsrName()[userId], totalSum)
		i++
	}

	fmt.Println(t.Render())
}
