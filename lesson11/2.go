package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		{
		    "manufacturer": "Apple",
		    "model": "Mcbook Air",
		    "technicalCharacteristics": {
		        "CPU": "M2",
		        "memory": "8"
		    }
		}
	*/
	type technicalCharacteristics struct {
		CPU    string
		Memory int
	}
	type PavelPhone struct {
		Manufacturer             string `json:"manuf"`
		Model                    string `json:"model"`
		TechnicalCharacteristics technicalCharacteristics
	}
	var newPhone PavelPhone = PavelPhone{Manufacturer: "Apple", Model: "Mcbook Air", TechnicalCharacteristics: technicalCharacteristics{CPU: "M2", Memory: 8}}
	fmt.Printf("%+v\n", newPhone)
	bytes, err := json.Marshal(newPhone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	//Unmarshall

	var myNewPhone PavelPhone
	err = json.Unmarshal(bytes, &myNewPhone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myNewPhone)
}
