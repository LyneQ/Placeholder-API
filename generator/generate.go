package generator

import (
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Generate(c *gin.Context) {

	// all image parameters
	size := c.Param("size") // Ex: "300x200"
	label := c.Request.URL.Query().Get("label")
	fontSize := c.Request.URL.Query().Get("font_size")
	fontStyle := c.Request.URL.Query().Get("font_style")
	fontColor := c.Request.URL.Query().Get("font_color")
	backgroundColor := c.Request.URL.Query().Get("background_color")

	parts := strings.Split(size, "x")
	if len(parts) != 2 {
		c.String(http.StatusBadRequest, "Invalid size format. Use /placeholder/300x200")
		return
	}

	width, err1 := strconv.Atoi(parts[0])
	height, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || width <= 0 || height <= 0 || width > 9000 || height > 9000 {
		c.String(http.StatusBadRequest, "Invalid dimensions.")
		return
	}

	dc := gg.NewContext(width, height)
	if backgroundColor == "" || len(backgroundColor) > 8 || len(backgroundColor) < 6 {
		backgroundColor = "#FFFFFF"
	}
	dc.SetHexColor(backgroundColor) // background
	dc.Clear()

	//use custom fonts
	var fontStylePath = "generator/fonts/" + fontStyle + ".ttf"
	if _, err := os.Stat(fontStylePath); err != nil {
		//panic(err)
		fontStylePath = "generator/fonts/Inter.ttf"
	}
	parsedFontSize, err := strconv.ParseFloat(fontSize, 64)
	err = dc.LoadFontFace(fontStylePath, parsedFontSize)
	if err != nil {
		panic(err)
	}

	// Write text on middle of image
	if fontColor == "" || len(fontColor) > 8 || len(fontColor) < 6 {
		fontColor = "#000000"
	}
	dc.SetHexColor(fontColor)
	if label == "" {
		label = strconv.Itoa(width) + "x" + strconv.Itoa(height)
	}
	dc.DrawStringAnchored(label, float64(width)/2, float64(height)/2, 0.5, 0.5)

	c.Header("Content-Type", "image/png")
	c.Status(http.StatusOK)
	dc.EncodePNG(c.Writer)
}
