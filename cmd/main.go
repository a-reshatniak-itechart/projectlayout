package main

import (
	"database/sql"
	"net/http"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
	"github.com/a-reshatniak-itechart/projectlayout/internal/app/controller"
	httppkg "github.com/a-reshatniak-itechart/projectlayout/internal/app/http"
	"github.com/a-reshatniak-itechart/projectlayout/internal/app/localfs"
	"github.com/a-reshatniak-itechart/projectlayout/internal/app/sqlite"
)

func main() {
	// TODO: init config
	var cfg Config

	userRepo := initUserRepo(cfg)
	userController := controller.NewUser(userRepo)
	userHandler := httppkg.NewUserHandler(userController)

	httppkg.RegisterRouter(userHandler)

	err := http.ListenAndServe(cfg.ListenAddr, nil)
	if err != nil {
		panic(err)
	}
}

func initUserRepo(cfg Config) app.UserRepository {
	switch cfg.UserRepoType {
	case "sqlite":
		conn, err := sql.Open("sqlite", cfg.UserSqliteDsn)
		if err != nil {
			panic(err)
		}

		return sqlite.NewUserRepository(conn)
	case "local_fs":
		return localfs.NewUserRepository(cfg.UserLocalFsDir)
	default:
		panic("unknown user repo type")
	}
}
