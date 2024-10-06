package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/demo/subdomain"
)

func (c *DemoSubdomainCommander) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.subdomainService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You delete product: "+product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}

	products := c.subdomainService.List()
	products[idx] = products[len(products)-1]
	products[len(products)-1] = subdomain.Subdomain{}
	products = products[:len(products)-1]

}
