package img

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

// GenerateBlankImage generates a fixed-size rectangle and saves it to the buffer.
// If no parameters are passed, the rectangle will default to 250x250.
// You can choose between two colors: black and white. If you don't want to choose a color,
// pass in an empty string and the default color, black, will be selected.
func GenerateBlankImage(width, height int, clr string) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: choiceColor(clr)}, image.Point{}, draw.Src)

	var img image.Image = newImg
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return nil, err
	}

	return buffer, nil
}

// choiceColor gets a color in the signature. Supports 4 colors.
// Returns the color depending on the strictly specified string.
func choiceColor(clr string) color.Color {
	switch clr {
	case "white":
		return color.White
	case "blue":
		return color.RGBA{R: 79, G: 152, B: 202}
	case "green":
		return color.RGBA{R: 80, G: 216, B: 144}
	default:
		return color.Black
	}
}
