// 代码生成时间: 2025-08-02 03:06:06
package main

import (
    "fmt"
    "os/exec"
    "strings"
    "time"
    "github.com/kataras/iris/v12"
)

// PerformanceMonitor 结构体用于封装监控数据
type PerformanceMonitor struct {
    // 可以在这里添加更多的监控数据字段
}

// getSystemLoad 获取系统的负载信息
func getSystemLoad() (string, error) {
    // 使用uptime命令获取系统负载
    cmd := exec.Command("uptime")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(output)), nil
}

// getCPUUsage 获取CPU使用率
func getCPUUsage() (string, error) {
    // 使用top命令获取CPU使用率
    cmd := exec.Command("top", "-b", "-n", "1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    lines := strings.Split(strings.TrimSpace(string(output)), "\
")
    for _, line := range lines {
        if strings.HasPrefix(line, "Cpu(s)") {
            return line, nil
        }
    }
    return "", fmt.Errorf("CPU usage not found")
}

// getMemoryUsage 获取内存使用情况
func getMemoryUsage() (string, error) {
    // 使用free命令获取内存使用情况
    cmd := exec.Command("free\, "-m")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    lines := strings.Split(strings.TrimSpace(string(output)), "\
")
    for i, line := range lines {
        if i == 1 { // 跳过标题行
            return line, nil
        }
    }
    return "", fmt.Errorf("Memory usage not found")
}

// handleMonitorRequest 处理监控请求
func handleMonitorRequest(ctx iris.Context) {
    systemLoad, err := getSystemLoad()
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": "Failed to get system load",
        })
        return
    }
    cpuUsage, err := getCPUUsage()
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": "Failed to get CPU usage",
        })
        return
    }
    memoryUsage, err := getMemoryUsage()
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": "Failed to get memory usage",
        })
        return
    }
    
    ctx.JSON(iris.Map{
        "systemLoad": systemLoad,
        "cpuUsage": cpuUsage,
        "memoryUsage": memoryUsage,
    })
}

func main() {
    app := iris.New()
    app.Get("/monitor", handleMonitorRequest)
    
    // 设置监听端口
    app.Listen(":8080")
}
