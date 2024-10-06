package subdomain

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Edit(inputMessage *tgbotapi.Message) {

	arg_string := inputMessage.CommandArguments()

	args := strings.SplitN(arg_string, " ", 2)
	if len(args) != 2 {
		SendMsgErrorFormat(c, inputMessage, arg_string)
		return
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil {
		SendMsgErrorFormat(c, inputMessage, arg_string)
		return
	}

	NewTitle := args[1]
	if NewTitle == "" {
		SendMsgErrorFormat(c, inputMessage, arg_string)
		return
	}

	err = c.subdomainService.Edit(idx, NewTitle)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			err.Error(),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You edit product: "+strconv.Itoa(idx)+" "+NewTitle,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}

}

func SendMsgErrorFormat(c *DemoSubdomainCommander, inputMessage *tgbotapi.Message, arg_string string) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You input incorrect format arguments: "+arg_string+"/n; use correct format: 1 one",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}

}
