package plugin

import (
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
	// 包内静态变量
	gormSpanKey = "__gorm_span"
)

// db.Use(&plugin.OpentracingPlugin{})

//OpentracingPlugin .tracing callback
type OpentracingPlugin struct{}

//Name .
func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

//Initialize .
func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func before(db *gorm.DB) {
	// 先从父级spans生成子span ---> 这里命名为gorm，但实际上可以自定义
	// 自己喜欢的operationName
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")

	// 利用db实例去传递span
	db.InstanceSet(gormSpanKey, span)

	return
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		// 不存在就直接抛弃掉
		return
	}

	// 断言进行类型转换
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	// <---- 一定一定一定要Finsih掉！！！
	defer span.Finish()

	// Error
	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}

	// sql --> 写法来源GORM V2的日志
	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	return
}
