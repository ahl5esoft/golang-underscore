package underscore

import (
	"testing"
)

func Test_Md5(t *testing.T) {
	if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("wrong")
	}
}

func Test_UUID(t *testing.T) {
	uuid := UUID()
	if len(uuid) != 32 {
		t.Error("wrong")
	}

	t.Log(uuid)
}
