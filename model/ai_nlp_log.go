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


Table: ai_nlp_log
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] msg_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
[ 2] req_content                                    text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 3] ai_ret                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 4] ai_ret_err                                     text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 5] callback_code                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 6] callback_ret                                   text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[ 9] send_id                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[10] ai_uid                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 53,    "msg_id": 62,    "req_content": "nVOeyCQsFnGTTIxkguumbiWAI",    "ai_ret": 97,    "ai_ret_err": "yCmoXwrLhujNVIpWrvovHWbSQ",    "callback_code": 79,    "callback_ret": "bsBjeBGqIPGjPWHLnwOlqgXgM",    "created_at": "2061-10-23T18:18:01.735381226+08:00",    "updated_at": "2269-09-28T08:18:45.463621251+08:00",    "send_id": "yoJhMNewDYLMosrDBIvkVMWop",    "ai_uid": "KAXWDqYieOwTKwHyfPdfRsxBk"}



*/

// AiNlpLog struct is a row record of the ai_nlp_log table in the  database
/*
type AiNlpLog struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] msg_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
    MsgID int64 `gorm:"column:msg_id;type:bigint;default:0;" json:"msg_id" xml:"msg_id" db:"msg_id" protobuf:"int64,1,opt,name=msg_id"`
    //[ 2] req_content                                    text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    ReqContent string `gorm:"column:req_content;type:text;" json:"req_content" xml:"req_content" db:"req_content" protobuf:"string,2,opt,name=req_content"`
    //[ 3] ai_ret                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    AiRet int32 `gorm:"column:ai_ret;type:tinyint;default:0;" json:"ai_ret" xml:"ai_ret" db:"ai_ret" protobuf:"int32,3,opt,name=ai_ret"`
    //[ 4] ai_ret_err                                     text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    AiRetErr string `gorm:"column:ai_ret_err;type:text;" json:"ai_ret_err" xml:"ai_ret_err" db:"ai_ret_err" protobuf:"string,4,opt,name=ai_ret_err"`
    //[ 5] callback_code                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    CallbackCode int32 `gorm:"column:callback_code;type:int;default:0;" json:"callback_code" xml:"callback_code" db:"callback_code" protobuf:"int32,5,opt,name=callback_code"`
    //[ 6] callback_ret                                   text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    CallbackRet string `gorm:"column:callback_ret;type:text;" json:"callback_ret" xml:"callback_ret" db:"callback_ret" protobuf:"string,6,opt,name=callback_ret"`
    //[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,7,opt,name=created_at"`
    //[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-01-01 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,8,opt,name=updated_at"`
    //[ 9] send_id                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    SendID string `gorm:"column:send_id;type:varchar;" json:"send_id" xml:"send_id" db:"send_id" protobuf:"string,9,opt,name=send_id"`
    //[10] ai_uid                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    AiUID string `gorm:"column:ai_uid;type:varchar;" json:"ai_uid" xml:"ai_uid" db:"ai_uid" protobuf:"string,10,opt,name=ai_uid"`

}
*/

var ai_nlp_logTableInfo = &TableInfo{
	Name: "ai_nlp_log",
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
			Name:               "msg_id",
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
			GoFieldName:        "MsgID",
			GoFieldType:        "int64",
			JSONFieldName:      "msg_id",
			ProtobufFieldName:  "msg_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "req_content",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "ReqContent",
			GoFieldType:        "string",
			JSONFieldName:      "req_content",
			ProtobufFieldName:  "req_content",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "ai_ret",
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
			GoFieldName:        "AiRet",
			GoFieldType:        "int32",
			JSONFieldName:      "ai_ret",
			ProtobufFieldName:  "ai_ret",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "ai_ret_err",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "AiRetErr",
			GoFieldType:        "string",
			JSONFieldName:      "ai_ret_err",
			ProtobufFieldName:  "ai_ret_err",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "callback_code",
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
			GoFieldName:        "CallbackCode",
			GoFieldType:        "int32",
			JSONFieldName:      "callback_code",
			ProtobufFieldName:  "callback_code",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "callback_ret",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "CallbackRet",
			GoFieldType:        "string",
			JSONFieldName:      "callback_ret",
			ProtobufFieldName:  "callback_ret",
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

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "send_id",
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
			GoFieldName:        "SendID",
			GoFieldType:        "string",
			JSONFieldName:      "send_id",
			ProtobufFieldName:  "send_id",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "ai_uid",
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
			GoFieldName:        "AiUID",
			GoFieldType:        "string",
			JSONFieldName:      "ai_uid",
			ProtobufFieldName:  "ai_uid",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiNlpLog) TableName() string {
	return "ai_nlp_log"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiNlpLog) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiNlpLog) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiNlpLog) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiNlpLog) TableInfo() *TableInfo {
	return ai_nlp_logTableInfo
}
