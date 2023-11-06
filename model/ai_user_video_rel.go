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


Table: ai_user_video_rel
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] user_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
[ 2] video_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
[ 3] video_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] video_sort                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 14,    "user_id": 43,    "video_type": 30,    "video_id": "gJjJLZdXnLhpSMGvECtniUwVA",    "video_sort": 24,    "creater": "pkYhxmQJqxRrUATMULRdDQoej",    "upadter": "VgSCrMJsrBCpkZkjVxkWDEvrS",    "created_at": "2179-05-03T05:53:09.203704207+08:00"}



*/

// AiUserVideoRel struct is a row record of the ai_user_video_rel table in the  database
/*
type AiUserVideoRel struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] user_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
    UserID int64 `gorm:"column:user_id;type:bigint;" json:"user_id" xml:"user_id" db:"user_id" protobuf:"int64,1,opt,name=user_id"`
    //[ 2] video_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
    VideoType int32 `gorm:"column:video_type;type:tinyint;default:1;" json:"video_type" xml:"video_type" db:"video_type" protobuf:"int32,2,opt,name=video_type"`
    //[ 3] video_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    VideoID string `gorm:"column:video_id;type:varchar;" json:"video_id" xml:"video_id" db:"video_id" protobuf:"string,3,opt,name=video_id"`
    //[ 4] video_sort                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    VideoSort int32 `gorm:"column:video_sort;type:int;default:0;" json:"video_sort" xml:"video_sort" db:"video_sort" protobuf:"int32,4,opt,name=video_sort"`
    //[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,5,opt,name=creater"`
    //[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,6,opt,name=upadter"`
    //[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,7,opt,name=created_at"`

}
*/

var ai_user_video_relTableInfo = &TableInfo{
	Name: "ai_user_video_rel",
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
			Name:               "video_type",
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
			GoFieldName:        "VideoType",
			GoFieldType:        "int32",
			JSONFieldName:      "video_type",
			ProtobufFieldName:  "video_type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "video_id",
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
			GoFieldName:        "VideoID",
			GoFieldType:        "string",
			JSONFieldName:      "video_id",
			ProtobufFieldName:  "video_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "video_sort",
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
			GoFieldName:        "VideoSort",
			GoFieldType:        "int32",
			JSONFieldName:      "video_sort",
			ProtobufFieldName:  "video_sort",
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
func (a *AiUserVideoRel) TableName() string {
	return "ai_user_video_rel"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiUserVideoRel) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiUserVideoRel) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiUserVideoRel) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiUserVideoRel) TableInfo() *TableInfo {
	return ai_user_video_relTableInfo
}
