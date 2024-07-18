// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Base service
//
// Command:
// $ goa gen mai.today/api/design

package base

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// 基地
type Service interface {
	// ReceiveCreateBase implements receiveCreateBase.
	ReceiveCreateBase(context.Context, *ReceiveCreateBasePayload) (res *CreateBaseResult, err error)
	// CreateBase implements CreateBase.
	CreateBase(context.Context, *CreateBasePayload) (res *CreateBaseResult, err error)
	// ReceiveDeleteBase implements receiveDeleteBase.
	ReceiveDeleteBase(context.Context, *ReceiveDeleteBasePayload) (res *DeleteBaseResult, err error)
	// DeleteBase implements DeleteBase.
	DeleteBase(context.Context, *DeleteBasePayload) (res *DeleteBaseResult, err error)
	// ReceiveUpdateBaseInfo implements receiveUpdateBaseInfo.
	ReceiveUpdateBaseInfo(context.Context, *ReceiveUpdateBaseInfoPayload) (res *UpdateBaseInfoResult, err error)
	// UpdateBaseInfo implements UpdateBaseInfo.
	UpdateBaseInfo(context.Context, *UpdateBaseInfoPayload) (res *UpdateBaseInfoResult, err error)
	// ReceiveReorderBaseNavStates implements receiveReorderBaseNavStates.
	ReceiveReorderBaseNavStates(context.Context, *ReceiveReorderBaseNavStatesPayload) (res *ReorderBaseNavStateResult, err error)
	// ReorderBaseNavStates implements ReorderBaseNavStates.
	ReorderBaseNavStates(context.Context, *ReorderBaseNavStatesPayload) (res *ReorderBaseNavStateResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "api"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Base"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [8]string{"receiveCreateBase", "CreateBase", "receiveDeleteBase", "DeleteBase", "receiveUpdateBaseInfo", "UpdateBaseInfo", "receiveReorderBaseNavStates", "ReorderBaseNavStates"}

// 基地資訊
type BaseInfo struct {
	// 識別碼
	ID string
	// 基地識別碼
	BaseID string
	// 主題颜色
	Color string
	// 標識圖片的 URL
	Logo string
	// 名稱
	Name string
	// 刪除時間
	DeletedAt *string
}

// 基地導覽列狀態
type BaseNavState struct {
	// 識別碼
	ID string
	// 基地識別碼
	BaseID string
	// 順序
	Index int
	// 刪除時間
	DeletedAt *string
}

// 指令
type Command struct {
	// 類型
	Type string
}

// CreateBasePayload is the payload type of the Base service CreateBase method.
type CreateBasePayload struct {
	JWT *string
	// 主題颜色
	Color string
	// 順序
	Index int
	// 標識圖片的 URL
	Logo string
	// 名稱
	Name string
}

// CreateBaseResult is the result type of the Base service receiveCreateBase
// method.
type CreateBaseResult struct {
	Command *Command
	// 時間戳記
	Timestamp int64
	// 資料
	Data *CreateBaseResultData
}

type CreateBaseResultData struct {
	// 識別碼
	ID string
	// 基地資訊
	Info *BaseInfo
	// 基地導覽列狀態
	NavState *BaseNavState
}

// DeleteBasePayload is the payload type of the Base service DeleteBase method.
type DeleteBasePayload struct {
	JWT *string
	// 識別碼
	ID string
}

// DeleteBaseResult is the result type of the Base service receiveDeleteBase
// method.
type DeleteBaseResult struct {
	Command *Command
	// 時間戳記
	Timestamp int64
	// 資料
	Data *DeleteBaseResultData
}

type DeleteBaseResultData struct {
	// 基地識別碼
	BaseID string
}

// ReceiveCreateBasePayload is the payload type of the Base service
// receiveCreateBase method.
type ReceiveCreateBasePayload struct {
	JWT *string
	// 使用者識別碼
	Channel string
}

// ReceiveDeleteBasePayload is the payload type of the Base service
// receiveDeleteBase method.
type ReceiveDeleteBasePayload struct {
	JWT *string
	// 使用者識別碼
	Channel string
}

// ReceiveReorderBaseNavStatesPayload is the payload type of the Base service
// receiveReorderBaseNavStates method.
type ReceiveReorderBaseNavStatesPayload struct {
	JWT *string
	// 使用者識別碼
	Channel string
}

// ReceiveUpdateBaseInfoPayload is the payload type of the Base service
// receiveUpdateBaseInfo method.
type ReceiveUpdateBaseInfoPayload struct {
	JWT *string
	// 使用者識別碼
	Channel string
}

// ReorderBaseNavStateResult is the result type of the Base service
// receiveReorderBaseNavStates method.
type ReorderBaseNavStateResult struct {
	Command *Command
	// 時間戳記
	Timestamp int64
	// 資料
	Data []*BaseNavState
}

// ReorderBaseNavStatesPayload is the payload type of the Base service
// ReorderBaseNavStates method.
type ReorderBaseNavStatesPayload struct {
	JWT *string
	// 識別碼
	ID string
	// 新的順序
	NewIndex int
}

// UpdateBaseInfoPayload is the payload type of the Base service UpdateBaseInfo
// method.
type UpdateBaseInfoPayload struct {
	JWT *string
	// 識別碼
	ID string
	// 主題颜色
	Color string
	// 標識圖片的 URL
	Logo string
	// 名稱
	Name string
}

// UpdateBaseInfoResult is the result type of the Base service
// receiveUpdateBaseInfo method.
type UpdateBaseInfoResult struct {
	Command *Command
	// 時間戳記
	Timestamp int64
	// 資料
	Data *BaseInfo
}

// MakeInvalidToken builds a goa.ServiceError from an error.
func MakeInvalidToken(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid token", false, false, false)
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not found", false, false, false)
}

// MakePermissionDenied builds a goa.ServiceError from an error.
func MakePermissionDenied(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "permission denied", false, false, false)
}