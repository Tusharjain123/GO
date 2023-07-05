package controller

import (
	"context"
	"fmt"
	"log"
	"os/user"
)

func inserData(users user.User) {
	inserted, err := collection.InsertOne(context.Background(), users)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted)
}
