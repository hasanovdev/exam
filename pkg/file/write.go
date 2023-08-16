package file

import (
	"encoding/json"
	"os"
)

func Write(fileName string, objectMap map[string]interface{}) error {
	var objects []interface{}
	for _, val := range objectMap {
		objects = append(objects, val)
	}
	body, err := json.MarshalIndent(objects, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
