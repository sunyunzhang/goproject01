package lib

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func GetOs() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	// Wait功能将等待，直到进程结束
	proc.Wait()

	// 在进程终止后，*os.ProcessState 包含有关进程运行的简单信息
	fmt.Printf("PID: %dn", proc.ProcessState.Pid())
	fmt.Printf("进程耗费时间: %dmsn", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("成功退出: %tn", proc.ProcessState.Success())
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

// func Getlll() {
// 	fmt.Println("PID")
// }
