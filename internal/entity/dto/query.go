package dto

import "gorm.io/gorm"

type QueryOption interface {
	Apply(*gorm.DB)
}

type QueryOptionFunc func(*gorm.DB)

func (f QueryOptionFunc) Apply(db *gorm.DB) {
	f(db)
}

// QueryWithID 查询ID
func QueryWithID(id int) QueryOption {
	return QueryOptionFunc(func(db *gorm.DB) {
		db.Where("id = ?", id)
	})
}

// QueryWithAPIToken 查询API Token
func QueryWithAPIToken(token int) QueryOption {
	return QueryOptionFunc(func(db *gorm.DB) {
		db.Where("api_token = ?", token)
	})
}

// QueryWithOpenID 查询OpenID
func QueryWithOpenID(openid string) QueryOption {
	return QueryOptionFunc(func(db *gorm.DB) {
		db.Where("open_id = ?", openid)
	})
}
