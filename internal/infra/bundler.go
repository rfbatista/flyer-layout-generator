package infra

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"go.uber.org/zap"
)

func NewBundler(logger *zap.Logger) (*Bundler, error) {
	assetsDirPath := fmt.Sprintf("%s/dist/web", FindProjectRoot())
	return &Bundler{assetsPath: assetsDirPath, logger: logger}, nil
}

type BundlerPageInfo struct {
	Name        string
	JSPath      string
	CSSPath     string
	JSName      string
	CSSName     string
	EntryPoints []string
}

type BuildResult struct {
	JS           string
	CSS          string
	JSPath       string
	CSSPath      string
	JSName       string
	CSSName      string
	Dependencies []string
}

type BundlerPageParams struct {
	Name        string
	EntryPoints []string
}

type Bundler struct {
	logger     *zap.Logger
	assetsPath string
	pages      []BundlerPageInfo
}

func (b *Bundler) AddPage(pa BundlerPageParams) (BundlerPageInfo, error) {
	hashName := strconv.Itoa(rand.Int())
	cssName := hashName + ".css"
	jsName := hashName + ".js"
	cssPath := b.assetsPath + "/" + cssName
	jsPath := b.assetsPath + "/" + jsName
	p := BundlerPageInfo{
		EntryPoints: pa.EntryPoints,
		Name:        pa.Name,
		JSName:      jsName,
		CSSName:     cssName,
		CSSPath:     cssPath,
		JSPath:      jsPath,
	}
	b.pages = append(b.pages, p)
	return p, nil
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bundler) Build() error {
	_, err := os.Stat(b.assetsPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(b.assetsPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if b.assetsPath == "" {
			return errors.New("empty assets directory path")
		}
		b.logger.Debug("cleaning assets folder")
		err := RemoveContents(b.assetsPath)
		if err != nil {
			return err
		}
	}
	for _, page := range b.pages {
		b.logger.Debug("building page", zap.String("name", page.Name))
		_, err := ServerBundler(b.assetsPath, page)
		if err != nil {
			return err
		}
	}
	return nil
}

func ServerBundler(assetsDirPath string, page BundlerPageInfo) (BuildResult, error) {
	opts := esbuild.BuildOptions{
		EntryPoints: page.EntryPoints,
		Platform:    esbuild.PlatformNode,
		Bundle:      true,
		Write:       false,
		Outdir:      assetsDirPath,
		Metafile:    false,
		AssetNames:  fmt.Sprintf("%s/[name]", strings.TrimPrefix(assetsDirPath, "/")),
		Loader: map[string]esbuild.Loader{ // for loading images properly
			".png":   esbuild.LoaderFile,
			".svg":   esbuild.LoaderFile,
			".jpg":   esbuild.LoaderFile,
			".jpeg":  esbuild.LoaderFile,
			".gif":   esbuild.LoaderFile,
			".bmp":   esbuild.LoaderFile,
			".woff2": esbuild.LoaderFile,
			".woff":  esbuild.LoaderFile,
			".ttf":   esbuild.LoaderFile,
			".eot":   esbuild.LoaderFile,
		},
	}
	result := esbuild.Build(opts)
	if len(result.Errors) > 0 {
		fileLocation := "unknown"
		lineNum := "unknown"
		if result.Errors[0].Location != nil {
			fileLocation = result.Errors[0].Location.File
			lineNum = result.Errors[0].Location.LineText
		}
		return BuildResult{}, fmt.Errorf(
			"%s <br>in %s <br>at %s",
			result.Errors[0].Text,
			fileLocation,
			lineNum,
		)
	}

	var br BuildResult
	for _, file := range result.OutputFiles {
		if strings.HasSuffix(file.Path, ".js") {
			br.JS = string(file.Contents)
		} else if strings.HasSuffix(file.Path, ".css") {
			br.CSS = string(file.Contents)
		}
	}
	br.CSSName = page.CSSName
	err := os.WriteFile(page.CSSPath, []byte(br.CSS), 0o644)
	if err != nil {
		return BuildResult{}, err
	}
	err = os.WriteFile(page.JSPath, []byte(br.JS), 0o644)
	if err != nil {
		return BuildResult{}, err
	}
	br.JS = page.JSPath
	br.CSS = page.CSSPath
	return br, nil
}
