package bot

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func botStart() {
	ctx := context.Background()
	botToken := "8542753002:AAH8H_hIBac4u4ZOKFBLcX_DvMSGbVTn4d8"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	for {
		select {
		case upd := <-updates:
			if upd.Message != nil {
				chatId := tu.ID(upd.Message.Chat.ID)

				_, _ = bot.CopyMessage(ctx,
					tu.CopyMessage(
						chatId,
						chatId,
						upd.Message.MessageID,
					),
				)
			}
		case <-ctx.Done():
			break
		}
	}

}
