package model

import (
	"time"
)

type Partnership struct {
	Id      int
	Person1 int
	Person2 int
	Start   time.Time
	End     time.Time
	Title   string
}

type Child struct {
	Id       int
	ChildId  int
	ParentId int
}
