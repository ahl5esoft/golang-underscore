package underscore

import "reflect"

// Chain is 初始化
func Chain(src interface{}) IEnumerable {
	return chainFromRV(
		reflect.ValueOf(src),
	)
}

func chainFromArrayOrSlice(srcRV reflect.Value, size int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					ok = index < size
					if ok {
						valueRV = srcRV.Index(index)
						keyRV = reflect.ValueOf(index)
						index++
					}

					return
				},
			}
		},
	}
}

func chainFromMap(srcRV reflect.Value, size int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			keyRVs := srcRV.MapKeys()
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					ok = index < size
					if ok {
						valueRV = srcRV.MapIndex(keyRVs[index])
						keyRV = keyRVs[index]
						index++
					}

					return
				},
			}
		},
	}
}

func chainFromRV(rv reflect.Value) IEnumerable {
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		return chainFromArrayOrSlice(rv, rv.Len())
	case reflect.Map:
		return chainFromMap(rv, rv.Len())
	default:
		if iterator, ok := rv.Interface().(IEnumerator); ok {
			return enumerable{
				Enumerator: func() IEnumerator {
					return iterator
				},
			}
		}

		return enumerable{
			Enumerator: func() IEnumerator {
				return nullEnumerator{
					Src: rv,
				}
			},
		}
	}
}
