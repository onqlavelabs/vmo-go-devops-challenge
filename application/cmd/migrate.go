package cmd

import (
    "github.com/sirupsen/logrus"
    "os"

    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
    "github.com/dinhtp/vmo-go-devops-challenge/application/logger"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "programming challenge migrate command",
    Run:   RunMigrateCommand,
}

func init() {
    migrationCmd.AddCommand(migrateCmd)
}

func RunMigrateCommand(cmd *cobra.Command, args []string) {
    mongoDb := database.NewMongoDatabase(viper.GetString("mongoDb"))

    db, err := mongoDb.Connect()
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "cmd",
            "method":  "RunMigrateCommand",
            "dsn":     viper.GetString("mongoDb"),
        }).WithError(err).Error("connect mongo database instance failed")
        os.Exit(1)
    }

    err = database.Migrate(db)
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "cmd",
            "method":  "RunMigrateCommand",
            "dsn":     viper.GetString("mongoDb"),
        }).WithError(err).Error("migrate mongo collection failed")
        os.Exit(1)
    }

    defer mongoDb.Disconnect()
    logger.Log.Info("migration successfully")
}
