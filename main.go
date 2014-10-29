package main

import "fmt"
import "syscall"
import "io/ioutil"

func main() {
	r1, _, err := syscall.RawSyscall6(syscall.SYS_CLONE, uintptr(syscall.SIGCHLD), 0, 0, 0, 0, 0)
	if err != 0 {
		fmt.Println(err)
	}

	if r1 != 0 {
		fmt.Println("Hello child")

		syscall.Chroot("/tmp")
		files, _ := ioutil.ReadDir("/")
		for _, f := range files {
			fmt.Println(f.Name())
		}
		return
	}

    	fmt.Println("Hello parent")
	files, _ := ioutil.ReadDir("/")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
