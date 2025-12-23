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
func payHandler(w http.ResponseWriter, r *http.Request) {

	httpReqBody, err := io.ReadAll(r.Body)
		if err != nil {
		fmt.Println("failed to read http body", err )
		return 
	}

	paymentAmount, err := strconv.Atoi(string(httpReqBody))

	if err != nil {
		fmt.Println("failed convert str to int", err)
		return 
	}
	mtx.Lock()
	if money - paymentAmount >= 0  {
		money =- paymentAmount
		fmt.Println("Оплата проведена, остаток на money:", money)
		w.Write([]byte("Оплата проведена"))
	}  else {
		fmt.Println("Не хватает денег для проведения операции")
		
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("failed to read http body", err )
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
		fmt.Println("Не хватает денег для проведения операции")
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