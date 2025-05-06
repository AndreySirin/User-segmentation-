package main

import (
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/server"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/storage"
)

const (
	dbname   = "database"
	user     = "admin"
	password = "secret"
	address  = "ps:5432"
)

func main() {
	lg := logger.NewLogger()
	lg.Info("Start server")

	db, err := storage.New(lg, user, password, dbname, address)
	if err != nil {
		lg.Error("ошибка при подключении к базе", err)
	}
	defer db.Close()

	segm := storage.NewSegment(db)
	subsc := storage.NewSubscription(db)
	us := storage.NewUser(db)

	SEGMservice := service.NewSegment(lg, segm)
	SUBSCsirvice := service.NewSubscriptionService(lg, subsc)
	USERservice := service.NewUserService(lg, us)

	HUNDL := server.NewHNDL(lg, USERservice, SEGMservice, SUBSCsirvice)

	srv := server.NewServer(lg, ":8080", HUNDL)
	srv.Start()
}
