package dao

import "model"

type ProfileDAO struct {
}

func GetProfileDAO() *ProfileDAO {
	return &ProfileDAO{}
}

func (dao *ProfileDAO) GetProfile(email string) (model.User, error) {
	user, error := GetUserDAO().GetUserByEmail(email)
	if error != nil {
		return emptyUser(), error
	}
	return user, nil
}

func (dao *ProfileDAO) UpdateProfile(user model.User) (model.User, error) {
	return emptyUser(), nil
}

func (dao *ProfileDAO) DeleteProfile(user model.User) error {
	return nil
}
