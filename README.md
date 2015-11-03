```                __
                  /\ \                                                       
 __  __    ___    \_\ \     __   _ __   ____    ___    ___   _ __    __	         __     ___
/\ \/\ \ /' _ `\  /'_  \  /'__`\/\  __\/ ,__\  / ___\ / __`\/\  __\/'__`\      /'_ `\  / __`\
\ \ \_\ \/\ \/\ \/\ \ \ \/\  __/\ \ \//\__, `\/\ \__//\ \ \ \ \ \//\  __/  __ /\ \L\ \/\ \L\ \
 \ \____/\ \_\ \_\ \___,_\ \____\\ \_\\/\____/\ \____\ \____/\ \_\\ \____\/\_\\ \____ \ \____/
  \/___/  \/_/\/_/\/__,_ /\/____/ \/_/ \/___/  \/____/\/___/  \/_/ \/____/\/_/ \/___L\ \/___/
                                                                                 /\____/
                                                                                 \_/__/
```

Underscore.go
==========================================

like <a href="http://underscorejs.org/">underscore.js</a>, but for Go

## Installation

    $ go get github.com/ahl5esoft/golang-underscore

## Lack

* SortBy
* Chain

## Documentation

<a name="all" />
### All(source, predicate)

__Arguments__

* `source` - array or map
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

* `source` - array or map
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
	// wrong
}

if res {
	// wrong
}
```

<a name="any" />
### Any(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(interface{}, interface{}) (bool, error)

__Return__

* bool, error

__Examples__

```go
arr := []int{ 1, 3 }
res, _ := Any(arr, func (n, _ interface{}) (bool, error) {
	return n.(int) % 2 == 0, nil	
})
if res {
	// wrong
}
```

<a name="anyBy" />
### AnyBy(source, properties)

__Arguments__

* `source` - array or map
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
res, err := AnyBy(arr, map[string]interface{}{
	"Id": 2,
	"Name": "two",
})
if err != nil {
	// wrong
}

if !res {
	// wrong
}
```

<a name="clone" />
### Clone()

__Return__

* interface{}, error

__Examples__

```go
arr := []int{ 1, 2, 3 }
v, _ := Clone(arr)
dstArr, _ := v.([]int)
dstArr[0] = 11
if arr[0] == dstArr[0] {
	// wrong
}
```

<a name="find" />
### Find(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(interface{}, interface{}) (bool, error)

__Return__

* interface{}, error

__Examples__

```go
arr := []int{ 1, 2, 3, 4 }
res, _ := Find(arr, func (n, _ interface{}) (bool, error) {
	return n.(int) % 2 == 0, nil
})
if res == nil {
	// wrong
} else {
	v, ok := res.(int)
	if !(ok && v == 2) {
		// wrong
	}
}
```

<a name="findBy" />
### FindBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
	TestModel{ 3, "three" },
}
res, err := FindBy(arr, map[string]interface{}{
	"Id": 1,
})
if err != nil || res == nil {
	// wrong
}

m, ok := res.(TestModel)
if !(ok && m.Name == "one") {
	// wrong
}
```

<a name="group" />
### Group(source, keySelector)

__Arguments__

* `source` - array or map
* `keySelector` - func(interface{}, interface{}) (interface{}, error)

__Return__

* map[interface{}][]interface{}, error

__Examples__

```go
dict, _ := Group([]int{ 1, 2, 3, 4, 5 }, func (item, _ interface{}) (interface{}, error) {
	if item.(int) % 2 == 0 {
		return "even", nil
	}
	return "odd", nil
})
group, ok := dict["even"]
if !(ok && len(group) == 2) {
	// wrong
}
```

<a name="groupBy" />
### GroupBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* map[interface{}][]interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "a" },
	TestModel{ 2, "a" },
	TestModel{ 3, "b" },
	TestModel{ 4, "b" },
}
dict, err := GroupBy(arr, "Name")
if !(err == nil && len(dict) == 2) {
	// wrong
}
```

<a name="index" />
### Index(source, indexSelector)

__Arguments__

* `source` - array or map
* `indexSelector` - func(interface{}, interface{}) (interface{}, error)

__Return__

* map[interface{}]interface{}, error

__Examples__

```go
res, _ := Index([]string{ "a", "b" }, func (item, _ interface{}) (interface{}, error) {
	return item, nil
})
str, ok := res["a"].(string)
if !(ok && str == "a") {
	// wrong
}
```

<a name="indexBy" />
### IndexBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* map[interface{}]interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "a" },
	TestModel{ 2, "a" },
	TestModel{ 3, "b" },
	TestModel{ 4, "b" },
}
dict, err := IndexBy(arr, "Name")
if !(err == nil && len(dict) == 2) {
	// wrong
}
```

