package main

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "example/moviecrud/config"
    "example/moviecrud/controllers"
    "example/moviecrud/middleware"
    "example/moviecrud/repository"
    "example/moviecrud/routes"
    "example/moviecrud/services"

    "github.com/gorilla/mux"
    "go.uber.org/zap"
)

func main() {
    // Logger
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    sugar := logger.Sugar()

    // DB
    config.ConnectDB()

    // Repositories
    userRepo := repository.NewUserMongoRepo(config.GetCollection("users"))

    // Services
    userService := services.NewUserService(userRepo)
    authService := services.NewAuthService(userRepo)

    // Controllers
    userCtrl := controllers.NewUserController(userService)
    authCtrl := controllers.NewAuthController(authService)

    // Router
    r := mux.NewRouter()

    // Middlewares
    r.Use(middleware.Recovery(sugar))
    r.Use(middleware.CORS())
    r.Use(middleware.Logger(sugar))

    // Routes
    routes.RegisterUserRoutes(r, userCtrl)
    routes.RegisterAuthRoutes(r, authCtrl)

    // Server
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    go func() {
        sugar.Info("Server starting on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            sugar.Fatal("Server failed:", err)
        }
    }()

    // Graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    <-c

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        sugar.Fatal("Server forced to shutdown:", err)
    }
    sugar.Info("Server stopped gracefully")
}