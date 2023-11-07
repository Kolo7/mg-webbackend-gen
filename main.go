package main

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/kolo7/mg-webbackend-gen/config"
	"github.com/kolo7/mg-webbackend-gen/database"
	"github.com/kolo7/mg-webbackend-gen/dbmeta"
	"github.com/kolo7/mg-webbackend-gen/generator"
	"github.com/kolo7/mg-webbackend-gen/options"
	"github.com/logrusorgru/aurora"
)

func main() {
	// 初始化彩色终端输出
	options.Au = aurora.NewAurora(!*options.NoColorOutput)
	dbmeta.InitColorOutput(options.Au)

	// 从包体取出模板
	options.BaseTemplates = packr.New("gen", "./template")

	// 检查必要参数设置
	if !options.CheckOptions() {
		return
	}

	// 初始化数据库
	err := database.InitSchema()
	if err != nil {
		return
	}

	// 初始化config，用于后续代码生成
	conf := dbmeta.NewConfig(config.LoadTemplate)
	initialize(conf)

	tableName, err := database.GetTableNames()
	if err != nil {
		return
	}
	excludeDbTables, err := database.GetExcludeTableName()
	if err != nil {
		return
	}
	options.TableInfos = dbmeta.LoadTableInfo(database.GetDB(),
		tableName, excludeDbTables, conf)
	if len(options.TableInfos) == 0 {
		fmt.Print(options.Au.Red("No tables loaded\n"))
		return
	}

	fmt.Printf("Generating code for the following tables (%d)\n", len(options.TableInfos))
	i := 0
	for tableName := range options.TableInfos {
		fmt.Printf("[%d] %s\n", i, tableName)
		i++
	}
	conf.TableInfos = options.TableInfos
	conf.ContextMap["tableInfos"] = options.TableInfos

	if *options.Verbose {
		listTemplates()
	}
	err = generator.Generate(conf)
	if err != nil {
		return
	}
}

func listTemplates() {
	for i, file := range options.BaseTemplates.List() {
		fmt.Printf("   [%d] [%s]\n", i, file)
	}
}

func initialize(conf *dbmeta.Config) {

	// load fragments if specified
	if options.FragmentsDir != nil && *options.FragmentsDir != "" {
		conf.LoadFragments(*options.FragmentsDir)
	}

	conf.SQLType = *options.SqlType
	conf.SQLDatabase = *options.SqlDatabase

	conf.AddJSONAnnotation = *options.AddJSONAnnotation
	conf.AddGormAnnotation = *options.AddGormAnnotation
	conf.AddDBAnnotation = *options.AddDBAnnotation
	conf.JSONNameFormat = *options.JsonNameFormat
	conf.Verbose = *options.Verbose
	conf.OutDir = *options.OutDir
	conf.Overwrite = *options.Overwrite
	conf.LineEndingCRLF = *options.Windows

	conf.SQLConnStr = *options.SqlConnStr
	conf.ServerPort = *options.ServerPort
	conf.ServerHost = *options.ServerHost
	conf.ServerScheme = *options.ServerScheme
	conf.ServerListen = *options.ServerListen

	conf.Module = *options.Module
	conf.ModelPackageName = *options.ModelPackageName
	conf.ModelFQPN = *options.Module + "/" + *options.ModelPackageName

	conf.DaoPackageName = *options.DaoPackageName
	conf.DaoFQPN = *options.Module + "/" + *options.DaoPackageName

	conf.APIPackageName = *options.ApiPackageName
	conf.APIFQPN = *options.Module + "/" + *options.ApiPackageName

	conf.FileNamingTemplate = *options.FileNamingTemplate
	conf.ModelNamingTemplate = *options.ModelNamingTemplate
	conf.FieldNamingTemplate = *options.FieldNamingTemplate

	conf.JSONNameFormat = strings.ToLower(conf.JSONNameFormat)
	conf.XMLNameFormat = strings.ToLower(conf.XMLNameFormat)
	conf.ProtobufNameFormat = strings.ToLower(conf.ProtobufNameFormat)
}
