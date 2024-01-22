package desensitization

import "strings"

func DesensitizationEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		// 如果@符号不存在
		return email
	}
	if atIndex == 1 {
		//如果邮箱@前只有一位
		return "*" + email[1:]
	}

	// 获取邮箱前缀和域名
	username := email[:atIndex]
	domain := email[atIndex:]

	// 脱敏处理
	maskedUsername := string(username[0]) + strings.Repeat("*", len(username)-2) + string(username[len(username)-1])

	// 组合脱敏后的邮箱地址
	maskedEmail := maskedUsername + domain
	return maskedEmail
}
func DesensitizationTel(tel string) string {
	//13941624270
	if len(tel) != 11 {
		return tel
	}
	return tel[:3] + "****" + tel[7:]
}
