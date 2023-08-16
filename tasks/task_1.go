package tasks

import (
	"exam/helper"
	"exam/models"
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
)

func makeOrderSl() []models.Order {
	orders, _ := helper.ReadOrders()

	mySl := make([]models.Order, 0, len(orders))
	for _, order := range orders {
		mySl = append(mySl, order)
	}

	sort.SliceStable(mySl, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02 15:04:05", orders[i].DateTime)
		dateJ, _ := time.Parse("2006-01-02 15:04:05", orders[j].DateTime)

		return dateI.After(dateJ)
	})
	return mySl
}

func Task_1() {
	orders := makeOrderSl()
	usersMap := helper.GetUsrName()
	productsMap := helper.GetPrName()

	for _, order := range orders {
		color.Yellow("date: %s", order.DateTime)
		fmt.Println("user name:", usersMap[order.UserId])
		fmt.Println("status:", order.Status)
		fmt.Println("summ count:", order.SumCount)
		for _, item := range order.OrderItems {
			fmt.Printf("Product name: %s\nCount: %d\nPrice: %d\n", productsMap[item.ProductId], item.Count, item.TotalPrice)
		}

		fmt.Println("Total summ:", order.Sum)
		fmt.Println("------------------------------------------------------")
	}
}
