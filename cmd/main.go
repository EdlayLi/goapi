package main

import (
	"apigo/configs"
	"apigo/internal/auth"
	"apigo/internal/link"
	"apigo/internal/user"
	"apigo/pkg/db"
	"apigo/pkg/middlewere"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	auth.NewAuthHundler(router, auth.AuthHundlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHundler(router, link.LinkHundlerDeps{
		LinkRepository: linkRepository,
	})

	//Midlewares
	stack := middlewere.Chein(
		middlewere.Cors,
		middlewere.Logging,
	)

	server := &http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
