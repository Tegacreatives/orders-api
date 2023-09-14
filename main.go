package main

import (
	"github.com/Tegacreatives/go-server/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main(){
	 app := application.New()

	 ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	 defer cancel()

	 err := app.Start(ctx)
	 if err != nil {
		fmt.Println("Failed to start app:", err)
	 }
}