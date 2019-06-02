package underscore

var (
	benchmarkSize = 1000000
)

type testModel struct {
	ID   int
	Name string
}

type testNestedModel struct {
	testModel

	Age int
}
