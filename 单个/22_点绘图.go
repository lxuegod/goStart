package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	//新建一个NRGBA颜色体系的图片
	//图片大小为 100*100  即宽度和高度都是100像素
	imageT := image.NewNRGBA(image.Rect(0, 0, 100, 100))

	//设置画点所需的色彩，这里设置的是纯红色
	colorT := color.NRGBA{0xFF, 0x00, 0x00, 0xFF}

	//用两重循环逐行进行绘制，i代表纵坐标，j代表横坐标
	for i := 0; i < 100; i = i + 2 {
		for j := 0; j < 100; j = j + 2 {
			imageT.Set(int(j), int(i), colorT)
		}
	}

	//将图片保存为png格式的图片文件
	//保存到桌面的文件夹
	fileT, errT := os.Create(`E:\桌面\dekstop\新建文件夹\test.png`)

	if errT != nil {
		fmt.Printf("创建图像文件时发生错误：%v\n", errT.Error())
	}

	defer fileT.Close()

	png.Encode(fileT, imageT)
}
