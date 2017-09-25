// The jped command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output

package ch10

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
)

func ToJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
