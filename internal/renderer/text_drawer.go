package renderer

import (
	"algvisual/internal/entities"
	"algvisual/internal/infra/config"
	"flag"
	"fmt"
	"image"
	"image/color"
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
	c *config.AppConfig,
) *TextDrawer {
	return &TextDrawer{c: c}
}

type TextDrawer struct {
	c    *config.AppConfig
	font *truetype.Font
	face font.Face
}

func (t *TextDrawer) LoadFonts() error {
	//fontBytes, err := os.ReadFile(fmt.Sprintf("%s/Roboto/Roboto-Black.ttf", t.c.FontsFolderPath))
	//f, err := truetype.Parse(fontBytes)
	//if err != nil {
	//	return err
	//}
	//t.font = f
	return nil
}

func (t *TextDrawer) addLabel(img *image.RGBA, x, y int, text string, size int32) error {
	fontSize := float64(size)
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
			var textWidth int
			if currentText == "" {
				textWidth = font.MeasureString(face, splitstr).Ceil()
			} else {
				textWidth = font.MeasureString(face, fmt.Sprintf("%s %s", currentText, splitstr)).Ceil()
			}
			if lineWidth+textWidth+x+50 < IMAGE_WIDTH {
				// stay on existing row
				if currentText == "" {
					currentText = splitstr
				} else {
					currentText = fmt.Sprintf("%s %s", currentText, splitstr)
				}
				lineWidth += textWidth
			} else {
				// move to new row
				lineWidth = 0
				linesTextx = append(linesTextx, currentText)
				currentText = ""
				currentText = splitstr
				lineWidth += textWidth
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

func (t *TextDrawer) FindTextSizeToFillContainer(text string, container entities.Container) int {
	size := 1
	for {
		w, h, l := t.measureText(float64(size), text, container)
		fmt.Println(l, w)
		if h > int(container.Height()) {
			return size
		}
		if w > int(container.Width()) {
			return size
		}
		size += 1
	}
}

func (t *TextDrawer) measureText(
	size float64,
	text string,
	container entities.Container,
) (int, int, int) {
	fontSize := float64(size)
	opts := truetype.Options{}
	opts.Size = fontSize
	face := truetype.NewFace(t.font, &opts)
	lines := strings.Split(text, "\r")
	totalLines := len(lines)
	maxWidth := 0
	for _, line := range lines {
		lineWidth := 0
		currentText := ""
		splitStrings := strings.Split(line, " ")
		for _, splitstr := range splitStrings {
			textWidth := font.MeasureString(face, splitstr).Ceil()
			lineWidth = textWidth + font.MeasureString(face, currentText).Ceil()
			if lineWidth < int(container.Width()) {
				if currentText == "" {
					currentText = splitstr
				} else {
					currentText = fmt.Sprintf("%s %s", currentText, splitstr)
				}
			} else {
				if maxWidth < lineWidth {
					maxWidth = lineWidth
				}
				lineWidth = 0
				totalLines += 1
				currentText = splitstr
			}
		}
	}
	height := totalLines * (face.Metrics().Ascent.Ceil() + face.Metrics().Descent.Ceil())
	return maxWidth, height, totalLines
}
