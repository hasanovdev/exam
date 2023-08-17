package main

import (
	"exam/config"
	"exam/controller"
	"exam/models"
	"exam/storage/memory"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := memory.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	fmt.Println("tasks number: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11")
	var num int
	fmt.Print("Enter task number: ")
	fmt.Scan(&num)
	if num == 1 {
		con.Task_1()
	} else if num == 2 {
		con.Task_2()
	} else if num == 3 {
		con.Task_3()
	} else if num == 4 {
		con.Task_4()
	} else if num == 5 {
		con.Task_5()
	} else if num == 6 {
		con.Task_6()
	} else if num == 7 {
		con.Task_7()
	} else if num == 8 {
		con.Task_8()
	} else if num == 9 {
		con.Task_9()
	} else if num == 10 {
		con.Task_10()
	} else if num == 11 {
		err := con.OrderPayment(&models.OrderPayment{OrderId: "c8ffa9a0-4a9d-4dc5-b4f0-f3603014a02c"})
		if err != nil {
			fmt.Println(err)
		}
	}
}
