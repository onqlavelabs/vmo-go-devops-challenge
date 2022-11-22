package cmd

import (
    "github.com/dinhtp/vmo-go-devops-challenge/application/server"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "programming challenge serve command",
    Run:   RunServeCommand,
}

func init() {
    rootCmd.AddCommand(serveCmd)

    serveCmd.Flags().StringP("address", "", ":8080", "service address")
    serveCmd.Flags().StringP("mongoUri", "", ":8080", "service address")

    _ = viper.BindPFlag("address", serveCmd.Flags().Lookup("address"))
    _ = viper.BindPFlag("mongoUri", serveCmd.Flags().Lookup("mongoUri"))
}

func RunServeCommand(cmd *cobra.Command, args []string) {
    //db, err := database.NewMongoDatabase(viper.GetString("mongoUri")).Connect()
    //if err != nil {
    //    panic(err)
    //}

    server.NewEchoServer(viper.GetString("address"), nil).Serve()
}
