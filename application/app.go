package application

import(
	"net/http"
	"fmt"
	"context"
)

type App struct{
	router http.Handler
}

func New() *App{
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error{
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Failed to connect to server: %w", err)
		}
		return nil
}