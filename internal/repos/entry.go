package repos

import (
	"Zeus/pkg/client"

	"gorm.io/gorm"
)

type (
	entry struct {
		mysql *gorm.DB
		gorm  InterGormClient
	}

	EntryInter interface {
		MySQL() *gorm.DB
		Gorm() InterGormClient
		User() UserRepoInter
	}
)

func NewEntryRepo() EntryInter {
	mysql := client.InitDB()
	return &entry{
		mysql: mysql,
		gorm:  NewGormClient(mysql),
	}
}

func (e entry) MySQL() *gorm.DB       { return e.mysql }
func (e entry) Gorm() InterGormClient { return e.gorm }
func (e entry) User() UserRepoInter   { return NewRepoUser(e.mysql) }
