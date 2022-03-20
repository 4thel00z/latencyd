package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/4thel00z/libhttp"
)

var (
	startFlag = flag.Int64("start", 0,
		"Left value for random sleep time interval.\n"+
			"In Milliseconds as int64. Defaults to 0.",
	)

	endFlag = flag.Int64("end", 1000,
		"Right value for random sleep time interval.\n"+
			"In Milliseconds as int64. Defaults to 1000ms.",
	)

	fixedFlag = flag.Int64("fixed",
		1000,
		"Value for sleeping a fixed time interval.\n"+
			"In Milliseconds as int64. Defaults to 1000ms.",
	)
)

func genFixedLatencyHandler(fixed int64) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {
		time.Sleep(time.Millisecond * time.Duration(fixed))
		return req.Response(map[string]string{
			"message": fmt.Sprintf("slept for %v", fixed),
		})

	}
}

func genRandomLatencyHandler(start, end int64) libhttp.Service {
	// Fix dumb user inputs
	left := math.Min(float64(start), float64(end))
	right := math.Max(float64(start), float64(end))
	delta := right - left

	return func(req libhttp.Request) libhttp.Response {

		sleepyTime := time.Millisecond * (time.Duration(rand.Intn(int(delta))) + time.Duration(start))

		log.Println("sleeping now for", sleepyTime))
		time.Sleep(sleepyTime)
		return req.Response(map[string]string{
			"message": fmt.Sprintf("slept for %v", sleepyTime),
		})
	}

}

func main() {
	flag.Parse()
	log.Println("Starting latencyd with following configuration:")
	log.Println("fixed:	", *fixedFlag, "ms")
	log.Println("start:	", *startFlag, "ms")
	log.Println("end:	", *endFlag, "ms")
	router := libhttp.Router{}
	router.GET("/fixed", genFixedLatencyHandler(*fixedFlag))
	router.GET("/random", genRandomLatencyHandler(*startFlag, *endFlag))

	svc := router.Serve().
		Filter(libhttp.ErrorFilter).
		Filter(libhttp.H2cFilter)
	srv, err := libhttp.Listen(svc, ":8000")
	if err != nil {
		panic(err)
	}

	addr := srv.Listener().Addr()
	log.Printf("👋  Listening on %v\n", addr)
	log.Printf("Curl fixed endpoint via:	curl '%v/fixed'\n", addr)
	log.Printf("Curl random endpoint via:	curl '%v/random'\n", addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
