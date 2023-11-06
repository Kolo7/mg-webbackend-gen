package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kolo7/mg-webbackend-gen/config"
	"github.com/kolo7/mg-webbackend-gen/dbmeta"
	"github.com/kolo7/mg-webbackend-gen/options"
)

func Generate(conf *dbmeta.Config) error {
	modelDir := filepath.Join(*options.OutDir, *options.ModelPackageName)
	daoDir := filepath.Join(*options.OutDir, *options.DaoPackageName)
	apiDir := filepath.Join(*options.OutDir, *options.ApiPackageName)
	err := initDir(modelDir, daoDir, apiDir)
	if err != nil {
		return err
	}

	var daoTmpl *dbmeta.GenTemplate
	var daoInitTmpl *dbmeta.GenTemplate
	var modelTmpl *dbmeta.GenTemplate
	//var modelBaseTmpl *dbmeta.GenTemplate

	if *options.AddGormAnnotation {
		if daoTmpl, err = config.LoadTemplate("dao_gorm.go.tmpl"); err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
			return err
		}
		if daoInitTmpl, err = config.LoadTemplate("dao_gorm_init.go.tmpl"); err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
			return err
		}
	}
	if modelTmpl, err = config.LoadTemplate("model.go.tmpl"); err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
		return err
	}
	for tableName, tableInfo := range options.TableInfos {
		if len(tableInfo.Fields) == 0 {
			if *options.Verbose {
				fmt.Printf("[%d] Table: %s - No Fields Available\n", tableInfo.Index, tableName)
			}
			continue
		}
		modelInfo := conf.CreateContextForTableFile(tableInfo)
		modelFile := filepath.Join(modelDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(modelTmpl, modelInfo, modelFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}
		if *options.DaoGenerate {
			// write dao
			outputFile := filepath.Join(daoDir, CreateGoSrcFileName(tableName))
			err = conf.WriteTemplate(daoTmpl, modelInfo, outputFile)
			if err != nil {
				fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
				return err
			}
		}
	}
	data := map[string]interface{}{}

	if *options.DaoGenerate {
		err = conf.WriteTemplate(daoInitTmpl, data, filepath.Join(daoDir, "dao_base.go"))
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			os.Exit(1)
		}
	}
	// err = conf.WriteTemplate(modelBaseTmpl, data, filepath.Join(modelDir, "model_base.go"))
	// if err != nil {
	// 	fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
	// 	os.Exit(1)
	// }
	data = map[string]interface{}{
		"deps":        "go list -f '{{ join .Deps  \"\\n\"}}' .",
		"CommandLine": conf.CmdLine,
		"Config":      conf,
	}
	if *options.RunGoFmt {
		GoFmt(conf.OutDir)
	}
	return nil
}

func initDir(modelDir, daoDir, apiDir string) error {
	err := os.MkdirAll(*options.OutDir, 0777)
	if err != nil && !*options.Overwrite {
		fmt.Print(options.Au.Red(fmt.Sprintf("unable to create outDir: %s error: %v\n", *options.OutDir, err)))
		return err
	}
	err = os.MkdirAll(modelDir, 0777)
	if err != nil && !*options.Overwrite {
		fmt.Print(options.Au.Red(fmt.Sprintf("unable to create modelDir: %s error: %v\n", modelDir, err)))
		return err
	}
	if *options.DaoGenerate {
		err = os.MkdirAll(daoDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}

	if *options.RestAPIGenerate {
		err = os.MkdirAll(apiDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create apiDir: %s error: %v\n", apiDir, err)))
			return err
		}
	}
	return nil
}

// createGoSrcFileName ensures name doesnt clash with go naming conventions like _test.go
func createGoSrcFileName(tableName string) string {
	name := dbmeta.Replace(*options.FileNamingTemplate, tableName)
	// name := inflection.Singular(tableName)

	if strings.HasSuffix(name, "_test") {
		name = name[0 : len(name)-5]
		name = name + "_tst"
	}
	return name + ".go"
}

// CreateGoSrcFileName ensures name doesnt clash with go naming conventions like _test.go
func CreateGoSrcFileName(tableName string) string {
	name := dbmeta.Replace(*options.FileNamingTemplate, tableName)
	// name := inflection.Singular(tableName)

	if strings.HasSuffix(name, "_test") {
		name = name[0 : len(name)-5]
		name = name + "_tst"
	}
	return name + ".go"
}

// GoFmt exec gofmt for a code dir
func GoFmt(codeDir string) (string, error) {
	args := []string{"-s", "-d", "-w", "-l", codeDir}
	cmd := exec.Command("gofmt", args...)

	cmdLineArgs := strings.Join(args, " ")
	fmt.Printf("gofmt %s\n", cmdLineArgs)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("error calling protoc: %T %v\n", err, err)))
		fmt.Print(options.Au.Red(fmt.Sprintf("%s\n", stdoutStderr)))
		return "", err
	}

	return string(stdoutStderr), nil
}
