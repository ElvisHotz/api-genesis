package utils

import (
	"errors"
	"teste-api/src/models"
	"time"
)

func moedaSifra(moeda string) string {
	switch moeda {
	case "USD":
		return "$"
	case "BRL":
		return "R$"
	case "EUR":
		return "€"
	case "BTC":
		return "BTC"
	default:
		return ""
	}
}

func moedaValida(moeda string) (err error) {
	value := moedaSifra(moeda)
	if value == "" {
		return errors.New("moeda infomada invalido")
	}
	return nil
}

func Conversao(exchange *models.Exchanges) error {
	var err error

	err = moedaValida(exchange.From)
	if err != nil {
		return err
	}

	err = moedaValida(exchange.To)
	if err != nil {
		return err
	}

	err = permiteConversao(exchange.From, exchange.To)
	if err != nil {
		return err
	}

	if exchange.Rate == 0 {
		return errors.New("valor de conversao igual a 0 (zero)")
	}

	exchange.Value = exchange.Amount * exchange.Rate
	exchange.Sifra = moedaSifra(exchange.To)
	exchange.DtCreated = time.Now()
	return nil
}

func permiteConversao(from, to string) error {
	/*
	* Conversões:
	* De Dólar para Real;

	* De Euro para Real;

	* De Real para Euro;
	* De Real para Dólar;


	* De BTC para Dolar;
	* De BTC para Real;
	 */
	var err error
	switch from {
	case "USD", "EUR":

		if to != "BRL" {
			err = errors.New("conversao nao permitida (BRL)")
		}

	case "BRL":

		if to != "USD" && to != "EUR" {
			err = errors.New("conversao nao permitida (USD ou EUR)")
		}

	case "BTC":

		if to != "USD" && to != "BRL" {
			err = errors.New("conversao nao permitida (USD ou BRL)")
		}

	default:
		err = errors.New("conversao nao permitida")
	}
	return err
}
