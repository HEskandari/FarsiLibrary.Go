package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/heskandari/farsilibrary.go/date"
	"log"
	"net/http"
	"time"
)

type AppRoutes struct {
	Route       string
	Description string
}

var myRouter = &mux.Router{}
var appRoutes = []AppRoutes{
	{"/", "Route page"},
	{"/now", "returns current date in PersianDate format"},
	{"/topersian/{date}", "converts a gregorian date to persian date"},
	{"/togregorian/{date}", "converts a persian date to gregorian date."},
}

func main() {
	handleRequests()
}

type ResponseData struct {
	Result string
}

func handleNow(w http.ResponseWriter, r *http.Request) {

	pd := date.ToPersianDate(time.Now())
	resp := ResponseData{
		Result: pd.Format(date.WrittenFormat),
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func handleToPersian(w http.ResponseWriter, r *http.Request) {
	resp := ResponseData{}
	vars := mux.Vars(r)
	requestedDate := vars["date"]

	parsed, err := time.Parse("2006-01-02", requestedDate)
	if err != nil {
		resp.Result = fmt.Sprintf("failed to parse the date '%s'. Expected YYYY-MM-DD format.", requestedDate)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	pd := date.ToPersianDate(parsed)
	resp.Result = pd.Format(date.WrittenFormat)
	_ = json.NewEncoder(w).Encode(resp)
}

func handleToGregorian(w http.ResponseWriter, r *http.Request) {
	resp := ResponseData{}
	vars := mux.Vars(r)
	requestedDate := vars["date"]

	parsed, err := date.ParseWithSeparator(requestedDate, '-')
	if err != nil {
		resp.Result = fmt.Sprintf("failed to parse the date '%s'. Expected YYYY-MM-DD format.", requestedDate)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	gd := date.ToGregorianDate(parsed)
	resp.Result = gd.String()
	_ = json.NewEncoder(w).Encode(resp)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(appRoutes)
}

func handleRequests() {
	myRouter = mux.NewRouter()
	myRouter.StrictSlash(true)

	myRouter.HandleFunc("/", handleHome)
	myRouter.HandleFunc("/now", handleNow)
	myRouter.HandleFunc("/topersian/{date}", handleToPersian)
	myRouter.HandleFunc("/togregorian/{date}", handleToGregorian)

	log.Fatal(http.ListenAndServe(":9900", myRouter))
}
