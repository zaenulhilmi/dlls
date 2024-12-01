package main

import (
	"dlls/contracts"
	"dlls/handlers"
	"dlls/middlewares"
	"dlls/repositories"
	"dlls/services"
	"dlls/utils"
	"fmt"
	"net/http"
)

func main() {
	port := "8010"
	fmt.Printf("Starting server on http://localhost:%s\n", port)

	mux := http.NewServeMux()

	userRepository := repositories.NewMemUserRepository()
	actionRepository := repositories.NewMemActionRepository()
	jwt := utils.NewJWT()
	hasher := utils.NewHasher()

	actionService := services.NewActionService(userRepository, actionRepository, 10)
	subscriptionService := services.NewSubscriptionService(userRepository)
	authService := services.NewAuthService(userRepository, hasher, jwt)
	userService := services.NewUserService(userRepository)

	seedUsers(userRepository, hasher)

	authHandler := handlers.NewAuthHandler(authService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	actionHandler := handlers.NewActionHandler(actionService, userService)

	mux.HandleFunc("POST /api/v1/login", authHandler.Login)
	mux.HandleFunc("POST /api/v1/signup", authHandler.SignUp)

	mux.HandleFunc("POST /api/v1/subscribe", middlewares.AuthMiddleware(subscriptionHandler.Subscribe, jwt))

	mux.HandleFunc("POST /api/v1/like", middlewares.AuthMiddleware(actionHandler.Like, jwt))
	mux.HandleFunc("POST /api/v1/pass", middlewares.AuthMiddleware(actionHandler.Pass, jwt))
	mux.HandleFunc("GET /api/v1/preview-next-profile", middlewares.AuthMiddleware(actionHandler.PreviewNextProfile, jwt))

	corsMux := enableCORS(mux)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsMux)

	if err != nil {
		fmt.Println("Error starting server")
		fmt.Println(err)
	} else {
		fmt.Printf("Server running on port %s\n", port)
	}
}

func seedUsers(userRepository contracts.UserRepository, hasher contracts.Hasher) {

	for i := 0; i < 100; i++ {
		user := contracts.User{
			ID:           fmt.Sprintf("user-%d", i),
			Email:        fmt.Sprintf("email-%d@example.com", i),
			PasswordHash: hasher.Hash("password"),
			Name:         fmt.Sprintf("name-%d", i),
		}

		userRepository.Save(user)
	}

}

func enableCORS(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})
}
