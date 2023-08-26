package main

import (
	"context"
	"fmt"
	"github.com/tomato3713/nullwiki/pkg/ent"
	"github.com/tomato3713/nullwiki/pkg/wiki"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	client, err := ent.Open("mysql", getDbConnStr())
	if err != nil {
		logger.Warn("failed opening connection to mysql: %v", err)
		panic(err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Warn("failed creating schema resources: %v", err)
	}
	logger.Info("the wiki program is starting")

	_ = wiki.Config{
		Logger:   logger,
		Context:  &ctx,
		DbClient: client,
	}
}

func getDbConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDR"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_TABLENAME"))
}
