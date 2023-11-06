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


Table: ai_play_guide
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] guide_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
[ 2] content                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] guide_jump                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 5] video_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
[ 6] video_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 8] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 9] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-10-11 00:00:00]
[10] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-10-11 00:00:00]


JSON Sample
-------------------------------------
{    "id": 58,    "guide_type": 75,    "content": "sKkJkniqjuJjbVfAKQXllulKr",    "guide_jump": "rYeghHngxqXSvkfLtaIwWuMcx",    "online_status": 4,    "video_type": 56,    "video_id": "yjnWbRIFdLGiZcjVmoMGOrUuL",    "creater": "AXxsRUqjRfITFFAsXCLbtIurS",    "upadter": "WqeVpVAjirQMqDtIJOGkbYRcl",    "created_at": "2270-01-05T08:24:26.512744944+08:00",    "updated_at": "2026-10-13T19:27:37.56802612+08:00"}



*/

// AiPlayGuide struct is a row record of the ai_play_guide table in the  database
/*
type AiPlayGuide struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] guide_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
    GuideType int32 `gorm:"column:guide_type;type:tinyint;default:1;" json:"guide_type" xml:"guide_type" db:"guide_type" protobuf:"int32,1,opt,name=guide_type"`
    //[ 2] content                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Content string `gorm:"column:content;type:varchar;" json:"content" xml:"content" db:"content" protobuf:"string,2,opt,name=content"`
    //[ 3] guide_jump                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    GuideJump string `gorm:"column:guide_jump;type:varchar;" json:"guide_jump" xml:"guide_jump" db:"guide_jump" protobuf:"string,3,opt,name=guide_jump"`
    //[ 4] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,4,opt,name=online_status"`
    //[ 5] video_type                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
    VideoType int32 `gorm:"column:video_type;type:tinyint;default:1;" json:"video_type" xml:"video_type" db:"video_type" protobuf:"int32,5,opt,name=video_type"`
    //[ 6] video_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    VideoID string `gorm:"column:video_id;type:varchar;" json:"video_id" xml:"video_id" db:"video_id" protobuf:"string,6,opt,name=video_id"`
    //[ 7] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,7,opt,name=creater"`
    //[ 8] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,8,opt,name=upadter"`
    //[ 9] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-10-11 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-10-11 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,9,opt,name=created_at"`
    //[10] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-10-11 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-10-11 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,10,opt,name=updated_at"`

}
*/

var ai_play_guideTableInfo = &TableInfo{
	Name: "ai_play_guide",
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
			Name:               "guide_type",
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
			GoFieldName:        "GuideType",
			GoFieldType:        "int32",
			JSONFieldName:      "guide_type",
			ProtobufFieldName:  "guide_type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "content",
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
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "guide_jump",
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
			GoFieldName:        "GuideJump",
			GoFieldType:        "string",
			JSONFieldName:      "guide_jump",
			ProtobufFieldName:  "guide_jump",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
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
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "updated_at",
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
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiPlayGuide) TableName() string {
	return "ai_play_guide"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiPlayGuide) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiPlayGuide) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiPlayGuide) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiPlayGuide) TableInfo() *TableInfo {
	return ai_play_guideTableInfo
}
