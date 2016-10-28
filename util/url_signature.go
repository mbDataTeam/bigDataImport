package util

import "github.com/stretchr/signature"

const privateKey = "!QAZ@WSX#$%^&*()"
//validate sign url
func ValidateSignUrl(requestUrl string) bool  {
	validate,_:= signature.ValidateSignature("Get",requestUrl,"",privateKey)
	return  validate
}

func GetSignUrl(requestUrl string) string {
	
	signedUrl,_ := signature.GetSignedURL("GET", requestUrl, "", privateKey)
	//signature.GetSignature("GET", requestUrl, "ABC123", "ABC123-private")
		//signature.GetSignedURL("Get",requestUrl,"body", privateKey)
	return signedUrl
}

func GetSign(requestUrl string) string {
	signed,_ := signature.GetSignature("GET", requestUrl, "", privateKey)
	return signed
}
