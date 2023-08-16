package file

import (
	"encoding/json"
	"exam/models"
	"log"
	"os"

	"github.com/spf13/cast"
)

func Read(fileName string) (map[string]interface{}, error) {

	var (
		objects   []interface{}
		objectMap = make(map[string]interface{})
	)
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, object := range objects {
		obj := cast.ToStringMap(object)
		objectMap[cast.ToString(obj["id"])] = object
	}

	return objectMap, nil
}

func ReadOrder(fileName string) ([]models.Order, error) {
	orders := []models.Order{}

	data, err := os.ReadFile(fileName)
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
