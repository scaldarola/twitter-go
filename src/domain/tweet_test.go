package domain_test

import (
	"github.com/scaldarola/twitter/src/domain"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T){
	tweet := domain.NewTweet("scalda", "Este es mi tweet")

	text := tweet.String()

	expected := "0 @scalda: Este es mi tweet"

	if text != expected {
		t.Errorf("Se esperaba " + expected + " pero fue " + text)
	}
}

func TestCanGetAStringFromATweet(t *testing.T){
	tweet := domain.NewTweet("scalda", "Este es mi tweet")

	text := tweet.String()

	expected := "0 @scalda: Este es mi tweet"

	if text != expected {
		t.Errorf("Se esperaba " + expected + " pero fue " + text)
	}
}

func TestImageTweetPrintsUserTextAndURL(t *testing.T){
	tweet := domain.NewImageTweet("scalda", "Esta es la foto", "https://www.google.com.ar/url?sa=i&source=images&cd=&ved=2ahUKEwidir3f_4vfAhVKipAKHRb7D0cQjRx6BAgBEAU&url=https%3A%2F%2Flosandes.com.ar%2Farticle%2Fview%3Fslug%3Dboca-palmeiras-hora-y-donde-verlo-por-tv&psig=AOvVaw3MKfMeRNVLvBuoF5ma7FVw&ust=1544212963357934")

	text := tweet.PrintableTweet()

	expectedtext := "0 @scalda: Esta es la foto https://www.google.com.ar/url?sa=i&source=images&cd=&ved=2ahUKEwidir3f_4vfAhVKipAKHRb7D0cQjRx6BAgBEAU&url=https%3A%2F%2Flosandes.com.ar%2Farticle%2Fview%3Fslug%3Dboca-palmeiras-hora-y-donde-verlo-por-tv&psig=AOvVaw3MKfMeRNVLvBuoF5ma7FVw&ust=1544212963357934"

	if text != expectedtext{
		t.Errorf("El tweet no es el esperado")
	}
}


func TestQuoteTweet_PrintableTweet(t *testing.T) {
	quotedTw := domain.NewQuotedTweet("scalda", "jaja que gracioso", domain.NewTweet("sangrexeneize", "riber te fuiste a la b"))

	text := quotedTw.PrintableTweet()

	expected := "0 @scalda: jaja que gracioso " + "------0 @sangrexeneize: riber te fuiste a la b------"

	if text != expected {
		t.Errorf("El tweet esperado era %s pero fue %s", expected, text)
	}
}