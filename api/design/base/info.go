package base

import (
	. "goa.design/goa/v3/dsl"
	. "mai.today/api/design/dsl"
)

const (
	BaseInfos = "baseInfos"

	color  = "color"
	info   = "info"
	infoId = "infoId"
	logo   = "logo"
	name   = "name"
)

var BaseInfo = Type("BaseInfo", func() {
	Description("基地資訊")
	AttributeID()
	attributeBaseID()
	attributeColor()
	attributeLogo()
	attributeName()
	AttributeDeletedAt()
	Required(ID, baseID, color, logo, name)
	exampleBaseInfo()
})

var updateBaseInfoResult = CommandResult("updateBaseInfo", BaseInfo)

func AttributeBaseInfos() {
	Attribute(BaseInfos, ArrayOf(BaseInfo), func() {
		exampleBaseInfos()
	})
}

func attributeColor() {
	Attribute(color, String, func() {
		Description("主題颜色")
		MinLength(1)
		MaxLength(16)
		Example("#FF5733")
	})
}

func attributeInfo() {
	Attribute(info, BaseInfo, func() {
		Description("基地資訊")
	})
}

// func attributeInfoId() {
// 	Attribute(infoId, String, func() {
// 		Description("基地資訊識別碼")
// 		FormatID()
// 		ExampleID()
// 	})
// }

func attributeLogo() {
	Attribute(logo, String, func() {
		Description("標識圖片的 URL")
		MinLength(1)
		MaxLength(1024)
		Example("http://example.com/logo.png")
	})
}

func attributeName() {
	Attribute(name, String, func() {
		Description("名稱")
		MinLength(1)
		MaxLength(128)
		Example("基地 A")
	})
}

func exampleBaseInfo() {
	Example("ExampleBaseInfo", func() {
		Value(Val{
			baseID:    "EbZvovBKbFogJeRem5oj2S",
			color:     "#FF5733",
			ID:        "EbZvovBKbFogJeRem5oj2S",
			logo:      "http://example.com/logo.png",
			name:      "基地 A",
			DeletedAt: nil,
		})
	})
}

func exampleBaseInfos() {
	Example("ExampleBaseInfos", func() {
		Value([]Val{
			{
				baseID:    "EbZvovBKbFogJeRem5oj2S",
				color:     "#FF5733",
				ID:        "EbZvovBKbFogJeRem5oj2S",
				logo:      "http://example.com/logo.png",
				name:      "基地 A",
				DeletedAt: nil,
			},
			{
				baseID:    "EbZvovBKbFogJeRem5oj3S",
				color:     "#FF5733",
				ID:        "EbZvovBKbFogJeRem5oj3S",
				logo:      "http://example.com/logo.png",
				name:      "基地 B",
				DeletedAt: nil,
			},
			{
				baseID:    "EbZvovBKbFogJeRem5oj4S",
				color:     "#FF5733",
				ID:        "EbZvovBKbFogJeRem5oj4S",
				logo:      "http://example.com/logo.png",
				name:      "基地 C",
				DeletedAt: nil,
			},
		})
	})
}

func updateBaseInfo(methodName string) {
	const sum = "更新基地資訊"
	var resType = updateBaseInfoResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			AttributeID()

			attributeColor()
			attributeLogo()
			attributeName()

			Required(ID, color, logo, name)
		})

		Result(resType)
		HTTP(func() {
			PUT("/base/{id}/info")
			ResponseDefaults()
			ResponseBadRequest()
			ResponseNotFound()
			ResponsePermissionDenied()
		})
	})
}
