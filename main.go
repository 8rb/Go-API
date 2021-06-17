package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/8rb/Go-API/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println("Running on port:", port)
	data := readCsvFile("./indicadoresrural2018.csv")

	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	router.HandleFunc("/", service.IndexRoute)
	router.HandleFunc("/indicators", service.GetAllIndicators(data)).Methods("GET")
	router.HandleFunc("/indicators/{name}", service.GetIndicatorByName(data)).Methods("GET")
	router.HandleFunc("/indicators/{indicator1}/{indicator2}", service.CompareTwoIndicators(data)).Methods("GET")
	router.HandleFunc("/kmeans/{indicator1}/{indicator2}", service.KmeansTwoIndicators(data)).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
