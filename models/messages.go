package models

import (
	"log"
	"strconv"
	"strings"
)

type Message struct {
	MessageID  int64  `gorm:"column:MessageID" json:"id,omitempty"`
	ChatID     int64  `gorm:"column:ChatID"`
	ToUserID   int64  `gorm:"column:ToUserID" json:"to_user_id,omitempty"`
	FromUserID int64  `gorm:"column:FromUserID" json:"from_user_id,omitempty"`
	Content    string `gorm:"column:Content" json:"content,omitempty"`
	CreateTime string `gorm:"column:CreateTime" json:"create_time,omitempty"`
}

func QueryMessageListByTokenAndUserID(token string, chatID string) ([]Message, error) {
	tableName := token + "-chathistory"
	ChatID, _ := strconv.ParseInt(chatID, 10, 64)

	var messageList []Message
	err := DB.Table(tableName).Select("MessageID, ToUserID, FromUserID, Content, CreateTime").Where("ChatID = ?", ChatID).
		Order("MessageID ASC").Find(&messageList).Error

	return messageList, err
}

func QueryMessageListCountByTokenAndUserID(token string, chatID string) (int, error) {
	tableName := token + "-chathistory"
	ChatID, _ := strconv.ParseInt(chatID, 10, 64)

	var Count int
	err := DB.Table(tableName).Select("COUNT(*)").Where("ChatID = ?", ChatID).Find(&Count).Error

	return Count, err
}

func InsertOwnerMessageList(token string, MessageID int, chatID string, content string, time string) {
	tableName := token + "-chathistory"
	ChatID, _ := strconv.ParseInt(chatID, 10, 64)
	strArray := strings.Split(token, "-")
	userID := strArray[0]
	UserID, _ := strconv.ParseInt(userID, 10, 64)

	message := Message{
		MessageID:  int64(MessageID + 1),
		ChatID:     ChatID,
		ToUserID:   ChatID,
		FromUserID: UserID,
		Content:    content,
		CreateTime: time,
	}

	err := DB.Table(tableName).Create(message).Error
	log.Println(err)
}

func InsertToUserIDMessageList(token string, MessageID int, chatID string, content string, time string) {
	name, _ := QueryUserInfoNameByID(chatID)
	tableName := chatID + "-" + name + "-chathistory"

	ChatID, _ := strconv.ParseInt(chatID, 10, 64)
	strArray := strings.Split(token, "-")
	userID := strArray[0]
	UserID, _ := strconv.ParseInt(userID, 10, 64)

	message := Message{
		MessageID:  int64(MessageID + 1),
		ChatID:     UserID,
		ToUserID:   UserID,
		FromUserID: ChatID,
		Content:    content,
		CreateTime: time,
	}

	err := DB.Table(tableName).Create(message).Error
	log.Println(err)
}
