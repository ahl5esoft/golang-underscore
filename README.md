Underscore.go
==========================================

like underscore.js, but for Go

## Documentation

<a name="all" />
### All(source, predicate)

__Arguments__

* `source` - Array or Map
* `predicate` - func(interface{}, interface{}) (bool, error)

__Examples__

```go
arr := []int{ 2, 4 }
res, _ := All(arr, func (n, _ interface{}) (bool, error) {
	return n.(int) % 2 == 0, nil	
})
if !res {
	// wrong
}
```