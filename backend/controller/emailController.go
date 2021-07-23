package controller

import (
	"encoding/json"
	"fmt"
)

func SendEmailBySendGrid(msg []byte) {
	var data EmailData
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
