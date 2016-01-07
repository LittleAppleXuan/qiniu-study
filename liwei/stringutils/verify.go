package stringutils

import (
	"regexp"
)

func IsEmail(emailAddress string) bool {
	reg := regexp.MustCompile(`^(\w)+([\.|-]\w+)*@(\w)+((\.\w+)+)$`)
	return reg.MatchString(emailAddress)
}

func IsPhone(phone string) bool {
	reg := regexp.MustCompile(`^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$`)
	return reg.MatchString(phone)
}

//18位身份证校验
func IsIDCard(IDCard string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{4}$`)
	return reg.MatchString(IDCard)
}

//15位身份证校验
func IsIDCard15(IDCard string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$/`)
	return reg.MatchString(IDCard)
}
