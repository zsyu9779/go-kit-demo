package service

type IUserService interface {
	GetName(userId int) string
}
type UserService struct {

}

func (u *UserService)GetName(userId int) string  {
	if userId == 111 {
		return "aaa"
	}
	return "bbb"
}