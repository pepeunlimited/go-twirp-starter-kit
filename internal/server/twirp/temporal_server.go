package twirp

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/mennanov/fmutils"
	"github.com/pepeunlimited/go-twirp-starter-kit/internal/saga"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/enums"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/resources"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/services"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/fm"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/tmp"
	"github.com/twitchtv/twirp"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// TemporalServer auto generate:
// => CMD + SHIFT + A
// => Go: Generate Interaface Stubs
// => server TemporalServer services.TemporalService
type TemporalServer struct {
	temporalClient client.Client
}

//w, err := server.temporalClient.DescribeWorkflowExecution(ctx, request.Id, request.WorkflowRunId)
//if err != nil {
//	return nil, twirp.InternalError(err.Error())
//}

// Basic example of the Temporal key components: workflow, activity and worker
func (server TemporalServer) CreateWithdraw(ctx context.Context, request *services.CreateWithdrawRequest) (*resources.Withdrawal, error) {
	workflowID := string(uuid.New().String())
	options := client.StartWorkflowOptions{
		ID:                                       workflowID,
		TaskQueue:                                saga.SagaTaskQueue,
		WorkflowExecutionErrorWhenAlreadyStarted: false,
	}
	workflowRun, err := server.temporalClient.ExecuteWorkflow(ctx, options, saga.SagaWorkflow, request)
	if err != nil {
		var svcerr serviceerror.ServiceError
		if errors.As(err, &svcerr) {
			return nil, tmp.Twerr(svcerr)
		} else {
			return nil, twirp.InternalErrorWith(err)
		}
	}
	// synchronously => waiting the result of the workflow (workflowRun.Get(ctx, ptr))
	// var resourcesWithdrawal *resources.Withdrawal
	// if err := workflowRun.Get(ctx, &resourcesWithdrawal); err != nil {
	//	return nil, twirp.InternalError(err.Error())
	// }
	// return resourcesWithdrawal, nil
	// asynchronously => not waiting the result of the workflow
	createdAt := time.Now().UTC()
	return &resources.Withdrawal{
		Id:        workflowID,
		Status:    enums.WithdrawalStatus_WITHDRAWAL_STATUS_CREATED,
		Amount:    request.Amount,
		UserId:    request.UserId,
		CreatedAt: createdAt.Format(time.RFC3339),
		UpdatedAt: createdAt.Format(time.RFC3339),
		Workflow: &resources.Workflow{
			WorkflowId:    workflowID,
			WorkflowRunId: workflowRun.GetRunID(),
			Status:        enums.WorkflowStatus_WORKFLOW_STATUS_UNSPECIFIED,
			Namespace:     "",
		},
	}, nil
}

func (server TemporalServer) GetWithdrawal(ctx context.Context, request *services.GetWithdrawalRequest) (*resources.Withdrawal, error) {
	if fm.IsNilOrPathEmpty(request.FieldMask) {
		// set all possible fields to paths
		request.FieldMask = &fieldmaskpb.FieldMask{
			Paths: []string{
				"id",
				"status",
				"amount",
				"created_at",
				"updated_at",
				"user_id",
				"workflow",
			},
		}
	}
	// => m proto.Message
	withdrawalWithMaskedFieldMask := &resources.Withdrawal{}
	// normalize and validate the field mask before using it.
	request.FieldMask.Normalize()
	if !request.FieldMask.IsValid(withdrawalWithMaskedFieldMask) {
		return nil, twirp.InvalidArgumentError("field_mask", "invalid_field_mask")
	}
	// IMPORTANT: sort the fieldMask paths
	sort.StringSlice(request.FieldMask.Paths).Sort()
	// check if should execute external database join queries or external service query
	withWorkflow := fm.PathContains(request.FieldMask.Paths, "workflow")
	var resourcesWithdrawal *resources.Withdrawal
	switch u := request.Paramenter.(type) {
	case *services.GetWithdrawalRequest_Query:
		if q, err := server.temporalClient.QueryWorkflow(ctx, u.Query.WorkflowId, u.Query.WorkflowRunId, saga.QuerySagaWorkflow); err != nil {
			var svcerr serviceerror.ServiceError
			if errors.As(err, &svcerr) {
				return nil, tmp.Twerr(svcerr)
			} else {
				return nil, twirp.InternalErrorWith(err)
			}
		} else {
			if err := q.Get(&resourcesWithdrawal); err != nil {
				return nil, twirp.InternalErrorWith(err)
			}
		}
		if withWorkflow {
			if w, err := server.temporalClient.DescribeWorkflowExecution(ctx, u.Query.WorkflowId, u.Query.WorkflowRunId); err != nil {
				var svcerr serviceerror.ServiceError
				if errors.As(err, &svcerr) {
					return nil, tmp.Twerr(svcerr)
				} else {
					return nil, twirp.InternalErrorWith(err)
				}
			} else {
				resourcesWithdrawal.Workflow.WorkflowId = w.WorkflowExecutionInfo.Execution.WorkflowId
				resourcesWithdrawal.Workflow.WorkflowRunId = w.WorkflowExecutionInfo.Execution.RunId
				resourcesWithdrawal.Workflow.Status = enums.WorkflowStatus(w.WorkflowExecutionInfo.Status)
				//resourcesWithdrawal.Workflow.Namespace = // DescribeWorkflowExecution doesn't contain `Namespace` property. `QueryWorkflow` returns it
			}
		}
		// verify permission for the withdrawal
		if u.Query.UserId != resourcesWithdrawal.UserId {
			return nil, twirp.NewError(twirp.PermissionDenied, "You don't have permission to access resources.Withdrawal on this server")
		}
	case *services.GetWithdrawalRequest_Query2:
		// query from the external rpc endpoint
		return nil, twirp.NewError(twirp.Unimplemented, "unimplemented for GetWithdrawalRequest_Query2")
	default:
		return nil, twirp.NewError(twirp.InvalidArgument, "not supported parameter")
	}
	// redact the request according to the provided field mask.
	fmutils.Filter(resourcesWithdrawal, request.FieldMask.Paths)
	// now that the request is vetted we can merge it with the user entity.
	proto.Merge(withdrawalWithMaskedFieldMask, resourcesWithdrawal)
	return withdrawalWithMaskedFieldMask, nil
}

