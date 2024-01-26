package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api"
	"github.com/gokch/cafe_manager/db"
	"github.com/gokch/cafe_manager/service"
	"github.com/gokch/cafe_manager/utilx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	rootCmd = cobra.Command{
		Use:   "cafe_manager",
		Short: "Cafe manage program",
		Long:  "Cafe manage program",
		Run:   rootRun,
	}

	dbAddress  string
	dbPort     string
	dbUserName string
	dbPassword string
	dbName     string
	port       string
)

func init() {
	fs := rootCmd.PersistentFlags()
	fs.StringVar(&dbAddress, "dbaddr", "localhost", "db address")
	fs.StringVar(&dbPort, "dbport", "3306", "db port")
	fs.StringVar(&dbUserName, "dbuser", "root", "db user name")
	fs.StringVar(&dbPassword, "dbpass", "1234", "db password")
	fs.StringVar(&dbName, "dbname", "cafe", "db name")
	fs.StringVar(&port, "port", "3000", "port")
}

func rootRun(cmd *cobra.Command, args []string) {
	// init db
	database, err := db.NewDB(30, db.ConnectFuncMysql(dbAddress, dbPort, dbUserName, dbPassword, dbName))
	if err != nil {
		log.Fatal().Err(err).Strs("args", []string{dbAddress, dbPort, dbUserName, dbPassword, dbName}).Msg("Failed to connect db. please check db connection")
	}

	// init service and api
	router := gin.Default()
	api.InitRouter(service.NewService(database), router)
	go func() {
		err = router.Run(":" + port)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to run router")
		}
	}()

	// graceful shutdown
	interrupt := utilx.HandleKillSig(func() {
		database.Close()
	})
	<-interrupt.C
}
