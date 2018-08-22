package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io"
	"log"
	"net/http"
	"sync"
)

var globalDB *mgo.Database
var account = "WeiZhe"
var in chan string
var out chan Result

type Result struct {
	Account string
	Result  float64
}

type currency struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount  float64       `bson:"amount"`
	Account string        `bson:"account"`
	Code    string        `bson:"code"`
}

func pay(w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		in <- account
		for {
			select {
			case result := <-out:
				fmt.Printf("%+v\n", result)
				wg.Done()
				return
			}
		}
	}(&wg)

	wg.Wait()
	io.WriteString(w, "ok")
}

func main() {
	in = make(chan string)
	out = make(chan Result)

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("can't connect mgo DB server")
	}
	globalDB = session.DB("logs")
	globalDB.C("bank").DropCollection()

	user := currency{Account: account, Amount: 100.00, Code: "USD"}
	err = globalDB.C("bank").Insert(&user)
	if err != nil {
		panic("insert error")
	}

	go func(in *chan string) {
		for {
			select {
			case account := <-*in:
				entry := currency{}
				err := globalDB.C("bank").Find(bson.M{"account": account}).One(&entry)
				if err != nil {
					panic(err)
				}

				entry.Amount = entry.Amount + 50.00
				err = globalDB.C("bank").UpdateId(entry.ID, &entry)
				if err != nil {
					panic("update error")
				}

				out <- Result{
					Account: account,
					Result:  entry.Amount,
				}

			}
		}

	}(&in)

	log.Println("Listen on server on 8080 port")
	http.HandleFunc("/", pay)
	http.ListenAndServe(":8080", nil)

}
