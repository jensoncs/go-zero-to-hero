package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var Header = []string{"SKU", "R200", "R520", "R100", "R650", "R460", "R680", "R150", "R260", "R050"}

type Availability struct {
	R200 int `json:"R200"`
	R520 int `json:"R520"`
	R100 int `json:"R100"`
	R650 int `json:"R650"`
	R460 int `json:"R460"`
	R680 int `json:"R680"`
	R150 int `json:"R150"`
	R260 int `json:"R260"`
	R050 int `json:"R050"`
}

type ProductAvailabilityService struct {
	Availability Availability `json: "Availability"`
	SkuCatalogID string       `json:"sku_catalog_id"`
}

func getJsonContents(fileName string) []ProductAvailabilityService {
	contents := make([]ProductAvailabilityService, 0)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(strings.NewReader(string(file))).Decode(&contents)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (a Availability) csvRecord() []string {
	var fields []string
	fields = append(fields, strconv.Itoa(a.R200))
	fields = append(fields, strconv.Itoa(a.R520))
	fields = append(fields, strconv.Itoa(a.R100))
	fields = append(fields, strconv.Itoa(a.R650))
	fields = append(fields, strconv.Itoa(a.R460))
	fields = append(fields, strconv.Itoa(a.R680))
	fields = append(fields, strconv.Itoa(a.R260))
	fields = append(fields, strconv.Itoa(a.R150))
	fields = append(fields, strconv.Itoa(a.R050))
	return fields
}

func main() {
	jsonData := getJsonContents("./complex.json")
	fmt.Println(jsonData)
	file, err := os.Create("itemToQuantityMapper.csv")
	checkError("Can't create file", err)
	writer := csv.NewWriter(file)
	writer.Write(Header)

	for _, v := range jsonData {
		if strings.HasSuffix(v.SkuCatalogID, "genericsm") {
			continue
		}
		var fields []string
		fields = append(fields, v.SkuCatalogID)
		fields = append(fields, v.Availability.csvRecord()...)
		writer.Write(fields)
	}
	writer.Flush()
}
