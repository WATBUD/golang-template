package authentication

import (
	. "goa.design/goa/v3/dsl"
)

const (
	firebaseIdToken = "firebaseIdToken"
)

var ServiceName = Service("authentication", func() {
	Description("認證")

	signIn("SignIn")
}).Name

func attributeFirebaseIdToken() {
	Attribute(firebaseIdToken, String, func() {
		Description("名稱")
		MinLength(100)
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U")
	})
}

func signIn(methodName string) {
	Method(methodName, func() {
		Meta("openapi:summary", "登入")
		Error("token_error")

		Payload(func() {
			attributeFirebaseIdToken()
			Required(firebaseIdToken)
		})

		HTTP(func() {
			POST("/signin")

			Response(func() {
				Code(StatusOK)
				Description("登入完成")
			})

			Response("token_error", func() {
				Code(StatusBadRequest)
				Description("token 錯誤")
			})
		})
	})
}
