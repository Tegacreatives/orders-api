package main

import (
	"github.com/Tegacreatives/go-server/application"
	"context"
	"fmt"
)

func main(){
	 app := application.New()

	 err := app.Start(context.TODO())
	 if err != nil {
		fmt.Println("Failed to start app:", err)
	 }
}