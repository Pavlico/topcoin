package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `gorm:"not null;default:''"`
	Body      string `gorm:"not null;default:''"`
}

type JsonPost struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DbServiceStruct struct {
	Database *sql.DB
}

type DbService interface {
	Save(body string, title string) (string, error)
	Update(id string, title string) (string, error)
	Get(id string) ([]byte, error)
	Exists(id string) bool
}

func Initialize() (DbServiceStruct, error) {
	db, err := Connect()
	if err != nil {
		return DbServiceStruct{}, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return DbServiceStruct{}, pingErr
	}
	return DbServiceStruct{Database: db}, nil
}

func dsn() string {
	c := conf.GetDbCredentials()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, c.Hostname, c.DbName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	return db, err
}

func (dss *DbServiceStruct) Save(symbol string, rank int, score float32) error {
	tx, err := dss.Database.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	if symbol == "" {
		return errors.New("Wrong params")
	}
	_, err = tx.Exec("UPDATE topcoins SET rank = ? score = ? WHERE symbol = ?",
		rank, score, symbol)
	if err != nil {
		return err
	}
	return nil
}

func (dss *DbServiceStruct) GetAll() error {
	tx, err := dss.Database.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Exec("SELECT symbol,rank,score FROM topcoins")
	if err != nil {
		return err
	}
	return nil
}
