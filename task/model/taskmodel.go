package model

type Task struct {
	Id        int32  `json:"id"`
	Title     string `json:"title"`
	Completed byte   `json:"completed"`
	CreateTs  string `json:"create_ts"`
}
