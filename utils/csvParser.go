package utils

import (
	"CodingTask/model"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var recordsHolder RecordsHolder

func init() {
	promotions, err := os.Open("promotions/promotions.csv")
	if err != nil {
		panic(err)
	}

	defer promotions.Close()

	csvReader := csv.NewReader(promotions)
	records, err := csvReader.ReadAll()

	recordsHolder.Records = records
}

func ParseCsvLine(id int) (model.Promotion, error) {
	if id < 1 || id > len(recordsHolder.Records) {
		return model.Promotion{}, errors.New(fmt.Sprintf(" id %v out of range", id))
	}
	id--
	parsedPrice, err := strconv.ParseFloat(recordsHolder.Records[id][1], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", recordsHolder.Records[id][2][:19])

	if err != nil {
		log.Fatal(err.Error())
	}

	promotion := model.Promotion{
		Id:             recordsHolder.Records[id][0],
		Price:          parsedPrice,
		ExpirationDate: parsedTime,
	}

	return promotion, nil
}

type RecordsHolder struct {
	Records [][]string
}
