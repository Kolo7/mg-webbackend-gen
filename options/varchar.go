package options

import (
	"github.com/droundy/goopt"
	"github.com/gobuffalo/packr/v2"
	"github.com/kolo7/mg-webbackend-gen/dbmeta"
	"github.com/logrusorgru/aurora"
)

var (
	SqlType          = goopt.String([]string{"--sqltype"}, "mysql", "sql database type such as [ mysql, mssql, postgres, sqlite, etc. ]")
	SqlConnStr       = goopt.String([]string{"-c", "--connstr"}, "nil", "database connection string")
	SqlDatabase      = goopt.String([]string{"-d", "--database"}, "nil", "Database to for connection")
	SqlTable         = goopt.String([]string{"-t", "--table"}, "", "Table to build struct from")
	ExcludeSQLTables = goopt.String([]string{"-x", "--exclude"}, "", "Table(s) to exclude")
	TemplateDir      = goopt.String([]string{"--templateDir"}, "", "Template Dir")
	FragmentsDir     = goopt.String([]string{"--fragmentsDir"}, "", "Code fragments Dir")
	//SaveTemplateDir  = goopt.String([]string{"--save"}, "", "Save templates to dir")

	ModelPackageName    = goopt.String([]string{"--model"}, "model", "name to set for model package")
	ModelNamingTemplate = goopt.String([]string{"--model_naming"}, "{{FmtFieldName .}}", "model naming template to name structs")
	//FieldNamingTemplate = goopt.String([]string{"--field_naming"}, "{{FmtFieldName (stringifyFirstChar .) }}", "field naming template to name structs")
	FileNamingTemplate = goopt.String([]string{"--file_naming"}, "{{.}}", "file_naming template to name files")

	DaoPackageName = goopt.String([]string{"--dao"}, "dao", "name to set for dao package")
	ApiPackageName = goopt.String([]string{"--api"}, "api", "name to set for api package")
	//GrpcPackageName = goopt.String([]string{"--grpc"}, "grpc", "name to set for grpc package")
	OutDir = goopt.String([]string{"--out"}, ".", "output dir")
	//Module        = goopt.String([]string{"--module"}, "example.com/example", "module path")
	Overwrite = goopt.Flag([]string{"--overwrite"}, []string{"--no-overwrite"}, "Overwrite existing files (default)", "disable overwriting files")
	//Windows       = goopt.Flag([]string{"--windows"}, []string{}, "use windows line endings in generated files", "")
	NoColorOutput = goopt.Flag([]string{"--no-color"}, []string{}, "disable color output", "")

	//ContextFileName  = goopt.String([]string{"--context"}, "", "context file (json) to populate context with")
	MappingFileName = goopt.String([]string{"--mapping"}, "", "mapping file (json) to map sql types to golang/protobuf etc")
	//ExecCustomScript = goopt.String([]string{"--exec"}, "", "execute script for custom code generation")

	//AddJSONAnnotation = goopt.Flag([]string{"--json"}, []string{"--no-json"}, "Add json annotations (default)", "Disable json annotations")
	JsonNameFormat = goopt.String([]string{"--json-fmt"}, "snake", "json name format [snake | camel | lower_camel | none]")

	//AddXMLAnnotation = goopt.Flag([]string{"--xml"}, []string{"--no-xml"}, "Add xml annotations (default)", "Disable xml annotations")
	XmlNameFormat = goopt.String([]string{"--xml-fmt"}, "snake", "xml name format [snake | camel | lower_camel | none]")

	AddGormAnnotation = goopt.Flag([]string{"--gorm"}, []string{}, "Add gorm annotations (tags)", "")
	//AddProtobufAnnotation = goopt.Flag([]string{"--protobuf"}, []string{}, "Add protobuf annotations (tags)", "")
	//ProtoNameFormat = goopt.String([]string{"--proto-fmt"}, "snake", "proto name format [snake | camel | lower_camel | none]")
	//GogoProtoImport = goopt.String([]string{"--gogo-proto"}, "", "location of gogo import ")

	//AddDBAnnotation = goopt.Flag([]string{"--db"}, []string{}, "Add db annotations (tags)", "")
	//UseGureguTypes = goopt.Flag([]string{"--guregu"}, []string{}, "Add guregu null types", "")

	//CopyTemplates    = goopt.Flag([]string{"--copy-templates"}, []string{}, "Copy regeneration templates to project directory", "")
	//ModGenerate      = goopt.Flag([]string{"--mod"}, []string{}, "Generate go.mod in output dir", "")
	//MakefileGenerate = goopt.Flag([]string{"--makefile"}, []string{}, "Generate Makefile in output dir", "")
	//ServerGenerate  = goopt.Flag([]string{"--server"}, []string{}, "Generate server app output dir", "")
	DaoGenerate = goopt.Flag([]string{"--generate-dao"}, []string{}, "Generate dao functions", "")
	//ProjectGenerate = goopt.Flag([]string{"--generate-proj"}, []string{}, "Generate project readme and gitignore", "")
	RestAPIGenerate = goopt.Flag([]string{"--rest"}, []string{}, "Enable generating RESTful api", "")
	RunGoFmt        = goopt.Flag([]string{"--run-gofmt"}, []string{}, "run gofmt on output dir", "")

	ServerListen = goopt.String([]string{"--listen"}, "", "listen address e.g. :8080")
	ServerScheme = goopt.String([]string{"--scheme"}, "http", "scheme for server url")
	ServerHost   = goopt.String([]string{"--host"}, "localhost", "host for server")
	ServerPort   = goopt.Int([]string{"--port"}, 8080, "port for server")
	//SwaggerVersion      = goopt.String([]string{"--swagger_version"}, "1.0", "swagger version")
	//SwaggerBasePath     = goopt.String([]string{"--swagger_path"}, "/", "swagger base path")
	//SwaggerTos          = goopt.String([]string{"--swagger_tos"}, "", "swagger tos url")
	//SwaggerContactName  = goopt.String([]string{"--swagger_contact_name"}, "Me", "swagger contact name")
	//SwaggerContactURL   = goopt.String([]string{"--swagger_contact_url"}, "http://me.com/terms.html", "swagger contact url")
	//SwaggerContactEmail = goopt.String([]string{"--swagger_contact_email"}, "me@me.com", "swagger contact email")

	Verbose = goopt.Flag([]string{"-v", "--verbose"}, []string{}, "Enable verbose output", "")

	//NameTest = goopt.String([]string{"--name_test"}, "", "perform name test using the --model_naming or --file_naming options")

	BaseTemplates *packr.Box
	TableInfos    map[string]*dbmeta.ModelInfo
	Au            aurora.Aurora
)
