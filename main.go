package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)


var money = 1000
var bank = 0
var mtx = sync.Mutex{}

var ErrInsufficientFunds  = "Не хватает денег для проведения операции"
const ErrMsgFailedToReadBody = "failed to read http body"
func payHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			errMsg := fmt.Sprintf("%s: %v", ErrMsgFailedToReadBody, err)
			fmt.Println(errMsg)
			w.Write([]byte(errMsg))
			return 
	}

	paymentAmount, err := strconv.Atoi(string(httpReqBody))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("failed convert str to int", err)
		return 
	}

	mtx.Lock()
	defer mtx.Unlock()
	if money - paymentAmount >= 0  {
		money -= paymentAmount
		fmt.Println("Оплата проведена, остаток на money:", money)
		w.Write([]byte("Оплата проведена"))
	}  else {
		w.WriteHeader(http.StatusPaymentRequired)
		fmt.Println(ErrInsufficientFunds)
		w.Write([]byte(ErrInsufficientFunds))
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errMsg := fmt.Sprintf("%s: %v", ErrMsgFailedToReadBody, err)
		fmt.Println(errMsg)
		w.Write([]byte(errMsg))
		return 
	}

	saveAmount, err := strconv.Atoi(string(httpReqBody))
	if err != nil {
		fmt.Println("failed to convert str to int", err)
		return 
	}

	mtx.Lock()
	if money >= saveAmount  {
		money -= saveAmount
		bank += saveAmount
		w.Write([]byte("Банк успешно пополнен"))
		fmt.Println("сумма на money:", money)
		fmt.Println("сумма на bank:", bank)
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
		fmt.Println(ErrInsufficientFunds)
	}
	mtx.Unlock()

}

func main() {
	
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("сервер запушен")
	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("произошла ошибка при запуске сервера", err.Error())
	} else {
		fmt.Println("Не хватает денег, чтобы положить в копилку")
	}

}