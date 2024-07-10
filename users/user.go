package users

import (
	"fmt"
	"godesde0/models"
	"time"
)

func CreateUser() {
	u := new(models.User)
	u.AddUser(10, "Oscar", time.Now(), true)
	fmt.Println(u)
}
