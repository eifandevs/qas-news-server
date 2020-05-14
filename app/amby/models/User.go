package models

import(
    "github.com/eifandevs/amby/repo"
    // "github.com/jinzhu/gorm"
    "log"
    "math/rand"
    "time"
)

type UserInfo struct {
    Mail string `json:"mail"`
}

type UserToken struct {
    AccessToken string `json:"access_token"`
}

type User struct {
    Mail string `gorm:"primary_key"`
    AccessToken string `gorm:"unique"`
    AccessTokenExpire string
}

type PostUserRequest struct {
    Item  UserInfo `json:"data"`
}

type PostUserResponse struct {
    Item  UserToken `json:"data"`
}

func CreateUser(userinfo UserInfo) (User, error) {
    db := repo.Connect("development")
    defer db.Close()

    users := []User{}
    if err := db.Where("mail = ?", userinfo.Mail).Find(&users).Error; err != nil {
        return User{}, err
    }

    if len(users) == 0 {
        accessToken := createToken()
        newUser := User{Mail: userinfo.Mail, AccessToken: accessToken, AccessTokenExpire: createExpireDate()}
        if err := db.Create(&newUser).Error; err != nil {
            return User{}, err
        }
        return newUser, nil
    } else {
        log.Println("already exist.")
        return users[0], nil
    }
}

func createToken() string {
    n := 40
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func createExpireDate() string {
  jst, _ := time.LoadLocation("Asia/Tokyo")
  now := time.Now()
  // 現在+90日
  return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, jst).Add(24 * time.Hour * 90).String()
}