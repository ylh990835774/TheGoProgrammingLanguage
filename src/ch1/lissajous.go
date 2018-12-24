// Lissajous generates GIF animations of random Lissajous figures.
// Lissajous生成随机Lissajous图形的GIF动画
// go run lissajous.go > out1.gif
package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// 外层for循序64次, 每次都会生成一个单独的动画帧
	// 生成了一个包含两种颜色的201*201大小的图片, 白色和黑色
	// 所有像素点都会被默认设置为其零值(也就是调色板palette里的第0个值), 这里设置的是白色
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 设置了两个偏振值
			// x轴偏振使用sin函数
			// y轴偏振也是正弦波, 但相对x轴是一个0-3的随机值
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// 渲染黑色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 设置80ms延时
		anim.Image = append(anim.Image, img)
	}

	// 结果写入到输出流
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
