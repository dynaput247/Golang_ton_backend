package main

import (
	"time"

	route "github.com/amitshekhariitbhu/go-backend-clean-architecture/api/route"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/cmd/bot"
	chat "github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/bot"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()


	tgBot := bot.TgBot{
		Bot: bot.InitBot(cfg.BotToken),
		BotID: cfg.BotID,
		ChannelID: cfg.ChannelID,
		AdminUserID: cfg.AdminUserID,
		TmaURL: cfg.TmaURL,
	}

	go func() {
		tgBot.Bot.Start()
	}()

	defer tgBot.Bot.Stop()
	registerBotHandlers(tgBot)

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)	
}

func registerBotHandlers(tgBot bot.TgBot) {
	commandHandler := chat.NewCommandHandler(&tgBot)
	tgBot.Bot.Handle("/start", commandHandler.StartHandler)
}
