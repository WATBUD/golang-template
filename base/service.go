package base

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"mai.today/api/gen/base"
	"mai.today/authentication"
	mongorepo "mai.today/base/mongo"
	"mai.today/core/entity"
	"mai.today/database/mongodb"
	usercontext "mai.today/foundation/context/user"
	"mai.today/realtime"
)

type Repository interface {
	FindMemberByBaseID(ctx context.Context, id string) (entity *entity.BaseMember, err error)
	FindNavStatesByUserID(ctx context.Context, userID string) ([]*entity.BaseNavState, error)
	InsertOne(ctx context.Context, member *entity.BaseMember, info *entity.BaseInfo, navState *entity.BaseNavState) (err error)
	SoftDeleteByID(ctx context.Context, id string) (err error)
	UpdateInfoByBaseID(ctx context.Context, info *entity.BaseInfo) (bool, error)
	UpdateNavStates(ctx context.Context, items []*entity.BaseNavState) error
}

type Service interface {
	base.Auther
	base.Service
}

var (
	// once ensures instance initialization is performed exactly once.
	once sync.Once

	// instance holds the singleton BaseService instance.
	instance Service
)

// Instance returns a singleton instance of BaseService.
func Instance() Service {
	once.Do(func() {
		instance = newBaseService()
	})
	return instance
}

func newBaseService() Service {
	return BaseService{
		authentication.Instance(),
		realtime.Instance(),
		mongorepo.NewBaseRepository(mongodb.Instance()),
	}
}

type BaseService struct {
	base.Auther
	realtime realtime.Realtime
	repo     Repository
}

func (sv BaseService) CreateBase(ctx context.Context, pl *base.CreateBasePayload) (res *base.CreateBaseResult, err error) {
	uid, _ := usercontext.GetUserID(ctx)

	member := &entity.BaseMember{
		UserID: uid,
	}
	info := &entity.BaseInfo{
		Name:  pl.Name,
		Logo:  pl.Logo,
		Color: pl.Color,
	}
	navState := &entity.BaseNavState{
		Index:  pl.Index,
		UserID: uid,
	}

	if err := sv.repo.InsertOne(ctx, member, info, navState); err != nil {
		return nil, err
	}

	res = &base.CreateBaseResult{
		Command: &base.Command{
			Type: "createBase",
		},
		Data: &base.CreateBaseResultData{
			ID:       info.BaseID,
			Info:     toInfoResult(info),
			NavState: toNavStateResult(navState),
		},
		Timestamp: time.Now().UnixNano(),
	}

	_, err = sv.realtime.PublishToUser(ctx, uid, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sv BaseService) DeleteBase(ctx context.Context, pl *base.DeleteBasePayload) (res *base.DeleteBaseResult, err error) {
	uid, _ := usercontext.GetUserID(ctx)

	ok, err := sv.hasPermission(ctx, pl.ID, uid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, base.MakeNotFound(errors.New("base id not found"))
		}
		return nil, err
	} else if !ok {
		return nil, base.MakePermissionDenied(errors.New("no permission to delete base"))
	}

	if err := sv.repo.SoftDeleteByID(ctx, pl.ID); err != nil {
		return nil, err
	}

	res = &base.DeleteBaseResult{
		Command: &base.Command{
			Type: "deleteBase",
		},
		Data: &base.DeleteBaseResultData{
			BaseID: pl.ID,
		},
		Timestamp: time.Now().UnixNano(),
	}

	r, err := sv.realtime.PublishToUser(ctx, uid, res)
	fmt.Print(r)
	if err != nil {
		return nil, err
	}

	return
}

func (sv BaseService) ReorderBaseNavStates(ctx context.Context, pl *base.ReorderBaseNavStatesPayload) (res *base.ReorderBaseNavStateResult, err error) {
	uid, _ := usercontext.GetUserID(ctx)

	navStates, err := sv.repo.FindNavStatesByUserID(ctx, uid)
	if err != nil {
		return nil, err
	}

	sv.reorderBaseNavStates(pl.ID, pl.NewIndex, navStates)
	if err = sv.repo.UpdateNavStates(ctx, navStates); err != nil {
		return nil, err
	}

	res = &base.ReorderBaseNavStateResult{
		Command: &base.Command{
			Type: "reorderBaseNavState",
		},
		Data:      toNavStateResults(navStates),
		Timestamp: time.Now().UnixNano(),
	}

	r, err := sv.realtime.PublishToUser(ctx, uid, res)
	fmt.Print(r)
	if err != nil {
		return nil, err
	}

	return
}

