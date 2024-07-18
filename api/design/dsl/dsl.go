package dsl

import (
	"strings"

	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

const (
	DeletedAt        = "deleted_at"
	ID               = "id"
	InvalidToken     = "invalid token"
	JWT              = "jwt"
	NotFound         = "not found"
	PermissionDenied = "permission denied"
	Timestamp        = "timestamp"
	UserIDChannel    = "channel"

	command           = "command"
	commandResultData = "data"
	commandType       = "type"
)

var commandEnum = Type("Command", func() {
	Description("指令")
	Attribute(commandType, String, func() {
		Description("類型")
		Enum(
			"sync",
			"createBase",
			"deleteBase",
			"reorderBaseNavState",
			"updateBaseInfo",
		)
	})
	Required(commandType)
})

func AsyncMethod(name, sum string, resType any, args ...any) {
	var n = "receive" + strings.ToUpper(name[:1]) + name[1:]
	var p = "/{channel}/" + strings.ToLower(name[:1]) + name[1:]
	Method(n, func() {
		MetaSummary("[WebSocket] " + sum)
		Payload(func() {
			AttributeJWT()
			if len(args) == 0 {
				AttributeUserIDChannel()
				Required(UserIDChannel)
			} else {
				if f, ok := args[len(args)-1].(func()); ok {
					f()
				}
			}
		})
		Result(resType)
		HTTP(func() {
			GET(p)
			ResponseSwitchingProtocols()
			ResponseOK()
		})
	})
}

func AttributeDateTime(name, desc string) {
	Attribute(name, String, func() {
		Description(desc)
		Format(FormatDateTime)
		Example("2024-01-01T11:53:54Z")
	})
}

func AttributeDeletedAt() {
	AttributeDateTime(DeletedAt, "刪除時間")
}

func AttributeID() {
	Attribute(ID, String, func() {
		Description("識別碼")
		FormatID()
		ExampleID()
	})
}

func AttributeJWT() {
	Token(JWT, String)
}

func AttributeTimestamp() {
	Attribute(Timestamp, Int64, func() {
		Description("時間戳記")
		FormatTimestamp()
		ExampleTimestamp()
	})
}

func AttributeUserIDChannel() {
	Attribute(UserIDChannel, String, func() {
		Description("使用者識別碼")
		FormatID()
		ExampleID()
	})
}

func CommandResult(cmd string, dataType any, args ...any) expr.UserType {
	return Type(cmd+"Result", func() {
		Description("WebSocket 的通用資料負載")
		attributeCommand(cmd)
		AttributeTimestamp()
		Attribute(commandResultData, dataType, func() {
			Description("資料")
			if len(args) == 1 {
				if f, ok := args[len(args)-1].(func()); ok {
					f()
				}
			}
		})
		Required(command, Timestamp, commandResultData)
	})
}

func ErrorInvalidToken() {
	Error(InvalidToken, func() {
		Description("token 無效")
		Example(Val{"message": "token 過期"})
	})
}

func ErrorNotFound() {
	Error(NotFound, func() {
		Description("資源不存在")
		Example(Val{"message": "資源不存在"})
	})
}

func ErrorPermissionDenied() {
	Error(PermissionDenied, func() {
		Description("沒有權限")
		Example(Val{"message": "缺乏存取所請求資源的必要權限"})
	})
}

func ExampleID() {
	Example("669071b8a650a662b82285ca")
}

func ExampleTimestamp() {
	Example(1719188007000000000)
}

func FormatID() {
	MinLength(24)
	MaxLength(24)
}

func FormatTimestamp() {
	Minimum(1719188007000000000)
}

func MetaSummary(value string) {
	Meta("openapi:summary", value)
}

func ResponseBadRequest() {
	Response(func() {
		Code(StatusBadRequest)
		Description("參數錯誤")
		Body(Empty)
	})
}

func ResponseConflict(description string) {
	Response(func() {
		Code(StatusConflict)
		Description(description)
		Body(Empty)
	})
}

// ResponseDefaults defines commonly used response functions for convenience
// This function should be called as the first response in an HTTP handler
func ResponseDefaults() {
	// ResponseOK() should be called first within ResponseDefaults
	ResponseOK()

	ResponseInvalidToken()
	ResponseInternalServerError()
}

func ResponseInternalServerError() {
	Response(func() {
		Code(StatusInternalServerError)
		Description("伺服器發生錯誤")
		Body(Empty)
	})
}

func ResponseNotFound() {
	Response(NotFound, func() {
		Code(StatusNotFound)
	})
}

func ResponseOK() {
	Response(func() {
		Code(StatusOK)
		Description("操作成功")
	})
}

func ResponsePermissionDenied() {
	Response(PermissionDenied, func() {
		Code(StatusForbidden)
	})
}

func ResponseSwitchingProtocols() {
	Response(func() {
		Description("此 API 僅用於 WebSocket 協議")
		Code(StatusSwitchingProtocols)
		Body(Empty)
	})
}

func ResponseInvalidToken() {
	Response(InvalidToken, func() {
		Code(StatusUnauthorized)
	})
}

func attributeCommand(ex string) {
	Attribute(command, commandEnum, func() {
		Example(func() {
			Value(Val{commandType: ex})
		})
	})
}
