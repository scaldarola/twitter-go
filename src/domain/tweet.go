package domain

type Tweet interface {
	GetText() string
	GetId() int
	GetUser() string
	SetId(id int)
	PrintableTweet() string
}

