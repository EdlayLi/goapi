package main

import (
	"apigo/configs"
	"apigo/internal/auth"
	"apigo/internal/link"
	"apigo/internal/stat"
	"apigo/internal/user"
	"apigo/pkg/db"
	"apigo/pkg/event"
	"apigo/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	//Middlewares
	stack := middleware.Chain(
		middleware.Cors,
		middleware.Logging,
	)

	return stack(router)
}

func main() {
	app := App()

	server := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
