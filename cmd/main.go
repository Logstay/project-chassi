package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/logstay/project-chassi/internal/endpoint"
	"github.com/logstay/project-chassi/internal/service"
	trans "github.com/logstay/project-chassi/internal/transport/http"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// init config databse
func init() {

	viper.SetConfigName("database.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	errViper := viper.ReadInConfig()
	if errViper != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", errViper))
	}

}

const (
	banner = ` 
 ___       ________  ________  ________  _________  ________      ___    ___ 
|\  \     |\   __  \|\   ____\|\   ____\|\___   ___\\   __  \    |\  \  /  /|
\ \  \    \ \  \|\  \ \  \___|\ \  \___|\|___ \  \_\ \  \|\  \   \ \  \/  / /
 \ \  \    \ \  \\\  \ \  \  __\ \_____  \   \ \  \ \ \   __  \   \ \    / / 
  \ \  \____\ \  \\\  \ \  \|\  \|____|\  \   \ \  \ \ \  \ \  \   \/  /  /  
   \ \_______\ \_______\ \_______\____\_\  \   \ \__\ \ \__\ \__\__/  / /    
    \|_______|\|_______|\|_______|\_________\   \|__|  \|__|\|__|\___/ /     
                                 \|_________|                   \|___|/      
                                                                                                                                                   
`
)

func main() {

	logrus.Info(banner)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger)
	defer level.Info(logger)

	var db *sqlx.DB
	{
		var err error
		db, err = sqlx.Open(viper.GetString("DB_TYPE"), viper.GetString("DB_HOST"))
		if err != nil {
			os.Exit(-1)
		}
	}

	var (
		context    context.Context
		services   = service.NewServiceFactory(db, logger)
		endpoint   = endpoint.MakeEndpoints(services, logger)
		serverHTTP = trans.NewService(context, &endpoint, &logger)
		httpAddr   = flag.String("http.addr", ":1707", "HTTP listen address")
		err        = make(chan error)
	)

	go func() {
		server := &http.Server{
			Addr:         *httpAddr,
			Handler:      serverHTTP,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		err <- server.ListenAndServe()
	}()

	fatal := level.Error(logger).Log("exit", <-err)
	if fatal != nil {
		logrus.Error(fatal)
	}

}
