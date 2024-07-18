package sync

import (
	. "goa.design/goa/v3/dsl"
	"mai.today/api/design/base"
	. "mai.today/api/design/dsl"
)

const (
	lastTimestamp = "lastTimestamp"
)

var ServiceName = Service("Sync", func() {
	Description("資料同步")
	ErrorInvalidToken()
	sync("sync")
}).Name

var syncResult = CommandResult("sync", syncPayloadData)

var syncPayloadData = Type("SyncPayloadData", func() {
	base.AttributeBaseInfos()
	base.AttributeBaseNavState()
})

func attributeLastTimestamp() {
	Attribute(lastTimestamp, Int64, func() {
		Description("上次同步的時間戳記")
		FormatTimestamp()
		ExampleTimestamp()
	})
}

func sync(methodName string) {
	const sum = "請求同步"
	var resType = syncResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			attributeLastTimestamp()
		})
		Result(resType)
		HTTP(func() {
			POST("/sync")
			ResponseDefaults()
		})
	})
}
