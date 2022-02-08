package utils

import (
	"os"
	"runtime"
	"syscall"
)

var stdFile *os.File

func Rewrite(name string) error {
	if runtime.GOOS != "linux" {
		return nil
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	stdFile = file // 把文件句柄保存到全局变量，避免被 GC 回收

	err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		file.WriteString(err.Error())
		file.Close()
		return err
	}
	// 内存回收前关闭文件描述符
	runtime.SetFinalizer(stdFile, func(fd *os.File) {
		fd.Close()
	})

	return nil
}
