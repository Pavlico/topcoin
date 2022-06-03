package database

import (
	"database/sql"
	"log"

	"github.com/Pavlico/topcoin/services/database/pkg/conf"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
	"github.com/go-sql-driver/mysql"
)

type DbServiceStruct struct {
	Database *sql.DB
}

type DbService interface {
	Save(coinData []dataTypes.CoinData) error
	GetAll() ([]dataTypes.CoinData, error)
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
	cfg := mysql.Config{
		User:   c.Username,
		Passwd: c.Password,
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: c.DbName,
	}
	log.Println(cfg.FormatDSN())
	return cfg.FormatDSN()
	// return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, c.Hostname, c.DbName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	return db, err
}

func (dss *DbServiceStruct) Save(coinData []dataTypes.CoinData) error {
	tx, err := dss.Database.Begin()

	if err != nil {
		return err
	}
	tx.Exec("TRUNCATE TABLE topcoins;")
	for _, v := range coinData {
		_, err = tx.Exec("INSERT INTO topcoins(`rank`, `score`, `symbol`) VALUES (?, ?,? )", v.Rank, v.Score, v.Symbol)
	}
	err = tx.Commit()

	if err != nil {
		defer tx.Rollback()
		return err
	}
	return nil
}

func (dss *DbServiceStruct) GetAll() ([]dataTypes.CoinData, error) {
	var coins []dataTypes.CoinData
	tx, err := dss.Database.Begin()
	defer tx.Rollback()
	if err != nil {
		return coins, err
	}

	rows, err := tx.Query("SELECT symbol,rank,score FROM topcoins")
	defer rows.Close()
	if err != nil {
		return coins, err
	}

	for rows.Next() {
		var coin dataTypes.CoinData
		if err := rows.Scan(&coin.Symbol, &coin.Score, &coin.Rank); err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}
	return coins, nil
}
