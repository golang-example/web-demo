package user

type UserRequest struct {
	UserName 	string 		`json:"userName"`
	Pwd        	string 		`json:"pwd"`
}

type UserResponse struct {
	Code		int		     `json:"code"`
	Msg 		string 		 `json:"msg"`
}
