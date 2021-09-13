package repositories

import (
	"database/sql"
	"fmt"
	"teste-api/src/models"

	_ "github.com/go-sql-driver/mysql"
)

type exchangeRepository struct {
	db *sql.DB
}

func NewExchangeRespository(db *sql.DB) *exchangeRepository {
	return &exchangeRepository{db: db}
}

func (r exchangeRepository) New(exchange models.Exchanges) (int, error) {
	statement, err :=
		r.db.Prepare("insert into exchanges(amount,fromExchange,toExchange,rate,dtCreated) values (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(exchange.Amount, exchange.From, exchange.To, exchange.Rate, exchange.DtCreated)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}

func (r exchangeRepository) GetAll(to string, from string) ([]models.Exchanges, error) {
	var fromExchange = ""
	var toExchange = ""
	var rows *sql.Rows
	var err error

	if from != "" {
		fromExchange = fmt.Sprintf("%%%s%%", from)
	}
	if to != "" {
		toExchange = fmt.Sprintf("%%%s%%", to)
	}
	if fromExchange != "" && toExchange != "" {
		rows, err = r.db.Query("select id,amount,fromExchange,toExchange,rate,dtCreated from exchanges where fromExchange like ? and toExchange like ?", fromExchange, toExchange)
	} else if toExchange != "" {
		rows, err = r.db.Query("select id,amount,fromExchange,toExchange,rate,dtCreated from exchanges where toExchange like ?", toExchange)
	} else if fromExchange != "" {
		rows, err = r.db.Query("select id,amount,fromExchange,toExchange,rate,dtCreated from exchanges where fromExchange like ?", fromExchange)
	} else {
		rows, err = r.db.Query("select id,amount,fromExchange,toExchange,rate,dtCreated from exchanges")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var exchanges []models.Exchanges
	for rows.Next() {
		var exchange models.Exchanges
		if err = rows.Scan(
			&exchange.Id,
			&exchange.Amount,
			&exchange.From,
			&exchange.To,
			&exchange.Rate,
			&exchange.DtCreated,
		); err != nil {
			return nil, err
		}
		exchanges = append(exchanges, exchange)
	}
	return exchanges, nil
}
