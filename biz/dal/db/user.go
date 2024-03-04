package db

import (
	"go_gin/biz/dal"
	"go_gin/biz/entity/user"
)

func CreateUsers(users []*entity.User) error {
	return dal.Mysql.Create(users).Error
}

func FindUserByNameOrPhoneNumber(userName string, phoneNumber string) ([]*entity.User, error) {
	res := make([]*entity.User, 0)
	if err := dal.Mysql.Where(dal.Mysql.Or("user_name = ?", userName).
		Or("phone_number = ?", phoneNumber)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}