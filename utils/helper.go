package utils

import (
	"encoding/json"
	"gorm.io/datatypes"
	"strconv"
	"strings"
)

func SplitStringAndMarshal(data string, sep string) ([]byte, error) {
	items := strings.Split(data, sep)

	var cleaned []string

	for _, item := range items {
		cleaned = append(cleaned, strings.Trim(item, " "))
	}

	return json.Marshal(cleaned)
}

func FromJsonDataTypeToSliceString(data datatypes.JSON) ([]string, error) {
	var sliceString []string

	if err := json.Unmarshal(data, &sliceString); err != nil {
		return nil, err
	}

	return sliceString, nil
}

func ToJsonB(data string, separator string) ([]byte, error) {
	genres, err := SplitStringAndMarshal(data, separator)

	if err != nil {
		return nil, err
	}

	return genres, nil
}

func UintToString(data uint) string {
	return strconv.Itoa(int(data))
}

func StringToUint(data string) (uint, error) {
	datum, err := strconv.Atoi(data)

	if err != nil {
		return 0, err
	}

	return uint(datum), nil
}
