package model

type Status int

const (
	CONFIRMED Status = iota
	PENDING
	DECLINED
	CANCELED
)
