// 代码生成时间: 2025-09-04 19:28:11
// batch_image_resizer.go
// 一个使用GOLANG和IRIS框架的批量图片尺寸调整器
package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"

    "github.com/kataras/iris/v12"
    "github.com/nfnt/resize"
)

// ImageResizingConfig 定义了图片尺寸调整的配置
type ImageResizingConfig struct {
    Width, Height int
}

// resizeImage 调整图片尺寸
func resizeImage(config ImageResizingConfig, src, dest string) error {
    srcImageFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcImageFile.Close()

    srcImage, _, err := image.Decode(srcImageFile)
    if err != nil {
        return err
    }

    resizedImage := resize.Resize(uint(config.Width), uint(config.Height), srcImage, resize.Lanczos3)

    destImageFile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destImageFile.Close()

    if err := jpeg.Encode(destImageFile, resizedImage, nil); err != nil {
        return err
    }

    return nil
}

// processDirectory 处理目录中的所有图片
func processDirectory(config ImageResizingConfig, dir string) error {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return err
    }
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        if !strings.HasSuffix(file.Name(), ".jpg") && !strings.HasSuffix(file.Name(), ".jpeg") {
            continue
        }

        sourcePath := filepath.Join(dir, file.Name())
        destPath := filepath.Join(dir, fmt.Sprintf("resized_%s", file.Name()))

        if err := resizeImage(config, sourcePath, destPath); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    app := iris.New()
    config := ImageResizingConfig{Width: 800, Height: 600} // 默认配置

    // 设置路由处理图片目录的处理
    app.Post("/resize", func(ctx iris.Context) {
        dir := ctx.FormValue("dir")
        if dir == "" {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.Writef("Directory path is required")
            return
        }

        if err := processDirectory(config, dir); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.Writef("Error resizing images: %v", err)
            return
        }

        ctx.StatusCode(iris.StatusOK)
        ctx.Writef("All images in %s have been resized successfully", dir)
    })

    // 启动服务器
    app.Listen(":8080")
}
