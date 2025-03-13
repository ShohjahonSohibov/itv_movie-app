package models

type Filter struct {
	LimitOffsetSort
}

type LimitOffsetSort struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Page      int    `json:"page"`
}
