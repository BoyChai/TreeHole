package utils

import (
	"fmt"
	"strconv"
)

func GetUint(value string) uint {
	parseUint, err := strconv.ParseUint(fmt.Sprint(value), 10, 64)
	if err != nil {
		fmt.Println("类型转换错误:", err.Error())
	}
	return uint(parseUint)
}
