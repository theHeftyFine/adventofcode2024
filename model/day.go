package model

type DayRunner interface {
	Parts() []func() int
}
