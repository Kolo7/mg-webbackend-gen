# mg-webbackend-gen

mg-webbackend-gen目的是开发一个代码生成工具。工具适用于数据库表，这里特指 mysql，表的特征是没有外键关联的配置表，但表之间可以有隐性的关联关系。字段类型主要是字符串和数字。表和所有的字段都有详细的备注，严格的类型定义，具体的数据库要求，以及特殊字段的处理在后面详细介绍。

### 安装
> GOPROXY=https://goproxy.io,direct go install github.com/kolo7/mg-webbackend-gen@latest

### 试用

```bash
mg-webbackend-gen --connstr="root:E2FU0O1b@tcp(0.0.0.0:3306)/aiinteract?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4" \
--database=aiinteract \
--generate-dao \
--generate-service \
--generate-controller \
--rest \
--controller=server \
--gorm \
--json-fmt=lower_camel \
--overwrite \
--out=./example \
--module="example"
```

### 设计

代码生成的最终标准是，能够直接在极速开发平台框架下的代码正常运行。并且满足以下几点功能实现：

- model 层：model 层生成和数据库直接相关的代码，主要是结构体，包含了和数据库相对应的名称、tag、tablename 方法，转 api 对象的方法。
- dao 层：包括文件名，包名的生成，基于*Dao类的新增方法，新增方法是数据库操作。包括了查询（列表、列表相关的 count、byID）、新增、修改、删除，方法实现上遵循gorm框架通用实现，全量更新、真删除，新增使用自增 id。
- service 层：service 层实现要考虑实现基础的列表、新增、修改、删除。所有的方法参数使用 api 对象，对象分为 req 和 resp。列表功能考虑分页参数。
- controller 层：controller 层负责进行 service 层的调用，以及http response 的返回处理，参数绑定。
- router 层：按照 crud 的命名习惯，对所有的controller层进行路由绑定。
- api层：api 层包含了和外部交互的数据对象，包括请求对象 req，返回对象resp，请求对象要进行基本的格式校验，格式检验是可选的、例如非空、非负值等。返回对象要支持 json 格式的格式化定义，一定要包含 json tag，字段序列化后是小驼峰命名。
- 测试代码：dao层包含基础测试代码，service 层不做测试，controller 层包含 httptest 基础测试，只需要测试返回结果是 200 即可。测试代码在 docker-compose的基础环境下可以正常运行。
- 接口文档：根据绑定的路由，以及controller的请求参数和响应返回生成一份 markdown 文档，文档包含接口基本介绍，路由 uri，method，请求参数表格，响应 json 代码块。### 数据库要求

**必须要求**
数据库的表必须是全小写，下划线分隔单词，只包含小写字母和下划线。
必须要有 ID 字段，bigint类型，字段自增，主键。
**特殊字段处理**
特殊字段都是可选的，如果存在则进行相应的特殊处理
- created_at：创建时间，datetime 类型，非空
- updated_at：更新时间，datetime 类型，非空
- created_by：创建人，varchar 类型，非空，默认空字符串
- updated_by：操作人，varchar 类型，非空，默认空字符串
- online_status：上下线状态，tinyint类型，非空，默认 0，0 下线，1 上线
- audit_status：审核状态，tinyint类型，非空，默认 0
