package main

import "fmt"
import "syscall"

func main() {
    r1, _, err1 := syscall.RawSyscall6(syscall.SYS_CLONE, uintptr(syscall.SIGCHLD), 0, 0, 0, 0, 0)
    if err1 != 0 {
        //runtime_AfterFork()
        fmt.Println("Hello child error")
        fmt.Println(err1)
        return
    }
    if r1 != 0 {
        fmt.Println("Hello child")
        return
    }
  

    fmt.Println("Hello, 世界")
}
