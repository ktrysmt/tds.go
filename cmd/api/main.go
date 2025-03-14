package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"tds.go/config"
	"tds.go/internal/infrastructure"
	"tds.go/internal/infrastructure/database"
	"tds.go/internal/infrastructure/logger"
	"tds.go/internal/infrastructure/middleware"
	"tds.go/internal/presentation"
	"tds.go/internal/usecase"
)

func main() {
	// 設定のロード
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// ロガーの初期化
	logger.Init("development")

	// ルーターの設定
	router := mux.NewRouter()

	// ミドルウェアの適用
	router.Use(middleware.Logger)
	router.Use(middleware.Auth)

	// 依存関係の注入とハンドラーの設定
	db := database.NewPostgresDB(cfg)
	userRepo := infrastructure.NewPostgresUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := presentation.NewUserHandler(userUseCase)

	// ルーティング
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// サーバーの起動
	logger.Info("Server starting on port " + cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, router)
}
