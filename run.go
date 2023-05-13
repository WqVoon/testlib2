package testlib2

import (
	"fmt"

	"github.com/wqvoon/testlib1"
)

func WrapperRegister(k, v string) {
	fmt.Println("testlib2@v0.0.2 WrapperRegister")
	testlib1.Register(k, v)
}

func WrapperGetAll() {
	fmt.Println("testlib2@v0.0.2 WrapperGetAll")
	testlib1.GetAll()
}
