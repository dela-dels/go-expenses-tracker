package utils

import (
	"encoding/json"
	"log"
)

type GormError struct {
	Code    int    `json:"Number"`
	Message string `json:"Message"`
}

var gormErrorStruct GormError

func ConvertGormErrorToStruct(gormError error) GormError {
	byteError, err := json.Marshal(gormError)

	if err != nil {
		log.Fatalf("Could not marshall gorm error into custom struct. failed with : %v", err)
	}

	err = json.Unmarshal(byteError, &gormErrorStruct)
	if err != nil {
		return GormError{}
	}

	return gormErrorStruct
}
