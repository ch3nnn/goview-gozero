package main

import (
	"strings"

	"github.com/ch3nnn/goview-gozero/pkg/config"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// SqlDsn TODO 数据库连接地址
const (
	Host              = "localhost"
	Port              = 3306
	User              = "root"
	Password          = "123456"
	DBName            = "goview"
	OutPath           = "../service/screen/internal/dal" // TODO 相对执行`go run`时的路径, 会自动创建目录
	GenerateTableName = "screen_project, screen_datum"   // TODO 要生成的表名

	Type = "mysql" // 数据库类型 postgres mysql
)

func main() {
	databaseConf := config.DatabaseConf{
		Host:     Host,
		Port:     Port,
		Username: User,
		Password: Password,
		DBName:   DBName,
		SSLMode:  "disable",
		Type:     Type,
	}
	db := databaseConf.NewDriver()

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 模型包名路径
		OutPath:      OutPath + "/query",
		ModelPkgPath: OutPath + "/model",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
		// // 生成单元测试
		// WithUnitTest: true,
	})
	// 设置目标 db
	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64, 兼容protobuf
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(
		map[string]func(columnType gorm.ColumnType) (dataType string){
			"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int2":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int4":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int8":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int32" }, // mysql 约定适应 smallint 等价于 int32, protobuf enum 为 int32
			"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"tinyint": func(columnType gorm.ColumnType) (dataType string) {
				// mysql tinyint(1) 类型 改为 bool
				ct, _ := columnType.ColumnType()
				if strings.HasPrefix(ct, "tinyint(1)") {
					return "bool"
				}
				return "int64"
			},
			// 统一日期类型为 sql.NullTime
			"datetime": func(columnType gorm.ColumnType) (dataType string) { return "*time.Time" },
			"json":     func(columnType gorm.ColumnType) (dataType string) { return "datatypes.JSON" },
		})

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	// jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `balance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	// })
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("update_at", func(tag field.GormTag) field.GormTag {
		return field.GormTag{"column": []string{"update_at"}, "type": []string{"int unsigned"}, "": []string{"autoUpdateTime"}}
	})
	autoCreateTimeField := gen.FieldGORMTag("create_at", func(tag field.GormTag) field.GormTag {
		return field.GormTag{"column": []string{"create_at"}, "type": []string{"int unsigned"}, "": []string{"autoCreateTime"}}
	})
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	softDeleteField := gen.FieldType("delete_at", "gorm.DeletedAt")

	softIsDeleteField := gen.FieldType("is_del", "soft_delete.DeletedAt")
	softIsDeleteTag := gen.FieldGORMTag("is_del", func(tag field.GormTag) field.GormTag {
		return field.GormTag{"column": []string{"is_del"}, "": []string{"softDelete:flag"}}
	})

	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField, softIsDeleteField, softIsDeleteTag}

	// 模型自定义选项组
	// fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField}
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	// allModel := g.GenerateAllTable(fieldOpts...)
	// g.ApplyBasic(allModel...)

	// 根据 tableName 生成对应 model
	var models []interface{}
	for _, tableName := range strings.Split(GenerateTableName, ",") {
		models = append(models, g.GenerateModel(strings.TrimSpace(tableName), fieldOpts...))
	}

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(models...)

	g.Execute()
}