package domain

import (
"strconv"
"time"
)
type TextTweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

type ImageTweet struct {
	TextTweet
	Url string
}

type QuoteTweet struct {
	TextTweet
	quotedTweet Tweet
}

func NewTweet(us string, tx string) *TextTweet{
	var p = time.Now()
	//var id,_ := service.PublishTweet()
	tw:= TextTweet{User :us, Text: tx, Date: &p}
	return &tw
}

func NewImageTweet(us string, tx string, url string) *ImageTweet{
	var p = time.Now()
	imgtw := ImageTweet{TextTweet: TextTweet{User :us, Text: tx, Date: &p}, Url: url}
	return &imgtw
}

func NewQuotedTweet(us string, tx string, tw Tweet) *QuoteTweet{
	var p = time.Now()
	quotw := QuoteTweet{TextTweet: TextTweet{User: us, Text: tx, Date: &p}, quotedTweet: tw}
	return &quotw
}

func (tw *TextTweet) String() string{
	return strconv.Itoa(tw.Id) + " @" + tw.User + ": " + tw.Text
}

func (tw *TextTweet) GetText() string{
	return tw.Text
}

func (tw *TextTweet) GetUser() string{
	return tw.User
}

func (tw *TextTweet) GetId() int{
	return tw.Id
}

func (tw *TextTweet) SetId(id int){
	tw.Id = id
}

func (tw *TextTweet) PrintableTweet() string{
	return strconv.Itoa(tw.Id) + " @" + tw.User + ": " + tw.Text
}

func (tw *ImageTweet) PrintableTweet() string{
	return strconv.Itoa(tw.Id) + " @" + tw.User + ": " + tw.Text + " " + tw.Url
}

func (tw *QuoteTweet) PrintableTweet() string{
	return strconv.Itoa(tw.Id) + " @" + tw.User + ": " + tw.Text + " ------" + tw.quotedTweet.PrintableTweet() + "------"
}

type Stringer interface {
	String() string
}