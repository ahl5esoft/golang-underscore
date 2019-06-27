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

# Underscore.go [![GoDoc](https://godoc.org/github.com/ahl5esoft/golang-underscore?status.svg)](https://godoc.org/github.com/ahl5esoft/golang-underscore) [![Go Report Card](https://goreportcard.com/badge/github.com/ahl5esoft/golang-underscore)](https://goreportcard.com/report/github.com/ahl5esoft/golang-underscore) ![Version](https://img.shields.io/badge/version-2.0.0-green.svg)
like <a href="http://underscorejs.org/">underscore.js</a> and C# LINQ, but for Go

## Installation
    $ go get github.com/ahl5esoft/golang-underscore

## Update
	$ go get -u github.com/ahl5esoft/golang-underscore

## Lack
* Except、ExceptBy、Last、Order、OrderBy、Reject、RejectBy、Reverse、ReverseBy、Sort、SortBy、ThenBy
* IQuery性能差，将来会逐步用IEnumerable替代

## Documentation

### API
* [`Aggregate`](#aggregate)
* [`All`](#all), [`AllBy`](#allBy)
* [`Any`](#any), [`AnyBy`](#anyBy)
* [`Chain`](#chain)
* [`Count`](#count)
* [`Distinct`](#distinct), [`DistinctBy`](#distinctBy)
* [`Each`](#each)
* [`Filter`](#where), [`FilterBy`](#whereBy)
* [`Find`](#find), [`FindBy`](#findBy)
* [`FindIndex`](#findIndex), [`FindIndexBy`](#findIndexBy)
* [`First`](#first)
* [`Group`](#group), [`GroupBy`](#groupBy)
* [`Index`](#index), [`IndexBy`](#indexBy)
* [`IsArray`](#isArray)
* [`IsMatch`](#isMatch)
* [`Keys`](#keys)
* [`Last`](#last)
* [`Map`](#select), [`MapBy`](#selectBy)
* [`MapMany`](#selectMany), [`MapManyBy`](#selectManyBy)
* [`Object`](#object)
* [`Property`](#property), [`PropertyRV`](#propertyRV)
* [`Range`](#range)
* [`Reduce`](#aggregate)
* [`Reject`](#reject), [`RejectBy`](#rejectBy)
* [`Reverse`](#reverse), [`ReverseBy`](#reverseBy)
* [`Select`](#select), [`SelectBy`](#selectBy)
* [`SelectMany`](#selectMany), [`SelectManyBy`](#selectManyBy)
* [`Size`](#count)
* [`Skip`](#skip)
* [`Sort`](#sort), [`SortBy`](#sortBy)
* [`Take`](#take)
* [`Uniq`](#distinct), [`UniqBy`](#distinctBy)
* [`Values`](#values)
* [`Where`](#where), [`WhereBy`](#whereBy)

<a name="aggregate" />

### Aggregate(memo, fn) IEnumerable

__Arguments__

* `memo` - anyType
* `iterator` - func(memo, element or value, key or index) memo

__Examples__

```go
dst := make([]int, 0)
Chain([]int{1, 2}).Aggregate(make([]int, 0), func(memo []int, n, _ int) []int {
	memo = append(memo, n)
	memo = append(memo, n+10)
	return memo
}).Value(&dst)
// dst = [1 11 2 12]
```

__Same__

* `Reduce`

<a name="all" />

### All(predicate) bool

__Arguments__

* `predicate` - func(element, index or key) bool

__Return__

* bool - all the values that pass a truth test `predicate`

__Examples__

```go
ok := Chain([]testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}).All(func(r testModel, _ int) bool {
	return r.ID == 1
})
// ok == true
```

<a name="allBy" />

### AllBy(properties) bool

__Arguments__

* `properties` - map[string]interface{}

__Return__

* bool - all the values that pass a truth test `predicate`

__Examples__

```go
ok := Chain([]testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "one"},
	{ID: 3, Name: "one"},
}).AllBy(map[string]interface{}{
	"name": "one",
})
// ok == true
```

<a name="any" />

### Any(predicate) bool

__Arguments__

* `predicate` - func(element or value, index or key) bool

__Return__

* bool - any of the values that pass a truth test `predicate`

__Examples__

```go
ok := Chain([]testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}).Any(func(r testModel, _ int) bool {
	return r.ID == 0
})
// ok == false
```

<a name="anyBy" />

### AnyBy(properties) bool

__Arguments__

* `properties` - map[string]interface{}

__Return__

* bool

__Examples__

```go
ok := Chain([]testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}).AnyBy(map[string]interface{}{
	"name": "two",
})
// ok == true
```

<a name="chain" />

### Chain(source) IEnumerable

__Arguments__

* `source` - array or map

__Examples__

```go
var dst int
Range(1, benchmarkSize, 1).Select(func(r, _ int) int {
	return -r
}).Where(func(r, _ int) bool {
	return r < -20
}).First().Value(&dst)
// dst = -21
```

<a name="count" />

### Count() int

__Examples__

```go
src := []string{"a", "b", "c"}
dst := Chain(src).Count()
// dst = 3
```

__Same__
* `Size`

<a name="distinct" />

### Distinct(selector) IEnumerable

__Arguments__

* `selector` - nil or func(element or value, index or key) anyType

__Examples__

```go
src := []int{1, 2, 1, 4, 1, 3}
dst := make([]int, 0)
Chain(src).Distinct(func(n, _ int) (int, error) {
	return n % 2, nil
}).Value(&dst)
// dst = [1 2]
```

__Same__

* `Uniq`

<a name="distinctBy" />

### DistinctBy(fieldName) IEnumerable

__Arguments__

* `fieldName` - string

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "a"},
	{ID: 2, Name: "a"},
	{ID: 3, Name: "a"},
}
dst := make([]testModel, 0)
Chain(src).DistinctBy("name").Value(&dst)
// dst = [{1 a}]
```

__Same__

* `UniqBy`

<a name="each" />

### Each(iterator)

__Arguments__

* `iterator` - func(element or value, index or key)

__Examples__

```go
arr := []testModel{
	{ID: 1, Name: "one"},
	{ID: 1, Name: "two"},
	{ID: 1, Name: "three"},
}
Chain(arr).Each(func(r testModel, i int) {
	if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
		// wrong
	}
})
```

<a name="find" />

### Find(predicate) IEnumerable

__Arguments__

* `predicate` - func(element or value, index or key) bool

__Examples__

```go
var dst int
Chain([]int{1, 2, 3}).Find(func(r, _ int) bool {
	return r == 2
}).Value(&dst)
// dst == 2
// or
var dst int
Chain([][]int{
	[]int{1, 3, 5, 7},
	[]int{2, 4, 6, 8},
}).Find(func(r []int, _ int) bool {
	return r[0]%2 == 0
}).Find(func(r, _ int) bool {
	return r > 6
}).Value(&dst)
// dst == 8
```

<a name="findBy" />

### FindBy(properties) IEnumerable

__Arguments__

* `properties` - map[string]interface{}

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
var dst testModel
Chain(src).FindBy(map[string]interface{}{
	"id": 2,
}).Value(&dst)
// dst == src[1]
```

<a name="findIndex" />

### FindIndex(predicate) int

__Arguments__

* `predicate` - func(element or value, index or key) bool

__Return__

* int - index

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
index := Chain(src).FindIndex(func(r testModel, _ int) bool {
	return r.Name == src[1].Name
})
// i == 1
```

<a name="findIndexBy" />

### FindIndexBy(properties) int

__Arguments__

* `properties` - map[string]interface{}

__Return__

* int - index

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
index := Chain(src).FindIndexBy(map[string]interface{}{
	"id": 1,
})
// index == 0
```

<a name="first" />

### First() IEnumerable

__Arguments__

* `predicate` - func(element or value, index or key) bool

__Examples__

```go
var dst int
Chain([]int{1, 2, 3}).First().Value(&dst)
// dst == 1
// or
var dst int
Chain([][]int{
	[]int{1, 3, 5, 7},
	[]int{2, 4, 6, 8},
}).First().First().Value(&dst)
// dst == 1
```

<a name="group" />

### Group(keySelector) IEnumerable

__Arguments__

* `keySelector` - func(element or value, index or key) anyType

__Examples__

```go
dst := make(map[string][]int)
Chain([]int{1, 2, 3, 4, 5}).Group(func(n, _ int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}).Value(&dst)
// dst = map[odd:[1 3 5] even:[2 4]]
```

<a name="groupBy" />

### GroupBy(fieldName) IEnumerable

__Arguments__

* `fieldName` - field name

__Examples__

```go
dst := make(map[string][]testModel)
Chain([]testModel{
	{ID: 1, Name: "a"},
	{ID: 2, Name: "a"},
	{ID: 3, Name: "b"},
	{ID: 4, Name: "b"},
}).GroupBy("Name").Value(&dst)
// dst = map[a:[{1 a} {2 a}] b:[{3 b} {4 b}]]
```

<a name="index" />

### Index(indexSelector) IEnumerable

__Arguments__

* `indexSelector` - func(element or value, index or key) anyType

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

### IndexBy(property) IEnumerable

__Arguments__

* `property` - string

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

### IsArray(element) bool

__Arguments__

* `element` - object

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

### IsMatch(element, properties) bool

__Arguments__

* `element` - object
* `properties` - map[string]interface{}

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

### Keys() IEnumerable

__Examples__

```go
src := []string{"aa", "bb", "cc"}
dst := make([]int, 0)
Chain(src).Keys().Value(&dst)
// dst = [0 1 2]

src := map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
}
dst := make([]int, 0)
Chain(src).Keys().Value(&dst)
// dst = [1 2 3 4]
```

<a name="last" />

### Last() IEnumerable

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

<a name="object" />

### Object() IEnumerable

__Examples__

```go
src := [][]interface{}{
	[]interface{}{"a", 1},
	[]interface{}{"b", 2},
}
dst := make(map[string]int)
Chain(src).Object().Value(&dst)
// dst = map[a:1 b:2]
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

### Range(start, stop, step) IEnumerable

__Arguments__

* `start` - int
* `stop` - int
* `step` - int

__Examples__

```go
var res []int
Range2(0, 0, 1).Value(&res)
// res = []

var res []int
Range2(0, 10, 0).Value(&res)
// panic

var res []int
Range2(4, 0, -1).Value(&res)
// res = [4 3 2 1]

var res []int
Range2(0, 2, 1).Value(&res)
// res = [0 1]

var res []int
Range2(0, 3, 2).Value(&res)
// res = [0 2]
```

<a name="reject" />

### Reject(predicate) IEnumerable

__Arguments__

* `predicate` - func(element or value, index or key) bool

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

### RejectBy(properties) IEnumerable

__Arguments__

* `properties` - map[string]interface{}

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

### Reverse(selector) IEnumerable

__Arguments__

* `selector` - func(element, key or index) anyType


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

### ReverseBy(fieldName) IEnumerable

__Arguments__

* `fieldName` - string

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

<a name="select" />

### Select(selector) IEnumerable

__Arguments__

* `selector` - func(element, index or key) anyType

__Examples__

```go
src := []string{"11", "12", "13"}
dst := make([]int, 0)
Chain(src).Select(func(s string, _ int) int {
	n, _ := strconv.Atoi(s)
	return n
}).Value(&dst)
// dst = [11 12 13]
```

__Same__

* `Map`

<a name="selectBy" />

### SelectBy(fieldName) IEnumerable

__Arguments__

* `fieldName` - string

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "two"},
	{ID: 3, Name: "three"},
}
dst := make([]string, 0)
Chain(src).SelectBy("name").Value(&dst)
// dst = [one two three]
```

__Same__

* `MapBy`

<a name="selectMany" />

### SelectMany(selector) IEnumerable

__Arguments__

* `selector` - func(element, index or key) anyType with array or slice

__Examples__

```go
src := [2]int{1, 2}
var dst []int
Chain(src).SelectMany(func(r, _ int) []int {
	return []int{r - 1, r + 1}
}).Value(&dst)
// dst = [0 2 1 3]
```

__Same__

* `MapMany`

<a name="mapManyBy" />

### MapManyBy(property) IEnumerable

__Arguments__

* `property` - string

__Examples__

```go
src := []testSelectManyModel{
	{Array: [2]int{1, 2}},
	{Array: [2]int{3, 4}},
}
var dst []int
Chain(src).SelectManyBy("Array").Value(&dst)
// res = [1 2 3 4]
```

__Same__

* `MapManyBy`

<a name="skip" />

### Skip(count) IEnumerable

__Arguments__

* `count` - int

__Examples__

```go
src := []int{1, 2, 3}
dst := make([]int, 0)
Chain(src).Skip(2).Value(&dst)
// dst = [3]
```

<a name="sort" />

### Sort(selector) IEnumerable

__Arguments__

* `selector` - func(element, key or index) anyType


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

### SortBy(fieldName) IEnumerable

__Arguments__

* `fieldName` - string

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

### Take(count) IEnumerable

__Arguments__

* `count` - int

__Examples__

```go
src := []int{1, 2, 3}
dst := make([]int, 0)
Chain(src).Take(1).Value(&dst)
// res = [1]
```

<a name="values" />

### Values() IEnumerable

__Examples__

```go
src := []string{"a", "b"}
dst := make([]string, 0)
Chain(src).Values().Value(&dst)
// dst = [a b]

src := map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
}
dst := make([]string, 0)
Chain(src).Values().Value(&dst)
// dst = [a b c d]
```

<a name="where" />

### Where(predicate) IEnumerable

__Arguments__

* `predicate` - func(element or value, index or key) bool

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "one"},
	{ID: 3, Name: "three"},
	{ID: 4, Name: "three"},
}
dst := make([]testModel, 0)
Chain(src).Where(func(r testModel, _ int) bool {
	return r.ID%2 == 0
}).Value(&dst)
// len(dst) == 2 && dst[0] == src[1] && dst[1] == src[3])
```

__Same__

* `Filter`

<a name="whereBy" />

### WhereBy(properties) IEnumerable

__Arguments__

* `properties` - map[string]interface{}

__Examples__

```go
src := []testModel{
	{ID: 1, Name: "one"},
	{ID: 2, Name: "one"},
	{ID: 3, Name: "three"},
	{ID: 4, Name: "three"},
}
dst := make([]testModel, 0)
Chain(src).WhereBy(map[string]interface{}{
	"Name": "one",
}).Value(&dst)
// len(dst) == 2 && dst[0] == src[0] && dst[1] == src[1]
```

__Same__

* `FilterBy`

## Release Notes
~~~
v2.0.0 (2019-06-27)
* 删除IQuery
* IEnumerable增加MapMany、MapManyBy、SelectMany、SelectManyBy
~~~

~~~
v1.6.0 (2019-06-21)
* IEnumerable增加Count、Size
* 删除FindLastIndex
~~~

~~~
v1.5.0 (2019-06-18)
* 增加Chain Benchmark
* IEnumerable增加Group、GroupBy
* 优化IEnumerable的Distinct、Enumerator、Index、Property、Select、Where
~~~

~~~
v1.4.0 (2019-06-15)
* Reduce、Take支持IEnumerable
* IEnumerable增加Aggregate、Skip
* IQuery删除Clone
* 优化IEnumerable的First、Index、Values
~~~

~~~
v1.3.0 (2019-06-09)
* FindIndex、FindIndexBy、Keys、Map、MapBy、Object、Uniq、UniqBy、Values支持IEnumerable
* IEnumerable增加Distinct、DistinctBy、Select、SelectBy
~~~

~~~
v1.2.0 (2019-06-04)
* Each、Filter、Where支持IEnumerable
~~~

~~~
v1.1.0 (2019-06-02)
* 增加IEnumerable、IEnumerator
* All、Any、Chain、Find、First、Range2、Value支持IEnumerable
~~~

~~~
v1.0.0 (2019-04-23)
* first edition
~~~