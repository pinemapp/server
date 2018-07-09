package models

import "strconv"

type Client struct {
	Model

	UserID       uint `gorm:"type:int references users(id)"`
	User         User
	Domain       string `gorm:"type:varchar(255)"`
	ClientID     string `gorm:"type:varchar(255);unique_index;not null"`
	ClientSecret string `gorm:"type:varchar(255);unique_index;not null"`
}

func (c *Client) GetID() string {
	return c.ClientID
}

func (c *Client) GetSecret() string {
	return c.ClientSecret
}

func (c *Client) GetDomain() string {
	return c.Domain
}

func (c *Client) GetUserID() string {
	return strconv.FormatUint(uint64(c.UserID), 10)
}
