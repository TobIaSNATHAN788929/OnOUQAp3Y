// 代码生成时间: 2025-10-08 22:08:58
package main

import (
    "image"
    "image/color"
    "image/jpeg"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)

// WatermarkService provides functionality to apply watermarks to images.
type WatermarkService struct {
    // Properties for watermark service can be added here
}

// NewWatermarkService creates a new instance of WatermarkService.
func NewWatermarkService() *WatermarkService {
    return &WatermarkService{}
}

// ApplyWatermark applies a watermark text to an image file.
func (s *WatermarkService) ApplyWatermark(inputImagePath, outputImagePath, watermarkText string, opacity float64) error {
    // Read the input image file
    imgFile, err := os.Open(inputImagePath)
    if err != nil {
        return err
    }
    defer imgFile.Close()

    // Decode the image
    img, _, err := image.Decode(imgFile)
    if err != nil {
        return err
    }

    // Create a new image with the same size as the original image
    newImg := image.NewRGBA(img.Bounds())
    draw.Draw(newImg, newImg.Bounds(), img, image.Point{0, 0}, draw.Src)

    // Set the watermark text and calculate its position
    fontColor := color.RGBA{R: 255, G: 255, B: 255, A: uint8(opacity * 255)}
    fontBounds := draw.GlyphBounds(newImg, font.New(font.Times, 12.0), []rune(watermarkText))
    watermarkX := img.Bounds().Dx() - fontBounds.Dx() - 10 // Position watermark at the bottom right corner
    watermarkY := img.Bounds().Dy() - fontBounds.Dy() - 10

    // Draw the watermark text on the new image
    draw.Draw(newImg, newImg.Bounds(), newImg, image.Point{0, 0}, draw.Over)
    draw.DrawString(newImg, watermarkText, image.Point{watermarkX, watermarkY}, fontColor, draw.Src)

    // Save the watermarked image to the output path
    file, err := os.Create(outputImagePath)
    if err != nil {
        return err
    }
    defer file.Close()

    jpeg.Encode(file, newImg, nil)
    return nil
}

func main() {
    // Create a new watermark service
    wmService := NewWatermarkService()

    // Apply watermark to an image
    err := wmService.ApplyWatermark(
        "input.jpg",
        "output.jpg",
        "Watermark Text",
        0.5, // Opacity of the watermark
    )
    if err != nil {
        log.Fatalf("Error applying watermark: %v", err)
    }
    log.Println("Watermark applied successfully")
}