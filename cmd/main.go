package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	path := strings.Replace(dir, "\\", "/", -1)
	fmt.Println(path)
	fmt.Println(path + "/NocalhostInstaller-v0.3.9.exe")
	cmd = exec.Command("cmd.exe", "/c", path+"/NocalhostInstaller-v0.3.9.exe")
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error info: %v", err)
	}
	fmt.Printf("output: %s", string(out))
}
