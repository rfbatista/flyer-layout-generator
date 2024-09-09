package main

import (
	"algvisual/internal/infrastructure/config"
	"fmt"
	"os"
	"path"

	"github.com/oov/psd"
)

func processLayer(filename string, layerName string, l *psd.Layer) error {
	if len(l.Layer) > 0 {
		for i, ll := range l.Layer {
			if err := processLayer(
				fmt.Sprintf("%s_%03d %s", filename, i, ll.SectionDividerSetting.Type),
				layerName+"/"+ll.Name, &ll); err != nil {
				return err
			}
		}
	}
	if !l.HasImage() {
		return nil
	}
	fmt.Printf("%s -> %s.png\n", layerName, filename)
	return nil
}

func main() {
	root := config.FindProjectRoot()
	filePath := path.Join(root, "./dist/files/teste")
	fmt.Printf("loading file from %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := psd.Decode(file, &psd.DecodeOptions{SkipMergedImage: true})
	if err != nil {
		panic(err)
	}
	for i, layer := range img.Layer {
		if err = processLayer(fmt.Sprintf("%03d", i), layer.Name, &layer); err != nil {
			panic(err)
		}
	}
}
