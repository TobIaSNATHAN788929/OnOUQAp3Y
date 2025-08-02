// 代码生成时间: 2025-08-02 17:51:05
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// 定义重命名规则结构体
type RenameRule struct {
    prefix string
    suffix string
}

func main() {
    // 定义命令行参数
    dirPtr := flag.String("dir", ".", "The directory containing files to rename")
    prefixPtr := flag.String("prefix", "", "Prefix to add to each filename")
    suffixPtr := flag.String("suffix", "", "Suffix to add to each filename")
    flag.Parse()

    // 创建重命名规则实例
    rule := RenameRule{
        prefix: *prefixPtr,
        suffix: *suffixPtr,
    }

    // 获取目录中的文件列表
    files, err := os.ReadDir(*dirPtr)
    if err != nil {
        log.Fatalf("Failed to read directory: %v", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            oldName := file.Name()
            newName := fmt.Sprintf("%s%s%s", rule.prefix, oldName, rule.suffix)

            // 生成新旧文件的完整路径
            oldPath := filepath.Join(*dirPtr, oldName)
            newPath := filepath.Join(*dirPtr, newName)

            // 重命名文件
            if err := os.Rename(oldPath, newPath); err != nil {
                log.Printf("Failed to rename file %s to %s: %v", oldName, newName, err)
                continue
            }
            fmt.Printf("Renamed '%s' to '%s'
", oldName, newName)
        }
    }
}
