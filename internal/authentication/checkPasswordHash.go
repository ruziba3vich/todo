package authentication

import "fmt"

func CheckPassword(password, hashshedPwd string) bool {
	newPwd, _ := HashPassword(password)
	fmt.Println(newPwd)
	fmt.Println(hashshedPwd)
	fmt.Println(hashshedPwd == newPwd)
	return newPwd == hashshedPwd;
}
