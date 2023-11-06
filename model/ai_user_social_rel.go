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


Table: ai_user_social_rel
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] user_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
[ 2] bind_user_id                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
[ 3] social_name                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] social_sort                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 74,    "user_id": 45,    "bind_user_id": 78,    "social_name": "qjfWGPgnkCFCouokQWEGOLrSa",    "social_sort": 8,    "creater": "mPUOgHlFvawfKkhWFVXYELiMm",    "upadter": "wcCqHgtJJhmEDNYJrmlNySiiB",    "created_at": "2233-09-17T10:59:25.308675714+08:00"}



*/

// AiUserSocialRel struct is a row record of the ai_user_social_rel table in the  database
/*
type AiUserSocialRel struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] user_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
    UserID int64 `gorm:"column:user_id;type:bigint;default:0;" json:"user_id" xml:"user_id" db:"user_id" protobuf:"int64,1,opt,name=user_id"`
    //[ 2] bind_user_id                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
    BindUserID int64 `gorm:"column:bind_user_id;type:bigint;default:0;" json:"bind_user_id" xml:"bind_user_id" db:"bind_user_id" protobuf:"int64,2,opt,name=bind_user_id"`
    //[ 3] social_name                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    SocialName string `gorm:"column:social_name;type:varchar;" json:"social_name" xml:"social_name" db:"social_name" protobuf:"string,3,opt,name=social_name"`
    //[ 4] social_sort                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    SocialSort int32 `gorm:"column:social_sort;type:int;default:0;" json:"social_sort" xml:"social_sort" db:"social_sort" protobuf:"int32,4,opt,name=social_sort"`
    //[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,5,opt,name=creater"`
    //[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,6,opt,name=upadter"`
    //[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,7,opt,name=created_at"`

}
*/

var ai_user_social_relTableInfo = &TableInfo{
	Name: "ai_user_social_rel",
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
			Name:               "user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int64",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "bind_user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "BindUserID",
			GoFieldType:        "int64",
			JSONFieldName:      "bind_user_id",
			ProtobufFieldName:  "bind_user_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "social_name",
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
			GoFieldName:        "SocialName",
			GoFieldType:        "string",
			JSONFieldName:      "social_name",
			ProtobufFieldName:  "social_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "social_sort",
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
			GoFieldName:        "SocialSort",
			GoFieldType:        "int32",
			JSONFieldName:      "social_sort",
			ProtobufFieldName:  "social_sort",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiUserSocialRel) TableName() string {
	return "ai_user_social_rel"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiUserSocialRel) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiUserSocialRel) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiUserSocialRel) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiUserSocialRel) TableInfo() *TableInfo {
	return ai_user_social_relTableInfo
}
