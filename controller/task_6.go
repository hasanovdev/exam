package controller

import (
	"fmt"
	"sort"
)

func (c *Controller) makeSlice6() []string {
	shopCartsMap := c.makeMap5()

	mySl := make([]string, 0, len(shopCartsMap))
	for prId := range shopCartsMap {
		mySl = append(mySl, prId)
	}

	sort.SliceStable(mySl, func(i, j int) bool {
		return shopCartsMap[mySl[i]] > shopCartsMap[mySl[j]]
	})

	return mySl
}

func (c *Controller) Task_6() {
	mySl := c.makeSlice6()
	shopCartsMap := c.makeMap5()

	i := 1
	for _, v := range mySl {
		fmt.Printf("%d. Name: %s, Count: %d\n", i, c.GetProduct()[v].Name, shopCartsMap[v])
		i++
		if i == 11 {
			return
		}
	}
}
