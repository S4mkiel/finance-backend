package utils

import (
	"log"
	"strconv"
	"time"
)

func PString(str string) *string {
	return &str
}

func StringIfNotNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func PStringIfNotNil(str *string) *string {
	if str != nil {
		if *str == "" {
			return nil
		}
	}
	return str
}

func PFloat64(f float64) *float64 {
	return &f
}

func PFloat64IfNotNil(f *float64) *float64 {
	if f != nil {
		if *f == 0 {
			return nil
		}
	}
	return f
}

func Float64IfNotNil(f float64) *float64 {
	if f == 0 {
		return nil
	}

	return &f
}

func TSFloat64(f string) *float64 {
	value, err := strconv.ParseFloat(f, 64)
	if err != nil {
		log.Println(err)
		return nil
	}

	if value == 0 {
		return nil
	}

	return &value
}

func PInt(i int) *int {
	return &i
}

func PIntIfNotNil(i *int) *int {
	if i != nil {
		if *i == 0 {
			return nil
		}
	}
	return i
}

func TSInt(i string) *int {
	value, err := strconv.Atoi(i)
	if err != nil {
		log.Println(err)
		return nil
	}

	if value == 0 {
		return nil
	}

	return &value
}

func IntIfNotNil(i int) *int {
	if i == 0 {
		return nil
	}

	return &i
}

func PUint(i uint) *uint {
	return &i
}

func PTime(t time.Time) *time.Time {
	return &t
}

func TSTime(dateStr string) *time.Time {
	layout := "2006-01-02 15:04"

	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &parsedTime
}

func PBool(b bool) *bool {
	return &b
}
