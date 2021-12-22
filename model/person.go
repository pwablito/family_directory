package model

type Person struct {
	Id            int
	Name          string
	Birthdate     string // TODO Change to time.Time
	Email         string
	Phone         string
	OwnerUsername string
}
