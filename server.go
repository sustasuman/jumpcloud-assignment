package main

import (
	"context"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thoas/stats"
)

type responseStats struct {
	Total   int     `json:"total"`
	Avarage float64 `json:"avarage"`
}

func HandleDbError(err error, ctx *gin.Context) {
	if err != nil {
		log.Panic(err)
	}
	ctx.HTML(505, "Error occured. Please contact admin.", gin.H{})
	ctx.AbortWithStatus(http.StatusInternalServerError)
	return
}

var Stats = stats.New()

func getStats(ctx *gin.Context) {
	var s = Stats.Data()
	var resTimeMs = math.Round(s.AverageResponseTimeSec*1000*1000*100) / 100
	actualStat := responseStats{Total: s.TotalCount, Avarage: resTimeMs}

	ctx.IndentedJSON(http.StatusOK, actualStat)
}

func shutdown(ctx *gin.Context) {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

//This function keeps the stats of the post request.
var registerStatHandler = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.URL.Path, "/hash") {
			beginning, recorder := Stats.Begin(c.Writer)
			c.Next()
			Stats.End(beginning, stats.WithRecorder(recorder))
		} else {
			c.Next()
		}
	}
}()

func StartServer() {
	//set the mode of run.
	gin.SetMode(viper.GetString("mode"))

	router := gin.Default()
	router.Use(registerStatHandler)
	//open db connection
	InitDb()

	//add the routes
	router.POST("/hash", PostHash)
	router.GET("/hash/:id", GetHash)
	router.GET("/stats", getStats)
	router.GET("/shutdown", shutdown)

	// define server info
	server := &http.Server{
		Addr:    viper.GetString("host"),
		Handler: router,
	}
	//router.Run(viper.GetString("host"))

	//go routine to run on background to listen server events.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//Listen to the os signal
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutingdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//close dbConnection
	defer CloseDb()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting....")

}
