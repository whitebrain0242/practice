package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type User struct { //i首字母必须大写
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 解析json的函数
func ParseUser(data []byte) (*User, error) {
	var cfg User
	err := json.Unmarshal(data, &cfg) //  第一个是字节数据,第二各是结构体指针,把json解压到结构体中
	if err != nil {
		fmt.Println("解析失败:", err)
		return nil, err
	}
	return &cfg, nil
}
func GetPort() (string, error) {
	//读取环境变量Port,得到字符串
	portEnv := os.Getenv("PORT")
	//如果没有使用默认端口
	if portEnv == "" {
		portEnv = "50051"
	}
	//字符串u转z数字
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		fmt.Printf("端口非法:%s\n", portEnv)
		return "", fmt.Errorf("端口非法:%s", portEnv)
	}
	return fmt.Sprintf("0.0.0.0:%d", port), nil //拼接监听地址
}
func main() {
	jsonStr := `{"name":"XiaoBai","age":19}`
	data := []byte(jsonStr) //因为unmarshal直接受字节切片
	cfg, err := ParseUser(data)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	fmt.Printf("解析结果:%+v\n", cfg)
	fmt.Printf("名字:%s\n", cfg.Name)
	fmt.Printf("年龄:%d\n", cfg.Age)
	addr, err := GetPort()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("n监听地址: ", addr)
	//http.ListenAndServe(addr, nil)

}