func (server TemporalServer) UpdateWithdrawal(ctx context.Context, request *services.UpdateWithdrawRequest) (*resources.Withdrawal, error) {
	if fm.IsNilOrPathEmpty(request.FieldMask) {
		return nil, twirp.RequiredArgumentError("field_mask")
	}
	updateWithMaskedFieldMask, err := server.GetWithdrawal(ctx, &services.GetWithdrawalRequest{
		Paramenter: &services.GetWithdrawalRequest_Query{
			Query: &services.QueryWithdrawal{
				WorkflowId:    request.Operation.Id,
				WorkflowRunId: request.Operation.Workflow.WorkflowRunId,
				UserId:        request.Operation.UserId,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	// normalize and validate the field mask before using it.
	request.FieldMask.Normalize()
	if !request.FieldMask.IsValid(updateWithMaskedFieldMask) {
		return nil, twirp.InvalidArgumentError("field_mask", "invalid_field_mask")
	}
	// IMPORTANT: sort the fieldMask paths
	sort.StringSlice(request.FieldMask.Paths).Sort()
	// Redact the request according to the provided field mask.
	fmutils.Filter(request.Operation, request.FieldMask.Paths)
	// Now that the request is vetted we can merge it with the bookmark entity.
	proto.Merge(updateWithMaskedFieldMask, request.Operation)
	withWorkflow := fm.PathContains(request.FieldMask.Paths, "workflow")
	if withWorkflow {
		switch updateWithMaskedFieldMask.Workflow.Status {
		case enums.WorkflowStatus_WORKFLOW_STATUS_CANCELED:
			if err := server.temporalClient.CancelWorkflow(ctx, updateWithMaskedFieldMask.Id, updateWithMaskedFieldMask.Workflow.WorkflowRunId); err != nil {
				var svcerr serviceerror.ServiceError
				if errors.As(err, &svcerr) {
					return nil, tmp.Twerr(svcerr)
				} else {
					return nil, twirp.InternalErrorWith(err)
				}
			}
		case enums.WorkflowStatus_WORKFLOW_STATUS_RUNNING:
			if result, err := server.temporalClient.ResetWorkflowExecution(ctx, &workflowservice.ResetWorkflowExecutionRequest{
				WorkflowExecution: &common.WorkflowExecution{
					WorkflowId: updateWithMaskedFieldMask.Id,
					RunId:      updateWithMaskedFieldMask.Workflow.WorkflowRunId,
				},
				Reason:                    "TemporalServer.UpdateWithdrawal re-run for workflow",
				Namespace:                 updateWithMaskedFieldMask.Workflow.Namespace,
				WorkflowTaskFinishEventId: 4,
			}); err != nil {
				var svcerr serviceerror.ServiceError
				if errors.As(err, &svcerr) {
					return nil, tmp.Twerr(svcerr)
				} else {
					return nil, twirp.InternalErrorWith(err)
				}
			} else {
				updateWithMaskedFieldMask.Workflow.WorkflowRunId = result.RunId
			}
		default:
			// ignore rest of status mutations
		}
	}
	return updateWithMaskedFieldMask, nil
}

func NewTemporalServer(temporalClient client.Client) TemporalServer {
	return TemporalServer{
		temporalClient: temporalClient,
	}
}
