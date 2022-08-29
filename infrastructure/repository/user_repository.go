package repository

import "github.com/Mo3g4u/go-ddd-sample/domain/model"

func (d *dbRepository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	if err := d.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *dbRepository) CreateUser(user *model.User) error {
	return d.db.Create(user).Error
}
