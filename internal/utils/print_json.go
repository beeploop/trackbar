package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error marshalling JSON data: ", err)
		return
	}

	fmt.Println(string(b))
}
