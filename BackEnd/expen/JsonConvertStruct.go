package expen

import (
	"encoding/json"
)

// json -> struct and struct -> json

// StructConvertJson struct->json
func StructConvertJson(Struct interface{}) (string, error) {
	res, err := json.Marshal(Struct)
	return string(res), err
}

// JsonConvertStruct json->struct
func JsonConvertStruct(Json []byte, Struct interface{}) error {
	return json.Unmarshal(Json, Struct)
}
