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


Table: ai_credit_gift
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] gift_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
[ 2] gift_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] gift_img                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] credit_level                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 5] gift_url                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 6] gift_tips                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] is_end                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 8] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 9] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[10] update_info                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[11] update_info_status                             tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[12] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[13] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[14] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[15] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 66,    "gift_id": 76,    "gift_name": "KPrqxKOvxeBRhlkWlxylipkki",    "gift_img": "PSvGfMfBmhbKAGJiVXchBVQWI",    "credit_level": 24,    "gift_url": "AtDKvObgFrsuCLIGMpdLPtmqo",    "gift_tips": "AxEToVlStHZesTOXYyOPFqZdc",    "is_end": 70,    "online_status": 46,    "audit_status": 93,    "update_info": "YkQlKfCtvnDcPuaTmZSitmqnW",    "update_info_status": 35,    "creater": "sGyQCiGCvVvrdqbJsrTNRYLEx",    "upadter": "bukODCxLnhIWjvykDiZihSUkc",    "created_at": "2300-02-07T03:59:01.794741123+08:00",    "updated_at": "2161-03-01T22:59:35.970550747+08:00"}



*/

// AiCreditGift struct is a row record of the ai_credit_gift table in the  database
/*
type AiCreditGift struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] gift_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
    GiftID int32 `gorm:"column:gift_id;type:int;default:0;" json:"gift_id" xml:"gift_id" db:"gift_id" protobuf:"int32,1,opt,name=gift_id"`
    //[ 2] gift_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    GiftName string `gorm:"column:gift_name;type:varchar;" json:"gift_name" xml:"gift_name" db:"gift_name" protobuf:"string,2,opt,name=gift_name"`
    //[ 3] gift_img                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    GiftImg string `gorm:"column:gift_img;type:varchar;" json:"gift_img" xml:"gift_img" db:"gift_img" protobuf:"string,3,opt,name=gift_img"`
    //[ 4] credit_level                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    CreditLevel int32 `gorm:"column:credit_level;type:tinyint;default:0;" json:"credit_level" xml:"credit_level" db:"credit_level" protobuf:"int32,4,opt,name=credit_level"`
    //[ 5] gift_url                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    GiftURL string `gorm:"column:gift_url;type:varchar;" json:"gift_url" xml:"gift_url" db:"gift_url" protobuf:"string,5,opt,name=gift_url"`
    //[ 6] gift_tips                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    GiftTips string `gorm:"column:gift_tips;type:varchar;" json:"gift_tips" xml:"gift_tips" db:"gift_tips" protobuf:"string,6,opt,name=gift_tips"`
    //[ 7] is_end                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    IsEnd int32 `gorm:"column:is_end;type:tinyint;default:0;" json:"is_end" xml:"is_end" db:"is_end" protobuf:"int32,7,opt,name=is_end"`
    //[ 8] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,8,opt,name=online_status"`
    //[ 9] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    AuditStatus int32 `gorm:"column:audit_status;type:tinyint;default:0;" json:"audit_status" xml:"audit_status" db:"audit_status" protobuf:"int32,9,opt,name=audit_status"`
    //[10] update_info                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UpdateInfo string `gorm:"column:update_info;type:varchar;" json:"update_info" xml:"update_info" db:"update_info" protobuf:"string,10,opt,name=update_info"`
    //[11] update_info_status                             tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    UpdateInfoStatus int32 `gorm:"column:update_info_status;type:tinyint;default:0;" json:"update_info_status" xml:"update_info_status" db:"update_info_status" protobuf:"int32,11,opt,name=update_info_status"`
    //[12] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,12,opt,name=creater"`
    //[13] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,13,opt,name=upadter"`
    //[14] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,14,opt,name=created_at"`
    //[15] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-01-01 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,15,opt,name=updated_at"`

}
*/

var ai_credit_giftTableInfo = &TableInfo{
	Name: "ai_credit_gift",
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
			Name:               "gift_id",
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
			GoFieldName:        "GiftID",
			GoFieldType:        "int32",
			JSONFieldName:      "gift_id",
			ProtobufFieldName:  "gift_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "gift_name",
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
			GoFieldName:        "GiftName",
			GoFieldType:        "string",
			JSONFieldName:      "gift_name",
			ProtobufFieldName:  "gift_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "gift_img",
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
			GoFieldName:        "GiftImg",
			GoFieldType:        "string",
			JSONFieldName:      "gift_img",
			ProtobufFieldName:  "gift_img",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
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
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "gift_url",
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
			GoFieldName:        "GiftURL",
			GoFieldType:        "string",
			JSONFieldName:      "gift_url",
			ProtobufFieldName:  "gift_url",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "gift_tips",
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
			GoFieldName:        "GiftTips",
			GoFieldType:        "string",
			JSONFieldName:      "gift_tips",
			ProtobufFieldName:  "gift_tips",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "is_end",
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
			GoFieldName:        "IsEnd",
			GoFieldType:        "int32",
			JSONFieldName:      "is_end",
			ProtobufFieldName:  "is_end",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "update_info",
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
			GoFieldName:        "UpdateInfo",
			GoFieldType:        "string",
			JSONFieldName:      "update_info",
			ProtobufFieldName:  "update_info",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
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
			ProtobufPos:        12,
		},

		{
			Index:              12,
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
			ProtobufPos:        13,
		},

		{
			Index:              13,
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
			ProtobufPos:        14,
		},

		{
			Index:              14,
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
			ProtobufPos:        15,
		},

		{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AiCreditGift) TableName() string {
	return "ai_credit_gift"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiCreditGift) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiCreditGift) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiCreditGift) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiCreditGift) TableInfo() *TableInfo {
	return ai_credit_giftTableInfo
}
