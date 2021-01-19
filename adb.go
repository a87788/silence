package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var (
		apks  string
		ip    string
		start int
		end   int
	)
	idr := "D:\\adb\\adb"
	for i := start; i <= end; i++ {
		kill := strings.Join([]string{idr, "kill-server"}, " ")
		conn := strings.Join([]string{idr, "connect", ip + ":", strconv.Itoa(i)}, " ")
		install := strings.Join([]string{idr, "install -r", apks}, " ")
		reboot := strings.Join([]string{idr, "reboot"}, " ")
		cmd1 := exec.Command("cmd", "/c", kill)
		cmd2 := exec.Command("cmd", "/c", conn)
		cmd3 := exec.Command("cmd", "/c", install)
		cmd4 := exec.Command("cmd", "/c", reboot)
		cmd1.Run()
		cmd2.Run()
		cmd3.Run()
		cmd4.Start()
		cmd2.Run()
	}

}
