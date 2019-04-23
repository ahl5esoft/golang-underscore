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

# Underscore.go [![GoDoc](https://godoc.org/github.com/ahl5esoft/golang-underscore?status.svg)](https://godoc.org/github.com/ahl5esoft/golang-underscore) [![Go Report Card](https://goreportcard.com/badge/github.com/ahl5esoft/golang-underscore)](https://goreportcard.com/report/github.com/ahl5esoft/golang-underscore) ![Version](https://img.shields.io/badge/version-1.0.0-green.svg)
like <a href="http://underscorejs.org/">underscore.js</a>, but for Go

## Installation
    $ go get github.com/ahl5esoft/golang-underscore

## Update
	$ go get -u github.com/ahl5esoft/golang-underscore

## Lack
* more...

## Suggest
	always using `chain`

## Documentation

### API
* [`All`](#all), [`AllBy`](#allBy)
* [`Any`](#any), [`AnyBy`](#anyBy)
* [`AsParallel`](#asParallel)
* [`Chain`](#chain)
* [`Clone`](#clone)
* [`Each`](#each)
* [`Find`](#find), [`FindBy`](#findBy)
* [`FindIndex`](#findIndex), [`FindIndexBy`](#findIndexBy)
* [`First`](#first)
* [`Group`](#group), [`GroupBy`](#groupBy)
* [`IsArray`](#isArray)
* [`IsMatch`](#isMatch)
* [`Keys`](#keys)
* [`Last`](#last)
* [`Map`](#map), [`MapBy`](#mapBy)
* [`MapMany`](#mapMany), [`MapManyBy`](#mapManyBy)
* [`Md5`](#md5)
* [`Object`](#object)
* [`Property`](#property), [`PropertyRV`](#propertyRV)
* [`Range`](#range)
* [`Reduce`](#reduce)
* [`Reject`](#reject), [`RejectBy`](#rejectBy)
* [`Reverse`](#reverse), [`ReverseBy`](#reverseBy)
* [`Size`](#size)
* [`Sort`](#sort), [`SortBy`](#sortBy)
* [`Take`](#take)
* [`Uniq`](#uniq), [`UniqBy`](#uniqBy)
* [`UUID`](#uuid)
* [`Value`](#value)
* [`Values`](#values)
* [`Where`](#where), [`WhereBy`](#whereBy)

<a name="all" />

### All(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element, index or key) bool

__Return__

* bool - all the values that pass a truth test `predicate`

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}
ok := All(src, func(r testModel, _ int) bool {
	return r.Id == 1
})
// ok == true
```

<a name="allBy" />

### AllBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* bool

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}
ok := AllBy(src, nil)
// ok == true

ok = AllBy(src, map[string]interface{}{
	"name": "a",
})
// ok == false

ok = AllBy(src, map[string]interface{}{
	"id": 1,
})
// ok == true
```

<a name="any" />

### Any(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element or value, index or key) bool

__Return__

* bool - any of the values that pass a truth test `predicate`

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
ok := Any(src, func(r testModel, _ int) bool {
	return r.Id == 0
})
// ok == false
```

<a name="anyBy" />

### AnyBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* bool

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
ok := AnyBy(src, map[string]interface{}{
	"Id": 0,
})
// ok == false

ok = AnyBy(src, map[string]interface{}{
	"id":   src[0].Id,
	"name": src[0].Name,
})
// ok == true
```

<a name="asParallel" />

### Chain(source).AsParallel()...

__Support__
* `Each`
* `Object`

__Examples__

```go
arr := []int{ 1, 2, 3 }
Chain(arr).AsParallel().Each(func (n, i int) {
	// code
})
```

<a name="chain" />

### Chain(source)

__Arguments__

* `source` - array or map

__Return__

* IQuery - a wrapped object, wrapped objects until value is called

__Examples__

```go
res := make(map[string][]int)
Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
	if n%2 == 0 {
		return "even"
	}

	return "old"
}).Value(&res)
// len(res) == 2 && ok == true
```

<a name="clone" />

### Clone()

__Return__

* interface{}

__Examples__

```go
arr := []int{1, 2, 3}
var duplicate []int
Chain(arr).Clone().Value(&duplicate)
// or
duplicate := Clone(arr)
ok := All(duplicate, func(n, i int) bool {
	return arr[i] == n
})
// ok == true
```

<a name="each" />

### Each(source, iterator)

__Arguments__

* `source` - array or map
* `iterator` - func(element or value, index or key)

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}
Each(src, func (r testModel, i int) {
	// coding
})
```

<a name="find" />

### Find(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element or value, index or key) bool

__Return__

* interface{} -- `source` elem

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var item testModel
Chain(src).Find(func(r testModel, _ int) bool {
	return r.ID == 1
}).Value(&item)
// or
item := Find(src, func(r testModel, _ int) bool {
	return r.ID == 1
})
// item == arr[0]
```

<a name="findBy" />

### FindBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* interface{} - `source` elem

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var item testModel
Chain(arr).FindBy(map[string]interface{}{
	"id": 2,
}).Value(&item)
// or
item := FindBy(arr, map[string]interface{}{
	"id": 2,
})
// item == arr[1]
```

<a name="findIndex" />

### FindIndex(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element or value, index or key) bool

__Return__

* int - index

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}
i := FindIndex(arr, func(r testModel, _ int) bool {
	return r.Name == arr[1].Name
})
// i == 1
```

<a name="findIndexBy" />

### FindIndexBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* int - index

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
i := FindIndexBy(arr, map[string]interface{}{
	"id": 1,
})
// i == 0
```

<a name="first" />

### First(source)

__Arguments__

* `source` - array or map

__Return__

* interface{}

__Examples__

```go
arr := []int{ 1, 2, 3 }
var res int
Chain(arr).First().Value(&res)
// or
res := First(arr).(int)
// res = 1

res := First(nil) 
// res == nil
```

<a name="group" />

### Group(source, keySelector)

__Arguments__

* `source` - array or map
* `keySelector` - func(element or value, index or key) anyType

__Return__

* interface{} - map[anyType][](element or value)

__Examples__

```go
src := []int{ 1, 2, 3, 4, 5 }
var res map[string][]int
Chain(src).Group(func (n, _ int) string {
	if n % 2 == 0 {
		return "even"
	}
	return "odd"
}).Value(&res)
// or
res := Group(src, func (n, _ int) string {
	if n % 2 == 0 {
		return "even"
	}
	return "odd"
}).(map[string][]int)
// res = map[odd:[1 3 5] even:[2 4]]
```

<a name="groupBy" />

### GroupBy(source, property)

__Arguments__

* `source` - array or map
* `property` - property name

__Return__

* interface{} - map[property type][](element or value)

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "a"},
	{ID: 2, Name: "a"},
	{ID: 3, Name: "b"},
	{ID: 4, Name: "b"},
}
var res map[string][]testModel
Chain(arr).GroupBy("name").Value(&res)
// or
res := GroupBy(arr, "name").(map[string][]testModel)
// res = map[a:[{{0} 1 a} {{0} 2 a}] b:[{{0} 3 b} {{0} 4 b}]]
```

<a name="index" />

### Index(source, indexSelector)

__Arguments__

* `source` - array or map
* `indexSelector` - func(element or value, index or key) anyType

__Return__

* interface{} - map[anyType](element or value)

__Examples__

```go
src := []string{ "a", "b" }
var res map[string]string
Chain(src).Index(func (r string, _ int) string {
	return r
}).Value(&res)
// or
res := Index(src, func (r string, _ int) string {
	return r
}).(map[string]string)
// res = map[a:a b:b]
```

<a name="indexBy" />

### IndexBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{} - map[propertyType](element or value)

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "a"},
	{ID: 2, Name: "a"},
	{ID: 3, Name: "b"},
	{ID: 4, Name: "b"},
}
var res map[int]testModel
Chain(arr).IndexBy("id").Value(&res)
// or
res := IndexBy(arr, "id").(map[int]testModel)
// res = map[1:{{0} 1 a} 2:{{0} 2 a} 3:{{0} 3 b} 4:{{0} 4 b}]
```

<a name="isArray" />

### IsArray(element)

__Arguments__

* `element` - object

__Return__

* bool

__Examples__

```go
if !IsArray([]int{}) {
	// wrong
}

if IsArray(map[string]int{}) {
	// wrong
}
```

<a name="isMatch" />

### IsMatch(element, properties)

__Arguments__

* `element` - object
* `properties` - map[string]interface{}

__Return__

* bool

__Examples__

```go
m := testModel{ 1, "one" }
ok := IsMatch(nil, nil)
// ok = false

ok = IsMatch(m, nil)
// ok = false

ok = IsMatch(m, map[string]interface{}{
	"id": m.Id,
	"name": "a",
})
// ok = false

ok = IsMatch(m, map[string]interface{}{
	"id": m.Id,
	"name": m.Name,
})
// ok = true
```

<a name="keys" />

### Keys()

__Arguments__

* `source` - map

__Return__

* interface{} - []keyType

__Examples__

```go
arr := []string{ "aa" }
v := Keys(arr)
// v = nil

dict := map[int]string{	
	1: "a",
	2: "b",
	3: "c",
	4: "d",
}
var res []int
Chain(dict).Keys().Value(&res)
// or
res := Keys(dict).([]int)
// res = [1 2 3 4]
```

<a name="last" />

### Last(source)

__Arguments__

* `source` - array or map

__Return__

* interface{} - last element of `source`

__Examples__

```go
arr := []int{1, 2, 3}
var res int
chain(arr).Last().Value(&res)
// or
res := Last(arr).(int)
// res = 3

dict := map[string]string{
	"a": "aa",
	"b": "bb",
}
var str string
Chain(dict).Last().Value(&str)
// or
str := Last(dict).(string)
// res = "aa" or "bb"
```

<a name="map" />

### Map(source, selector)

__Arguments__

* `source` - array or map
* `selector` - func(element, index or key) anyType

__Return__

* interface{} - an slice of property value

__Examples__

```go
arr := []string{ "11", "12", "13" }
var res []int
Chain(arr).Map(func (s string, _ int) int {
	n, _ := strconv.Atoi(s)
	return n
}).Value(&res)
// or
res := Map(arr, func (s string, _ int) int {
	n, _ := strconv.Atoi(s)
	return n
}).([]int)
// res = [11 12 13]
```

<a name="mapBy" />

### MapBy(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* interface{} - an slice of property value

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var res []string
Chain(arr).MapBy("name").Value(&res)
// or
res := MapBy(arr, "name").([]string)
// res = [one two three]
```

<a name="mapMany" />

### MapMany(source, selector)

__Arguments__

* `source` - array or map
* `selector` - func(element, index or key) anyType with array or slice

__Return__

* interface{} - an slice of property elem value

__Examples__

```go
src := []int{1, 2}
MapMany(src, func(r, _ int) int {
	return r // will panic because `r` is not array or slice
})
// or
Chain(src).MapMany(func(r, _ int) int {
	return r // will panic because `r` is not array or slice
})

var res []int
Chain(src).MapMany(func(r, _ int) []int {
	var temp []int
	Range(0, r, 1).Map(func(_, _ int) int {
		return r
	}).Value(&temp)
	return temp
}).Value(&res)
// or
res := MapMany(src, func(r, _ int) []int {
	var temp []int
	Range(0, r, 1).Map(func(_, _ int) int {
		return r
	}).Value(&temp)
	return temp
}).([]int)
// res = [1 2 2]
```

<a name="mapManyBy" />

### MapManyBy(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* interface{} - an slice of propery elem value

__Examples__

```go
src := []mapManyTestModel{
	{Slice: []string{"a", "b"}},
	{Slice: []string{"c", "d"}},
}
Chain(src).MapManyBy("Str") // will panic because `Str` property value is not array or slice
// or
MapManyBy(src, "Str") // will panic because `Str` property value is not array or slice

var res []string
Chain(src).MapManyBy("Slice").Value(&res)
// or
res := MapManyBy(src, "Slice").([]string)
// res = [a b c d]
```

<a name="md5" />

### Md5(plaintext)

__Arguments__

* `plaintext` - string

__Return__

* string - md5 string

__Examples__

```go
if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
	// wrong
}	
```

<a name="object" />

### Object(arr)

__Arguments__

* `arr` - array

__Return__

* map, error

__Examples__

```go
arr := []interface{}{
	[]interface{}{"a", 1},
	[]interface{}{"b", 2},
}
var res map[string]int
Chain(arr).Object().Value(&res)
// or
res := Object(arr).(map[string]int)
// res = map[b:2 a:1] or map[a:1 b:2]
```

<a name="property" />

### Property(name)

__Arguments__

* `name` - property name

__Return__

* func(interface{}) (interface{}, error)

__Examples__

```go
item := testModel{ 1, "one" }

getAge := Property("age")
_, err := getAge(item)
// err != nil

getName := Property("name")
name, err := getName(item)
// name = "one"
```

<a name="propertyRV" />

### Property(name)

__Arguments__

* `name` - property name

__Return__

* func(interface{}) (reflect.Value, error)

__Examples__

```go
item := testModel{ 1, "one" }

getAgeRV := PropertyRV("age")
_, err := getAgeRV(item)
// err != nil

getNameRV := PropertyRV("name")
nameRV, err := getNameRV(item)
// nameRV = reflect.ValueOf("one")
```

<a name="range" />

### Range(start, stop, step)

__Arguments__

* `start` - int
* `stop` - int
* `step` - int

__Return__

* IQuery - a wrapped object, wrapped objects until value is called

__Examples__

```go
var res []int
Range(0, 0, 1).Value(&res)
// res = []

var res []int
Range(0, 10, 0).Value(&res)
// res = []

var res []int
Range(10, 0, 1).Value(&res)
// res = []

var res []int
Range(0, 2, 1).Value(&res)
// res = [0 1]

var res []int
Range(0, 3, 2).Value(&res)
// res = [0 2]
```

<a name="reduce" />

### Reduce(source, iterator)

__Arguments__

* `source` - array
* `iterator` - func(memo, element or value, key or index) memo
* `memo` - anyType

__Return__

* interface{} - memo

__Examples__

```go
var res []int
Chain([]int{1, 2}).Reduce(func(memo []int, n, _ int) []int {
	memo = append(memo, n)
	memo = append(memo, n+10)
	return memo
}, make([]int, 0)).Value(&res)
// or
res := Reduce([]int{1, 2}, func(memo []int, n, _ int) []int {
	memo = append(memo, n)
	memo = append(memo, n+10)
	return memo
}, make([]int, 0)).([]int)
// res = [1 11 2 12]
```

<a name="reject" />

### Reject(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element or value, index or key) bool

__Return__

* interface{} - an array of all the values that without pass a truth test `predicate`

__Examples__

```go
arr := []int{1, 2, 3, 4}
var res []int
Chain(arr).Reject(func(n, i int) bool {
	return n%2 == 0
}).Value(&res)
// or
res := Reject(arr, func(n, i int) bool {
	return n%2 == 0
}).([]int)
// res = [1 3]
```

<a name="rejectBy" />

### RejectBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* interface{} - an array of all the values that without pass a truth test `properties`

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(arr).RejectBy(map[string]interface{}{
	"Id": 1,
}).Value(&res)
// or
res := RejectBy(arr, map[string]interface{}{
	"Id": 1,
}).([]testModel)
// res = [{{0} 2 two} {{0} 3 three}]
```

<a name="reverse" />

### Reverse(source, selector)

__Arguments__

* `source` - array or map
* `selector` - func(element, key or index) anyType

__Return__

* interface{} - an array of `source` that reversed

__Examples__

```go
arr := []testModel{
	{ID: 2, Name: "two"},
	{ID: 1, Name: "one"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(arr).Reverse(func(n testModel, _ int) int {
	return n.ID
}).Value(&res)
// or
res := Reverse(arr, func(n testModel, _ int) int {
	return n.ID
}).([]testModel)
// res = [{{0} 3 three} {{0} 2 two} {{0} 1 one}]
```

<a name="reverseBy" />

### ReverseBy(source, selector)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{} - an array of `source` that reversed

__Examples__

```go
arr := []testModel{
	{ID: 2, Name: "two"},
	{ID: 1, Name: "one"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(arr).ReverseBy("id").Value(&res)
// or
res := ReverseBy(arr, "id").([]testModel)
// res = [{{0} 3 three} {{0} 2 two} {{0} 1 one}]
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
l := Size(dict)
// l = 3
```

<a name="sort" />

### Sort(source, selector)

__Arguments__

* `source` - array or map
* `selector` - func(element, key or index) anyType

__Return__

* interface{} - an array of `source` that sorted

__Examples__

```go
arr := []testModel{
	{ID: 2, Name: "two"},
	{ID: 1, Name: "one"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(arr).Sort(func(n testModel, _ int) int {
	return n.ID
}).Value(&res)
// or
res := Sort(arr, func(n testModel, _ int) int {
	return n.ID
}).([]testModel)
// res = [{{0} 1 one} {{0} 2 two} {{0} 3 three}]
```

<a name="sortBy" />

### SortBy(source, property)

__Arguments__

* `source` - array or map
* `property` - string

__Return__

* interface{}

__Examples__

```go
arr := []testModel{
	{ID: 2, Name: "two"},
	{ID: 1, Name: "one"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(arr).SortBy("id").Value(&res)
// or
res := SortBy(arr, "id").([]testModel)
// res = [{{0} 1 one} {{0} 2 two} {{0} 3 three}]
```

<a name="take" />

### Take(source, count)

__Arguments__

* `source` - array or map
* `count` - int

__Return__

* interface{}

__Examples__

```go
arr := []int{1, 2, 3}
var res []int
Chain(arr).Take(1).Value(&res)
// or
res := Take(arr, 1).([]int)
// res = [1]
```

<a name="uniq" />

### Uniq(source, selector)

__Arguments__

* `source` - array
* `selector` - nil or func(element or value, index or key) anyType

__Return__

* interface{} - only the first occurence of each value is kept

__Examples__

```go
arr := []int{1, 2, 1, 4, 1, 3}
var res []int
Chain(arr).Uniq(func(n, _ int) int {
	return n % 2
}).Value(&res)
// or
res := Uniq(arr, func(n, _ int) int {
	return n % 2
}).([]int)
// res = [1 2]
```

<a name="uniqBy" />

### UniqBy(source, property)

__Arguments__

* `source` - array
* `property` - string

__Return__

* interface{}

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "one"},
	{ID: 3, Name: "one"},
}
var res []testModel
Chain(arr).UniqBy("name").Value(&res)
// or
res := UniqBy(arr, "Name").([]testModel)
// res = [{{0} 1 one}]
```

<a name="uuid" />

### UUID()

__Return__

* string - uuid string

__Examples__

```go
uuid := UUID()
//1a40272540e57d1c80e7b06042219d0c
```

<a name="value" />

### Value(result)

__Examples__

```go
arr := []int{1, 2, 1, 4, 1, 3}
var res map[string][]int
Chain(arr).Uniq(nil).Group(func(n, _ int) string {
	if n%2 == 0 {
		return "even"
	}

	return "old"
}).Value(&res)
// res = map[old:[1 3] even:[2 4]]
```

<a name="values" />

### Values(source)

__Arguments__

* `source` - map

__Return__

* interface{} - an array of `source`'s values

__Examples__

```go
src := map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
}
var res []string
Chain(src).Values().Value(&res)
// or
res := Values(src).([]string)
// res = [a b c d]
```

<a name="where" />

### Where(source, predicate)

__Arguments__

* `source` - array or map
* `predicate` - func(element or value, index or key) bool

__Return__

* interface{} - an array of all the values that pass a truth test `predicate`

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(src).Where(func(r testModel, i int) bool {
	return r.ID%2 == 0
}).Value(&res)
// or
res := Where(src, func(r testModel, i int) bool {
	return r.ID%2 == 0
}).([]testModel)
// res = [{{0} 2 two}]
```

<a name="whereBy" />

### WhereBy(source, properties)

__Arguments__

* `source` - array or map
* `properties` - map[string]interface{}

__Return__

* interface{} - an array of all the values that pass a truth test `properties`

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var res []testModel
Chain(src).WhereBy(map[string]interface{}{
	"name": "one",
}).Value(&res)
// or
res := WhereBy(src, map[string]interface{}{
	"name": "one",
}).([]testModel
// res = [{{0} 1 one}]
```

## Release Notes

~~~
v1.0.0 (2018-04-23)
* first edition
~~~