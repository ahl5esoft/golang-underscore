package underscore

import "reflect"

// Range is 生成区间数据
func Range(start, stop, step int) IQuery {
	return &query{
		Source: createRangeArray(start, stop, step),
	}
}

func createRangeArray(start, stop, step int) []int {
	arr := make([]int, 0)
	if step == 0 {
		return arr
	}

	if start == stop {
		return arr
	} else if start < stop {
		if step < 0 {
			return arr
		}
	} else {
		if step > 0 {
			return arr
		}
	}

	for ; start < stop; start = start + step {
		arr = append(arr, start)
	}
	return arr
}

// Range2 is 生成范围内的整数序列
func Range2(start, stop, step int) IEnumerable {
	if step == 0 {
		panic("step can not equal 0")
	}

	return enumerable{
		Enumerator: func() IEnumerator {
			current := start
			index := 0
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if step > 0 {
						ok = current < stop
					} else {
						ok = current > stop
					}
					if ok {
						valueRV = reflect.ValueOf(current)
						keyRV = reflect.ValueOf(index)
						current += step
						index++
					}

					return
				},
			}
		},
	}
}
