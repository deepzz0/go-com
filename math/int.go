package math

import (
	"fmt"
	"math"
)

func Int64ToInt(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = int64ToInt(x)
	return i, nil
}

func int64ToInt(x int64) {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", x))
}
