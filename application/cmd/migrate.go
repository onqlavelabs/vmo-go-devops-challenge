package cmd

import (
    "os"

    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
    "github.com/dinhtp/vmo-go-devops-challenge/application/logger"
    "github.com/dinhtp/vmo-go-devops-challenge/application/migration"
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
    mongoDb := database.NewMongoDatabase(viper.GetString("mongoUri"))

    db, err := mongoDb.Connect()
    if err != nil {
        logger.Log.WithError(err).Error("connect mongo database instance failed")
        os.Exit(1)
    }

    err = migration.Migrate(db)
    if err != nil {
        logger.Log.WithError(err).Error("migrate mongo collection failed")
        os.Exit(1)
    }

    defer mongoDb.Disconnect()
    logger.Log.Info("migration successfully")
}
