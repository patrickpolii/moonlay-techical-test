package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// call repo
	repo := NewDatamartRepository(db)

	// call service
	svc := NewDatamartService(repo)

	done := make(chan bool)
	ticker := time.NewTicker(time.Minute)

	// helper function to stop program, can be deleted later
	go func() {
		time.Sleep(2 * time.Minute) // wait for 10 seconds
		done <- true
	}()

	// Scheduler
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case <-ticker.C:
			// call service to generate datamart
			svc.GenerateDatamart1()

		}
	}

}
