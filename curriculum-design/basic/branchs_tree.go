/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : branchs_tree.go
#   Last Modified : 2018-12-18 13:54
#   Describe      : 绘制简单的分型树
#
# ====================================================*/

package basic

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	// "math/rand"
	"os"
	// "time"
)

// DrawBranchsTree use go to draw a branchs trree
// 使用 go  绘制一个简单的分行树
func DrawBranchsTree() {
	const (
		dx float64 = 1000
		dy float64 = 800
	)
	// 需要保存的文件

	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, 1000, 1000))

	// drawline(5, 5, dx-8, dy-10, func(x, y int) {
	//     img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 200})
	// })
	// draw(dx/2, dy, dx/2, dy/2, 255, img)

	// 左右都画一条竖线
	// for i := 0; i < dy; i++ {
	//     img.Set(0, i, color.Black)
	//     img.Set(dx-1, i, color.Black)
	//
	// }
	ddraw(500, 1000, 180, 300, img)

	imgcounter := 250
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 以PNG格式保存文件
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

// PutPixel describes a function expected to draw a point on a bitmap at (x, y) coordinates
type PutPixel func(x, y int)

// 返回绝对值
func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// https://github.com/akavel/polyclip-go/blob/9b07bdd6e0a784f7e5d9321bff03425ab3a98beb/polyutil/draw.go
// TODO: handle int overflow etc.
// drawline 划线
func drawline(x0, y0, x1, y1 int, brush PutPixel) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	var sx, sy int
	sx, sy = 1, 1

	if x0 >= x1 {
		sx = -1
	}

	if y0 >= y1 {
		sy = -1
	}

	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

// func draw(x0, y0, x1, y1 int, deep uint8, img *image.NRGBA) {
//     length := getLength(x0, y0, x1, y1)
//
//     if length > 300 {
//
//         rand.Seed(time.Now().Unix())
//         leftAngle := float64(rand.Intn(20)+40) / 100.0
//         rightAngle := float64(rand.Intn(20)+40) / 100.0
//         leftLength := float64(rand.Intn(20)+60) / 100.0
//         rightLength := float64(rand.Intn(20)+60) / 100.0
//
//         // left
//         x1L := x1 - int(length*leftLength*math.Sqrt(1.0/(1.0+leftAngle*leftAngle)))
//         y1L := y1 - int(leftAngle*length*leftLength*math.Sqrt(1/(1+leftAngle*leftAngle)))
//         log.Printf("x1L:%d,y1L:%d\n", x1L, y1L)
//         deep++
//         drawline(x1, y1, x1L, y1L, func(x, y float64) {
//             img.Set(x, y, color.RGBA{255, 0, 0, deep})
//         })
//
//         // right
//         x1R := x1 + int(length*rightLength*math.Sqrt(1/(1+rightAngle*rightAngle)))
//         y1R := y1 - int(rightAngle*length*rightLength*math.Sqrt(1/(1+rightAngle*rightAngle)))
//         log.Printf("x1R:%d,y1R:%d\n", x1R, y1R)
//         drawline(x1, y1, x1R, y1R, func(x, y int) {
//             img.Set(x, y, color.RGBA{255, 0, 0, deep})
//         })
//     }
// }

// Get the length between x0,y0 x1,y1
//func getLength(x0, y0, x1, y1 int) float64 {
//	return math.Sqrt(float64(math.Pow(float64(x0-x1), 2) + math.Pow(float64(y0-y1), 2)))
//}

func ddraw(x, y, angle, length float64, img *image.NRGBA) {
	if length <= 4 {
		return
	}

	var angleModifier float64 = 35
	x2 := x + length*math.Cos(toRadians(angle+angleModifier))
	y2 := y - length*math.Sin(toRadians(angle+angleModifier))
	x3 := x + length*math.Cos(toRadians(angle-angleModifier))
	y3 := y - length*math.Sin(toRadians(angle-angleModifier))
	log.Println(x2, y2, x3, y3)
	drawline(int(x), int(y), int(x2), int(y2), func(x, y int) {
		img.Set(int(x), int(y), color.RGBA{255, 0, 0, 255})
	})
	drawline(int(x), int(y), int(x3), int(y3), func(x, y int) {
		img.Set(int(x), int(y), color.RGBA{255, 0, 0, 255})
	})

	ddraw(x2, y2, angle+angleModifier, length*0.6, img)
	ddraw(x3, y3, angle-angleModifier, length*0.6, img)
}

func toRadians(angle float64) float64 {
	return angle / 360 * math.Pi
}
