package decimal

type Sign int

const (
	Positive = Sign(1)
	Negative = Sign(-1)
)

type Compare int

const (
	Less    = Compare(-1)
	Equal   = Compare(0)
	Greater = Compare(1)
)
