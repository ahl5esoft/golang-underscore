package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Order(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		arr := []testModel{
			{ID: 2, Name: "two"},
			{ID: 1, Name: "one"},
			{ID: 3, Name: "three"},
		}
		var res []int
		Chain(arr).Order(func(n testModel, _ int) int {
			return n.ID
		}).Map(func(r testModel, _ int) int {
			return r.ID
		}).Value(&res)
		assert.Len(
			t,
			res,
			len(arr),
		)
		assert.EqualValues(
			t,
			res,
			[]int{1, 2, 3},
		)
	})

	t.Run("chain", func(t *testing.T) {
		arr := []testModel{
			{ID: 2, Name: "two"},
			{ID: 1, Name: "one"},
			{ID: 3, Name: "three"},
		}
		var res []int
		Chain(arr).Where(func(r testModel, _ int) bool {
			return r.ID > 5
		}).Order(func(r testModel, _ int) int {
			return r.ID
		}).Map(func(r testModel, _ int) int {
			return r.ID
		}).Value(&res)
	})
}

func Test_OrderBy(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []string
	Chain(arr).OrderBy("id").Map(func(r testModel, _ int) string {
		return r.Name
	}).Value(&res)
	assert.Len(t, res, 3)
	assert.EqualValues(
		t,
		res,
		[]string{"one", "two", "three"},
	)
}

func Test_OrderByFields(t *testing.T) {
	// 创建测试数据
	users := []struct {
		Name  string
		Age   int
		Score int
	}{{
		Name:  "Alice",
		Age:   25,
		Score: 85,
	}, {
		Name:  "Bob",
		Age:   30,
		Score: 90,
	}, {
		Name:  "Charlie",
		Age:   25,
		Score: 95,
	}, {
		Name:  "David",
		Age:   30,
		Score: 80,
	}}

	// 按年龄升序，然后按分数降序排序
	var result []struct {
		Name  string
		Age   int
		Score int
	}

	Chain(users).OrderByFields("Age", "Score").Value(&result)

	// 验证排序结果
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "Alice", result[0].Name)   // 年龄25，分数85
	assert.Equal(t, "Charlie", result[1].Name) // 年龄25，分数95
	assert.Equal(t, "David", result[2].Name)   // 年龄30，分数80
	assert.Equal(t, "Bob", result[3].Name)     // 年龄30，分数90
}

func Test_OrderByFields_WithSingleField(t *testing.T) {
	// 创建测试数据
	users := []struct {
		Name string
		Age  int
	}{{
		Name: "Alice",
		Age:  25,
	}, {
		Name: "Bob",
		Age:  30,
	}, {
		Name: "Charlie",
		Age:  20,
	}}

	// 按年龄升序排序
	var result []struct {
		Name string
		Age  int
	}

	Chain(users).OrderByFields("Age").Value(&result)

	// 验证排序结果
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "Charlie", result[0].Name) // 年龄20
	assert.Equal(t, "Alice", result[1].Name)   // 年龄25
	assert.Equal(t, "Bob", result[2].Name)     // 年龄30
}

func Test_OrderByFields_WithThreeFields(t *testing.T) {
	// 创建测试数据
	products := []struct {
		Category string
		Price    float64
		Rating   float64
	}{{
		Category: "Electronics",
		Price:    500.0,
		Rating:   4.5,
	}, {
		Category: "Clothing",
		Price:    100.0,
		Rating:   4.8,
	}, {
		Category: "Electronics",
		Price:    600.0,
		Rating:   4.7,
	}, {
		Category: "Clothing",
		Price:    150.0,
		Rating:   4.5,
	}}

	// 按类别升序，然后按价格升序，最后按评分降序排序
	var result []struct {
		Category string
		Price    float64
		Rating   float64
	}

	Chain(products).OrderByFields("Category", "Price", "Rating").Value(&result)

	// 验证排序结果
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "Clothing", result[0].Category)    // 类别Clothing，价格100，评分4.8
	assert.Equal(t, "Clothing", result[1].Category)    // 类别Clothing，价格150，评分4.5
	assert.Equal(t, "Electronics", result[2].Category) // 类别Electronics，价格500，评分4.5
	assert.Equal(t, "Electronics", result[3].Category) // 类别Electronics，价格600，评分4.7
}
