package chat

import (
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/cmd/bot"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/utils"
	tele "gopkg.in/telebot.v3"
)

func NewCommandHandler(bot *bot.TgBot) *Handler {
	return &Handler{bot: bot}
}

const welcomeMessageTemplateRu = `

Чтобы начать работу – нажми на кнопку ниже 👇🏻

`

const welcomeMessageTemplateEn = `

To get started, click on the button below 👇🏻
`

func (h *Handler) StartHandler(ctx tele.Context) error {
	menu := &tele.ReplyMarkup{}
	tmaButton := &tele.Btn{Text: "Запустить / Launch 🃏", WebApp: &tele.WebApp{URL: h.bot.TmaURL} }

	menu.Inline(
		menu.Row(*tmaButton),
	)

	return ctx.Send(utils.SumStrings(
		ctx.Sender().FirstName,
		", приветствую! ♥️",
		welcomeMessageTemplateRu,
		ctx.Sender().FirstName,
		", hello! ♥️",
		welcomeMessageTemplateEn), menu)
}