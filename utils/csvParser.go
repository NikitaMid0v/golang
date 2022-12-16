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

func GetRecordsHolder() RecordsHolder {
	return recordsHolder
}

func ParseCsvLine(records [][]string, id int) (model.Promotion, error) {
	fmt.Println(len(recordsHolder.Records))
	if id < 1 || id > len(recordsHolder.Records) {
		return model.Promotion{}, errors.New(fmt.Sprintf(" id %v out of range", id))
	}
	id--
	parsedPrice, err := strconv.ParseFloat(records[id][1], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", records[id][2][:19])

	if err != nil {
		log.Fatal(err.Error())
	}

	promotion := model.Promotion{
		Id:             records[id][0],
		Price:          parsedPrice,
		ExpirationDate: parsedTime,
	}

	return promotion, nil
}

type RecordsHolder struct {
	Records [][]string
}
