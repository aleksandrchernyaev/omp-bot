package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	if args == "" {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Error, you do'nt input name of product",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("DemoSubdomainCommander.New: error sending reply message to chat - %v", err)
		}
		return
	}

	c.subdomainService.Add(args)
	products := c.subdomainService.List()

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You add new product: "+strconv.Itoa(len(products)-1)+" "+products[len(products)-1].Title,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}

}
