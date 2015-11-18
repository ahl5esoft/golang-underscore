```                __
                  /\ \                                                       
 __  __    ___    \_\ \     __   _ __   ____    ___    ___   _ __    __	         __     ___
/\ \/\ \ /' _ `\  /'_  \  /'__`\/\  __\/ ,__\  / ___\ / __`\/\  __\/'__`\      /'_ `\  / __`\
\ \ \_\ \/\ \/\ \/\ \ \ \/\  __/\ \ \//\__, `\/\ \__//\ \ \ \ \ \//\  __/  __ /\ \L\ \/\ \L\ \
 \ \____/\ \_\ \_\ \___,_\ \____\\ \_\\/\____/\ \____\ \____/\ \_\\ \____\/\_\\ \____ \ \____/
  \/___/  \/_/\/_/\/__,_ /\/____/ \/_/ \/___/  \/____/\/___/  \/_/ \/____/\/_/ \/___L\ \/___/
                                                                                 /\____/
                                                                                 \_/__/
																			--by 0890.ahl
```

Underscore.go
==========================================

like <a href="http://underscorejs.org/">underscore.js</a>, but for Go

## Installation

    $ go get github.com/ahl5esoft/golang-underscore

## Documentation

### API
* [`All`](#all), [`AllBy`](#allBy)
* [`Any`](#any), [`AnyBy`](#anyBy)
* [`Chain`](#chain)
* [`Clone`](#clone)
* [`Find`](#find), [`FindBy`](#findBy)
* [`Group`](#group), [`GroupBy`](#groupBy)
* [`Index`](#index), [`IndexBy`](#indexBy)
* [`Map`](#map)
* [`Pluck`](#pluck)
* [`Reduce`](#reduce)
* [`Select`](#select), [`SelectBy`](#selectBy)
* [`Size`](#size)
* [`Sort`](#sort), [`SortBy`](#sortBy)
* [`Uniq`](#uniq), [`UniqBy`](#uniqBy)

<a name="all" />
### All(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element, index or key) (bool, error)

__Return__

* bool - all the values that pass a truth test `predicate`
* error

__Examples__

```go
arr := []int{ 2, 4 }
res, _ := All(arr, func (n, _ int) (bool, error) {
	return n % 2 == 0, nil	
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
* `predicate` - func(element, index or key) (bool, error)

__Return__

* bool - any of the values that pass a truth test `predicate`
* error

__Examples__

```go
arr := []int{ 1, 3 }
res, _ := Any(arr, func (n, _ int) (bool, error) {
	return n % 2 == 0, nil	
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

<a name="chain" />
### Chain(source)

__Arguments__

* `source` - array or map

__Return__

* interface{} - a wrapped object, wrapped objects until value is called
* error

__Examples__

```go
v, _ := Chain([]int{ 1, 2, 1, 4, 1, 3 }).Uniq(nil).Group(func (n, _ int) (string, error) {
	if (n % 2 == 0) {
		return "even", nil
	}

	return "old", nil
}).Value()
res, ok := v.(map[string][]int)
if !(ok && len(res) == 2) {
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
* `predicate` - func(element, index or key) (bool, error)

__Return__

* interface{}, error

__Examples__

```go
arr := []int{ 1, 2, 3, 4 }
res, _ := Find(arr, func (n, _ int) (bool, error) {
	return n % 2 == 0, nil
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
* `keySelector` - func(element, index or key) (anyType, error)

__Return__

* interface{} - map[anyType][]element
* error

__Examples__

```go
v, _ := Group([]int{ 1, 2, 3, 4, 5 }, func (n, _ int) (string, error) {
	if n % 2 == 0 {
		return "even", nil
	}
	return "odd", nil
})
dict, ok := v.(map[string][]int)
if !(ok && len(dict["even"]) == 2) {
	// wrong
}
```

<a name="groupBy" />
### GroupBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{} - map[propertyType][]element
* error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "a" },
	TestModel{ 2, "a" },
	TestModel{ 3, "b" },
	TestModel{ 4, "b" },
}
v, err := GroupBy(arr, "Name")
if err != nil {
	// wrong
}

dict, ok := v.(map[string][]TestModel)
if !(ok && len(dict) == 2) {
	// wrong
}
```

<a name="index" />
### Index(source, indexSelector)

__Arguments__

* `source` - array or map
* `indexSelector` - func(element, index or key) (anyType, error)

__Return__

* interface{} - map[anyType]element
* error

__Examples__

```go
v, _ := Index([]string{ "a", "b" }, func (item string, _ int) (string, error) {
	return item, nil
})
res, ok := v.(map[string]string)
if !(ok && res["a"] == "a") {
	// wrong
}
```

<a name="indexBy" />
### IndexBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{} - map[propertyType]element
* error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "a" },
	TestModel{ 2, "a" },
	TestModel{ 3, "b" },
	TestModel{ 4, "b" },
}
res, err := IndexBy(arr, "Name")
if err != nil {
	// wrong
}

dict, ok := res.(map[string]TestModel)
if !(ok && len(dict) == 2) {
	// wrong
}
```

<a name="map" />
### Map(source, selector)

__Arguments__

* `source` - array
* `selector` - func(element, index or key) (anyType, error)

__Return__

* interface{} - []anyType
* error

__Examples__

```go
arr := []string{ "11", "12", "13" }
v, err := Map(arr, func (s string, _ int) (int, error) {
	return strconv.Atoi(s)
})
if err != nil {
	// wrong
}

res, ok := v.([]int)
if !(ok && len(res) == len(arr)) {
	// wrong
}

for i, s := range arr {
	n, _ := strconv.Atoi(s)
	if n != res[i] {
		// wrong
	}
}
```

<a name="pluck" />
### Pluck(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* interface{} - an array of property type
* error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
	TestModel{ 3, "three" },
}
v, err := Pluck(arr, "Name")
if err != nil {
	// wrong
}

res, ok := v.([]string)
if !(ok && len(res) == len(arr)) {
	// wrong
}

for i := 0; i < 3; i++ {
	if res[i] != arr[i].Name {
		// wrong
	}
}
```

<a name="reduce" />
### Reduce(source, iterator)

__Arguments__

* `source` - array
* `iterator` - func(memo, element, key or index) (memo, error)
* `memo` - anyType

__Return__

* interface{} - memo
* error

__Examples__

```go
v, err := Reduce([]int{ 1, 2 }, func (memo []int, n, _ int) ([]int, error) {
	memo = append(memo, n)
	memo = append(memo, n + 10)
	return memo, nil
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
* `predicate` - func(element, index or key) (bool, error)

__Return__

* interface{} - an array of all the values that pass a truth test `predicate`
* error

__Examples__

```go
arr := []int{ 1, 2, 3, 4 }
v, _ := Select(arr, func (n, i int) (bool, error) {
	return n % 2 == 0, nil
})
res, ok := v.([]int)
if !(ok && len(res) == 2) {
	// wrong
}

if !(res[0] == 2 && res[1] == 4) {
	// wrong
}
```

<a name="selectBy" />
### SelectBy(source, properties)

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
v, err := SelectBy(arr, map[string]interface{}{
	"Id": 1,
})
if err != nil {
	// wrong
}

res, ok := v.([]TestModel)
if !(ok && len(res) == 1 && res[0] == arr[0]) {
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
### Sort(source, selector)

__Arguments__

* `source` - array or map
* `selector` - func(element, key or index) (anyType, error)

__Return__

* interface{} - an array of `source` that sorted
* error

__Examples__

```go
arr := []int{ 1, 2, 3, 5 }
v, _ := Sort([]int{ 5, 3, 2, 1 }, func (n, _ int) (int, error) {
	return n, nil
})
res, ok := v.([]int)
if !(ok && len(res) == len(arr)) {
	// wrong
}

for i, n := range arr {
	if res[i] != n {
		// wrong
	}
}
```

<a name="sortBy" />
### SortBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 3, "three" },
	TestModel{ 1, "one" },
	TestModel{ 2, "two" },
}
v, _ := SortBy(arr, "Id")
res, ok := v.([]TestModel)
if !(ok && len(res) == len(arr)) {
	// wrong
}

if !(res[0].Id < res[1].Id && res[1].Id < res[2].Id) {
	// wrong
}
```

<a name="uniq" />
### Uniq(source, selector)

__Arguments__

* `source` - array
* `selector` - nil or func(element, index or key) (anyType, error)

__Return__

* interface{} - only the first occurence of each value is kept
* error

__Examples__

```go
v, _ := Uniq([]int{ 1, 2, 1, 4, 1, 3 }, func (n, _ int) (int, error) {
	return n % 2, nil
})
res, ok := v.([]int)
if !(ok && len(res) == 2) {
	// wrong
}
```

<a name="uniqBy" />
### UniqBy(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* interface{}, error

__Examples__

```go
arr := []TestModel{
	TestModel{ 1, "one" },
	TestModel{ 2, "one" },
	TestModel{ 3, "one" },
}
v, _ := UniqBy(arr, "Name")
res, ok := v.([]TestModel)
if !(ok && len(res) == 1) {
	// wrong
}
```