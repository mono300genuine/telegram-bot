package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestRaydiumService(t *testing.T) {
	ctx := context.Background()

	convey.Convey("TestRaydiumService", t, func() {
		svc := NewRaydiumService()

		convey.Convey("GetAllPools", func() {
			pools, err := svc.GetPools(ctx, PoolType_ALL, 1)
			convey.So(err, convey.ShouldBeNil)

			res, _ := json.Marshal(pools.Data)
			fmt.Println(string(res))
		})
	})
}
