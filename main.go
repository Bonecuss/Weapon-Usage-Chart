package main

import (
	"fmt"
	"image"
	"image/color"

	// "image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/font"

	"golang.org/x/image/draw"

	"io/ioutil"

	"github.com/goki/freetype"
	"github.com/goki/freetype/truetype"
)

func loadimgs() []image.Image {
	var wepimgs []image.Image

	file1, _ := os.Open("sns.png")
	defer file1.Close()
	img1, _, _ := image.Decode(file1)
	wepimgs = append(wepimgs, img1)

	file2, _ := os.Open("ds.png")
	defer file2.Close()
	img2, _, _ := image.Decode(file2)
	wepimgs = append(wepimgs, img2)

	file3, _ := os.Open("gs.png")
	defer file3.Close()
	img3, _, _ := image.Decode(file3)
	wepimgs = append(wepimgs, img3)

	file4, _ := os.Open("ls.png")
	defer file4.Close()
	img4, _, _ := image.Decode(file4)
	wepimgs = append(wepimgs, img4)

	file5, _ := os.Open("h.png")
	defer file5.Close()
	img5, _, _ := image.Decode(file5)
	wepimgs = append(wepimgs, img5)

	file6, _ := os.Open("hh.png")
	defer file6.Close()
	img6, _, _ := image.Decode(file6)
	wepimgs = append(wepimgs, img6)

	file7, _ := os.Open("l.png")
	defer file7.Close()
	img7, _, _ := image.Decode(file7)
	wepimgs = append(wepimgs, img7)

	file8, _ := os.Open("gl.png")
	defer file8.Close()
	img8, _, _ := image.Decode(file8)
	wepimgs = append(wepimgs, img8)

	file9, _ := os.Open("t.png")
	defer file9.Close()
	img9, _, _ := image.Decode(file9)
	wepimgs = append(wepimgs, img9)

	file10, _ := os.Open("s.png")
	defer file3.Close()
	img10, _, _ := image.Decode(file10)
	wepimgs = append(wepimgs, img10)

	file11, _ := os.Open("m.png")
	defer file11.Close()
	img11, _, _ := image.Decode(file11)
	wepimgs = append(wepimgs, img11)

	file12, _ := os.Open("lbg.png")
	defer file12.Close()
	img12, _, _ := image.Decode(file12)
	wepimgs = append(wepimgs, img12)

	file13, _ := os.Open("hbg.png")
	defer file13.Close()
	img13, _, _ := image.Decode(file13)
	wepimgs = append(wepimgs, img13)

	file14, _ := os.Open("b.png")
	defer file14.Close()
	img14, _, _ := image.Decode(file14)
	wepimgs = append(wepimgs, img14)

	return wepimgs
}

func fontInit() {
	fontFile := "./font/DejaVuSans.ttf"
	// font writing stuff
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	ttf, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	fnt = *ttf
}

var fnt truetype.Font

func drawLabel(dst *image.RGBA, text string, xPos int, yPos int, fontSize float64) *image.RGBA {
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(&fnt)
	c.SetFontSize(fontSize)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetHinting(font.HintingNone)

	drawSize := int(c.PointToFixed(10) >> 6)

	d := &font.Drawer{
		Dst: dst,
		Src: image.White,
		Face: truetype.NewFace(&fnt, &truetype.Options{
			Size:    10,
			DPI:     72,
			Hinting: font.HintingFull,
		}),
	}

	// c.SetHinting(font.HintingFull)
	for _, text := range strings.Split(text, "\n") {
		//fmt.Println(text)
		//fmt.Println(d.MeasureString(text).Ceil())
		c.SetSrc(image.White)

		adj := (100 - int(d.MeasureString(text).Floor())) / 2
		_, _ = c.DrawString(text, freetype.Pt((xPos+adj), yPos+drawSize))
		drawSize += int(c.PointToFixed(10) >> 6)
	}
	return dst
}

