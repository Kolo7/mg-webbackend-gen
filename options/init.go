package options

import "github.com/droundy/goopt"

func init() {
	// Setup goopts
	goopt.Description = func() string {
		return "ORM and RESTful API generator for SQl databases"
	}
	goopt.Version = "v0.9.27 (08/04/2020)"
	goopt.Summary = `gen [-v] --sqltype=mysql --connstr "user:password@/dbname" --database <databaseName> --module=example.com/example [--json] [--gorm] [--guregu] [--generate-dao] [--generate-proj]
git fetch up
           sqltype - sql database type such as [ mysql, mssql, postgres, sqlite, etc. ]

`

	// Parse options
	goopt.Parse(nil)
}
