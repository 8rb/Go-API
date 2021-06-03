package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/8rb/Go-API/model"
	"github.com/gorilla/mux"
)

var data = [][]string{}

var columns = map[string]int{
	"D_DPTO":            0,
	"D_PROV":            1,
	"D_DIST":            2,
	"D_DREUGEL":         3,
	"RURAL_PMM":         4,
	"RURAL_PMM_MUJE1":   5,
	"RURAL_PMM_MUJE2":   6,
	"RURAL_PMM_HOME1":   7,
	"RURAL_PMM_HOME2":   8,
	"RURAL_PMMA_MUJE1":  9,
	"RURAL_PMMA_MUJE2":  10,
	"RURAL_PMMA_HOME1":  11,
	"RURAL_PMMA_HOME2":  12,
	"RURAL_PMM_MUJDOC":  13,
	"RURAL_PMM_HOMDOC":  14,
	"RURAL_PMMA_MUJDOC": 15,
	"RURAL_PMMA_HOMDOC": 16,
	"RURAL_CRFA":        17,
	"RURAL_SRE":         18,
	"RURAL_ST":          19,
	"RURAL_CRFA_MUJE1":  20,
	"RURAL_CRFA_MUJE2":  21,
	"RURAL_CRFA_HOME1":  22,
	"RURAL_CRFA_HOME2":  23,
	"RURAL_SRE_MUJE1":   24,
	"RURAL_SRE_MUJE2":   25,
	"RURAL_SRE_HOME1":   26,
	"RURAL_SRE_HOME2":   27,
	"RURAL_MSE_MUJDOC":  28,
	"RURAL_MSE_HOMDOC":  29,
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome the my GO API!")
}

func getIndicatorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colName := vars["name"]
	if colName == os.DevNull {
		return
	}
	response := []model.OneField{}
	for _, row := range data {
		object := model.OneField{Field1: row[columns[colName]]}
		response = append(response, object)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getAllIndicators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := []model.AllFields{}
	for _, row := range data {
		object := model.AllFields{
			D_DPTO:            row[columns["D_DPTO"]],
			D_PROV:            row[columns["D_PROV"]],
			D_DIST:            row[columns["D_DIST"]],
			D_DREUGEL:         row[columns["D_DREUGEL"]],
			RURAL_PMM:         row[columns["RURAL_PMM"]],
			RURAL_PMM_MUJE1:   row[columns["RURAL_PMM_MUJE1"]],
			RURAL_PMM_MUJE2:   row[columns["RURAL_PMM_MUJE2"]],
			RURAL_PMM_HOME1:   row[columns["RURAL_PMM_HOME1"]],
			RURAL_PMM_HOME2:   row[columns["RURAL_PMM_HOME2"]],
			RURAL_PMMA_MUJE1:  row[columns["RURAL_PMMA_MUJE1"]],
			RURAL_PMMA_MUJE2:  row[columns["RURAL_PMMA_MUJE2"]],
			RURAL_PMMA_HOME1:  row[columns["RURAL_PMMA_HOME1"]],
			RURAL_PMMA_HOME2:  row[columns["RURAL_PMMA_HOME2"]],
			RURAL_PMM_MUJDOC:  row[columns["RURAL_PMM_MUJDOC"]],
			RURAL_PMM_HOMDOC:  row[columns["RURAL_PMM_HOMDOC"]],
			RURAL_PMMA_MUJDOC: row[columns["RURAL_PMMA_MUJDOC"]],
			RURAL_PMMA_HOMDOC: row[columns["RURAL_PMMA_HOMDOC"]],
			RURAL_CRFA:        row[columns["RURAL_CRFA"]],
			RURAL_SRE:         row[columns["RURAL_SRE"]],
			RURAL_ST:          row[columns["RURAL_ST"]],
			RURAL_CRFA_MUJE1:  row[columns["RURAL_CRFA_MUJE1"]],
			RURAL_CRFA_MUJE2:  row[columns["RURAL_CRFA_MUJE2"]],
			RURAL_CRFA_HOME1:  row[columns["RURAL_CRFA_HOME1"]],
			RURAL_CRFA_HOME2:  row[columns["RURAL_CRFA_HOME2"]],
			RURAL_SRE_MUJE1:   row[columns["RURAL_SRE_MUJE1"]],
			RURAL_SRE_MUJE2:   row[columns["RURAL_SRE_MUJE2"]],
			RURAL_SRE_HOME1:   row[columns["RURAL_SRE_HOME1"]],
			RURAL_SRE_HOME2:   row[columns["RURAL_SRE_HOME2"]],
			RURAL_MSE_MUJDOC:  row[columns["RURAL_MSE_MUJDOC"]],
			RURAL_MSE_HOMDOC:  row[columns["RURAL_MSE_HOMDOC"]],
		}
		response = append(response, object)
	}
	json.NewEncoder(w).Encode(response)
}

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
	data = readCsvFile("./indicadoresrural2018.csv")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/indicators", getAllIndicators).Methods("GET")
	router.HandleFunc("/indicators/{name}", getIndicatorByName).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
