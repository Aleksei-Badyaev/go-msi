package tpls

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/mattn/go-zglob"
	"github.com/observiq/go-msi/manifest"
)

func Find(srcDir string) ([]string, error) {
	glob := filepath.Join(srcDir, "*.wxs")
	return zglob.Glob(glob)
}

func GenerateTemplate(wixFile *manifest.WixManifest, src string, out string) error {
	tpl, err := template.ParseFiles(src)
	if err != nil {
		return err
	}

	fileWriter, err := os.Create(out)
	if err != nil {
		return err
	}
	defer fileWriter.Close()
	err = tpl.Execute(fileWriter, wixFile)
	if err != nil {
		return err
	}
	return nil
}
