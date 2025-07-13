package user

type UserResponse struct {
    ID    *uint  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UsersResponse struct {
    Users []UserResponse `json:"users"`
}