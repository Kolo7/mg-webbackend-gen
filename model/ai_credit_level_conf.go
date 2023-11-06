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


Table: ai_credit_level_conf
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] credit_level                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
[ 2] level_name                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] level_score                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 4] logo_url                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 5] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 6] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 8] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 27,    "credit_level": 26,    "level_name": "EumPmHgMpFBeQBRDsZvLNGDoZ",    "level_score": 81,    "logo_url": "BjqadnclCgSnvDumUdlKywPom",    "online_status": 15,    "creater": "OxNiLfKtvZmkcVpEGeUlDCwFY",    "upadter": "TUOtHATmIJDopUPIDUgiHuBsr",    "created_at": "2290-06-23T22:05:51.743362745+08:00"}



*/

// AiCreditLevelConf struct is a row record of the ai_credit_level_conf table in the  database
/*
type AiCreditLevelConf struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] credit_level                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
    CreditLevel int32 `gorm:"column:credit_level;type:tinyint;default:1;" json:"credit_level" xml:"credit_level" db:"credit_level" protobuf:"int32,1,opt,name=credit_level"`
    //[ 2] level_name                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    LevelName string `gorm:"column:level_name;type:varchar;" json:"level_name" xml:"level_name" db:"level_name" protobuf:"string,2,opt,name=level_name"`
    //[ 3] level_score                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    LevelScore int32 `gorm:"column:level_score;type:int;default:0;" json:"level_score" xml:"level_score" db:"level_score" protobuf:"int32,3,opt,name=level_score"`
    //[ 4] logo_url                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    LogoURL string `gorm:"column:logo_url;type:varchar;" json:"logo_url" xml:"logo_url" db:"logo_url" protobuf:"string,4,opt,name=logo_url"`
    //[ 5] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,5,opt,name=online_status"`
    //[ 6] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,6,opt,name=creater"`
    //[ 7] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,7,opt,name=upadter"`
    //[ 8] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,8,opt,name=created_at"`

}
*/

var ai_credit_level_confTableInfo = &TableInfo{
	Name: "ai_credit_level_conf",
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
			Name:               "credit_level",
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
			GoFieldName:        "CreditLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "credit_level",
			ProtobufFieldName:  "credit_level",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "level_name",
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
			GoFieldName:        "LevelName",
			GoFieldType:        "string",
			JSONFieldName:      "level_name",
			ProtobufFieldName:  "level_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "level_score",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LevelScore",
			GoFieldType:        "int32",
			JSONFieldName:      "level_score",
			ProtobufFieldName:  "level_score",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "logo_url",
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
			GoFieldName:        "LogoURL",
			GoFieldType:        "string",
			JSONFieldName:      "logo_url",
			ProtobufFieldName:  "logo_url",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "online_status",
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
			GoFieldName:        "OnlineStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "online_status",
			ProtobufFieldName:  "online_status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "creater",
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
			GoFieldName:        "Creater",
			GoFieldType:        "string",
			JSONFieldName:      "creater",
			ProtobufFieldName:  "creater",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "upadter",
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
			GoFieldName:        "Upadter",
			GoFieldType:        "string",
			JSONFieldName:      "upadter",
			ProtobufFieldName:  "upadter",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiCreditLevelConf) TableName() string {
	return "ai_credit_level_conf"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiCreditLevelConf) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiCreditLevelConf) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiCreditLevelConf) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiCreditLevelConf) TableInfo() *TableInfo {
	return ai_credit_level_confTableInfo
}
