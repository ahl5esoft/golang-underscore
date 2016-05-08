package underscore

import (
	"testing"
)

func TestParseJson(t *testing.T) {
	str := `["a","b"]`
	var arr []string
	err := ParseJson(str, &arr)
	if !(err == nil && len(arr) == 2) {
		t.Error("wrong")
	}
}

func TestToJson(t *testing.T) {
	b := true
	v, _ := ToJson(b)
	if v != "true" {
		t.Error("bool fail")
		return
	}

	str := "a"
	v, _ = ToJson(str)
	if v != str {
		t.Error("string fail")
		return
	}

	v, _ = ToJson(1)
	if v != "1" {
		t.Error("int fail")
		return
	}

	arr := []int{ 1, 2, 3 }
	v, _ = ToJson(arr)
	if v != "[1,2,3]" {
		t.Error("array fail")
		return
	}

	obj := TestModel{ 1, "name" }
	v, _ = ToJson(obj)
	if v != `{"Id":1,"Name":"name"}` {
		t.Error("obj fail")
		return
	}
}

func TestMd5(t *testing.T) {
	if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("wrong")
	}	
}

func TestUUID(t *testing.T) {
	uuid := UUID()
	if len(uuid) != 32 {
		t.Error("wrong")
	}

	t.Log(uuid)
}