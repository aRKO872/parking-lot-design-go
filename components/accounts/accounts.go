package accounts

type Account interface {
	Login() string
	Logout() bool
	ResetPassword() string
}