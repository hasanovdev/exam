package controller

import (
	"exam/config"
	"exam/storage"
)

type Controller struct {
	Cfg  *config.Config
	Strg storage.StorageI
}

func NewController(cfg *config.Config, storage storage.StorageI) *Controller {
	return &Controller{
		Cfg:  cfg,
		Strg: storage,
	}
}