func (sv BaseService) UpdateBaseInfo(ctx context.Context, pl *base.UpdateBaseInfoPayload) (res *base.UpdateBaseInfoResult, err error) {
	uid, _ := usercontext.GetUserID(ctx)

	ok, err := sv.hasPermission(ctx, pl.ID, uid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, base.MakeNotFound(errors.New("base info id not found"))
		}
		return nil, err
	} else if !ok {
		return nil, base.MakePermissionDenied(errors.New("no permission to update base info"))
	}

	info := &entity.BaseInfo{
		BaseID: pl.ID,
		Name:   pl.Name,
		Logo:   pl.Logo,
		Color:  pl.Color,
	}

	if _, err := sv.repo.UpdateInfoByBaseID(ctx, info); err != nil {
		return nil, err
	}

	res = &base.UpdateBaseInfoResult{
		Command: &base.Command{
			Type: "updateBaseInfo",
		},
		Data:      toInfoResult(info),
		Timestamp: time.Now().UnixNano(),
	}

	r, err := sv.realtime.PublishToUser(ctx, uid, res)
	fmt.Print(r)
	if err != nil {
		return nil, err
	}

	return
}

func (sv BaseService) ReceiveCreateBase(_ context.Context, _ *base.ReceiveCreateBasePayload) (res *base.CreateBaseResult, err error) {
	return nil, errors.New("not implemented")
}

func (sv BaseService) ReceiveDeleteBase(_ context.Context, _ *base.ReceiveDeleteBasePayload) (res *base.DeleteBaseResult, err error) {
	return nil, errors.New("not implemented")
}

func (sv BaseService) ReceiveUpdateBaseInfo(_ context.Context, _ *base.ReceiveUpdateBaseInfoPayload) (res *base.UpdateBaseInfoResult, err error) {
	return nil, errors.New("not implemented")
}

func (sv BaseService) ReceiveReorderBaseNavStates(_ context.Context, _ *base.ReceiveReorderBaseNavStatesPayload) (res *base.ReorderBaseNavStateResult, err error) {
	return nil, errors.New("not implemented")
}

func (sv BaseService) hasPermission(ctx context.Context, baseID, userID string) (bool, error) {
	member, err := sv.repo.FindMemberByBaseID(ctx, baseID)
	if err != nil {
		return false, err
	}

	return member.UserID == userID, nil
}

func (sv BaseService) reorderBaseNavStates(id string, newIndex int, items []*entity.BaseNavState) {
	var targetItem *entity.BaseNavState
	var targetIndex int
	for i, item := range items {
		if item.ID == id {
			targetItem = item
			targetIndex = i
			break
		}
	}

	// Remove the target item from its current position.
	items = append(items[:targetIndex], items[targetIndex+1:]...)

	// Insert the target item at the new position.
	items = append(items[:newIndex-1], append([]*entity.BaseNavState{targetItem}, items[newIndex-1:]...)...)

	// Adjust the indices.
	for i := range items {
		items[i].Index = i + 1
	}
}

func toInfoResult(v *entity.BaseInfo) *base.BaseInfo {
	return &base.BaseInfo{
		ID:     v.ID,
		BaseID: v.BaseID,
		Color:  v.Color,
		Logo:   v.Logo,
		Name:   v.Name,
	}
}

func toNavStateResult(v *entity.BaseNavState) *base.BaseNavState {
	return &base.BaseNavState{
		ID:     v.ID,
		BaseID: v.BaseID,
		Index:  v.Index,
	}
}

func toNavStateResults(v []*entity.BaseNavState) []*base.BaseNavState {
	results := make([]*base.BaseNavState, len(v))
	for i, navState := range v {
		results[i] = toNavStateResult(navState)
	}
	return results
}
