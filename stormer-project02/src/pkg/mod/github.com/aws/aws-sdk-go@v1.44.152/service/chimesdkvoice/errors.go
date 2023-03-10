// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chimesdkvoice

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeResourceLimitExceededException for service response error code
	// "ResourceLimitExceededException".
	ErrCodeResourceLimitExceededException = "ResourceLimitExceededException"

	// ErrCodeServiceFailureException for service response error code
	// "ServiceFailureException".
	ErrCodeServiceFailureException = "ServiceFailureException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottledClientException for service response error code
	// "ThrottledClientException".
	ErrCodeThrottledClientException = "ThrottledClientException"

	// ErrCodeUnauthorizedClientException for service response error code
	// "UnauthorizedClientException".
	ErrCodeUnauthorizedClientException = "UnauthorizedClientException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":          newErrorAccessDeniedException,
	"BadRequestException":            newErrorBadRequestException,
	"ConflictException":              newErrorConflictException,
	"ForbiddenException":             newErrorForbiddenException,
	"NotFoundException":              newErrorNotFoundException,
	"ResourceLimitExceededException": newErrorResourceLimitExceededException,
	"ServiceFailureException":        newErrorServiceFailureException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
	"ThrottledClientException":       newErrorThrottledClientException,
	"UnauthorizedClientException":    newErrorUnauthorizedClientException,
}
