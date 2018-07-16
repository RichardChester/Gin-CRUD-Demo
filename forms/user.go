package forms

type CreateUserCommand struct {
  Username 		string  `json:"username" binding:"required"`
  Password		string  `json:"password" binding:"required"`
  Age 				int	    `json:"age" binding:"required"`
  Sex			    string	`json:"sex" binding:"required"`
}

type UpdateUserCommand struct {
  Username 		string  `json:"username" binding:"required"`
  Password		string  `json:"password" binding:"required"`
  Age 				int	    `json:"age" binding:"required"`
  Sex			    string	`json:"sex" binding:"required"`
}
