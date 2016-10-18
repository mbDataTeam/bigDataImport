package util

import "github.com/stretchr/signature"

const privateKey  = "!QAZ@WSX#$%^&*()"
//validate sign url
func ValidateSignUrl(requestUrl string) bool  {
	if validate,error:= signature.ValidateSignature("Get",requestUrl,"",privateKey); error != nil{
		return  validate
	}
	return false
}
