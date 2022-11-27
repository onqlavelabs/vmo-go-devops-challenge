package cmd

import (
    "os"

    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
    "github.com/dinhtp/vmo-go-devops-challenge/application/logger"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var seedCmd = &cobra.Command{
    Use:   "seed",
    Short: "programming challenge seed command",
    Run:   RunSeedCommand,
}

func init() {
    migrationCmd.AddCommand(seedCmd)
}

func RunSeedCommand(cmd *cobra.Command, args []string) {
    mongoDb := database.NewMongoDatabase(viper.GetString("mongoDb"))

    db, err := mongoDb.Connect()
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "cmd",
            "method":  "RunSeedCommand",
            "dsn":     viper.GetString("mongoDb"),
        }).WithError(err).Error("connect mongo database instance failed")
        os.Exit(1)
    }

    err = database.Seed(db)
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "cmd",
            "method":  "RunSeedCommand",
            "dsn":     viper.GetString("mongoDb"),
        }).WithError(err).Error("seed mongo collection failed")
        os.Exit(1)
    }

    defer mongoDb.Disconnect()
    logger.Log.Info("seeding data successfully")
}
