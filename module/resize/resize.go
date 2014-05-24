package resize

import (
	"PushKids/module/utility"
	"github.com/nfnt/resize"
	"image/jpeg"
	"os"
)

func ResizeMe(pic string) {
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
