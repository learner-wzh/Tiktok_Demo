package models

type UserModel struct {
	UserID            int64  `gorm:"column:ID"`            // 用户id
	UserName          string `gorm:"column:Name"`          // 用户名称
	UserPwd           string `gorm:"column:Pwd"`           // 用户密码
	UserFollowCount   int64  `gorm:"column:FollowCount"`   // 用户关注数
	UserFollowerCount int64  `gorm:"column:FollowerCount"` // 用户粉丝数
}

func (value UserModel) TableName() string {
	return "users"
}

func UserRegister(Name string, Pwd string) (bool, int64) {

	var user UserModel

	DB.Select("ID, `Name`").Where("`Name` = ?", Name).First(&user)

	if user.UserID == 0 {

		InsertUser := UserModel{UserName: Name, UserPwd: Pwd}
		DB.Create(&InsertUser)

		return true, InsertUser.UserID
	} else {
		return false, 0
	}

}

func UserLogin(Name string, Pwd string) UserModel {

	var user UserModel

	DB.Select("ID, `Name`, Pwd").Where("`Name` = ? AND Pwd = ?", Name, Pwd).Take(&user)

	return user
}
