package image
import (
	"github.com/noverguo/util"
	"image/jpeg"
	"os"
	"github.com/nfnt/resize"
)

// 改变图像大小
func ResizeImage(inPath, outPath string, width, height uint) {
	file, err := os.Open(inPath)
	util.CheckErr(err)

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	util.CheckErr(err)

	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)
	out, err := os.Create(outPath)
	util.CheckErr(err)
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
