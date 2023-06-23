package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
	"github.com/vanyovan/mini-wallet.git/internal/handler"
	"github.com/vanyovan/mini-wallet.git/internal/middleware"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repo.NewUserRepo(db)
	walletRepo := repo.NewWalletRepo(db)

	WalletUsecase := usecase.NewWalletService(userRepo, walletRepo)
	userUsecase := usecase.NewUserService(userRepo)

	handler := handler.NewHandler(userUsecase, WalletUsecase)

	router := chi.NewRouter()

	router.Method(http.MethodPost, "/api/v1/init", http.HandlerFunc(handler.HandleInitWallet))

	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthenticateUser(&userUsecase))
		r.Method(http.MethodPost, "/api/v1/wallet", http.HandlerFunc(handler.HandleEnableWallet))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on", server.Addr)
	server.ListenAndServe()
}
