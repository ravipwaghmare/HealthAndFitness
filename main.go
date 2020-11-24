package main

import (
	"HealthAndFitness/database"
	"HealthAndFitness/server"
	"fmt"
	"io"
	"os"

	"github.com/labstack/gommon/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("Welcome to HealthAndFitness App")

	logger := lumberjack.Logger{
		Filename:   "./HealthAndFitness.log",
		MaxAge:     10,
		MaxBackups: 10,
		Compress:   true, // disabled by default
		LocalTime:  true,
	}

	mw := io.MultiWriter(os.Stdout, &logger)
	log.SetOutput(mw)

	log.Print("CONNECT TO DATABASE")
	db, err := database.Connect()
	if err != nil {
		log.Panic("COULDN'T CONNECT TO DATABASE " + err.Error())
	}

	log.Print("START UP SERVER")
	_, _ = server.New(db, &logger)

}
