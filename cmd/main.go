package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
	"github.com/vanyovan/mini-wallet.git/internal/handler"
	"github.com/vanyovan/mini-wallet.git/internal/middleware"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
	"github.com/vanyovan/mini-wallet.git/internal/repo/wrapper"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", getAppRootDirectory()+"/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlWrapper := wrapper.NewSqlWrapper(db)
	userRepo := repo.NewUserRepo(db)
	walletRepo := repo.NewWalletRepo(db)
	transactionRepo := repo.NewTransactionRepo(db)

	walletUsecase := usecase.NewWalletService(userRepo, walletRepo)
	userUsecase := usecase.NewUserService(userRepo)
	transactionUsecase := usecase.NewTransactionService(userRepo, walletRepo, transactionRepo, sqlWrapper)

	handler := handler.NewHandler(userUsecase, walletUsecase, transactionUsecase)

	router := chi.NewRouter()

	router.Method(http.MethodPost, "/api/v1/init", http.HandlerFunc(handler.HandleInitWallet))

	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthenticateUser(&userUsecase))
		r.Method(http.MethodPost, "/api/v1/wallet", http.HandlerFunc(handler.HandleEnableWallet))                 //enable wallet
		r.Method(http.MethodGet, "/api/v1/wallet", http.HandlerFunc(handler.HandleViewWallet))                    //view wallet
		r.Method(http.MethodGet, "/api/v1/wallet/transactions", http.HandlerFunc(handler.HandleViewTransaction))  //view transaction
		r.Method(http.MethodPost, "/api/v1/wallet/deposits", http.HandlerFunc(handler.HandleDepositWallet))       //deposit wallet
		r.Method(http.MethodPost, "/api/v1/wallet/withdrawals", http.HandlerFunc(handler.HandleWithdrawalWallet)) //withdraw wallet
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on", server.Addr)
	server.ListenAndServe()
}

func getAppRootDirectory() string {
	projectName := regexp.MustCompile(`^(.*mini-wallet)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
