package data

import "fmt"

type user struct {
	ID       int
	userName string
}

type users []*user

var usersList = users{
	&user{ID: 1,
		userName: "Emre"},
}

func GetAllUsers() users {
	return usersList
}

func (user user) String() string {
	return fmt.Sprintf("{%v %v}", user.ID, user.userName)
}

type ErrorUserNotFound struct {
	userId int
}

func (err *ErrorUserNotFound) Error() string {
	return fmt.Sprintf("Could not found user. User id : %v", err.userId)
}

func GetUser(userId int) (*user, error) {
	for _, user := range usersList {
		if userId == user.ID {
			return user, nil
		}
	}

	return nil, &ErrorUserNotFound{userId}
}

func (user *user) GetUserName() string {
	return user.userName
}
