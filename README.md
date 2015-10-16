Underscore.go
==========================================

like <a href="http://underscorejs.org/">underscore.js</a>, but for Go

## Documentation

<a name="all" />
### All(source, predicate)

__Arguments__

* `source` - Array or Map
* `predicate` - func(interface{}, interface{}) (bool, error)

__Return__

* bool, error

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

<a name="allBy" />
### AllBy(source, properties)

__Arguments__

* `source` - Array or Map
* `properties` - map[string]interface{}

__Return__

* bool, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
	TestModel{ 3, "three" },
}
res, err := AllBy(arr, map[string]interface{}{
	"Name": "a",
})
if err != nil {
	t.Error(err)
}

if res {
	t.Error("wrong result")
}
```