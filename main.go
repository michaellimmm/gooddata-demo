package main

import (
	"context"
	"github/michaellimmm/gooddata-demo/internal/adapter"
	"github/michaellimmm/gooddata-demo/internal/repositories"
	"github/michaellimmm/gooddata-demo/internal/usecases"
	"github/michaellimmm/gooddata-demo/pkg/gooddata"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env", err)
		return
	}

	mongoDBUri := os.Getenv("MONGODB_URI")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		slog.Error("failed to initiate mongo client", err)
		return
	}
	defer mongoClient.Disconnect(ctx)

	dbName := os.Getenv("MONGODB_NAME")
	db := mongoClient.Database(dbName)

	repo := repositories.NewRepositories(db)

	gooddataBaseUrl := os.Getenv("GOODDATA_BASEURL")
	gooddataApiKey := os.Getenv("GOODDATA_ACCESSTOKEN")
	gooddataApi, err := gooddata.NewGooddataAPI(gooddataBaseUrl, gooddataApiKey)
	if err != nil {
		slog.Error("failed to initialize gooddataAPI", err)
		return
	}

	usecases := usecases.NewUsecases(*repo, gooddataApi)

	slog.Info("server is start")
	httpServer := adapter.NewHttpServer(usecases)

	go func() {
		if err := httpServer.Run(":8080"); err != nil {
			slog.Error("server run return error", err)
		}
	}()

	<-ctx.Done()

	if err := httpServer.Stop(context.TODO()); err != nil {
		slog.Error("failed to stop server", err)
	}

	slog.Info("server is shutdown")
}
