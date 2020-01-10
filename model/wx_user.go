package model

import (
	"time"
)

type WxUser struct {
	Openid     string    `gorm:"type:varchar(255);column:openid;primary_key"`
	Nickname   string    `gorm:"type:varchar(255);column:nickname"`
	Avatarurl  string    `gorm:"type:varchar(255);column:avatarurl"`
	Gender     int       `gorm:"type:tinyint(4);column:gender"`
	Sessionkey string    `gorm:"type:varchar(100);column:sessionkey"`
	Province   string    `gorm:"type:varchar(10);column:province"`
	Ctiy       string    `gorm:"type:varchar(10);column:city"`
	Country    string    `gorm:"type:varchar(20);column:country"`
	CreatedAt  time.Time `gorm:"column:create_time"`
	UpdatedAt  time.Time `gorm:"column:last_login_time"`
}

func (u WxUser) TableName() string {
	return "wx_user"
}
