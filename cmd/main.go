package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/logstay/project-church-service/config"
	"github.com/logstay/project-church-service/internal/endpoint"
	"github.com/logstay/project-church-service/internal/service"
	trans "github.com/logstay/project-church-service/internal/transport/http"
	"github.com/sirupsen/logrus"
)

const (
	banner = `             __               __           .__                         .__     
_____________  ____    |__| ____   _____/  |_    ____ |  |__  __ _________   ____ |  |__  
\____ \_  __ \/  _ \   |  |/ __ \_/ ___\   __\ _/ ___\|  |  \|  |  \_  __ \_/ ___\|  |  \ 
|  |_> >  | \(  <_> )  |  \  ___/\  \___|  |   \  \___|   Y  \  |  /|  | \/\  \___|   Y  \
|   __/|__|   \____/\__|  |\___  >\___  >__|    \___  >___|  /____/ |__|    \___  >___|  /
|__|               \______|    \/     \/            \/     \/                   \/     \/ 

`
)

func main() {

	logrus.Info(banner)

	// initialize our OpenCensus configuration and defer a clean-up
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
		// Connect to the "ordersdb" database
		db, err = sqlx.Open(config.GetDBType(), config.GetPostgresConnectionString())
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
