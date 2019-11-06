package server

import (
	"context"
	"log"
	"reflect"
	"unicode"
	"unicode/utf8"
)

// Precompute the reflect type for error. Can't use error directly
// because Typeof takes an empty interface value. This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

func validateArgs(mtype reflect.Type, reportErr bool, mname string) (argType reflect.Type, ok bool) {
	// 需要有四个参数: receiver, Context, args, *reply.
	if mtype.NumIn() != 4 {
		if reportErr {
			log.Println("method", mname, "has wrong number of ins:", mtype.NumIn())
		}
		return nil, false
	}

	// 第一个参数必须是context.Context
	ctxType := mtype.In(1)
	if !ctxType.Implements(typeOfContext) {
		if reportErr {
			log.Println("method", mname, " must use context.Context as the first parameter")
		}
		return nil, false
	}

	// 第二个参数arg
	argType = mtype.In(2)
	if !isExportedOrBuiltinType(argType) {
		if reportErr {
			log.Println(mname, "parameter type not exported:", argType)
		}
		return nil, false
	}

	// 第三个参数是返回值，必须是指针类型的
	replyType := mtype.In(3)
	if replyType.Kind() != reflect.Ptr {
		if reportErr {
			log.Println("method", mname, "reply type not a pointer:", replyType)
		}
		return nil, false
	}

	// 返回值的类型必须是可导出的
	if !isExportedOrBuiltinType(replyType) {
		if reportErr {
			log.Println("method", mname, "reply type not exported:", replyType)
		}
		return nil, false
	}

	return argType, true
}

func validateReply(mtype reflect.Type, reportErr bool, mname string) (returnType reflect.Type, ok bool) {
	// 必须有一个返回值
	if mtype.NumOut() != 1 {
		if reportErr {
			log.Println("method", mname, "has wrong number of outs:", mtype.NumOut())
		}
		return nil, false
	}

	// 返回值类型必须是error
	if returnType = mtype.Out(0); returnType != typeOfError {
		if reportErr {
			log.Println("method", mname, "returns", returnType.String(), "not error")
		}
		return nil, false
	}

	return returnType, true
}

type Method map[string]*methodType

//过滤符合规则的方法，从net.rpc包抄的
func suitableMethods(typ reflect.Type, reportErr bool) Method {
	methods := make(Method)

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name

		// 方法必须是可导出的
		if method.PkgPath != "" {
			continue
		}

		var (
			argType   reflect.Type
			replyType reflect.Type
			ok        bool
		)
		// 验证参数合法性
		if argType, ok = validateArgs(mtype, reportErr, mname); !ok {
			continue
		}

		// 验证返回值
		if replyType, ok = validateReply(mtype, reportErr, mname); !ok {
			continue
		}

		methods[mname] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
	}

	return methods
}

func newValue(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()).Interface()
	}

	return reflect.New(t).Interface()
}
