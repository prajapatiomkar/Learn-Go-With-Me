package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS", 299, "LCO.in", "abc123", []string{"webdev", "js"}},
		{"MERN", 399, "LCO.in", "pqr321", []string{"full-stack", "js"}},
		{"Angular", 499, "LCO.in", "xyz456", nil},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //prefix 2nd "" if we pass value it will print every line
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFormate := []byte(`
	{
		"name": "React JS",
		"Price": 299,
		"Platform": "LCO.in",
		"tags": [
				"webdev",
				"js"
		]
}
	`)
	var lcoCourse course 

	checkValid := json.Valid(jsonDataFormate)
	
	if checkValid{
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFormate,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("Json was not valid")
	}
	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormate,&myOnlineData)

	fmt.Printf("%#v\n",myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("key is %v value is %v and type is %T\n", k,v,v)
	}
}
