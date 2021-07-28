package presenter

import (
	"github.com/marlonflying/chatroom/chatservice/entity"
)

//User data
type User struct {
	ID    entity.ID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
