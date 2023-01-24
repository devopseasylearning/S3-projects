// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package qldbsessioniface provides an interface to enable mocking the Amazon QLDB Session service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package qldbsessioniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/qldbsession"
)

// QLDBSessionAPI provides an interface to enable mocking the
// qldbsession.QLDBSession service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon QLDB Session.
//	func myFunc(svc qldbsessioniface.QLDBSessionAPI) bool {
//	    // Make svc.SendCommand request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := qldbsession.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockQLDBSessionClient struct {
//	    qldbsessioniface.QLDBSessionAPI
//	}
//	func (m *mockQLDBSessionClient) SendCommand(input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockQLDBSessionClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type QLDBSessionAPI interface {
	SendCommand(*qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error)
	SendCommandWithContext(aws.Context, *qldbsession.SendCommandInput, ...request.Option) (*qldbsession.SendCommandOutput, error)
	SendCommandRequest(*qldbsession.SendCommandInput) (*request.Request, *qldbsession.SendCommandOutput)
}

var _ QLDBSessionAPI = (*qldbsession.QLDBSession)(nil)
