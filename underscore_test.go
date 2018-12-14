package underscore

// TestModel is 测试模型
type TestModel struct {
	TestNestedModel

	ID   int
	Name string
}

// TestNestedModel is 嵌套模型
type TestNestedModel struct {
	Age int
}
