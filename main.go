package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgibs750/lynx/api/router"
	"github.com/dgibs750/lynx/config"
	"github.com/dgibs750/lynx/util/validator"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	c := config.New()
	logger := log.Default()
	v := validator.New()

	// Capture connection propterties
	cfg := mysql.Config{
		User:                 c.DB.Username,
		Passwd:               c.DB.Password,
		Net:                  c.DB.Net,
		Addr:                 c.DB.Addr,
		DBName:               c.DB.DBName,
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	r := router.New(logger, v, db)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		logger.Printf("Shutting down server %v", s.Addr)

		ctx, cancel := context.WithTimeout(context.Background(), c.Server.TimeoutIdle)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			logger.Fatalf("Server shutdown failure : %v", err)
		}

		pingErr := db.Ping()
		if pingErr != nil {
			logger.Fatalf("DB connection closing failure : %v", pingErr)
		}

		close(closed)
	}()

	logger.Printf("Starting server %v", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server startup failure : %v", err)
	}

	<-closed
}

// a := app.New()
// w := a.NewWindow("Hello")

// lowerLetters := true
// upperLetters := true
// symbols := true
// nums := true

// length := 12
// numsAmount := 3
// symbolsAmount := 3

// bindUpper := binding.BindBool(&upperLetters)
// bindLower := binding.BindBool(&lowerLetters)
// bindSymbols := binding.BindBool(&symbols)
// bindNums := binding.BindBool(&nums)

// pwMesage := widget.NewLabel("Hello Fyne!")
// generateButton := widget.NewButton("Generate new password", func() {
// 	password, err := generator.GeneratePassword(length, lowerLetters, upperLetters, symbols, nums, numsAmount, symbolsAmount)
// 	if err != nil {
// 		errorText := fmt.Sprintf("Error : %v", err)
// 		pwMesage.SetText(" Error : " + errorText)
// 		log.Fatalf("Failed to generate password : %v\n", err)
// 	}
// 	pwMesage.SetText("Random password : " + password)
// })

// upperLettersCheck := widget.NewCheckWithData("A-Z", bindUpper)
// lowerLettersCheck := widget.NewCheckWithData("a-z", bindLower)
// numsCheck := widget.NewCheckWithData("0-9", bindNums)
// symbolsCheck := widget.NewCheckWithData("~!@#$%^&*+`-=,.", bindSymbols)

// w.SetContent(container.NewVBox(
// 	pwMesage,
// 	generateButton,
// 	upperLettersCheck,
// 	lowerLettersCheck,
// 	numsCheck,
// 	symbolsCheck,
// ))

// w.ShowAndRun()
// }
