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


Table: ai_user_credit
[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
[ 1] user_uid                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 2] ai_uid                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 3] act_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
[ 4] user_credit                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]


JSON Sample
-------------------------------------
{    "id": 98,    "user_uid": "BkqBsqcXkBDfDaBGAanamFsKR",    "ai_uid": "bimORjwMLJnVxFWWRmGDRWiof",    "act_id": 0,    "user_credit": 76,    "creater": "PuKoxidlbETxbStDAZYxvxBUI",    "upadter": "yVLPvZpaPipGjkmaxYlmmZwjC",    "created_at": "2200-11-20T14:34:45.41693282+08:00",    "updated_at": "2140-05-08T06:26:39.079335626+08:00"}



*/

// AiUserCredit struct is a row record of the ai_user_credit table in the  database
/*
type AiUserCredit struct {
    //[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
    ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id" xml:"id" db:"id" protobuf:"int64,0,opt,name=id"`
    //[ 1] user_uid                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    UserUID string `gorm:"column:user_uid;type:varchar;" json:"user_uid" xml:"user_uid" db:"user_uid" protobuf:"string,1,opt,name=user_uid"`
    //[ 2] ai_uid                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    AiUID string `gorm:"column:ai_uid;type:varchar;" json:"ai_uid" xml:"ai_uid" db:"ai_uid" protobuf:"string,2,opt,name=ai_uid"`
    //[ 3] act_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
    ActID int64 `gorm:"column:act_id;type:bigint;" json:"act_id" xml:"act_id" db:"act_id" protobuf:"int64,3,opt,name=act_id"`
    //[ 4] user_credit                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
    UserCredit int32 `gorm:"column:user_credit;type:int;" json:"user_credit" xml:"user_credit" db:"user_credit" protobuf:"int32,4,opt,name=user_credit"`
    //[ 5] creater                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Creater string `gorm:"column:creater;type:varchar;" json:"creater" xml:"creater" db:"creater" protobuf:"string,5,opt,name=creater"`
    //[ 6] upadter                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
    Upadter string `gorm:"column:upadter;type:varchar;" json:"upadter" xml:"upadter" db:"upadter" protobuf:"string,6,opt,name=upadter"`
    //[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:2023-01-01 00:00:00;" json:"created_at" xml:"created_at" db:"created_at" protobuf:"google.protobuf.Timestamp,7,opt,name=created_at"`
    //[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [2023-01-01 00:00:00]
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:2023-01-01 00:00:00;" json:"updated_at" xml:"updated_at" db:"updated_at" protobuf:"google.protobuf.Timestamp,8,opt,name=updated_at"`

}
*/

var ai_user_creditTableInfo = &TableInfo{
	Name: "ai_user_credit",
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
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "act_id",
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
			GoFieldName:        "ActID",
			GoFieldType:        "int64",
			JSONFieldName:      "act_id",
			ProtobufFieldName:  "act_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "user_credit",
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
			GoFieldName:        "UserCredit",
			GoFieldType:        "int32",
			JSONFieldName:      "user_credit",
			ProtobufFieldName:  "user_credit",
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
	},
}

// TableName sets the insert table name for this struct type
func (a *AiUserCredit) TableName() string {
	return "ai_user_credit"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AiUserCredit) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AiUserCredit) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AiUserCredit) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AiUserCredit) TableInfo() *TableInfo {
	return ai_user_creditTableInfo
}
