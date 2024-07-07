package renderer

import (
	"algvisual/internal/infra"
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "../../testdata/luxisr.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

func NewTextDrawer(
	c *infra.AppConfig,
) *TextDrawer {
	return &TextDrawer{c: c}
}

type TextDrawer struct {
	c    *infra.AppConfig
	font *truetype.Font
	face font.Face
}

func (t *TextDrawer) LoadFonts() error {
	fontBytes, err := os.ReadFile(fmt.Sprintf("%s/Roboto/Roboto-Black.ttf", t.c.FontsFolderPath))
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
	t.font = f
	return nil
}

func (t *TextDrawer) addLabel(img *image.RGBA, x, y int, text string) error {
	fontSize := float64(35)
	opts := truetype.Options{}
	opts.Size = fontSize
	face := truetype.NewFace(t.font, &opts)
	c := freetype.NewContext()
	fg := image.NewUniform(color.RGBA{R: 0x00, G: 0x33, B: 0x66, A: 0xff})
	c.SetFontSize(fontSize)
	c.SetFont(t.font)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)
	// textWidth := font.MeasureString(face, text).Ceil()
	textHeight := face.Metrics().Ascent.Ceil() + face.Metrics().Descent.Ceil()
	// pt := freetype.Pt(x, y)
	temp := strings.Split(text, "\r")
	// totalRows := len(temp)
	numRows := 0
	IMAGE_WIDTH := img.Rect.Dx()
	// IMAGE_HEIGHT := img.Rect.Dy()
	for _, textInLine := range temp {
		// totalRows := 1
		lineWidth := 0
		currentText := ""
		var linesTextx []string
		splitStrings := strings.Split(textInLine, " ")
		for _, splitstr := range splitStrings {
			strWidth := font.MeasureString(face, splitstr).Ceil()
			if lineWidth+strWidth+x+50 < IMAGE_WIDTH {
				// stay on existing row
				lineWidth += strWidth
				currentText = fmt.Sprintf("%s %s", currentText, splitstr)
			} else {
				// move to new row
				lineWidth = 0
				linesTextx = append(linesTextx, currentText)
				currentText = ""
				currentText = fmt.Sprintf("%s %s", currentText, splitstr)
				lineWidth += strWidth
			}
		}
		if currentText != "" {
			linesTextx = append(linesTextx, currentText)
		}
		for _, splitstr := range linesTextx {
			if splitstr == "" {
				continue
			}
			pt := freetype.Pt(x, y+(numRows*textHeight))
			_, err := c.DrawString(splitstr, pt)
			if err != nil {
				return err
			}
			numRows += 1
		}
	}
	return nil
}

func (t *TextDrawer) WriteMultilineText(texts []string, face font.Face) {
}
