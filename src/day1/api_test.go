package day1

import (
	"math"
	"testing"
)

func TestApi(t *testing.T) {
	t.Log(math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64, math.MaxUint8, math.MaxUint16, math.MaxUint32, math.MaxFloat32, math.MaxFloat64)
	t.Log(math.MinInt8, math.MinInt16, math.MinInt32, math.MinInt64)
}
