package authentication

import (
	. "goa.design/goa/v3/dsl"
	. "mai.today/api/design/dsl"
)

const (
	firebaseIDToken = "firebaseIdToken"
)

var ServiceName = Service("Authentication", func() {
	Description("認證")
	ErrorInvalidToken()
	signIn("SignIn")
}).Name

func attributeFirebaseIDToken() {
	Attribute(firebaseIDToken, String, func() {
		Description("Firebase Id Token")
		MinLength(100)
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U")
	})
}

func signIn(methodName string) {
	Method(methodName, func() {
		Meta("openapi:summary", "登入")
		Error("token_error")
		NoSecurity()

		Payload(func() {
			attributeFirebaseIDToken()
			Required(firebaseIDToken)
		})

		HTTP(func() {
			POST("/signin")

			Response(func() {
				Code(StatusOK)
				Description("登入完成")
				Body(Empty)
			})

			Response("token_error", func() {
				Code(StatusBadRequest)
				Description("token 錯誤")
			})
		})
	})
}
