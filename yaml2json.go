/*
Package yaml2json provides functionality to transform YAML to JSON.
*/
package yaml2json

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Transform transforms an YAML input and transforms it to JSON.
func Transform(p []byte) ([]byte, error) {
	var yamlData interface{}
	if err := yaml.Unmarshal(p, &yamlData); err != nil {
		return nil, err
	}
	jsonData, err := transform(yamlData)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonData)
}

func transform(yamlData interface{}) (interface{}, error) {
	switch yamlData.(type) {
	case map[interface{}]interface{}:
		jsonData := make(map[string]interface{})
		for key, yamlValue := range yamlData.(map[interface{}]interface{}) {
			jsonValue, err := transform(yamlValue)
			if err != nil {
				return nil, err
			}
			switch key.(type) {
			case string:
				jsonData[key.(string)] = jsonValue
			case int:
				jsonData[strconv.Itoa(key.(int))] = jsonValue
			default:
				return nil, fmt.Errorf("yaml2json: unexpected key type %T for %v", key, key)
			}
		}
		return jsonData, nil
	case []interface{}:
		yamlDataSlice := yamlData.([]interface{})
		jsonDataSlice := make([]interface{}, len(yamlDataSlice))
		for i, yamlValue := range yamlDataSlice {
			jsonValue, err := transform(yamlValue)
			if err != nil {
				return nil, err
			}
			jsonDataSlice[i] = jsonValue
		}
		return jsonDataSlice, nil
	default:
		return yamlData, nil
	}
}
