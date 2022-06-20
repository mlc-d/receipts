package utils

import (
	"fmt"
	"recibosV2/errorh"
	"strconv"
)

func StringValue(input interface{}) string {
	s := fmt.Sprint(input)
	return s
}

func Uint8Value(input interface{}) uint8 {
	res, err := strconv.Atoi(StringValue(input))
	errorh.Handle(err)
	return uint8(res)
}

func Int64Value(input interface{}) int64 {
	res, err := strconv.Atoi(StringValue(input))
	errorh.Handle(err)
	return int64(res)
}

func Float32Value(input interface{}) float32 {
	res, err := strconv.ParseFloat(fmt.Sprintf("%f", input), 32)
	errorh.Handle(err)
	return float32(res)
}
