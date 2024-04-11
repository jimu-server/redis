package cache

import "time"

const (
	// Prefix 缓存前缀
	Prefix = "app"
	// Separator 缓存分隔符
	Separator = ":"

	// PhoneLoginKey 用户短信登录验证码key
	PhoneLoginKey = Prefix + Separator + "phone_login"

	// PhoneSecureKey 更换手机绑定验证码key
	PhoneSecureKey = Prefix + Separator + "phone_secure"

	// PhoneCodeTime 短信验证码登录时长
	PhoneCodeTime = 30 * time.Second
)
