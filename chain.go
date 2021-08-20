package underscore

import (
	"reflect"
)

// Chain is 创建枚举器
func Chain(src interface{}) IEnumerable {
	return chainFromValue(
		reflect.ValueOf(src),
	)
}

func chainFromArrayOrSlice(srcValue reflect.Value, size int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			return &enumerator{
				MoveNextFunc: func() (valueValue reflect.Value, keyValue reflect.Value, ok bool) {
					ok = index < size
					if ok {
						valueValue = srcValue.Index(index)
						keyValue = reflect.ValueOf(index)
						index++
					}

					return
				},
			}
		},
	}
}

func chainFromMap(srcValue reflect.Value, size int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			keyValues := srcValue.MapKeys()
			return &enumerator{
				MoveNextFunc: func() (valueValue reflect.Value, keyValue reflect.Value, ok bool) {
					ok = index < size
					if ok {
						valueValue = srcValue.MapIndex(keyValues[index])
						keyValue = keyValues[index]
						index++
					}

					return
				},
			}
		},
	}
}

func chainFromValue(value reflect.Value) IEnumerable {
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		return chainFromArrayOrSlice(
			value,
			value.Len(),
		)
	case reflect.Map:
		return chainFromMap(
			value,
			value.Len(),
		)
	default:
		if value.IsValid() {
			if iterator, ok := value.Interface().(IEnumerator); ok {
				return enumerable{
					Enumerator: func() IEnumerator {
						return iterator
					},
				}
			}

			if value.Kind() == reflect.Ptr {
				return chainFromValue(
					value.Elem(),
				)
			}
		}

		return enumerable{
			Enumerator: func() IEnumerator {
				return nullEnumerator{
					Src: value,
				}
			},
		}
	}
}
