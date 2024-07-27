package converters

import (
	"encoding/json"
	"log"
)

func MapToString(mapData map[string]string) string {
	str, err := json.Marshal(mapData)
	if err != nil {
		log.Fatal("Could not convert this map data to string")
	}
	return string(str)
}

func StringToMap(stringData string) map[string]string {
	var out map[string]string
	_ = json.Unmarshal([]byte(stringData), &out)
	// if err != nil {
	// 	log.Fatal("Could not convert this string data to map;The error is", err.Error())
	// }
	return out
}

func ListMapToString(listMapData map[string][]string) string {
	str, err := json.Marshal(listMapData)
	if err != nil {
		log.Fatal("Could not convert this list-map data to string.")
	}
	return string(str)
}

func StringToListMap(stringData string) map[string][]string {
	var out map[string][]string
	_ = json.Unmarshal([]byte(stringData), &out)
	// if err != nil {
	// 	log.Fatal("Could not convert this string data to map;The error is", err.Error())
	// }
	return out
}
