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


Table: chat_keyword_material
[ 0] id                                             unsigned bigint      null: false  primary: true   isArray: false  auto: false  col: unsigned bigint len: -1      default: []
[ 1] uuids                                          text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 2] match_keywords                                 text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 3] material_type                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 4] material_text                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 5] time_length                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 6] audio_url                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] image_height                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 8] image_width                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 9] image_url                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[10] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[11] create_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[12] update_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[13] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[14] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": "SGaxPFPjovRKGwDqsqMDrXSWd",    "uuids": "eOLDBysAivJnynbPbYVeWlbhY",    "match_keywords": 77,    "material_type": "TdvPGJlIgESGCqtBrvyopvvRn",    "material_text": 63,    "time_length": "hLpmHmfBvogSoWXgbebqZTFCt",    "audio_url": 45,    "image_height": 26,    "image_width": "bSDeedVJKQFIBehGGrlHWeGdP",    "image_url": 52,    "online_status": "NZvywfcICZwnMUThhNkpcBhrS",    "create_by": "mfKJoGLrKOiLiYNPVCdUAJRlD",    "update_by": "2238-06-19T03:25:16.419154811+08:00",    "created_at": "2219-05-22T10:50:24.564374601+08:00"}



*/

// ChatKeywordMaterial struct is a row record of the chat_keyword_material table in the  database
/*
type ChatKeywordMaterial struct {
    //[ 1] uuids                                          text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    Uuids sql.NullString `gorm:"column:uuids;type:text;" json:"uuids" xml:"uuids" db:"uuids" protobuf:"string,1,opt,name=uuids"`
    //[ 2] match_keywords                                 text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    MatchKeywords sql.NullString `gorm:"column:match_keywords;type:text;" json:"match_keywords" xml:"match_keywords" db:"match_keywords" protobuf:"string,2,opt,name=match_keywords"`
    //[ 3] material_type                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    MaterialType int32 `gorm:"column:material_type;type:tinyint;default:0;" json:"material_type" xml:"material_type" db:"material_type" protobuf:"int32,3,opt,name=material_type"`
    //[ 4] material_text                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    MaterialText string `gorm:"column:material_text;type:varchar;" json:"material_text" xml:"material_text" db:"material_text" protobuf:"string,4,opt,name=material_text"`
    //[ 5] time_length                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    TimeLength int32 `gorm:"column:time_length;type:int;default:0;" json:"time_length" xml:"time_length" db:"time_length" protobuf:"int32,5,opt,name=time_length"`
    //[ 6] audio_url                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    AudioURL string `gorm:"column:audio_url;type:varchar;" json:"audio_url" xml:"audio_url" db:"audio_url" protobuf:"string,6,opt,name=audio_url"`
    //[ 7] image_height                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    ImageHeight int32 `gorm:"column:image_height;type:int;default:0;" json:"image_height" xml:"image_height" db:"image_height" protobuf:"int32,7,opt,name=image_height"`
    //[ 8] image_width                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    ImageWidth int32 `gorm:"column:image_width;type:int;default:0;" json:"image_width" xml:"image_width" db:"image_width" protobuf:"int32,8,opt,name=image_width"`
    //[ 9] image_url                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    ImageURL string `gorm:"column:image_url;type:varchar;" json:"image_url" xml:"image_url" db:"image_url" protobuf:"string,9,opt,name=image_url"`
    //[10] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,10,opt,name=online_status"`
    //[11] create_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    CreateBy string `gorm:"column:create_by;type:varchar;" json:"create_by" xml:"create_by" db:"create_by" protobuf:"string,11,opt,name=create_by"`
    //[12] update_by                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UpdateBy string `gorm:"column:update_by;type:varchar;" json:"update_by" xml:"update_by" db:"update_by" protobuf:"string,12,opt,name=update_by"`
    //[13] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,13,opt,name=created_at"`
    //[14] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-01-01 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,14,opt,name=updated_at"`

}
*/

var chat_keyword_materialTableInfo = &TableInfo{
	Name: "chat_keyword_material",
	Columns: []*ColumnInfo{

		{
			Index:              1,
			Name:               "uuids",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Uuids",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "uuids",
			ProtobufFieldName:  "uuids",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "match_keywords",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "MatchKeywords",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "match_keywords",
			ProtobufFieldName:  "match_keywords",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "material_type",
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
			GoFieldName:        "MaterialType",
			GoFieldType:        "int32",
			JSONFieldName:      "material_type",
			ProtobufFieldName:  "material_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "material_text",
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
			GoFieldName:        "MaterialText",
			GoFieldType:        "string",
			JSONFieldName:      "material_text",
			ProtobufFieldName:  "material_text",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "time_length",
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
			GoFieldName:        "TimeLength",
			GoFieldType:        "int32",
			JSONFieldName:      "time_length",
			ProtobufFieldName:  "time_length",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "audio_url",
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
			GoFieldName:        "AudioURL",
			GoFieldType:        "string",
			JSONFieldName:      "audio_url",
			ProtobufFieldName:  "audio_url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "image_height",
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
			GoFieldName:        "ImageHeight",
			GoFieldType:        "int32",
			JSONFieldName:      "image_height",
			ProtobufFieldName:  "image_height",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "image_width",
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
			GoFieldName:        "ImageWidth",
			GoFieldType:        "int32",
			JSONFieldName:      "image_width",
			ProtobufFieldName:  "image_width",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "image_url",
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
			GoFieldName:        "ImageURL",
			GoFieldType:        "string",
			JSONFieldName:      "image_url",
			ProtobufFieldName:  "image_url",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},

		{
			Index:              11,
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
			ProtobufPos:        12,
		},

		{
			Index:              12,
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
			ProtobufPos:        13,
		},

		{
			Index:              13,
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
			ProtobufPos:        14,
		},

		{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ChatKeywordMaterial) TableName() string {
	return "chat_keyword_material"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ChatKeywordMaterial) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ChatKeywordMaterial) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ChatKeywordMaterial) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ChatKeywordMaterial) TableInfo() *TableInfo {
	return chat_keyword_materialTableInfo
}
