package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	str := "curl --fail -L \"https://codingcorp-generic.pkg.coding.net/nocalhost/nhctl/NocalhostInstaller.exe?version=v0.3.9\" -o NocalhostInstaller-v0.3.9.exe"
	cmd := exec.Command("bash", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("curl error info: %v", err)
	}
	fmt.Printf("curl output: %s", string(out))
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	fmt.Println(dir)
	fmt.Println(dir + "\\NocalhostInstaller-v0.3.9.exe")
	cmd = exec.Command("cmd.exe", "/c", dir+"\\NocalhostInstaller-v0.3.9.exe")
	stdout, stderr, err := RunWithRollingOut(cmd)
	if err != nil {
		fmt.Printf("error info: %v", err)
	}
	fmt.Printf("stdout: %s, stderr: %s", stdout, stderr)
}

func RunWithRollingOut(cmd *exec.Cmd) (string, string, error) {
	fmt.Printf("Running command: %s\n", cmd.Args)
	var stdoutLog strings.Builder
	var stderrLog strings.Builder
	stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return stdoutLog.String(), stderrLog.String(), err
	}
	if err = cmd.Start(); err != nil {
		_ = cmd.Process.Kill()
		return "", "", err
	}
	defer stdoutPipe.Close()
	defer stderrPipe.Close()
	stdoutReader := bufio.NewReaderSize(stdoutPipe, 1024)
	stderrReader := bufio.NewReaderSize(stderrPipe, 1024)
	go func() {
		for {
			line, isPrefix, err := stdoutReader.ReadLine()
			if err != nil {
				if err != io.EOF && !strings.Contains(err.Error(), "closed") {
					os.Stdout.WriteString(fmt.Sprintf("command log error: %v, log: %v\n", err, string(line)))
				}
				break
			}
			if len(line) != 0 && !isPrefix {
				os.Stdout.WriteString(string(line) + "\n")
				stdoutLog.WriteString(string(line) + "\n")
			}
		}
	}()
	go func() {
		for {
			line, isPrefix, err := stderrReader.ReadLine()
			if err != nil {
				if err != io.EOF && !strings.Contains(err.Error(), "closed") {
					os.Stderr.WriteString(fmt.Sprintf("command log error: %v, log: %v\n", err, string(line)))
				}
				break
			}
			if len(line) != 0 && !isPrefix {
				os.Stderr.WriteString(string(line) + "\n")
				stderrLog.WriteString(string(line) + "\n")
			}
		}
	}()
	_ = cmd.Wait()
	time.Sleep(2 * time.Second)
	if cmd.ProcessState.Success() && stderrLog.Len() == 0 {
		return stdoutLog.String(), "", nil
	} else {
		return stdoutLog.String(), stderrLog.String(), errors.New("exit code is not 0")
	}
}
