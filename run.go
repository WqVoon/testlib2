package testlib2

import (
	"fmt"

	"github.com/wqvoon/testlib1"
)

func WrapperRegister(k, v string) {
	fmt.Println("testlib2@v0.0.1 WrapperRegister")
	testlib1.Register(k, v)
}

func WrapperGetAll() {
	fmt.Println("testlib2@v0.0.1 WrapperGetAll")
	testlib1.GetAll()
}
