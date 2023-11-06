package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	_ = datatypes.JSON{}
)

/*
DB Table Details
-------------------------------------


Table: ai_sys_config
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] config_name                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 2] config_key                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] config_value                                   varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] remark                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 5] deleted                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 6] create_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
[ 8] update_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 9] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]


JSON Sample
-------------------------------------
{    "id": 36,    "config_name": "UfGYysJqotUDdWElFJgYRiwZm",    "config_key": "HbmtoOstyqYdgHxathvhWLQsQ",    "config_value": "BmZSgElSWvmitGSoPdBjsPVkm",    "remark": "vYChRcaVegfJrFFFtBoRKFAPn",    "deleted": 0,    "create_by": "ymBubaWFlyyEhZaZfmoKnbjSZ",    "create_time": "2207-08-27T11:24:12.547752998+08:00",    "update_by": "umLhtSJBuseqNnydtlBBQOtAa",    "update_time": "2171-03-29T23:55:14.155455568+08:00"}



*/

// AiSysConfig struct is a row record of the ai_sys_config table in the  database
/*
type AiSysConfig struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] config_name                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    ConfigName string `gorm:"column:config_name;type:varchar;" json:"config_name" xml:"config_name" db:"config_name" protobuf:"string,1,opt,name=config_name"`
    //[ 2] config_key                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    ConfigKey string `gorm:"column:config_key;type:varchar;" json:"config_key" xml:"config_key" db:"config_key" protobuf:"string,2,opt,name=config_key"`
    //[ 3] config_value                                   varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    ConfigValue string `gorm:"column:config_value;type:varchar;" json:"config_value" xml:"config_value" db:"config_value" protobuf:"string,3,opt,name=config_value"`
    //[ 4] remark                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Remark string `gorm:"column:remark;type:varchar;" json:"remark" xml:"remark" db:"remark" protobuf:"string,4,opt,name=remark"`
    //[ 5] deleted                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    Deleted int32 `gorm:"column:deleted;type:tinyint;default:0;" json:"deleted" xml:"deleted" db:"deleted" protobuf:"int32,5,opt,name=deleted"`
    //[ 6] create_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    CreateBy string `gorm:"column:create_by;type:varchar;" json:"create_by" xml:"create_by" db:"create_by" protobuf:"string,6,opt,name=create_by"`
    //[ 7] create_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
    CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;" json:"create_time" xml:"create_time" db:"create_time" protobuf:"uint64,7,opt,name=create_time"`
    //[ 8] update_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UpdateBy string `gorm:"column:update_by;type:varchar;" json:"update_by" xml:"update_by" db:"update_by" protobuf:"string,8,opt,name=update_by"`
    //[ 9] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
    UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;" json:"update_time" xml:"update_time" db:"update_time" protobuf:"uint64,9,opt,name=update_time"`

}
*/

var ai_sys_configTableInfo = &TableInfo{
	Name: "ai_sys_config",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "config_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "ConfigName",
			GoFieldType:        "string",
			JSONFieldName:      "config_name",
			ProtobufFieldName:  "config_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "config_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "ConfigKey",
			GoFieldType:        "string",
			JSONFieldName:      "config_key",
			ProtobufFieldName:  "config_key",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "config_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "ConfigValue",
			GoFieldType:        "string",
			JSONFieldName:      "config_value",
			ProtobufFieldName:  "config_value",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "remark",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "deleted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Deleted",
			GoFieldType:        "int32",
			JSONFieldName:      "deleted",
			ProtobufFieldName:  "deleted",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "create_by",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "CreateBy",
			GoFieldType:        "string",
			JSONFieldName:      "create_by",
			ProtobufFieldName:  "create_by",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "create_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "update_by",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       -1,
			GoFieldName:        "UpdateBy",
			GoFieldType:        "string",
			JSONFieldName:      "update_by",
			ProtobufFieldName:  "update_by",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiSysConfig) TableName() string {
	return "ai_sys_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiSysConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiSysConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiSysConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiSysConfig) TableInfo() *TableInfo {
	return ai_sys_configTableInfo
}
