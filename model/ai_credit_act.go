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


Table: ai_credit_act
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] act_name                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 2] start_time                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[ 3] end_time                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[ 4] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 5] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 6] update_info                                    text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 7] update_info_status                             tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 8] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 9] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 19,    "act_name": "GSBsyMCPQAYtbZNhnZTJtBZXN",    "start_time": "2223-07-26T19:33:45.704739899+08:00",    "end_time": "2194-04-15T13:09:42.985069339+08:00",    "audit_status": 82,    "online_status": 41,    "update_info": "xVPLnYGpoFKRMutRsFwYFNDpN",    "update_info_status": 53,    "creater": "NfgqSsQvBQfWZRuxITIpySeSU",    "upadter": "dNKFOSXkjyifrMijQvMNlAUUK",    "created_at": "2262-05-04T09:59:42.943370936+08:00",    "updated_at": "2206-10-02T15:50:44.238392456+08:00"}



*/

// AiCreditAct struct is a row record of the ai_credit_act table in the  database
/*
type AiCreditAct struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] act_name                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    ActName string `gorm:"column:act_name;type:varchar;" json:"act_name" xml:"act_name" db:"act_name" protobuf:"string,1,opt,name=act_name"`
    //[ 2] start_time                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    StartTime time.Time `gorm:"column:start_time;type:datetime;default:2023-01-01 00:00:00;" json:"start_time" xml:"start_time" db:"start_time" protobuf:"google.protobuf.Timestamp,2,opt,name=start_time"`
    //[ 3] end_time                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    EndTime time.Time `gorm:"column:end_time;type:datetime;default:2023-01-01 00:00:00;" json:"end_time" xml:"end_time" db:"end_time" protobuf:"google.protobuf.Timestamp,3,opt,name=end_time"`
    //[ 4] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    AuditStatus int32 `gorm:"column:audit_status;type:tinyint;default:0;" json:"audit_status" xml:"audit_status" db:"audit_status" protobuf:"int32,4,opt,name=audit_status"`
    //[ 5] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,5,opt,name=online_status"`
    //[ 6] update_info                                    text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
    UpdateInfo string `gorm:"column:update_info;type:text;" json:"update_info" xml:"update_info" db:"update_info" protobuf:"string,6,opt,name=update_info"`
    //[ 7] update_info_status                             tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    UpdateInfoStatus int32 `gorm:"column:update_info_status;type:tinyint;default:0;" json:"update_info_status" xml:"update_info_status" db:"update_info_status" protobuf:"int32,7,opt,name=update_info_status"`
    //[ 8] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,8,opt,name=creater"`
    //[ 9] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,9,opt,name=upadter"`
    //[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,10,opt,name=created_at"`
    //[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-01-01 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,11,opt,name=updated_at"`

}
*/

var ai_credit_actTableInfo = &TableInfo{
	Name: "ai_credit_act",
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
			Name:               "act_name",
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
			GoFieldName:        "ActName",
			GoFieldType:        "string",
			JSONFieldName:      "act_name",
			ProtobufFieldName:  "act_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "start_time",
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
			GoFieldName:        "StartTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "end_time",
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
			GoFieldName:        "EndTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "audit_status",
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
			GoFieldName:        "AuditStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "audit_status",
			ProtobufFieldName:  "audit_status",
			ProtobufType:       "int32",
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
			Name:               "update_info",
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
			GoFieldName:        "UpdateInfo",
			GoFieldType:        "string",
			JSONFieldName:      "update_info",
			ProtobufFieldName:  "update_info",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "update_info_status",
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
			GoFieldName:        "UpdateInfoStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "update_info_status",
			ProtobufFieldName:  "update_info_status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},

		{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiCreditAct) TableName() string {
	return "ai_credit_act"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiCreditAct) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiCreditAct) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiCreditAct) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiCreditAct) TableInfo() *TableInfo {
	return ai_credit_actTableInfo
}
