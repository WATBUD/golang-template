package folder_mod

import (
	"context"
	"folder_mod/internal/entities"
	"folder_mod/internal/interfaces/repository"
	"folder_mod/internal/usecases"
	"sync"
	"time"

	goa "goa.design/goa/v3/pkg"
	"mai.today/api/gen/folder"
	"mai.today/authentication"
	"mai.today/database/mongodb"
	"mai.today/realtime"
)

func NewFolderServices(repo usecases.FolderRepository) *FolderService {
	return &FolderService{repo: repo}
}

func (fs *FolderService) CreateFolder(ctx context.Context, p *folder.CreateFolderPayload) (*folder.CreateFolderResult, error) {
	newFolder := entities.Folder{
		ID:        p.Folder.FolderID,
		BaseID:    p.Folder.BaseID,
		ParentID:  p.Folder.ParentID,
		Position:  p.Folder.Position,
		Data:      entities.FolderData{Color: p.Folder.Data.Color, Name: p.Folder.Data.Name},
		Type:      p.Folder.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := fs.repo.CreateFolder(ctx, &newFolder)
	if err != nil {
		return nil, err
	}
	return &folder.CreateFolderResult{
		Command:   &folder.Command{Type: "Create"},
		Timestamp: newFolder.CreatedAt.Unix(),
		Data: &folder.Createfolderresultdata{
			FolderID:  newFolder.ID,
			BaseID:    newFolder.BaseID,
			ParentID:  newFolder.ParentID,
			Position:  newFolder.Position,
			CreatedAt: newFolder.CreatedAt.String(),
			UpdatedAt: newFolder.UpdatedAt.String(),
			Data:      &folder.FolderData{Color: newFolder.Data.Color, Name: newFolder.Data.Name},
			Type:      newFolder.Type,
		},
	}, nil
}

func (fs *FolderService) DeleteFolder(ctx context.Context, p *folder.DeleteFolderPayload) (*folder.DeleteFolderResult, error) {
	err := fs.repo.DeleteFolderByID(ctx, p.FolderID)
	if err != nil {
		return nil, err
	}
	return &folder.DeleteFolderResult{
		Command:   &folder.Command{Type: "Delete"},
		Timestamp: time.Now().Unix(),
		Data: &folder.Createfolderresultdata{
			FolderID: p.FolderID,
		},
	}, nil
}

func (fs *FolderService) UpdateFolder(ctx context.Context, p *folder.UpdateFolderPayload) (*folder.UpdateFolderResult, error) {

	_folder := entities.Folder{
		ID:        p.Folder.FolderID,
		BaseID:    p.Folder.BaseID,
		ParentID:  p.Folder.ParentID,
		Position:  p.Folder.Position,
		Data:      entities.FolderData{Color: p.Folder.Data.Color, Name: p.Folder.Data.Name},
		Type:      p.Folder.Type,
		UpdatedAt: time.Now(),
	}
	err := fs.repo.UpdateFolderData(ctx, &_folder)
	if err != nil {
		return nil, err
	}
	return &folder.UpdateFolderResult{
		Command:   &folder.Command{Type: "Update"},
		Timestamp: _folder.UpdatedAt.Unix(),
		Data: &folder.Createfolderresultdata{
			FolderID:  _folder.ID,
			BaseID:    _folder.BaseID,
			ParentID:  _folder.ParentID,
			Position:  _folder.Position, // No need to use pointer
			CreatedAt: _folder.CreatedAt.String(),
			UpdatedAt: _folder.UpdatedAt.String(),
			Data:      &folder.FolderData{Color: _folder.Data.Color, Name: _folder.Data.Name}, // No pointer
			Type:      _folder.Type,
		},
	}, nil
}

func (fs *FolderService) ListFolders(ctx context.Context, p *folder.ListFoldersPayload) (*folder.ListFoldersResult, error) {
	folders, err := fs.repo.GetFolders(ctx)
	if err != nil {
		return nil, err
	}
	var folderCollection []*folder.Createfolderresultdata
	for _, f := range folders {
		folderCollection = append(folderCollection, &folder.Createfolderresultdata{
			FolderID:  f.ID,
			BaseID:    f.BaseID,
			ParentID:  f.ParentID,
			Position:  f.Position,
			CreatedAt: f.CreatedAt.String(),
			UpdatedAt: f.UpdatedAt.String(),
			Data:      &folder.FolderData{Color: f.Data.Color, Name: f.Data.Name},
			Type:      f.Type,
		})
	}
	result := &folder.ListFoldersResult{
		Command:   nil,
		Timestamp: time.Now().Unix(),
		Data:      folderCollection,
	}
	return result, nil
}

// MakeInvalidToken builds a goa.ServiceError from an error.
func MakeInvalidToken(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid token", false, false, false)
}

func (fs *FolderService) ReceiveCreateFolder(context.Context, *folder.ReceiveCreateFolderPayload) (res *folder.CreateFolderResult, err error) {

	panic("not implemented")
}
func (fs *FolderService) ReceiveDeleteFolder(context.Context, *folder.ReceiveDeleteFolderPayload) (res *folder.DeleteFolderResult, err error) {
	panic("not implemented")
}
func (fs *FolderService) ReceiveUpdateFolder(context.Context, *folder.ReceiveUpdateFolderPayload) (res *folder.UpdateFolderResult, err error) {
	panic("not implemented")

}
func (fs *FolderService) ReceiveListFolders(context.Context, *folder.ReceiveListFoldersPayload) (res *folder.ListFoldersResult, err error) {
	panic("not implemented")
}

var (
	once     sync.Once
	instance *FolderService
)

func Instance() *FolderService {
	once.Do(func() {
		instance = newFolderService()
	})
	return instance
}

type FolderService struct {
	realtime realtime.Realtime
	folder.Auther
	repo usecases.FolderRepository
}

func newFolderService() *FolderService {
	return &FolderService{
		realtime.Instance(),
		authentication.Instance(),
		repository.NewMongoFolderRepository(mongodb.Instance()),
	}
}
