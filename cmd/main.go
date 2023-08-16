package main

import (
	"exam/config"
	"exam/controller"
	"exam/storage/memory"
)

func main() {
	cfg := config.Load()
	strg, err := memory.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	// con.Task_1()
	con.Task_8()
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