func main() {
	fontInit()

	intslice := [][]int{{50, 75, 175, 200}, {0, 59, 150, 25}, {100, 20, 50, 50}, {100, 150, 200, 560}, {15, 1, 0, 5}, {100, 25, 200, 5}, {50, 5, 5, 55},
		{20, 40, 50, 60}, {12, 17, 14, 15}, {100, 40, 60, 50}, {100}, {2, 2, 2, 2}, {50, 50, 50, 50}, {20, 25, 25, 5}}

	biggest := 0
	for i := 0; i < len(intslice); i++ {
		total := 0
		for j := 0; j < len(intslice[i]); j++ {
			total += intslice[i][j]
		}
		if total > biggest {
			biggest = total
		}
	}

	file, _ := os.Open("base.png")
	defer file.Close()
	img, _, _ := image.Decode(file)

	m := image.NewRGBA(img.Bounds())
	draw.Draw(m, m.Bounds(), img, image.Point{0, 0}, draw.Src)
	b := m.Bounds()

	wepimgs := loadimgs()

	numweapons := 14

	barwidth := 25
	barspacing := 46
	totalwidth := 14 * barspacing
	maxbarheight := 480
	offsetx := (b.Max.X - totalwidth + barwidth) / 2

	wepbg := color.RGBA{117, 104, 70, 255}
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, image.Rect(0, maxbarheight+6, b.Size().X, 540), &image.Uniform{wepbg}, image.Point{0, 0}, draw.Src)
	draw.Draw(m, image.Rect(0, maxbarheight+2, b.Size().X, maxbarheight+4), &image.Uniform{white}, image.Point{0, 0}, draw.Src)

	breakpoints10 := []int{10, 100, 1000, 10000, 100000, 1000000}
	breakpoints5 := []int{5, 50, 500, 5000, 50000, 500000, 5000000}
	redline := color.RGBA{255, 0, 0, 255}
	for i := len(breakpoints10) - 1; i > 0; i-- {
		fmt.Println(biggest % breakpoints10[i])
		if breakpoints10[i] <= biggest && breakpoints10[i]%biggest != 0 {
			lineoffsety := int((1 - (float64(breakpoints10[i])/float64(biggest))*0.9) * float64(maxbarheight))
			draw.Draw(m, image.Rect(38, lineoffsety, b.Size().X-38, lineoffsety+2), &image.Uniform{redline}, image.Point{0, 0}, draw.Src)
			drawLabel(m, strconv.Itoa(breakpoints10[i]), -35, lineoffsety-5, 13)
			i = 0
			break
		}
	}
	for i := len(breakpoints5) - 1; i > 0; i-- {
		if breakpoints5[i] <= biggest && breakpoints5[i]%biggest != 0 {
			lineoffsety := int((1 - (float64(breakpoints5[i])/float64(biggest))*0.9) * float64(maxbarheight))
			draw.Draw(m, image.Rect(38, lineoffsety, b.Size().X-38, lineoffsety+2), &image.Uniform{redline}, image.Point{0, 0}, draw.Src)
			drawLabel(m, strconv.Itoa(breakpoints5[i]), -35, lineoffsety-5, 13)
			i = 0
			break
		}
	}

	barcolors := [][]int{{255, 80, 91, 255}, {155, 234, 241, 255}, {223, 214, 93, 255},
		{106, 192, 131, 255}, {195, 163, 210, 255}, {213, 167, 134, 255}, {158, 189, 255, 255},
		{231, 203, 215, 255}, {155, 234, 241, 255}, {170, 170, 170, 255}, {253, 87, 3, 255},
		{172, 213, 107, 255}, {249, 136, 156, 255}, {101, 230, 199, 255}}

	for i := 0; i < numweapons; i++ {
		offsety := maxbarheight
		prev := maxbarheight
		total := 0
		for j := 0; j < len(intslice[i]); j++ {
			total += intslice[i][j]
		}
		r := barcolors[i][0]
		g := barcolors[i][1]
		b := barcolors[i][2]
		for j := 0; j < len(intslice[i])-1; j++ {
			if r >= 40 {
				r -= 40
			}
			if g >= 40 {
				g -= 40
			}
			if b > 40 {
				b -= 40
			}
		}
		barcolor := color.RGBA{uint8(r), uint8(g), uint8(b), 255}

		for j := 0; j < len(intslice[i]); j++ {
			if j < len(intslice[i]) {
				subspercent := float64(intslice[i][j]) / float64(total)
				subslength := float64(maxbarheight-5) * subspercent
				screenpercent := float64(total) / float64(biggest)
				offsety = int(float64(prev) - subslength*screenpercent*0.9)
				if offsety == maxbarheight && intslice[i][j] > 0 {
					offsety = maxbarheight - 1
				}
			} else {
				offsety = maxbarheight
			}

			draw.Draw(m, image.Rect(offsetx+(i*barspacing), offsety, (i*barspacing)+(offsetx+barwidth), prev), &image.Uniform{barcolor}, image.Point{0, 0}, draw.Src)
			if prev-offsety > 14 && len(intslice[i]) > 1 {
				//drawLabel(m, strconv.Itoa(int(float32(intslice[i][j])/float32(total)*float32(100)))+"%", offsetx+(i*barspacing)-38, (offsety+prev)/2-5, 10)
			}
			prev = offsety

			if r < 215 {
				r += 40
			}
			if g < 215 {
				g += 40
			}
			if b < 215 {
				b += 40
			}
			barcolor = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
		//color := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

		//draw.DrawMask(w, w.Bounds(), c, image.Point{0, 0}, draw.Over)

		draw.Draw(m, image.Rect(offsetx+(i*barspacing)-(48-barwidth)/2, 490, (i*barspacing)+(offsetx+48-(48-barwidth)/2), 540), wepimgs[i], image.Pt(0, 0), draw.Over)

		labeltext := strconv.Itoa(total)
		drawLabel(m, labeltext, offsetx+(i*barspacing)-40-2*(len(labeltext)-1), offsety-20, 16)

	}

	f, _ := os.Create("image.png")
	png.Encode(f, m)
}
