package model

import (
	"time"
)

type Marriage struct {
	Id      int
	Person1 int
	Person2 int
	Start   time.Time
	End     time.Time
}

type Child struct {
	Id        int
	ChildId   int
	Parent1Id int
	Parent2Id int
}
