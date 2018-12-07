package service

import (
	"fmt"

	"github.com/scaldarola/twitter/src/domain"
)

var tweet domain.Tweet

var s = make([]domain.Tweet, 0)

var m = make(map[string][]domain.Tweet)

var id int

func InitializeService() {
	s = make([]domain.Tweet, 0)
	m = make(map[string][]domain.Tweet)
	id = 0
	return
}

type TweetManager struct {
	Id int
	S []domain.Tweet
	M map[string][]domain.Tweet
	Tweet domain.Tweet
}

func NewTweetManager() TweetManager{
	s = make([]domain.Tweet, 0)
	m = make(map[string][]domain.Tweet)
	id = 0
	twm := TweetManager{Id: id,S: s, M: m }
	return twm
}

func (twm *TweetManager) PublishTweet(tw domain.Tweet) (int, error) {

	if tw.GetUser() == "" {
		return 0, fmt.Errorf("el usuario no existe")
	}
	if tw.GetText() == "" {
		return 0, fmt.Errorf("no se puede publicar un tweet sin texto")
	}
	if len(tw.GetText()) > 140 {
		return 0, fmt.Errorf("no se puede publicar un tweet de mas de 140 caracteres")
	}
	twm.Tweet = tw
	twm.Tweet.SetId(id)
	twm.S = append(twm.S, twm.Tweet)
	twm.M[twm.Tweet.GetUser()] = append(twm.M[twm.Tweet.GetUser()], twm.Tweet)
	id ++
	return twm.Tweet.GetId(), nil
}

func (twm *TweetManager) GetTweet() domain.Tweet {
	return twm.Tweet
}

func (twm *TweetManager) GetTweets() []domain.Tweet {
	return twm.S
}

func (twm *TweetManager) GetTweetByID(id int) domain.Tweet{
	for _, tw := range twm.S {
		if tw.GetId() == id{
			return tw
		}
	}
	return nil
}

func (twm *TweetManager) GetTweetsByUser(user string) []domain.Tweet{
	return twm.M[user]
}

func (twm *TweetManager) CountTweetsByUser(user string) int{
	i := 0

	for _, tw := range twm.S {
		if(tw.GetUser() == user){
			i++
		}
	}

	return i
}