<a name="map" />
### Map(source, selector)

__Arguments__

* `source` - array
* `selector` - func(interface{}, interface{}) (interface{}, error)

__Return__

* []interface{}, error

__Examples__

```go
arr := []string{ "a", "b", "c" }
res, _ := Map(arr, func (item, _ interface{}) (interface{}, error) {
	return item.(string) + "-", nil
})
if !(len(res) == len(arr) && res[0].(string) == "a-") {
	// wrong
}
```

<a name="pluck" />
### Pluck(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* []interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
	TestModel{ 3, "three" },
}
res, err := Pluck(arr, "Name")
if err != nil {
	// wrong
}

if len(res) != len(arr) {
	// wrong
}

for i := 0; i < 3; i++ {
	if res[i].(string) != arr[i].Name {
		// wrong
	}
}
```

<a name="reduce" />
### Reduce(source, iterator)

__Arguments__

* `source` - array
* `iterator` - func(memo, value, key interface{}) (interface{}, error), memo interface{}

__Return__

* interface{}, error

__Examples__

```go
v, err := Reduce([]int{ 1, 2 }, func (memo, value, _ interface{}) (interface{}, error) {
	arr := memo.([]int)
	num := value.(int)
	arr = append(arr, num)
	arr = append(arr, num + 10)
	return arr, nil
}, make([]int, 0))	
if err != nil {
	// wrong
}

res, ok := v.([]int)
if !(ok && len(res) == 4) {
	// wrong
}

if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
	// wrong
}
```

<a name="select" />
### Select(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(interface{}, interface{}) (bool, error)

__Return__

* []interface{}, error

__Examples__

```go
arr := []int{ 1, 2, 3, 4 }
res, _ := Select(arr, func (n, _ interface{}) (bool, error) {
	return n.(int) % 2 == 0, nil
})
if len(res) != 2 {
	// wrong
}

if !(res[0].(int) == 2 && res[1].(int) == 4) {
	// wrong
}
```

<a name="selectBy" />
### SelectBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* []interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
	TestModel{ 3, "three" },
}
res, err := SelectBy(arr, map[string]interface{}{
	"Id": 1,
})
if err != nil {
	// wrong
}

if !(len(res) == 1 && res[0].(TestModel) == arr[0]) {
	// wrong
}
```

<a name="size" />
### Size(source)

__Arguments__

* `source` - array or map

__Return__

* int

__Examples__

```go
dict := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}
if Size(dict) != len(dict) {
	// wrong
}
```

<a name="sort" />
### Sort(source, compare)

__Arguments__

* `source` - array or map
* `compare` - func(thisValue, thisKey, thatValue, thatKey interface{}) bool

__Return__

* []interface{}, error

__Examples__

```go
arr := []int{ 1, 2, 3, 5 }
res, _ := Sort([]int{ 5, 3, 2, 1 }, func (thisValue, _, thatValue, _ interface{}) bool {
	return thisValue.(int) < thatValue.(int)
})

for i, n := range arr {
	if res[i].(int) != n {
		// wrong
	}
}
```

<a name="uniq" />
### Uniq(source)

__Arguments__

* `source` - array

__Return__

* []interface{}, error

__Examples__

```go
res, _ := Uniq([]int{ 1, 2, 1, 4, 1, 3 })
if len(res) != 4 {
	// wrong
}
```

<a name="uniqBy" />
### UniqBy(source, selector)

__Arguments__

* `source` - array
* `selector` - func(interface{}, int) interface{}

__Return__

* []interface{}, error

__Examples__

```go
res, _ := UniqBy([]int{ 1, 2, 1, 4, 1, 3 }, func (item interface{}, _ int) interface{} {
	return item.(int) % 2
})
if len(res) != 2 {
	// wrong
}
```