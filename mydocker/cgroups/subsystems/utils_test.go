package subsystems

import (
	"fmt"
	"testing"
)

func TestFindCgroupMountpoint(t *testing.T) {
	v := FindCgroupMountpoint("memory")
	fmt.Println(v)
}
