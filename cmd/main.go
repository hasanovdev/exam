package main

import (
	"exam/config"
	"exam/controller"
	"exam/storage/memory"
	"fmt"
	"time"
)

func main() {
	cfg := config.Load()
	strg, err := memory.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	// con.Task_1()
	start := time.Now()
	con.Task_9()
	dur := time.Since(start)
	fmt.Println(dur)
	// res, err := con.Strg.User().GetList(&models.UserGetListRequest{
	// 	Offset: 1,
	// 	Limit:  100,
	// })

	// if err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	for _, v := range res.Users {
	// 		fmt.Println(*v)
	// 	}
	// }
}
