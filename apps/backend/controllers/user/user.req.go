package user

type CreateUserRequest struct {
    Name     string `json:"name" binding:"required,min=2,max=50"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
    ID    uint   `json:"id" binding:"required"`
    Name  string `json:"name" binding:"omitempty,min=2,max=50"`
    Email string `json:"email" binding:"omitempty,email"`
    Password string `json:"password" binding:"required,min=8,max=100"`
}