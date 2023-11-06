package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kolo7/mg-webbackend-gen/dbmeta"
	"github.com/kolo7/mg-webbackend-gen/options"
)

func LoadTemplate(filename string) (tpl *dbmeta.GenTemplate, err error) {
	baseName := filepath.Base(filename)
	// fmt.Printf("LoadTemplate: %s / %s\n", filename, baseName)

	if *options.TemplateDir != "" {
		fpath := filepath.Join(*options.TemplateDir, filename)
		var b []byte
		b, err = ioutil.ReadFile(fpath)
		if err == nil {

			absPath, err := filepath.Abs(fpath)
			if err != nil {
				absPath = fpath
			}
			// fmt.Printf("Loaded template from file: %s\n", fpath)
			tpl = &dbmeta.GenTemplate{Name: "file://" + absPath, Content: string(b)}
			return tpl, nil
		}
	}

	content, err := options.BaseTemplates.FindString(baseName)
	if err != nil {
		return nil, fmt.Errorf("%s not found internally", baseName)
	}
	if *options.Verbose {
		fmt.Printf("Loaded template from app: %s\n", filename)
	}

	tpl = &dbmeta.GenTemplate{Name: "internal://" + filename, Content: content}
	return tpl, nil
}
