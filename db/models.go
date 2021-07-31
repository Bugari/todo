package db

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Seq      int       `gorm:"index;"`
	Name     string
	Priority int `gorm:"default:3"`
	Due      *time.Time
	Done     *time.Time
	Tags     []*Tag `gorm:"many2many:task_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

const prioritiesCodes = "hml"

func (t *Task) SetPriority(priority string) {
	if len(priority) == 0 {
		t.Priority = 4
		return
	}
	var priorityIdx = strings.Index(prioritiesCodes, string([]rune(strings.ToLower(priority))[0]))
	if priorityIdx == -1 {
		priorityIdx = 4
		return
	}
	t.Priority = priorityIdx + 1
}

var prioritiesArr = []string{"High", "Medium", "Low", ""}

func (t *Task) GetPriority() string {
	return prioritiesArr[t.Priority-1]
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.updateTags()
	return nil
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
