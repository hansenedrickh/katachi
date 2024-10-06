package main

import (
	"os"

	"github.com/hansenedrickh/katachi/config"
	"github.com/hansenedrickh/katachi/dependencies"
	"github.com/hansenedrickh/katachi/server"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	cfg := config.Load()

	app := setupCli(cfg)

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatalf("application ends with an error: %v", err)
	}
}

func setupCli(cfg config.Config) *cli.App {
	app := cli.NewApp()
	app.Name = "Katachi"
	app.Version = "1.0"
	app.Commands = []*cli.Command{
		{
			Name:        "start",
			Description: "Start the server",
			Action: func(c *cli.Context) error {
				return server.Start(cfg)
			},
		},
		{
			Name:        "migrate",
			Description: "Run DB Migration",
			Action: func(c *cli.Context) error {
				return MigrateUp(cfg)
			},
		},
	}

	return app
}

func MigrateUp(cfg config.Config) error {
	db, err := dependencies.SetupDB(cfg.Database)
	if err != nil {
		logrus.Panicf("[Migration] Error When SetupDB: %v", err.Error())
		return err
	}

	dbConn, err := db.DB()
	if err != nil {
		logrus.Panicf("[Migration] Error When Getting DB Connection: %v", err.Error())
		return err
	}

	if err = goose.SetDialect("postgres"); err != nil {
		logrus.Panicf("[Migration] Error When Selecting Dialect to Postgres: %v", err.Error())
		return err
	}

	if err = goose.Up(dbConn, "migrations"); err != nil {
		logrus.Fatalf("[Migration] Error When Running Migrate Up: %v", err.Error())
		return err
	}

	logrus.Info("Migration Success!")
	return nil
}
