package storage

import "errors"

var (
	UserExists         = errors.New("user exists")
	UserNotFound       = errors.New("user or password is incorrect")
	FileNotFound       = errors.New("file not found or you haven't access")
	ShareExists        = errors.New("share exists")
	ShareUserNotExists = errors.New("user you share with not exists")
	NotOwner           = errors.New("you are not owner or audio not exists")
	NotAacFile         = errors.New("file is not Aac")
	WrongRefreshToken  = errors.New("token not found or expires in")
)
