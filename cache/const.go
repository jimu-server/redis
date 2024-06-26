package cache

import "time"

const (
	// Prefix 缓存前缀
	Prefix = "jimuos"
	// Separator 缓存分隔符
	Separator = ":"

	// PhoneLoginKey 用户短信登录验证码key
	PhoneLoginKey = Prefix + Separator + "phone_login" + Separator
	ForGetKey     = Prefix + Separator + "forget" + Separator

	// PhoneSecureKey 更换手机绑定验证码key
	PhoneSecureKey = Prefix + Separator + "phone_secure" + Separator
	// PhoneCodeTime 短信验证码登录时长
	PhoneCodeTime = 30 * time.Second

	// EmailVerifyKey  邮箱绑定验证缓存key
	EmailVerifyKey = Prefix + Separator + "email_verify" + Separator
	// EmailVerifyCodeTime 邮箱验证码缓存时长
	EmailVerifyCodeTime = 30 * time.Minute

	SettingCacheTime = 30 * time.Minute
)
