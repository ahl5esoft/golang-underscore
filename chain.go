package underscore

import "reflect"

// Chain will cause all future method calls to return wrapped objects
func Chain(source interface{}) IQuery {
	return &query{
		Source: source,
	}
}

// Chain2 is 初始化
func Chain2(src interface{}) IEnumerable {
	srcRV := reflect.ValueOf(src)
	switch srcRV.Kind() {
	case reflect.Array, reflect.Slice:
		return chainByArrayOrSlice(srcRV, srcRV.Len())
	case reflect.Map:
		return chainByMap(srcRV, srcRV.Len())
	default:
		if iterator, ok := src.(IEnumerator); ok {
			return enumerable{
				Enumerator: func() IEnumerator {
					return iterator
				},
			}
		}

		return enumerable{
			Enumerator: func() IEnumerator {
				return nullEnumerator{
					Src: srcRV,
				}
			},
		}
	}
}

func chainByArrayOrSlice(srcRV reflect.Value, size int) IEnumerable {
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

func chainByMap(srcRV reflect.Value, size int) IEnumerable {
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
