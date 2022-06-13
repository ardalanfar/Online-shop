package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowMembers(ctx context.Context) ([]entity.User, error) {
	var users []model.User

	//get all id,email,username,password
	cheek := s.Db.WithContext(ctx).Select("id", "email", "username", "password").Find(&users)
	if cheek.Error != nil {
		return nil, cheek.Error
	}

	usersEntities := make([]entity.User, len(users))

	for i := range users {
		usersEntities[i] = model.MapToUserEntity(users[i])
	}

	//return
	return usersEntities, nil
}

func (s DbConn) DeleteMember(ctx context.Context, userID uint) error {
	var user model.User

	//find user for delete
	cheekFind := s.Db.WithContext(ctx).Where("id = ?", userID).First(&user)
	if cheekFind.Error != nil {
		return cheekFind.Error
	}

	//delete username by id
	cheekDelete := s.Db.WithContext(ctx).Delete(&user)
	if cheekDelete.Error != nil {
		return cheekDelete.Error
	}
	//return
	return nil
}

func (s DbConn) ShowInfoMember(ctx context.Context, userID uint) (entity.User, error) {
	var user model.User

	//get "id", "email", "username" by username
	cheek := s.Db.WithContext(ctx).Select("users.id", "users.email", "users.username", "users.access", "users.is_verified").Where("id", userID).First(&user)
	if cheek.Error != nil {
		return entity.User{}, cheek.Error
	}
	//return
	return model.MapToUserEntity(user), nil
}
