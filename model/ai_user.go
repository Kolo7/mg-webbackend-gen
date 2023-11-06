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


Table: ai_user
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] user_uid                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 2] user_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] cover_img                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 4] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 5] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 6] update_info                                    text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
[ 7] update_info_status                             tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
[ 8] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 9] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 12,    "user_uid": "MZcEoilXkepFEuJrjnwZYjvDJ",    "user_name": "eKeALXQONAhcdtlswxFjtKVFc",    "cover_img": "NPEjxQGakkPwPAYLityivaiCq",    "online_status": 97,    "audit_status": 61,    "update_info": "TlQMndiMRVyaewmvEyTFPIQLu",    "update_info_status": 61,    "creater": "vTIvSSIDIBaAybXYeFsHwiJbZ",    "upadter": "NydUDVSnDqZaLXuQfIuQGBlrs",    "created_at": "2151-04-26T07:59:41.002105952+08:00",    "updated_at": "2297-09-07T19:16:30.546090777+08:00"}



*/

// AiUser struct is a row record of the ai_user table in the  database
/*
type AiUser struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] user_uid                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UserUID string `gorm:"column:user_uid;type:varchar;" json:"user_uid" xml:"user_uid" db:"user_uid" protobuf:"string,1,opt,name=user_uid"`
    //[ 2] user_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UserName string `gorm:"column:user_name;type:varchar;" json:"user_name" xml:"user_name" db:"user_name" protobuf:"string,2,opt,name=user_name"`
    //[ 3] cover_img                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    CoverImg string `gorm:"column:cover_img;type:varchar;" json:"cover_img" xml:"cover_img" db:"cover_img" protobuf:"string,3,opt,name=cover_img"`
    //[ 4] online_status                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    OnlineStatus int32 `gorm:"column:online_status;type:tinyint;default:0;" json:"online_status" xml:"online_status" db:"online_status" protobuf:"int32,4,opt,name=online_status"`
    //[ 5] audit_status                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
    AuditStatus int32 `gorm:"column:audit_status;type:tinyint;default:0;" json:"audit_status" xml:"audit_status" db:"audit_status" protobuf:"int32,5,opt,name=audit_status"`
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

var ai_userTableInfo = &TableInfo{
	Name: "ai_user",
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
			Name:               "user_uid",
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
			GoFieldName:        "UserUID",
			GoFieldType:        "string",
			JSONFieldName:      "user_uid",
			ProtobufFieldName:  "user_uid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "user_name",
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
			GoFieldName:        "UserName",
			GoFieldType:        "string",
			JSONFieldName:      "user_name",
			ProtobufFieldName:  "user_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "cover_img",
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
			GoFieldName:        "CoverImg",
			GoFieldType:        "string",
			JSONFieldName:      "cover_img",
			ProtobufFieldName:  "cover_img",
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
func (a *AiUser) TableName() string {
	return "ai_user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiUser) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiUser) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiUser) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiUser) TableInfo() *TableInfo {
	return ai_userTableInfo
}
