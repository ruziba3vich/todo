package authentication

func CheckPassword(password, hashshedPwd string) bool {
	newPwd, _ := HashPassword(password)
	return newPwd == hashshedPwd;
}
