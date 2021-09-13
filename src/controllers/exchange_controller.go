package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"teste-api/src/database"
	"teste-api/src/models"
	"teste-api/src/repositories"
	"teste-api/src/respostas"
	"teste-api/src/utils"

	"github.com/gorilla/mux"
)

func GetExchange(w http.ResponseWriter, r *http.Request) {
	//{amount}/{from}/{to}/{rate}
	parameters := mux.Vars(r)
	var exchange models.Exchanges

	amount, err := strconv.ParseFloat(parameters["amount"], 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	rate, err := strconv.ParseFloat(parameters["rate"], 64)

	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	exchange.Amount = amount
	exchange.Rate = rate
	exchange.From = parameters["from"]
	exchange.To = parameters["to"]
	err = utils.Conversao(&exchange)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NewExchangeRespository(db)
	_, err = repositorio.New(exchange)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK,
		struct {
			ValorConvertido float64 `json:"valorConvertido"`
			SimboloMoeda    string  `json:"simboloMoeda"`
		}{
			exchange.Value,
			exchange.Sifra,
		})

}
func GetExchanges(w http.ResponseWriter, r *http.Request) {

	to := strings.ToUpper(r.URL.Query().Get("to"))
	from := strings.ToUpper(r.URL.Query().Get("from"))

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repositorio := repositories.NewExchangeRespository(db)
	exchanges, erro := repositorio.GetAll(to, from)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, exchanges)

}
