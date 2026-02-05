package bot

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var waitInput = make(map[int64]bool)
var userInput = make(map[int64]string)

func botStart() {

	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	for upd := range updates {
		chatId := tu.ID(upd.Message.Chat.ID)
		if upd.Message != nil {
			if upd.Message.Text == "/start" {
				bot.SendMessage(ctx,
					tu.Message(chatId,
						"Hello!"),
				)
			}

			if upd.Message.Text == "/input" {

				message := tu.Message(chatId,
					"Отправьте данные в формате ГГГГ-ММ-ДД/МинВремяотправления/МаксВремяотправления/городОтправки/городПрибытия/времяПоиска/времяЗапроса")

				bot.SendMessage(ctx, message)

				waitInput[chatId.ID] = true

				continue

			}

			if upd.Message.Text == "/myInput" {

				message := tu.Message(chatId,
					userInput[chatId.ID])

				bot.SendMessage(ctx, message)

				continue

			}

			if waitInput[chatId.ID] == true {

				userInput[chatId.ID] = upd.Message.Text

				continue

			}
		}

	}

}
