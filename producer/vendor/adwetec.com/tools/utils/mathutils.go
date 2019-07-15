package utils

import (
	"math"
	"math/rand"
	"time"
)

func Round(f float64, n int) float64 {

	return f

	pow := math.Pow10(n)

	return math.Trunc((f+0.5/pow)*pow) / pow

	return f

}


// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {

	rand.Seed(time.Now().Unix())

	randNum := rand.Intn(max - min)

	randNum = randNum + min

	return randNum
}

