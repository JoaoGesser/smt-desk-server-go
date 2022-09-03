package taks

import (
	"smt-desk-server/pkg/dto/group"
	"smt-desk-server/pkg/dto/user"
)

type TaskDTO struct {
	InternalIdentification string
	ExternalIdentification string
	Title                  string
	Group                  group.GroupDTO
	User                   user.UserDTO
	TimesSpent             []TimeSpentDTO
}
