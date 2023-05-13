package testlib2

import "testing"

func TestWrapperGetAll(t *testing.T) {
	WrapperRegister("name", "hygao")
	WrapperGetAll()
}
