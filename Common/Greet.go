package Common

import (
	"fmt"

	"github.com/feiya1314/Auth"
)

func Greet(name string) string {
	fmt.Println("common module:" + Auth.Check(name))
	return "hello " + name
}
