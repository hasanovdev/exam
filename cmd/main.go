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

	con.Task_1()

}
