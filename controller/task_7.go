package controller

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (c *Controller) makeSlice7() []string {
	shopCartsMap := c.makeMap5()

	mySl := make([]string, 0, len(shopCartsMap))
	for prId := range shopCartsMap {
		mySl = append(mySl, prId)
	}

	sort.SliceStable(mySl, func(i, j int) bool {
		return shopCartsMap[mySl[i]] < shopCartsMap[mySl[j]]
	})

	return mySl
}

func (c *Controller) Task_7() {
	mySl := c.makeSlice7()
	shopCartsMap := c.makeMap5()

	t := table.NewWriter()
	color.Yellow("Top 10 ta eng past sotilayotgan mahsulotlar ro'yxati")

	t.AppendHeader(table.Row{"No", "name", "count"})

	i := 1
	for _, v := range mySl {
		t.AppendRow(table.Row{i, c.GetProduct()[v].Name, shopCartsMap[v]})
		i++
		if i == 11 {
			break
		}
	}

	fmt.Println(t.Render())
}
