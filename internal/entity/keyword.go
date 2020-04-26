package entity

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/pkg/txt"
)

// Keyword used for full text search
type Keyword struct {
	ID      uint   `gorm:"primary_key"`
	Keyword string `gorm:"type:varchar(64);index;"`
	Skip    bool
}

// NewKeyword registers a new keyword in database
func NewKeyword(keyword string) *Keyword {
	keyword = strings.ToLower(txt.Clip(keyword, txt.ClipKeyword))

	result := &Keyword{
		Keyword: keyword,
	}

	return result
}

// FirstOrCreate checks wether the keyword already exist in the database
func (m *Keyword) FirstOrCreate(db *gorm.DB) *Keyword {
	mutex.Db.Lock()
	defer mutex.Db.Unlock()

	if err := db.FirstOrCreate(m, "keyword = ?", m.Keyword).Error; err != nil {
		log.Errorf("keyword: %s", err)
	}

	return m
}
