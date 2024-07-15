package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/getnimbus/ultrago/u_logger"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/slongfield/pyfmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	tele "gopkg.in/telebot.v3"

	"raybot/internal/entity"
	"raybot/internal/service"
)

func NewPoolHandler(
	svc *service.RaydiumService,
) *PoolHandler {
	return &PoolHandler{
		svc: svc,
	}
}

type PoolHandler struct {
	svc *service.RaydiumService
}

func (h *PoolHandler) HelpHandler(c tele.Context) error {
	return c.Send(`☀️ Raydium Trader Bot ☀️
The premium bot is here to help you discovery all LP on Raydium without headache 🚀


/allpools - Show all liquidity pools
/concentratedpools - Show concentrated liquidity pools
/standardpools - Show standard liquidity pools
`)
}

func (h *PoolHandler) PoolHandler(ctx context.Context, c tele.Context, poolType service.PoolType) error {
	ctx, logger := u_logger.GetLogger(ctx)

	var (
		split = strings.Split(c.Message().Payload, " ")
		page  int
	)
	if len(split) > 0 {
		var err error
		page, err = strconv.Atoi(split[0])
		if err != nil {
			page = 1
		}
	}
	if len(split) > 1 {
		poolType = service.PoolType(split[1])
	}

	logger.Infof("User id %d - username %s requested pool data %s", c.Sender().ID, c.Sender().Username, string(poolType))

	// get pools data
	poolData, err := h.svc.GetPools(context.Background(), poolType, page)
	if err != nil {
		return c.Send("failed to get pools data")
	} else if len(poolData.Data) == 0 {
		return c.Send("no pools data")
	}

	// build message
	msg := h.BuildMessage(poolData)

	// build inline button
	btn := make([]tele.InlineButton, 0)
	if page > 1 {
		btn = append(btn, tele.InlineButton{Text: "Prev", Unique: "prev", Data: fmt.Sprintf("%d %s", page-1, poolType)})
	}
	if poolData.HasNextPage {
		btn = append(btn, tele.InlineButton{Text: "Next", Unique: "next", Data: fmt.Sprintf("%d %s", page+1, poolType)})
	}

	return c.Send(msg, &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			btn,
		},
	})
}

func (h *PoolHandler) BuildMessage(poolData *service.PoolQueryData) string {
	tw := table.NewWriter()
	p := message.NewPrinter(language.English)
	tw.AppendHeader(table.Row{"Pair", "Liquidity", "24h Volume", "24h Fee", "24h APR"})
	tw.AppendRows(lo.Map(poolData.Data, func(item *entity.Pool, _ int) table.Row {
		symbolMintA := item.MintA.Symbol
		if symbolMintA == "WSOL" {
			symbolMintA = "SOL"
		}
		symbolMintB := item.MintB.Symbol
		if symbolMintB == "WSOL" {
			symbolMintB = "SOL"
		}

		return table.Row{
			fmt.Sprintf("%s-%s", symbolMintA, symbolMintB),
			p.Sprintf("$%.0f", item.Tvl),
			p.Sprintf("$%.0f", item.Day.Volume),
			p.Sprintf("$%.0f", item.Day.VolumeFee),
			p.Sprintf("%.2f%%", item.Day.Apr),
		}
	}))

	msg := pyfmt.Must(`🌎 {formattedTime}

🚀 Raydium {poolType} Liquidity Pools - Page {page} 🚀

{poolTable}
`, map[string]interface{}{
		"formattedTime": time.Now().Format(time.DateTime),
		"poolType":      strings.Title(string(poolData.PoolType)),
		"page":          poolData.Page,
		"poolTable":     fmt.Sprintf("```%s```", tw.Render()),
	})

	return msg
}
