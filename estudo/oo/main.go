package main

import (
	"encoding/json"
	"fmt"
	"fosouzadev/oo/models"
)

func main() {
	user := new(models.User)

	fmt.Println(user)

	userJson, _ := json.Marshal(user)

	var otherUser models.User
	json.Unmarshal(userJson, &otherUser)

	fmt.Println(otherUser)
	fmt.Println(string(userJson))

	user.ForgotLogin()
}
