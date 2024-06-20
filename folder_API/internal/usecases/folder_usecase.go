package usecases

import (
	"context"
	"folder_API/internal/entities"
)

type FolderRepository interface {
	Create(ctx context.Context, folder *entities.Folder) error
	FindAll(ctx context.Context) ([]*entities.Folder, error)
	FindByID(ctx context.Context, id string) (*entities.Folder, error)
	Update(ctx context.Context, folder *entities.Folder) error
	Delete(ctx context.Context, id string) error
	UpdateIndex(ctx context.Context, id string, index int) error
	UpdateParent(ctx context.Context, id string, parentIndex int) error
}
