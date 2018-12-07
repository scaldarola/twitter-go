package main

import (
	"github.com/abiosoft/ishell"
	"github.com/scaldarola/twitter/src/domain"
	"github.com/scaldarola/twitter/src/service"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Ingrese su nombre de usuario \n")
	user := shell.ReadLine()
	twm := service.NewTweetManager()

	if user != "" {
		shell.Print("Type 'help' to know commands\n")
		shell.AddCmd(&ishell.Cmd{
			Name: "tweet",
			Help: "Publishes a tweet",
			Func: func(c *ishell.Context) {

				defer c.ShowPrompt(true)

				c.Print("Write your tweet: ")

				var tweet *domain.TextTweet = domain.NewTweet(user, c.ReadLine())

				twm.PublishTweet(tweet)

				c.Print("Tweet sent\n")

				return
			},
		})

			shell.AddCmd(&ishell.Cmd{
				Name: "quote",
				Help: "Lista tweets y permite citarlos, creando un nuevo tweet",
				Func: func(c *ishell.Context) {

					defer c.ShowPrompt(true)

					c.Print("Elije tu tweet: \n")

					tweets := twm.GetTweets()
					//i := 0
					for i := 0; i < len(tweets); i++ {
						c.Println(tweets[i].PrintableTweet())
					}

					c.Println("Ingrese el Id del tweet a citar")

					id, _ := strconv.Atoi(c.ReadLine())

					tweet := twm.GetTweetByID(id)

					if tweet != nil {

						c.Println("Ingrese el mensaje")

						var tweetp *domain.QuoteTweet = domain.NewQuotedTweet(user, c.ReadLine(), tweet)

						twm.PublishTweet(tweetp)

						c.Print("Tweet sent\n")

						//c.Println(tweet.PrintableTweet())

						return
					}

					c.Println("El tweet no existe")

					return
				},
			})

		shell.AddCmd(&ishell.Cmd{
			Name: "image",
			Help: "Permite twittear con imagen",
			Func: func(c *ishell.Context) {

				defer c.ShowPrompt(true)

				c.Print("Write your tweet: ")

				text := c.ReadLine()

				c.Println("Ingrese URL de imagen")

				var tweet *domain.ImageTweet = domain.NewImageTweet(user, text, c.ReadLine())

				twm.PublishTweet(tweet)

				c.Print("Tweet sent\n")

				return
			},
		})

			shell.AddCmd(&ishell.Cmd{
				Name: "showTweet",
				Help: "Shows a tweet",
				Func: func(c *ishell.Context) {

					defer c.ShowPrompt(true)

					tweet := twm.GetTweet()

					c.Println(tweet.PrintableTweet())

					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "list",
				Help: "Lista Todos los Tweets",
				Func: func(c *ishell.Context) {

					defer c.ShowPrompt(true)

					tweets := twm.GetTweets()
					//i := 0
					for i := 0; i < len(tweets); i++ {
						c.Println(tweets[i].PrintableTweet())
					}

					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "buscarid",
				Help: "Busca Tweet por ID",
				Func: func(c *ishell.Context) {

					defer c.ShowPrompt(true)

					id, _ := strconv.Atoi(c.ReadLine())

					tweet := twm.GetTweetByID(id)

					if tweet != nil {
						c.Println(tweet.PrintableTweet())
						return
					}

					c.Println("El tweet no existe")

					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "countByUser",
				Help: "cuenta Tweets del usuario ingresado",
				Func: func(c *ishell.Context) {
					defer c.ShowPrompt(true)

					c.Println("Ingrese el usuario que quiere buscar su cantidad de tweets")

					user := c.ReadLine()

					num := twm.CountTweetsByUser(user)

					c.Println("El usuario " + user + " tiene  " + strconv.Itoa(num) + "tweet(s) publicado(s).")

					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "buscarUser",
				Help: "lista Tweets del usuario ingresado",
				Func: func(c *ishell.Context) {
					defer c.ShowPrompt(true)

					c.Println("Ingrese el usuario que quiere buscar sus tweets")

					user := c.ReadLine()

					s := twm.GetTweetsByUser(user)

					c.Println("Mostrando tweets del usuario " + user)

					for _, tw := range s {
						c.Println(tw.PrintableTweet())
					}

					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "changeUser",
				Help: "Cambia de usuario",
				Func: func(c *ishell.Context) {
					c.Println("Ingrese el nuevo nombre de usuario")
					user = c.ReadLine()
					c.Println("Hola usuario " + user + "!")
					return
				},
			})

			shell.AddCmd(&ishell.Cmd{
				Name: "logout",
				Help: "Sale de Twitter",
				Func: func(c *ishell.Context) {

					user = ""
					c.Println("Chau!")
					shell.Close()
				},
			})
			shell.Run()
		} else {
			shell.Println("No se puede twittear. El user no es válido")
		}
}
