package user

type IUserRepo interface {
	FindUserByFireBaseUID(fireBaseUID string) *UserModel
	CreateUser(user UserModel) error
}

type IUser interface {
	FindUserByFireBaseUID(fireBaseUID string) *UserModel
	CreateUser(firebaseUID string) error
}
