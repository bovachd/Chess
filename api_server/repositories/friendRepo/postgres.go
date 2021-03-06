package friendRepo

import (
	"Chess/api_server/infrastructure"
	model "Chess/api_server/models"
	"Chess/api_server/repositories/accountRepo"
	"github.com/jinzhu/gorm"
)

type FriendRepository struct {
	userRepo model.UserRepository
}

func (FriendRepository) GetById(id int) (*model.Friend, error) {
	db := infrastructure.GetDB()

	var Friend model.Friend
	if err := db.First(&Friend, id).Error; err != nil {
		return nil, err
	}

	return &Friend, nil
}

func (fr *FriendRepository) GetAllFriendByUserId(userId int) ([]*model.User, error) {
	db := infrastructure.GetDB()

	var friends []*model.Friend
	var users []*model.User

	if err := db.Debug().Table("friends").Find(&friends, "user_id = ?", userId).Error; err != nil{
		return nil, err
	}

	for _, item := range friends{
		user, err := fr.userRepo.FindById(item.FriendId)

		if err != nil{
			infrastructure.ErrLog.Println(err)
		} else {
			users = append(users, user)
		}
	}

	if len(users) == 0{
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}

func (FriendRepository) DeleteFriendByFriendId(userId, friendId int) error {
	db := infrastructure.GetDB()

	if err := db.Table("friends").Delete(&model.Friend{}, "user_id = ? and friend_id = ?", userId, friendId).Error; err != nil{
		return err
	}
	return nil
}

func (FriendRepository) CreateNewFriend(friend *model.Friend) (*model.Friend, error) {
	db := infrastructure.GetDB()

	if err := db.Create(&friend).Error; err != nil {
		return nil, err
	}

	return friend, nil
}

func NewFriendRepository() model.FriendRepository {
	userRepo := accountRepo.NewUserRepository()
	return &FriendRepository{
		userRepo: userRepo,
	}
}