package main

import (
	"testing"
)

func TestValidJSON(t *testing.T) {
	user, err := ParseUser([]byte(`{"name":"Tom","age":20}`))
	if err != nil {
		t.Fatalf("出错了:%v", err)
	}
	if user.Name != "Tom" || user.Age != 20 {
		t.Errorf("解析错误:%v", user)
	}
}

// 测试 2：错误的 JSON（错误输入测试）
func TestInvalidJSON(t *testing.T) {
	_, err := ParseUser([]byte(`{"name":"Tom","age":"20"}`)) // age 故意写成字符串
	if err == nil {
		t.Errorf("期望报错，但居然解析成功了！")
	}
}
