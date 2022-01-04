package model

import (
	"time"
)

type Partnership struct {
	Id            int
	PartnerIds    []int
	Start         time.Time
	End           time.Time
	Title         string
	Notes         string
	OwnerUsername string
}

type Child struct {
	Id            int
	ChildId       int
	ParentId      int
	Notes         string
	OwnerUsername string
}
