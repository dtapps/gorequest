package gorequest

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
)

func getCmdIPV4() string {
	// 执行 ipconfig | findstr IPv4 命令
	cmd := exec.Command("cmd", "/c", "ipconfig", "|", "findstr", "IPv4")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			ipv4 := fields[len(fields)-1]
			if IsIPV4(ipv4) && IsIPv4Public(net.ParseIP(ipv4)) {
				return ipv4
			}
		}
	}

	return ""
}

func getCmdIPV6() string {
	// 执行 ipconfig | findstr IPv6 命令
	cmd := exec.Command("cmd", "/c", "ipconfig", "|", "findstr", "IPv6")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			ipv6 := fields[len(fields)-1]
			if IsIPV6(ipv6) && IsIPv6Public(net.ParseIP(ipv6)) {
				return ipv6
			}
		}
	}

	return ""
}
