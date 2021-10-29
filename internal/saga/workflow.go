package saga

import (
	"errors"
	"time"

	"github.com/twitchtv/twirp"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/enums"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/resources"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/services"

	temporalenumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// SagaWorkflow In the Temporal Go SDK programming model, a Workflow is an exportable
// function that adheres to a set of rules.
//
// We recommend Saga for long-lived transactions. In a microservices-based application,
// you expect interservice calls and communication with third-party systems.
// Therefore, it's best to design for eventual consistency: retry for recoverable errors and expose
// compensating events that eventually amend non-recoverable errors.
//
// Command:
// - ExecuteWorkflow
// - SignalWorkflow (Signals provide a mechanism to send data directly in to a running Workflow)
// - QueryWorkflow
//
// Logic:
// - selectors
//
// @see https://github.com/temporalio/documentation/blob/master/docs/go/workflows.md
// 		https://github.com/temporalio/documentation/blob/master/docs/go/selectors.md
// @param ctx workflow.Context similarly to the standard context.Context
// @param request *services.CreateWithdrawRequest is a custom parameter that can be used to pass data into the Workflow when it starts
// @return an err or a value, err
func SagaWorkflow(ctx workflow.Context, request *services.CreateWithdrawRequest) (*resources.Withdrawal, error) {
	logger := workflow.GetLogger(ctx)
	workflowInfo := workflow.GetInfo(ctx)
	createdAt := workflowInfo.WorkflowStartTime
	resourcesWithdrawal := &resources.Withdrawal{
		Id:        workflowInfo.WorkflowExecution.ID,
		Status:    enums.WithdrawalStatus_WITHDRAWAL_STATUS_RUNNING,
		Amount:    request.Amount,
		CreatedAt: createdAt.Format(time.RFC3339),
		UpdatedAt: createdAt.Format(time.RFC3339),
		UserId:    request.UserId,
		Workflow: &resources.Workflow{
			WorkflowId:    workflowInfo.WorkflowExecution.ID,
			WorkflowRunId: workflowInfo.WorkflowExecution.RunID,
			Status:        enums.WorkflowStatus_WORKFLOW_STATUS_UNSPECIFIED, // workflowInfo doesn't contain `Status` property 
			Namespace:     workflowInfo.Namespace,
		},
	}
	// set query handler for QuerySagaWorkflow
	err := workflow.SetQueryHandler(ctx, QuerySagaWorkflow, func(input []byte) (*resources.Withdrawal, error) {
		// TODO: how to query exact ctx.userId?
		// TODO: how to set current workflow_info_status?
		return resourcesWithdrawal, nil
	})
	if err != nil {
		resourcesWithdrawal.Status = enums.WithdrawalStatus_WITHDRAWAL_STATUS_FAILED
		logger.Info("workflow.SetQueryHandler occurred issue", "QuerySagaWorkflow", QuerySagaWorkflow, "Error", err)
		return nil, err
	}
	var a *Activities
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    3,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute, // Timeout options specify when to automatically timeout Actvitivy functions.
		RetryPolicy:         retrypolicy, // Optionally provide a customized RetryPolicy.
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	// data from the withdraw activity
	var withdrawData *string
	err = workflow.ExecuteActivity(ctx, a.Withdraw, request).Get(ctx, &withdrawData)
	// case nil ptr for return
	// err = workflow.ExecuteActivity(ctx, a.Withdraw, request).Get(ctx, nil)
	if err != nil {
		resourcesWithdrawal.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
		var apperr *temporal.ApplicationError
		if errors.As(err, &apperr) {
			switch apperr.Type() {
			case TwirpErrorCode:
				// decode twirp.ErrorCode
				var twerrCode twirp.ErrorCode
				if err := apperr.Details(&twerrCode); err != nil {
					return nil, err
				}
				switch twerrCode {
				case twirp.Internal:
					// TODO: do something when twerrCode is internal server error
				default:
					// unkwnerr of the twerrCode
				}
			default:
				// unkwnerr
			}
		}
		var cancelederr *temporal.CanceledError
		if errors.As(err, &cancelederr) {
			resourcesWithdrawal.Status = enums.WithdrawalStatus_WITHDRAWAL_STATUS_CANCELED
			return resourcesWithdrawal, nil
		}
		var timeoutErr *temporal.TimeoutError
		if errors.As(err, &timeoutErr) {
			switch timeoutErr.TimeoutType() {
			case temporalenumspb.TIMEOUT_TYPE_SCHEDULE_TO_START:
				// Handle TIMEOUT_TYPE_SCHEDULE_TO_START timeout.
			case temporalenumspb.TIMEOUT_TYPE_SCHEDULE_TO_CLOSE:
				// Handle TIMEOUT_TYPE_SCHEDULE_TO_CLOSE timeout.
			case temporalenumspb.TIMEOUT_TYPE_START_TO_CLOSE:
				// Handle TIMEOUT_TYPE_START_TO_CLOSE timeout.
			case temporalenumspb.TIMEOUT_TYPE_HEARTBEAT:
				// Handle TIMEOUT_TYPE_HEARTBEAT timeout.
			default:
			}
		}
		var terminatederr *temporal.TerminatedError
		if errors.As(err, &terminatederr) {
			// TODO: terminated error
		}
		var panicerr *temporal.PanicError
		if errors.As(err, &panicerr) {
			// TODO: panicerr panicErr.Error(), panicErr.StackTrace()
		}
		resourcesWithdrawal.Status = enums.WithdrawalStatus_WITHDRAWAL_STATUS_FAILED
		logger.Error("ExecuteActivity.Withdraw occurred issue", "Error", err)
		return nil, err
	}
	// set the withdraw data here.
	resourcesWithdrawal.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	resourcesWithdrawal.Status = enums.WithdrawalStatus_WITHDRAWAL_STATUS_COMPLETED
	return resourcesWithdrawal, nil
}
