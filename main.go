package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("a")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
    rows, err := db.Raw("show tables").Rows()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
	defer rows.Close()

	fmt.Printf("%+v", rows)

	type Note struct {
		ID           uint    `gorm:"primaryKey"`
		Name         string  `gorm:"type:varchar(255)"`
		Content      string  `gorm:"type:text"`
		Comment      string  `gorm:"type:text"`
		CreatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
		UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	}
	// db.Migrator().CreateTable(&Note{})
	// db.Migrator().DropTable(&Note{})
	if err := db.AutoMigrate(&Note{}); err != nil {
		fmt.Printf("failed %s", err)
	}

	type Tag struct {
		ID           uint    `gorm:"primaryKey"`
		NoteId       uint    `gorm:"not null"`
		Name         string  `gorm:"type:varchar(255)"`
		Value        string  `gorm:"type:varchar(255)"`
		CreatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
		UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	}
	// db.Migrator().DropTable(&Tag{})
	if err := db.AutoMigrate(&Tag{}); err != nil {
		fmt.Printf("failed %s", err)
	}
}
