package bot

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func InitBot(token string) *tele.Bot {
	pref := tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{Timeout: 10},
		Offline: true,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("error initializing bot %v", err)
	}

	return b

}