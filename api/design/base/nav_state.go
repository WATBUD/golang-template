package base

import (
	. "goa.design/goa/v3/dsl"
	. "mai.today/api/design/dsl"
)

const (
	BaseNavStates = "baseNavStates"

	index      = "index"
	navState   = "navState"
	navStateID = "navStateId"
	newIndex   = "newIndex"
)

var baseNavState = Type("BaseNavState", func() {
	Description("基地導覽列狀態")
	AttributeID()
	attributeBaseID()
	attributeIndex()
	AttributeDeletedAt()
	Required(ID, baseID, index)
})

var reorderBaseNavStateResult = CommandResult("reorderBaseNavState", ArrayOf(baseNavState), func() {
	exampleBaseNavStates()
})

func AttributeBaseNavState() {
	Attribute(BaseNavStates, ArrayOf(baseNavState), func() {
		exampleBaseNavStates()
	})
}

func attributeIndex() {
	Attribute(index, Int, func() {
		Description("順序")
		Minimum(1)
		Default(1)
		Example(0)
	})
}

func attributeNavState() {
	Attribute(navState, baseNavState, func() {
		Description("基地導覽列狀態")
	})
}

// func attributeNavStateID() {
// 	Attribute(navStateID, String, func() {
// 		Description("基地導覽列狀態識別碼")
// 		FormatID()
// 		ExampleID()
// 	})
// }

func attributeNewIndex() {
	Attribute(newIndex, Int, func() {
		Description("新的順序")
		Minimum(1)
		Default(1)
		Example(0)
	})
}

func exampleBaseNavStates() {
	Example("ExampleBaseNavStates", func() {
		Value([]Val{
			{
				baseID:    "EbZvovBKbFogJeRem5oj2S",
				ID:        "EbZvovBKbFogJeRem5oj7a",
				index:     0,
				DeletedAt: nil,
			},
			{
				baseID:    "EbZvovBKbFogJeRem5oj3S",
				ID:        "EbZvovBKbFogJeRem5oj7b",
				index:     1,
				DeletedAt: nil,
			},
			{
				baseID:    "EbZvovBKbFogJeRem5oj4S",
				ID:        "EbZvovBKbFogJeRem5oj7c",
				index:     2,
				DeletedAt: nil,
			},
		})
	})
}

func reorderBaseNavStates(methodName string) {
	const sum = "變更基地的順序"
	var resType = reorderBaseNavStateResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			AttributeID()
			attributeNewIndex()
			Required(ID, newIndex)
		})

		Result(resType)
		HTTP(func() {
			PUT("/base/nav-state/reorder")
			ResponseDefaults()
			ResponseBadRequest()
		})
	})
}
