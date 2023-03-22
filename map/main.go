package main

import (
	"encoding/json"
	"fmt"
)

var myMap = make(map[string]string)

func main() {
	fmt.Println("learn with go map basic : ")

	// define map

	sample := map[string]string{
		"name": "saiful",
		"dept": "cse",
		"dist": "tangail",
	}

	// iteration map with value and key

	for key, item := range sample {
		fmt.Printf("key = %s, value = %s\n", key, item)
	}

	// iteration for only key

	for key := range sample {
		fmt.Printf("key = %s\n", key)
	}

	// iteration for only value

	for _, val := range sample {
		fmt.Printf("value = %s\n", val)
	}

	// get all keys from a map

	result := getAllKeys(sample)

	fmt.Printf("all keys list are : %v\n", result)

	// calculate the length of map

	lenMap := len(sample)

	fmt.Printf("the length of map is : %d", lenMap)

	// there are two ways to map define

	// make mthod or map literal

	firstMap := map[string]string{}

	fmt.Println("the map before adding value", firstMap)

	// adding a value
	firstMap["name"] = "saiful"

	fmt.Println("the map after adding value", firstMap)

	// map declared with  make mathod

	secondmap := make(map[string]int)

	fmt.Println("the map before second map : ", secondmap)

	// adding value into second map

	secondmap["student"] = 3

	fmt.Println("the map after adding value for second map: ", secondmap)

	// delete map key

	delete(secondmap, "student")

	fmt.Println("secnod map after delete key ", secondmap)

	// key exists into map

	_, ok := secondmap["student"]

	if ok {
		fmt.Println("the key is exist ")
	} else {
		fmt.Println("key does not exist")
	}

	// map are referenced data type
	// when a map assing into another variable if change into original map then also reflect into

	thirdMap := map[string]string{}
	thirdMap["dept"] = "cse"

	copyMap := thirdMap

	thirdMap["rollNo"] = "ce-14029"

	fmt.Println("the third map is :", thirdMap)
	fmt.Println("the copy map is : ", copyMap)

	//

	go setValue("dept", "cse")

	go getValue("dept")

	//fmt.Println("the key value is = ", resVal)

	// map convert to json
	// map key value can be integer while json not

	simpleMap := map[int]string{}

	simpleMap[1] = "john"
	simpleMap[0] = "ragib"

	fmt.Println("simple map is :", simpleMap)

	// map to json

	j, _ := json.Marshal(simpleMap)
	fmt.Println("convert from map to json :", string(j))

	// convert json to map

	var jsonToMap map[string]string
	json.Unmarshal(j, &jsonToMap)
	fmt.Println("json to map :", jsonToMap)

	// create struct

	type employee struct {
		Name string
	}

	employeeMap := make(map[string]employee)
	employeeMap["1"] = employee{Name: "saiful"}

	fmt.Println("employee map : ", employeeMap)

	emplJson, err := json.Marshal(employeeMap)
	if err != nil {

	} else {
		fmt.Println("json employee maps : ", string(emplJson))
	}

	// json to map

	customJson := `{"1" : "john"}`
	var mapNew map[string]string

	json.Unmarshal([]byte(customJson), &mapNew)

	fmt.Println("convert json to map : ", mapNew)

	// employee struct

	type student struct {
		Name string
	}

	studentJson := `{"1":{"Name":"saiful"}}`
	fmt.Println("convert json to map : ", studentJson)
	var studentmap map[int]student

	json.Unmarshal([]byte(studentJson), &studentmap)
	fmt.Println("your student map is : ", studentmap)

	// if we unknown the json value type then using empty interface type for map like as generic

	rawData := `{"0":"rafi"}`

	var myData map[int]interface{}
	json.Unmarshal([]byte(rawData), &myData)
	fmt.Println("unknown data : ", myData)

	// delete key from map

	mapElements := map[string]int{
		"a": 2,
		"b": 3,
		"c": 4,
	}

	// At first check if key exist then we delete

	if _, ok := mapElements["a"]; ok {
		delete(mapElements, "a")
	}

	fmt.Println("the new map is after delete: ", mapElements)

}

func getValue(key string) string {
	return myMap[key]
}

func setValue(key string, val string) {
	myMap[key] = val
}

func getAllKeys(sam map[string]string) []string {
	var keysList []string

	for key := range sam {
		keysList = append(keysList, key)
	}

	return keysList
}
