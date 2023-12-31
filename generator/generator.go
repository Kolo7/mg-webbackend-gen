package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kolo7/mg-webbackend-gen/config"
	"github.com/kolo7/mg-webbackend-gen/dbmeta"
	"github.com/kolo7/mg-webbackend-gen/fileutils"
	"github.com/kolo7/mg-webbackend-gen/options"
)

func Generate(conf *dbmeta.Config) error {
	err := MkDirAll()
	if err != nil {
		return err
	}
	err = generateModel(conf)
	if err != nil {
		return err
	}
	err = generateDao(conf)
	if err != nil {
		return err
	}
	err = generateService(conf)
	if err != nil {
		return err
	}
	err = generateAPI(conf)
	if err != nil {
		return err
	}
	err = generateMain(conf)
	if err != nil {
		return err
	}
	err = generateController(conf)
	if err != nil {
		return err
	}

	if err = copyTemplatesToTarget(); err != nil {
		return err
	}

	err = generateDoc(conf)
	if err != nil {
		return err
	}

	if *options.RunGoFmt {
		GoFmt(conf.OutDir)
	}
	return nil
}

func MkDirAll() error {
	modelDir := filepath.Join(*options.OutDir, *options.ModelPackageName)
	daoDir := filepath.Join(*options.OutDir, *options.DaoPackageName)
	serviceDir := filepath.Join(*options.OutDir, *options.ServicePackageName)
	controllerDir := filepath.Join(*options.OutDir, *options.ControllerPackageName)
	mainDir := filepath.Join(*options.OutDir, "cmd")

	apiDir := filepath.Join(*options.OutDir, *options.ApiPackageName)
	templatesDir := filepath.Join(*options.OutDir, "templates")
	docDir := filepath.Join(*options.OutDir, "doc")

	err := os.MkdirAll(*options.OutDir, 0777)
	if err != nil && !*options.Overwrite {
		fmt.Print(options.Au.Red(fmt.Sprintf("unable to create outDir: %s error: %v\n", *options.OutDir, err)))
		return err
	}
	err = os.MkdirAll(modelDir, 0777)
	if err != nil && !*options.Overwrite {
		fmt.Print(options.Au.Red(fmt.Sprintf("unable to create modelDir: %s error: %v\n", *options.OutDir, err)))
		return err
	}
	if *options.DaoGenerate {
		err = os.MkdirAll(daoDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}
	if *options.ServiceGenerate {
		err = os.MkdirAll(serviceDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}
	if *options.RestAPIGenerate {
		err = os.MkdirAll(apiDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}
	if *options.ControllerGenerate {
		err = os.MkdirAll(controllerDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}

	if *options.FullGenerate {
		err = os.MkdirAll(mainDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create daoDir: %s error: %v\n", daoDir, err)))
			return err
		}
	}

	if *options.CopyTemplates {
		err = os.MkdirAll(templatesDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create templatesDir: %s error: %v\n", templatesDir, err)))
			return err
		}
	}

	if *options.DocGenerate {
		err = os.MkdirAll(docDir, 0777)
		if err != nil && !*options.Overwrite {
			fmt.Print(options.Au.Red(fmt.Sprintf("unable to create docDir: %s error: %v\n", docDir, err)))
			return err
		}
	}
	return nil
}

// generate model
func generateModel(conf *dbmeta.Config) error {
	var (
		err       error
		modelTmpl *dbmeta.GenTemplate
		modelDir  = filepath.Join(*options.OutDir, *options.ModelPackageName)
	)
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
	}
	return nil
}

func generateDao(conf *dbmeta.Config) error {
	if !*options.DaoGenerate {
		return nil
	}
	var (
		err         error
		daoTmpl     *dbmeta.GenTemplate
		daoInitTmpl *dbmeta.GenTemplate
		daoDir      = filepath.Join(*options.OutDir, *options.DaoPackageName)
	)
	if daoTmpl, err = config.LoadTemplate("dao_gorm.go.tmpl"); err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
		return err
	}
	if daoInitTmpl, err = config.LoadTemplate("dao_gorm_init.go.tmpl"); err != nil {
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
		// write dao
		outputFile := filepath.Join(daoDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(daoTmpl, modelInfo, outputFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}

	}
	data := map[string]interface{}{}

	err = conf.WriteTemplate(daoInitTmpl, data, filepath.Join(daoDir, "dao_base.go"))
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
		os.Exit(1)
	}
	return nil
}

func generateService(conf *dbmeta.Config) error {
	if !*options.ServiceGenerate {
		return nil
	}
	var (
		err         error
		serviceTmpl *dbmeta.GenTemplate
		serviceDir  = filepath.Join(*options.OutDir, *options.ServicePackageName)
	)
	if serviceTmpl, err = config.LoadTemplate("service.go.tmpl"); err != nil {
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
		serviceFile := filepath.Join(serviceDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(serviceTmpl, modelInfo, serviceFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}
	}
	return nil
}

func generateController(conf *dbmeta.Config) error {
	if !*options.ControllerGenerate {
		return nil
	}
	var (
		err            error
		controllerTmpl *dbmeta.GenTemplate
		apiDir         = filepath.Join(*options.OutDir, *options.ControllerPackageName)
	)
	if controllerTmpl, err = config.LoadTemplate("controller.go.tmpl"); err != nil {
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

		controllerFile := filepath.Join(apiDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(controllerTmpl, modelInfo, controllerFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}
	}
	return nil
}

func generateAPI(conf *dbmeta.Config) error {
	if !*options.RestAPIGenerate {
		return nil
	}
	var (
		err     error
		apiTmpl *dbmeta.GenTemplate
		apiDir  = filepath.Join(*options.OutDir, *options.ApiPackageName)
	)
	if apiTmpl, err = config.LoadTemplate("api.go.tmpl"); err != nil {
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

		restFile := filepath.Join(apiDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(apiTmpl, modelInfo, restFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}
	}
	return nil
}

func generateMain(conf *dbmeta.Config) error {
	if !*options.FullGenerate {
		return nil
	}
	var (
		err      error
		mainTmpl *dbmeta.GenTemplate
		mainDir  = filepath.Join(*options.OutDir, "cmd")
	)
	if mainTmpl, err = config.LoadTemplate("main.go.tmpl"); err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
		return err
	}

	mainFile := filepath.Join(mainDir, "main.go")

	data := map[string]interface{}{}
	list := []interface{}{}
	for tableName, tableInfo := range options.TableInfos {
		if len(tableInfo.Fields) == 0 {
			if *options.Verbose {
				fmt.Printf("[%d] Table: %s - No Fields Available\n", tableInfo.Index, tableName)
			}
			continue
		}
		modelInfo := conf.CreateContextForTableFile(tableInfo)
		list = append(list, modelInfo)
	}
	data["Modules"] = list
	err = conf.WriteTemplate(mainTmpl, data, mainFile)
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
		return err
	}
	return nil
}

func generateDoc(conf *dbmeta.Config) error {
	if !*options.DocGenerate {
		return nil
	}
	var (
		err         error
		docTmpl     *dbmeta.GenTemplate
		docInitTmpl *dbmeta.GenTemplate
		docDir      = filepath.Join(*options.OutDir, "doc")
	)
	if docTmpl, err = config.LoadTemplate("doc.md.tmpl"); err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
		return err
	}
	if docInitTmpl, err = config.LoadTemplate("doc_init.md.tmpl"); err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error loading template %v\n", err)))
		return err
	}

	docFile := filepath.Join(docDir, "doc.md")
	// 删除
	err = os.Remove(docFile)
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error remove file: %v\n", err)))
		return err
	}
	err = conf.AppendWriteTemplate(docInitTmpl, map[string]interface{}{}, docFile)
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
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

		err = conf.AppendWriteTemplate(docTmpl, modelInfo, docFile)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error writing file: %v\n", err)))
			return err
		}
	}
	return nil
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

func copyTemplatesToTarget() (err error) {
	if !*options.CopyTemplates {
		return nil
	}
	templatesDir := filepath.Join(*options.OutDir, "templates")
	fmt.Printf("Saving templates to %s\n", templatesDir)
	srcDir := "template"
	if options.TemplateDir != nil && *options.TemplateDir != "" {
		srcDir = *options.TemplateDir
	}
	err = fileutils.SaveAssets(srcDir, templatesDir, options.BaseTemplates)
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error saving: %v\n", err)))
	}
	return nil
}
