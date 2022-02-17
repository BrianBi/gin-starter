package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	test   Environment = &environment{value: "test"}
	pro    Environment = &environment{value: "pro"}
)

var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsTest() bool {
	return e.value == "test"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

// 小名方法, 限制当前包实现接口函数
func (e *environment) t() {}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n test:测试环境\n pro:正式环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "test":
		active = test
	case "pro":
		active = pro
	default:
		active = dev
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'dev' will be used.")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
