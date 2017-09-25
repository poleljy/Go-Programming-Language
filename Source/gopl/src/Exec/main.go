package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var cmdPath = "G:/Learn/GitHub/Go-Programming-Language/Source/gopl/bin/BuildModel"
var projDir = "Z:/sbobase/a2a113d3ca4baa09dcba1a5aefa52b1f/d19502f3f39c9c6b4eed3b9805ee1d71"

type HandlerFunc func(msg string)

type Executor struct {
	CmdPath        string
	Args           []string
	MessageHandler HandlerFunc

	cmd exec.Cmd
}

// NewExecutor returns an Executor
func NewExecutor(cmdPath string, args []string) Executor {
	exec := Executor{
		CmdPath: cmdPath,
		Args:    args,
	}
	return exec
}

func (e *Executor) LaunchCmd(projDir string, handler HandlerFunc) error {
	// 初始化日志文件
	infoFile, err := os.Create(filepath.Join(projDir, "log/rebuild.stdout.log"))
	defer infoFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	infoLog := log.New(infoFile, "", log.LstdFlags)

	errFile, err := os.Create(filepath.Join(projDir, "log/rebuild.stderr.log"))
	defer errFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	errLog := log.New(errFile, "", log.LstdFlags)

	infoLog.Println("INFO: launching command")

	// Set the commands arguments
	e.cmd.Path = cmdPath
	e.cmd.Args = append([]string{e.cmd.Path}, e.cmd.Args...)

	stdout, err := e.cmd.StdoutPipe()
	if err != nil {
		errLog.Println("ERROR:", err.Error())
		return err
	}

	stderr, err := e.cmd.StderrPipe()
	if err != nil {
		errLog.Println("ERROR:", err.Error())
		return err
	}

	// Start the process
	e.MessageHandler = handler
	if err := e.cmd.Start(); err != nil {
		err = fmt.Errorf("Failed to start command path=%q --- args=%q: %v", e.cmd.Path, e.cmd.Args, err)
		errLog.Println("ERROR:", err.Error())
		return err
	} else {
		go e.MessageHandler("[Started]\n")
	}

	// 循环读取输出流中的一行内容
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			infoLog.Printf("INFO:%s", line)

			go e.MessageHandler(line)
		}
	}()

	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			errLog.Printf("ERROR:%s", line)
		}
	}()

	if err := e.cmd.Wait(); err != nil {
		go e.MessageHandler("[Failed]\n")
		return err
	} else {
		go e.MessageHandler("[Finished]\n")
	}
	return nil
}

var (
	// finishedErr is the error message received when trying to kill and already
	// exited process.
	finishedErr = "os: process already finished"
)

// Shutdown sends an interrupt signal to the user process
func (e *Executor) ShutDown() error {
	if e.cmd.Process == nil {
		return fmt.Errorf("executor.shutdown error: no process found")
	}
	proc, err := os.FindProcess(e.cmd.Process.Pid)
	if err != nil {
		return fmt.Errorf("executor.shutdown failed to find process: %v", err)
	}
	if runtime.GOOS == "windows" {
		if err := proc.Kill(); err != nil && err.Error() != finishedErr {
			return err
		}
		return nil
	}
	if err = proc.Signal(os.Interrupt); err != nil && err.Error() != finishedErr {
		return fmt.Errorf("executor.shutdown error: %v", err)
	}
	return nil
}

func MessageHandler(msg string) {
	msg = msg[1 : len(msg)-2]
	if strings.HasSuffix(msg, `%`) {
		fmt.Printf("Progress:%s\n", msg)
	} else {
		fmt.Printf("Info:%s\n", msg)
	}
}

func main() {
	/*
		go func() {
			executor := NewExecutor(cmdPath, []string{})
			err := executor.LaunchCmd(projDir, MessageHandler)
			if err != nil {
				fmt.Printf("error in launching command: %v", err)
			}
		}()

		time.Sleep(5 * time.Second)
	*/

	executor := NewExecutor(cmdPath, []string{})
	err := executor.LaunchCmd(projDir, MessageHandler)
	if err != nil {
		fmt.Printf("error in launching command: %v", err)
	}
	time.Sleep(5 * time.Microsecond)
}
