package service_test

import (
	"testing"

	"github.com/scaldarola/twitter/src/domain"
	"github.com/scaldarola/twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet domain.TextTweet
	twm := service.NewTweetManager()
	user := "scaldarola"
	text := "hola mundo twitter"
	tweet = *domain.NewTweet(user, text)
	//var tweet string = "Mi primer Tweet"

	//tweet := tw

	var err error
	var id1 int
	id1, err = twm.PublishTweet(tweet)

	if err != nil && err.Error() != "el usuario no existe" {
		t.Error("Error controlado, el usuario no existe")

	}

	publishedTweet := twm.GetTweet()

	if !isValidTweet(t, publishedTweet, id1, user, text) {
		t.Error("El tweet no es valido")
		return
	}

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet domain.TextTweet
	twm := service.NewTweetManager()
	user := "scaldarola"
	text := ""
	tweet = *domain.NewTweet(user, text)

	//tweet := tw

	_, err := twm.PublishTweet(tweet)

	if err != nil && err.Error() != "no se puede publicar un tweet sin texto" {
		t.Errorf("Error!")

	}

}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	twm := service.NewTweetManager()
	var tweet domain.TextTweet

	user := "scaldarola"
	text := "aaaaaassdlajdsaidasjidjadjsfkjfkajdskajdakssakdaskdakdojfijfanckmkasmfdakfodkfodskfldsmvkdsnjsiafpjsdcjdslkñcmsdlfksojkfa`jdapijfpiajfdipasjdjsaipdjadpiajdfpias"
	tweet = *domain.NewTweet(user, text)

	//tweet := tw

	_, err := twm.PublishTweet(tweet)

	if err != nil && err.Error() != "no se puede publicar un tweet de mas de 140 caracteres" {
		t.Errorf("Error!")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweeT(t *testing.T) {


	twm := service.NewTweetManager()

	var tweet, secondTweet domain.TextTweet

	user := "scaldarola"
	text1 := "Hola mundo de twitter"
	tweet = *domain.NewTweet(user, text1)

	//user := "scaldarola"
	text2 := "Gano Boca"
	secondTweet = *domain.NewTweet(user, text2)

	id1,_ := twm.PublishTweet(tweet)
	id2,_ := twm.PublishTweet(secondTweet)

	publishedTweets := twm.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Se esperaban 2 tweets pero llegaron %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, &firstPublishedTweet,id1, user, text1) {
		t.Error("El tweet no es valido")
		return
	}

	if !isValidTweet(t, &secondPublishedTweet, id2, user, text1) {
		t.Error("El tweet no es valido")
		return
	}

}

func isValidTweet(t *testing.T, tw *domain.TextTweet, id int, user string, text string) bool {

	if tw.User != user && tw.Text != text {
		t.Errorf("El tweet esperado era %s: %s \n pero es %s: %s", user, text, tw.User, tw.Text)
		return false
	}

	if tw.Date == nil {
		t.Error("La fecha del tweet no puede ser nula")
		return false
	}

	return true

}

func TestCanRetrieveTweetById(t *testing.T){
	twm := service.NewTweetManager()

	var tweet domain.TextTweet
	var id int

	user:= "scaldarola"
	text:= "hola twitter"

	tweet = *domain.NewTweet(user,text)

	id,_ = twm.PublishTweet(tweet)

	publishedTweet := *twm.GetTweetByID(id)

	isValidTweet(t, &publishedTweet, id, user, text)

}


func TestCanCountTweetsSentByAnyUser(t *testing.T){
	twm := service.NewTweetManager()

	var tweet1, tweet2, tweet3 *domain.TextTweet

	user := "scaldarola"
	user2 := "nlopez"

	text := "hola"
	text2 := "comotan"

	tweet1 = domain.NewTweet(user, text)
	tweet2 = domain.NewTweet(user2, text)
	tweet3 = domain.NewTweet(user2, text2)

	twm.PublishTweet(*tweet1)
	twm.PublishTweet(*tweet2)
	twm.PublishTweet(*tweet3)

	count := twm.CountTweetsByUser(user)

	if count != 1{
		t.Errorf("Se esperaba 1 pero llego %d", count)
	}


}

func TestCanRetrieveTheTweetsSentByAnyUser(t *testing.T){
	twm := service.NewTweetManager()

	var tweet1, tweet2, tweet3 *domain.TextTweet

	user := "scaldarola"
	user2 := "nlopez"

	text := "hola"
	text2 := "comotan"

	tweet1 = domain.NewTweet(user, text)
	tweet2 = domain.NewTweet(user2, text)
	tweet3 = domain.NewTweet(user2, text2)

	twm.PublishTweet(*tweet1)
	twm.PublishTweet(*tweet2)
	twm.PublishTweet(*tweet3)

	tweets := twm.GetTweetsByUser(user2)

	if len(tweets) != 2{
		t.Errorf("Se esperaba 2 pero llego %d", len(tweets))
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, &firstPublishedTweet, firstPublishedTweet.Id, user2, text) {
		t.Error("El tweet no es valido")
		return
	}

	if !isValidTweet(t, &secondPublishedTweet, secondPublishedTweet.Id, user2, text2) {
		t.Error("El tweet no es valido")
		return
	}

}


