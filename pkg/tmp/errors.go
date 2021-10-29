package tmp

import (
	"log"

	"github.com/twitchtv/twirp"
	"go.temporal.io/api/serviceerror"
	"google.golang.org/grpc/codes"
)

func Twerr(svcerr serviceerror.ServiceError) twirp.Error {
	
	log.Println(svcerr.Status().Code())
	switch svcerr.Status().Code() {
	case codes.OK:
		return twirp.NewError(twirp.NoError, svcerr.Status().Message())
	case codes.Canceled:
		return twirp.NewError(twirp.Canceled, svcerr.Status().Message())
	case codes.Unknown:
		return twirp.NewError(twirp.Unknown, svcerr.Status().Message())
	case codes.InvalidArgument:
		return twirp.NewError(twirp.InvalidArgument, svcerr.Status().Message())
	case codes.DeadlineExceeded:
		return twirp.NewError(twirp.DeadlineExceeded, svcerr.Status().Message())
	case codes.NotFound:
		return twirp.NewError(twirp.NotFound, svcerr.Status().Message())
	case codes.AlreadyExists:
		return twirp.NewError(twirp.AlreadyExists, svcerr.Status().Message())
	case codes.PermissionDenied:
		return twirp.NewError(twirp.PermissionDenied, svcerr.Status().Message())
	case codes.ResourceExhausted:
		return twirp.NewError(twirp.ResourceExhausted, svcerr.Status().Message())
	case codes.FailedPrecondition:
		return twirp.NewError(twirp.FailedPrecondition, svcerr.Status().Message())
	case codes.Aborted:
		return twirp.NewError(twirp.Aborted, svcerr.Status().Message())
	case codes.Unimplemented:
		return twirp.NewError(twirp.Unimplemented, svcerr.Status().Message())
	case codes.Internal:
		return twirp.NewError(twirp.Internal, svcerr.Status().Message())
	case codes.Unavailable:
		return twirp.NewError(twirp.Unavailable, svcerr.Status().Message())
	case codes.DataLoss:
		return twirp.NewError(twirp.DataLoss, svcerr.Status().Message())
	case codes.Unauthenticated:
		return twirp.NewError(twirp.Unauthenticated, svcerr.Status().Message())
	default:
		return twirp.NewError(twirp.Unknown, svcerr.Status().Message())
	}
}
