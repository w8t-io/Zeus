package repos

import (
	"Zeus/pkg/client"
	"gorm.io/gorm"
)

type (
	entry struct {
		mysql *gorm.DB
	}

	EntryInter interface {
		MySQL() *gorm.DB
		User() UserRepoInter
	}
)

func NewEntryRepo() EntryInter {
	return &entry{
		mysql: client.InitDB(),
	}
}

func (e entry) MySQL() *gorm.DB     { return e.mysql }
func (e entry) User() UserRepoInter { return NewRepoUser(e.mysql) }
