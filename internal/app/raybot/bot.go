package raybot

import (
	"context"
	"time"

	tele "gopkg.in/telebot.v3"

	"raybot/internal/conf"
	"raybot/internal/handler"
	"raybot/internal/service"
)

var (
	selector = &tele.ReplyMarkup{}

	btnPrev = selector.Data("⬅", "prev")
	btnNext = selector.Data("➡", "next")
)

func NewRayBot(
	poolHandler *handler.PoolHandler,
) *RayBot {
	ctx := context.Background()
	pref := tele.Settings{
		Token:     conf.Config.BotToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/start", poolHandler.HelpHandler)
	b.Handle("/help", poolHandler.HelpHandler)
	b.Handle("/allpools", func(c tele.Context) error {
		return poolHandler.PoolHandler(ctx, c, service.PoolType_ALL)
	})
	b.Handle("/concentratedpools", func(c tele.Context) error {
		return poolHandler.PoolHandler(ctx, c, service.PoolType_CONCENTRATED)
	})
	b.Handle("/standardpools", func(c tele.Context) error {
		return poolHandler.PoolHandler(ctx, c, service.PoolType_STANDARD)
	})
	b.Handle(&btnPrev, func(c tele.Context) error {
		c.Message().Payload = c.Callback().Data
		return poolHandler.PoolHandler(ctx, c, "")
	})
	b.Handle(&btnNext, func(c tele.Context) error {
		c.Message().Payload = c.Callback().Data
		return poolHandler.PoolHandler(ctx, c, "")
	})

	return &RayBot{
		bot: b,
	}
}

type RayBot struct {
	bot *tele.Bot
}

func (rb *RayBot) Start() error {
	rb.bot.Start()
	return nil
}

func (rb *RayBot) Stop() error {
	rb.bot.Stop()
	return nil
}
