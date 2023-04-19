package main

import (
	"fmt"
	"mydocker/cgroups/subsystems"
	"os"
	"path"
)

func main() {
	memSubSys := subsystems.MemorySubSystem{}
	resConfig := subsystems.ResourceConfig{
		MemoryLimit: "100m",
	}

	testCgroup := "testmemlimit"

	if err := memSubSys.Set(testCgroup, &resConfig); err != nil {
		fmt.Printf("cgroup fail %v", err)
	}

	stat, _ := os.Stat(path.Join(subsystems.FindCgroupMountpoint("memory"), testCgroup))
	fmt.Printf("cgroup stats:%+v", stat)

	if err := memSubSys.Apply(testCgroup, os.Getpid()); err != nil {
		fmt.Printf("cgroup Apply %v", err)
	}

}
