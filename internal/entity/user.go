package entity

import "time"

// User represents the user entity for Gorm mappings from users table.
type User struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	Mobile        string    `gorm:"column:mobile" json:"mobile"`
	Name          string    `gorm:"column:name" json:"name"`
	Nickname      string    `gorm:"column:nickname" json:"nickname"`
	Email         string    `gorm:"column:email" json:"email"`
	Score         int       `gorm:"column:score" json:"score"`
	InviterID     int       `gorm:"column:inviter_id" json:"inviter_id"`
	Code          string    `gorm:"column:code" json:"code"`
	APIToken      string    `gorm:"column:api_token" json:"api_token"`
	City          string    `gorm:"column:city" json:"city"`
	Birthday      string    `gorm:"column:brithday" json:"birthday"`
	Sex           string    `gorm:"column:sex" json:"sex"`
	TotalScore    int       `gorm:"column:total_score" json:"total_score"`
	OpenID        string    `gorm:"column:openid" json:"openid"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	Addr          string    `gorm:"column:addr" json:"addr"`
	Password      string    `gorm:"column:password" json:"-"`
	RememberToken string    `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	XcxOpenID     string    `gorm:"column:xcx_openid" json:"xcx_openid"`
	UnionID       string    `gorm:"column:unionid" json:"unionid"`
	SessionKey    string    `gorm:"column:session_key" json:"session_key"`
	IsStaff       bool      `gorm:"column:is_staff" json:"is_staff"`
	IsRoot        bool      `gorm:"column:is_root" json:"is_root"`
	CheckCode     string    `gorm:"column:check_code" json:"check_code"`
	Status        int       `gorm:"column:status" json:"status"`
}
