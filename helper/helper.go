package helper

import (
	"encoding/json"
	"exam/models"
	"log"
	"os"
)

func GetPrName() map[string]string {
	products, _ := ReadProducts()
	res := make(map[string]string)
	for _, product := range products {
		res[product.Id] = product.Name
	}
	return res
}

func GetUsrName() map[string]string {
	users, _ := ReadUsers()
	res := make(map[string]string)
	for _, user := range users {
		res[user.Id] = user.FirstName + " " + user.LastName
	}
	return res
}

func ReadOrders() ([]models.Order, error) {
	var orders []models.Order

	data, err := os.ReadFile("./data/order.json")
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return orders, nil
}

func ReadProducts() ([]models.Product, error) {
	var products []models.Product

	data, err := os.ReadFile("./data/product.json")
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return products, nil
}

func ReadUsers() ([]models.User, error) {
	var users []models.User

	data, err := os.ReadFile("./data/user.json")
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return users, nil
}
