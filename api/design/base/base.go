package base

import (
	. "goa.design/goa/v3/dsl"
	. "mai.today/api/design/dsl"
)

const (
	baseID = "baseId"
)

var ServiceName = Service("Base", func() {
	Description("基地")
	ErrorInvalidToken()
	ErrorNotFound()
	ErrorPermissionDenied()

	createBase("CreateBase")
	deleteBase("DeleteBase")

	updateBaseInfo("UpdateBaseInfo")
	reorderBaseNavStates("ReorderBaseNavStates")
}).Name

var createBaseResult = CommandResult("createBase", createBaseResultData)

var createBaseResultData = Type("CreateBaseResultData", func() {
	AttributeID()
	attributeInfo()
	attributeNavState()
	Required(ID, info, navState)
})

var deleteBaseResult = CommandResult("deleteBase", deleteBaseResultData)

var deleteBaseResultData = Type("DeleteBaseResultData", func() {
	attributeBaseID()
	Required(baseID)
})

func attributeBaseID() {
	Attribute(baseID, String, func() {
		Description("基地識別碼")
		FormatID()
		ExampleID()
	})
}

func createBase(methodName string) {
	const sum = "創造一個新的基地"
	var resType = createBaseResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()

			attributeColor()
			attributeIndex()
			attributeLogo()
			attributeName()

			Required(color, logo, name)
		})
		Result(resType)
		HTTP(func() {
			POST("/base")
			ResponseDefaults()
			ResponseBadRequest()
		})
	})
}

func deleteBase(methodName string) {
	const sum = "删除指定的基地"
	var resType = deleteBaseResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			AttributeID()

			Required(ID)
		})
		Result(resType)
		HTTP(func() {
			DELETE("/base/{id}")
			ResponseDefaults()
			ResponseNotFound()
			ResponsePermissionDenied()
		})
	})
}
