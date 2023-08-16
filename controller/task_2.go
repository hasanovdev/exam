package controller

import (
	"exam/models"
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
)

func (c *Controller) makeOrderSl2(fromDate, toDate string) []models.Order {
	r, _ := c.Strg.Order().GetList(&models.OrderGetListRequest{Offset: 1, Limit: 1000})
	orders := r.Orders

	fDate, _ := time.Parse("2006-01-02", fromDate)
	tDate, _ := time.Parse("2006-01-02", toDate)
	mySl := make([]models.Order, 0, len(orders))
	for _, order := range orders {
		dateOrder, _ := time.Parse("2006-01-02 15:04:05", order.DateTime)
		if dateOrder.After(fDate) && dateOrder.Before(tDate) {
			mySl = append(mySl, order)
		}
	}

	sort.SliceStable(mySl, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02 15:04:05", orders[i].DateTime)
		dateJ, _ := time.Parse("2006-01-02 15:04:05", orders[j].DateTime)

		return dateI.Before(dateJ)
	})

	return mySl
}

func (c *Controller) Task_2() {
	var fDate, tDate string
	fmt.Print("Enter From Date: ")
	fmt.Scan(&fDate)
	fmt.Print("Enter To Date: ")
	fmt.Scan(&tDate)

	orders := c.makeOrderSl2(fDate, tDate)
	usersMap := c.GetUsrName()
	productsMap := c.GetProduct()

	for _, order := range orders {
		color.Yellow("date: %s", order.DateTime)
		fmt.Println("user name:", usersMap[order.UserId])
		fmt.Println("status:", order.Status)
		fmt.Println("summ count:", order.SumCount)
		for _, item := range order.OrderItems {
			fmt.Printf("Product name: %s\nCount: %d\nPrice: %d\n", productsMap[item.ProductId].Name, item.Count, item.TotalPrice)
		}

		fmt.Println("Total summ:", order.Sum)
		fmt.Println("------------------------------------------------------")
	}
}
