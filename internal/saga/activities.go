package saga

import (
	"context"
	"time"

	"github.com/google/uuid"
	//"github.com/twitchtv/twirp"
	//"go.temporal.io/sdk/temporal"

	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/services"
)

// Activites is the implementation of a particular task in the business logic.
// @see https://github.com/temporalio/documentation/blob/master/docs/go/activities.md

type Activities struct {
}

// @return an err or a value, err
func (a *Activities) Withdraw(ctx context.Context, request *services.CreateWithdrawRequest) (*string, error) {
	time.Sleep(20 * time.Second)
	// NOTICE: real application should request e.g. RPC/API endpoint

	//var twerr twirp.Error
	//if ok := errors.As(err, &twerr); ok {
	// 	twerr
	//	return nil, temporal.NewApplicationError(twerr.Msg(), TwirpErrorCode, twerr.Code())
	// }
	// unkwnerr
	// return nil, temporal.NewApplicationError(err.Error(), "unkwnerr")
	
	// success return
	txId := uuid.New().String()
	return &txId, nil
	
	// failure return
	//twerr := twirp.InternalError("internal server error")
	//return nil, temporal.NewApplicationError(twerr.Msg(), "twerr", twerr.Code())
}
