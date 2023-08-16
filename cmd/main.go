package main

import (
	"exam/tasks"
)

func main() {
	// cfg := config.Load()
	// strg, err := memory.NewConnectionJSON(&cfg)
	// if err != nil {
	// 	panic("Failed connect to json:" + err.Error())
	// }
	// con := controller.NewController(&cfg, strg)

	// resp, _ := con.Strg.Order().Create(&models.CreateOrder{
	// 	UserId:   "204ff9b0-3f4e-41b3-a436-3a1fce028fa6",
	// 	Sum:      43000,
	// 	SumCount: 3,
	// 	DateTime: "2023-08-14 13:10:32",
	// 	Status:   "1",
	// })

	// fmt.Println(*resp)

	tasks.Task_1()
}
