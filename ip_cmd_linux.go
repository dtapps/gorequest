package gorequest

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
)

func getCmdIPV4() string {
	// 执行 ifconfig | grep 'inet ' | awk '{print $2}' 命令
	cmd := exec.Command("bash", "-c", "ifconfig | grep 'inet ' | awk '{print $2}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv4 := scanner.Text()
		if IsIPV4(ipv4) && IsIPv4Public(net.ParseIP(ipv4)) {
			return ipv4
		}
	}

	return ""
}

func getCmdIPV6() string {
	// 执行 ip -6 addr | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}' 命令
	//cmd := exec.Command("bash", "-c", "ip -6 addr | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}'")
	//output, err := cmd.Output()
	// 执行 ip -6 addr | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} {print $3}' 命令
	cmd := exec.Command("bash", "-c", "ip -6 addr | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} {print $3}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv6 := scanner.Text()
		if IsIPV6(ipv6) && IsIPv6Public(net.ParseIP(ipv6)) {
			return ipv6
		}
	}

	return ""
}
