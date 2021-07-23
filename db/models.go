package db

import (
	"regexp"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string
	Done time.Time
	Tags []*Tag `gorm:"many2many:task_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	t.updateTags()
	return nil
}
func (t *Task) BeforeSave(tx *gorm.DB) (err error) {
	t.updateTags()
	return nil
}

var tagMatcher = regexp.MustCompile(`#\w+`)

func (task *Task) updateTags() {
	dbTags := []*Tag{}
	for _, tag := range tagMatcher.FindAllString(task.Name, -1) {
		print(tag)
		var dbTag Tag
		Conn.Where(Tag{Name: tag}).FirstOrInit(&dbTag)
		dbTags = append(dbTags, &dbTag)
	}
	task.Tags = dbTags
}

type Tag struct {
	gorm.Model
	Name  string
	Tasks []*Task `gorm:"many2many:task_tags"`
}
