package fileutils

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// SaveAssets will save the prepacked templates for local editing. File structure will be recreated under the output dir.
func SaveAssets(srcDir, outputDir string, ffs *embed.FS) error {
	fmt.Printf("SaveAssets: %v\n", outputDir)
	if outputDir == "" {
		outputDir = "."
	}

	outputDir = strings.TrimSuffix(outputDir, "/")

	if outputDir == "" {
		outputDir = "."
	}

	dirs, err := ffs.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		fileInfo, err := dir.Info()
		fileName := fmt.Sprintf("%s/%s", outputDir, fileInfo.Name())
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			file, err := ffs.Open(filepath.Join(srcDir, fileInfo.Name()))
			if err != nil {
				return err
			}
			err = WriteNewFile(fileName, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// WriteNewFile will attempt to write a file with the filename and path, a Reader and the FileMode of the file to be created.
// If an error is encountered an error will be returned.
func WriteNewFile(fpath string, in io.Reader) error {
	err := os.MkdirAll(filepath.Dir(fpath), 0775)
	if err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}

	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer func() {
		_ = out.Close()
	}()

	fmt.Printf("WriteNewFile: %s\n", fpath)

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", fpath, err)
	}
	return nil
}
