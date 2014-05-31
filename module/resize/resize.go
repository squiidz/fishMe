package resize

import (
	"PushKids/module/utility"
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func ResizeMe(pic string) {
	picture := strings.SplitAfter(pic, ".")
	ext := picture[len(picture)-1]

	if ext == "jpg" {
		ResizeJpeg(pic)
	} else if ext == "png" {
		ResizePng(pic)
	} else if ext == "gif" {
		ResizeGif(pic)
	} else {
		//os.Remove(pic)
	}
}

func ResizeJpeg(pic string) {
	log.Println("RESIZING !! JPEG")
	// open the Picture
	file, err := os.Open(pic)
	utility.ShitAppend(err)

	img, err := jpeg.Decode(file)
	utility.ShitAppend(err)
	file.Close()

	// Resize the Picture Data
	m := resize.Resize(565, 363, img, resize.Lanczos3)

	out, err := os.Create(pic)
	utility.ShitAppend(err)
	defer out.Close()

	// OverWrite the None Resizing Picture
	jpeg.Encode(out, m, nil)
}

func ResizePng(pic string) {
	// open the Picture
	file, err := os.Open(pic)
	utility.ShitAppend(err)

	img, err := png.Decode(file)
	utility.ShitAppend(err)
	file.Close()

	// Resize the Picture Data
	m := resize.Resize(565, 363, img, resize.Lanczos3)

	out, err := os.Create(pic)
	utility.ShitAppend(err)
	defer out.Close()

	// OverWrite the None Resizing Picture
	png.Encode(out, m)
}

func ResizeGif(pic string) {
	// open the Picture
	file, err := os.Open(pic)
	utility.ShitAppend(err)

	img, err := gif.Decode(file)
	utility.ShitAppend(err)
	file.Close()

	// Resize the Picture Data
	m := resize.Resize(565, 363, img, resize.Lanczos3)

	out, err := os.Create(pic)
	utility.ShitAppend(err)
	defer out.Close()

	// OverWrite the None Resizing Picture
	gif.Encode(out, m, nil)
}
