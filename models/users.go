package models

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User UserInfo `json:"user"`
}

type User struct {
	UserID   int64  `gorm:"column:UserID"` // 用户id
	UserName string `gorm:"column:Name"`   // 用户名称
	UserPwd  string `gorm:"column:Pwd"`    // 用户密码
}

func (value User) TableName() string {
	return "Users"
}

func UserRegister(Name string, Pwd string) (bool, int64) {

	var user User

	DB.Select("UserID, `Name`").Where("`Name` = ?", Name).First(&user)

	if user.UserID == 0 {

		InsertUser := User{UserName: Name, UserPwd: Pwd}
		DB.Create(&InsertUser)

		return true, InsertUser.UserID
	} else {
		return false, 0
	}

}

func UserLogin(Name string, Pwd string) User {

	var user User

	DB.Select("UserID, `Name`, Pwd").Where("`Name` = ? AND Pwd = ?", Name, Pwd).Take(&user)

	return user
}
