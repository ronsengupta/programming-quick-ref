package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func close(path string) {
	if err := syscall.Unmount(path, 0); err != nil {
		fmt.Errorf("Failed to unmount path (%v): %v", path, err)
	}

	if err := os.RemoveAll(path); err != nil {
		fmt.Errorf("Failed to delete path directory (%v): %v", path, err)
	}
}

func main() {
	fmt.Printf("--Mount test--\n")

	dev := "/tmp/foo/dev"

	if err := os.Mkdir(dev, 0777); err != nil {
		fmt.Errorf("Mkdir(%v) failed: %v", dev, err)
	}

	if err := syscall.Mount("", dev, "devtmpfs", syscall.MS_RDONLY, ""); err != nil {
		fmt.Printf("Couldn't mount /dev to %v: %v", dev, err)
	}

	defer close(dev)

	proc := "/tmp/foo/proc"

	if err := os.Mkdir(proc, 0777); err != nil {
		fmt.Errorf("Mkdir(%v) failed: %v", proc, err)
	}

	if err := syscall.Mount("", proc, "proc", syscall.MS_RDONLY, ""); err != nil {
		fmt.Errorf("Couldn't mount /proc to %v: %v", proc, err)
	}
	time.Sleep(5000 * time.Millisecond)

	defer close(proc)

}
