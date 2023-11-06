package options

import (
	"fmt"
	"strings"

	"github.com/droundy/goopt"
	"github.com/kolo7/mg-webbackend-gen/dbmeta"
)

func CheckOptions() bool {
	if *ServerListen == "" {
		*ServerListen = fmt.Sprintf(":%d", *ServerPort)
	}
	if *ServerScheme != "http" && *ServerScheme != "https" {
		*ServerScheme = "http"
	}

	if SqlConnStr == nil || *SqlConnStr == "" || *SqlConnStr == "nil" {
		fmt.Print(Au.Red("sql connection string is required! Add it with --connstr=s\n\n"))
		fmt.Println(goopt.Usage())
		return false
	}

	if SqlDatabase == nil || *SqlDatabase == "" || *SqlDatabase == "nil" {
		fmt.Print(Au.Red("Database can not be null\n\n"))
		fmt.Println(goopt.Usage())
		return false
	}

	if strings.HasPrefix(*ModelNamingTemplate, "'") && strings.HasSuffix(*ModelNamingTemplate, "'") {
		*ModelNamingTemplate = strings.TrimSuffix(*ModelNamingTemplate, "'")
		*ModelNamingTemplate = strings.TrimPrefix(*ModelNamingTemplate, "'")
	}

	if OutDir == nil || *OutDir == "" {
		*OutDir = "."
	}

	if *MappingFileName != "" {
		err := dbmeta.LoadMappings(*MappingFileName, *Verbose)
		if err != nil {
			fmt.Print(Au.Red(fmt.Sprintf("Error loading mappings file %s error: %v\n", *MappingFileName, err)))
			return false
		}
	}

	*JsonNameFormat = strings.ToLower(*JsonNameFormat)
	*XmlNameFormat = strings.ToLower(*XmlNameFormat)
	return true
}
