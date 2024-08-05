package response

import (
	"example.com/internal/infra/database/ent"
)

type UserResponse struct {
	ID   int    `json:"id,omitempty"`
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

func ToUseResponse(entity *ent.User) *UserResponse {
	if entity == nil {
		return nil
	}
	return &UserResponse{
		ID:   entity.ID,
		Age:  entity.Age,
		Name: entity.Name,
	}
}
