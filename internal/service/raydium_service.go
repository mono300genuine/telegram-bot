package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"github.com/getnimbus/ultrago/u_logger"

	"raybot/internal/entity"
	"raybot/pkg/caching"
)

type PoolType string

const (
	PoolType_ALL          PoolType = "all"
	PoolType_CONCENTRATED PoolType = "concentrated"
	PoolType_STANDARD     PoolType = "standard"
)

func NewRaydiumService() *RaydiumService {
	retryPolicy := failsafehttp.RetryPolicyBuilder().
		WithBackoff(time.Second, 60*time.Second).
		WithMaxRetries(3).
		OnRetryScheduled(func(e failsafe.ExecutionScheduledEvent[*http.Response]) {
			fmt.Println("Ping retry", e.Attempts(), "after delay of", e.Delay)
		}).
		Build()

	roundTripper := failsafehttp.NewRoundTripper(nil, retryPolicy)
	client := &http.Client{Transport: roundTripper}

	return &RaydiumService{
		client: client,
	}
}

type RaydiumService struct {
	client *http.Client
}

func (svc *RaydiumService) GetPools(ctx context.Context, poolType PoolType, page int) (*PoolQueryData, error) {
	ctx, logger := u_logger.GetLogger(ctx)

	queryData, err, hit := caching.MemoizeFunc("pools", map[string]interface{}{
		"poolType": poolType,
		"page":     page,
	}, func() (interface{}, error) {
		r, _ := http.NewRequest("GET", "https://api-v3.raydium.io/pools/info/list", nil)
		r.Header.Add("Content-Type", "application/json")
		q := r.URL.Query()
		q.Add("poolType", string(poolType))
		q.Add("poolSortField", "default")
		q.Add("sortType", "desc")
		q.Add("pageSize", "20")
		q.Add("page", strconv.Itoa(page))
		r.URL.RawQuery = q.Encode()

		resp, err := svc.client.Do(r)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var data = struct {
			ID      string `json:"id"`
			Success bool   `json:"success"`
			Data    struct {
				Count       int            `json:"count"`
				Data        []*entity.Pool `json:"data"`
				HasNextPage bool           `json:"hasNextPage"`
			} `json:"data"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, err
		}

		return &PoolQueryData{
			PoolType:    poolType,
			Page:        page,
			Data:        data.Data.Data,
			HasNextPage: data.Data.HasNextPage,
		}, nil
	})
	if err != nil {
		logger.Errorf("failed to get pools: %v", err)
		return nil, err
	}

	logger.Debugf("hit cache: %v", hit)

	return queryData.(*PoolQueryData), nil
}

type PoolQueryData struct {
	PoolType    PoolType
	Page        int
	Data        []*entity.Pool
	HasNextPage bool
}
