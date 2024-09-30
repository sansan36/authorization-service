package repository

import (
	"context"
	"time"

	commonv1 "github.com/sansan36/authorization-service/gen/common/v1"
	userv1 "github.com/sansan36/authorization-service/gen/user/v1"
	"github.com/sansan36/authorization-service/libs/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserList(context.Context, *userv1.User, *userv1.GetUserListRequest) ([]*userv1.User, *commonv1.StandardPaginationResponse, error)
	GetUser(context.Context, *userv1.User) (*userv1.User, error)
	AddUser(context.Context, *userv1.User) (*userv1.User, error)
	EditUser(context.Context, *userv1.User) (*userv1.User, error)
	RemoveUser(context.Context, *userv1.User) (*userv1.User, error)
}

type gormUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(dbMain *gorm.DB) *gormUserRepo {
	return &gormUserRepo{
		db: dbMain,
	}
}

func (repo *gormUserRepo) GetUserList(ctx context.Context, req *userv1.User, query *userv1.GetUserListRequest) ([]*userv1.User, *commonv1.StandardPaginationResponse, error) {
	usersORM := []*userv1.UserORM{}
	pagination := &commonv1.StandardPaginationResponse{}

	statement := repo.db.Model(&userv1.UserORM{}).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, pagination, err
		}
		statement = statement.Where(&reqORM)
	}

	statement = searchQuery(statement, &commonv1.StandardQuery{Search: query.Search}, []string{"user_name", "name", "email"})
	statement = PaginationCounterQuery(statement, &commonv1.StandardQuery{Page: query.Page, PageSize: query.PageSize}, pagination)

	err := statement.Find(&usersORM).Error
	if err != nil {
		logger.Errorln("Fail to get user list from DB")
		return nil, pagination, err
	}

	users := make([]*userv1.User, len(usersORM))
	for i, userORM := range usersORM {
		userPB, err := userORM.ToPB(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert response to PB")
			return nil, pagination, err
		}
		users[i] = &userPB
	}

	pagination.Found = int64(len(users))
	return users, pagination, nil
}

func (repo *gormUserRepo) GetUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Where("deleted_at IS NULL").First(userORM, userORM).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Errorln("Fail to get user from DB")
		}
		return nil, err
	}

	userPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}

	return &userPB, nil
}

func (repo *gormUserRepo) AddUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Create(userORM).Error
	if err != nil {
		logger.Errorln("Fail to add user to DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}

	return &resPB, nil
}

func (repo *gormUserRepo) EditUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Save(userORM).Error
	if err != nil {
		logger.Errorln("Fail to edit user in DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}

	return &resPB, nil
}

func (repo *gormUserRepo) RemoveUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Model(userORM).Update("deleted_at", time.Now()).Error
	if err != nil {
		logger.Errorln("Fail to soft delete user from DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}

	return &resPB, nil
}
