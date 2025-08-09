// 代码生成时间: 2025-08-10 05:35:05
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/kataras/iris/v12"
)

// TestDataGenerator 是测试数据生成器的主要结构体
type TestDataGenerator struct {
    // 可以添加更多属性以支持更复杂的测试数据生成逻辑
}

// GenerateUser 创建一个随机用户对象
func GenerateUser() User {
    return User{
        ID:        rand.Intn(10000),
        FirstName: randomString(10),
        LastName:  randomString(10),
        Email:     fmt.Sprintf("%s.%s@example.com", randomString(8), randomString(8)),
    }
}

// randomString 生成一个指定长度的随机字符串
func randomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyz"
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(result)
}

// User 定义了一个用户对象
type User struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
}

func main() {
    // 初始化 Iris
    app := iris.New()

    // 创建测试数据生成器实例
    testDataGenerator := TestDataGenerator{}

    // 设置路由，返回随机用户数据
    app.Get("/user", func(ctx iris.Context) {
        user := GenerateUser()
        if err := ctx.JSON(user); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Internal Server Error")
            return
        }
    })

    // 设置端口并运行 Iris 应用程序
    app.Listen(":8080")
}
