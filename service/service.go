package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/8rb/Go-API/model"
	"github.com/8rb/Go-API/worker"
	"github.com/gorilla/mux"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API!!\nGo to: /kmeans/{indicator1}/{indicator2} to get the results of the K-Means Algorithm")
}

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

func GetIndicatorByName(data [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func CompareTwoIndicators(data [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		indicator1 := vars["indicator1"]
		indicator2 := vars["indicator2"]
		fmt.Println(indicator1, indicator2)
		if indicator1 == os.DevNull || indicator2 == os.DevNull {
			return
		}
		fmt.Println(indicator1, indicator2)
		formattedData := []model.Indicator{}
		for _, row := range data {
			x, _ := strconv.Atoi(row[columns[indicator1]])
			y, _ := strconv.Atoi(row[columns[indicator2]])
			object := model.Indicator{
				Label: row[columns["D_DPTO"]],
				X:     x,
				Y:     y,
			}
			formattedData = append(formattedData, object)
		}
		formattedData = formattedData[1:]
		response := []model.Group{}
		prevLabel := formattedData[0].Label
		fmt.Println(prevLabel)
		var preData [][]int
		for i, object := range formattedData {

			var tuple []int
			x := object.X
			y := object.Y
			tuple = append(tuple, x)
			tuple = append(tuple, y)

			label := object.Label

			if prevLabel == label {
				preData = append(preData, tuple)
			} else {
				group := model.Group{}
				group.NAME = prevLabel
				group.DATA = preData
				if len(preData) != 0 {
					response = append(response, group)
					preData = nil
					preData = append(preData, tuple)
				} else {
					preData = append(preData, tuple)
				}
			}
			if i == len(formattedData)-1 {
				group := model.Group{}
				group.NAME = prevLabel
				group.DATA = preData
				response = append(response, group)
			}
			prevLabel = label
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func GetAllIndicators(data [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func workerCalculateFactor(c chan []model.Item, halfItems []model.Item) {
	for i := range halfItems {
		halfItems[i].Factor = halfItems[i].X * halfItems[i].Y
	}
}

func workerCalculateTotalFactor(c chan int, halfItems []model.Item) {
	halfFactor := 0
	for i := range halfItems {
		halfFactor += halfItems[i].Factor
	}
	c <- halfFactor
}

func KmeansTwoIndicators(data [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// First we get the variables from the URI
		indicator1 := vars["indicator1"]
		indicator2 := vars["indicator2"]
		fmt.Println(indicator1, indicator2)
		items := []model.Item{}
		groupId := 0
		// Then, we iterate over the data to assign the x and y coordinates
		// Factor starts at 0, since we will calculate it later
		// GroupId starts at 0 since we will calculate it later
		for _, row := range data {
			x, _ := strconv.Atoi(row[columns[indicator1]])
			y, _ := strconv.Atoi(row[columns[indicator2]])
			region := row[columns["D_DPTO"]]
			object := model.Item{
				X:       x,
				Y:       y,
				Region:  region,
				Factor:  0,
				GroupId: groupId,
			}
			items = append(items, object)
		}
		// Then we assign the "Factor" to each element
		// For better performance we divide the array
		// into two halves and process each half concurrently
		chanCF1 := make(chan []model.Item)
		chanCF2 := make(chan []model.Item)

		index := len(items) / 2
		items1 := items[:index]
		items2 := items[index:]

		go worker.CalculateFactor(chanCF1, items1)
		go worker.CalculateFactor(chanCF2, items2)

		items = nil
		items = append(items, items1...)
		items = append(items, items2...)

		// Later we proceed to calculate the total factor
		// Which is the sum of all the factors of each element

		chanTF := make(chan int)
		go worker.CalculateTotalFactor(chanTF, items1)
		go worker.CalculateTotalFactor(chanTF, items2)

		totalFactor := 0
		totalFactor = <-chanTF + <-chanTF

		fmt.Println("TOTAL FACTOR:", totalFactor)

		// With the total factor we calculate the clusterClassifier
		// Which will help us to divide the elements into clusters
		clusterClassifier := totalFactor / len(items) / 3
		fmt.Println("GROUP CLASSIFIER:", clusterClassifier)

		// Later we sort the items based on that factor
		sort.Slice(items, func(i, j int) bool {
			return items[i].Factor < items[j].Factor
		})

		formattedItems := []model.Item{}
		// we remove the empty data (the one that has 0's)
		for i := range items {
			if items[i].Factor != 0 {
				formattedItems = append(formattedItems, items[i])
			}
		}
		// From now on we will use the "formattedItems" slice/array
		// Now we assign the groups to each element based on the difference of its Factor
		for i := 0; i < len(formattedItems)-1; i++ {
			if formattedItems[i+1].Factor-formattedItems[i].Factor > clusterClassifier {
				groupId += 1
				clusterClassifier += clusterClassifier
			}
			formattedItems[i+1].GroupId = groupId
		}
		// Now we format the data for the front-end:

		//We define our data types
		response := []model.Group{}
		actualGroup := model.Group{}
		var actualData [][]int

		// We add the first element to the data group
		var firstTuple []int
		firstTuple = append(firstTuple, formattedItems[0].X)
		firstTuple = append(firstTuple, formattedItems[0].Y)

		actualData = append(actualData, firstTuple)

		// We start the formatting
		for i := 0; i < len(formattedItems)-1; i++ {
			var tuple []int
			tuple = append(tuple, formattedItems[i+1].X)
			tuple = append(tuple, formattedItems[i+1].Y)
			if formattedItems[i].GroupId == formattedItems[i+1].GroupId {
				actualData = append(actualData, tuple)
			} else {
				actualGroup.NAME = "Cluster " + strconv.Itoa(formattedItems[i].GroupId+1)
				actualGroup.DATA = actualData
				response = append(response, actualGroup)
				actualData = nil
				actualData = append(actualData, tuple)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
