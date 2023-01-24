// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package redshiftdataapiserviceiface provides an interface to enable mocking the Redshift Data API Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package redshiftdataapiserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
)

// RedshiftDataAPIServiceAPI provides an interface to enable mocking the
// redshiftdataapiservice.RedshiftDataAPIService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Redshift Data API Service.
//	func myFunc(svc redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI) bool {
//	    // Make svc.BatchExecuteStatement request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := redshiftdataapiservice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockRedshiftDataAPIServiceClient struct {
//	    redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI
//	}
//	func (m *mockRedshiftDataAPIServiceClient) BatchExecuteStatement(input *redshiftdataapiservice.BatchExecuteStatementInput) (*redshiftdataapiservice.BatchExecuteStatementOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockRedshiftDataAPIServiceClient{}
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
type RedshiftDataAPIServiceAPI interface {
	BatchExecuteStatement(*redshiftdataapiservice.BatchExecuteStatementInput) (*redshiftdataapiservice.BatchExecuteStatementOutput, error)
	BatchExecuteStatementWithContext(aws.Context, *redshiftdataapiservice.BatchExecuteStatementInput, ...request.Option) (*redshiftdataapiservice.BatchExecuteStatementOutput, error)
	BatchExecuteStatementRequest(*redshiftdataapiservice.BatchExecuteStatementInput) (*request.Request, *redshiftdataapiservice.BatchExecuteStatementOutput)

	CancelStatement(*redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error)
	CancelStatementWithContext(aws.Context, *redshiftdataapiservice.CancelStatementInput, ...request.Option) (*redshiftdataapiservice.CancelStatementOutput, error)
	CancelStatementRequest(*redshiftdataapiservice.CancelStatementInput) (*request.Request, *redshiftdataapiservice.CancelStatementOutput)

	DescribeStatement(*redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error)
	DescribeStatementWithContext(aws.Context, *redshiftdataapiservice.DescribeStatementInput, ...request.Option) (*redshiftdataapiservice.DescribeStatementOutput, error)
	DescribeStatementRequest(*redshiftdataapiservice.DescribeStatementInput) (*request.Request, *redshiftdataapiservice.DescribeStatementOutput)

	DescribeTable(*redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error)
	DescribeTableWithContext(aws.Context, *redshiftdataapiservice.DescribeTableInput, ...request.Option) (*redshiftdataapiservice.DescribeTableOutput, error)
	DescribeTableRequest(*redshiftdataapiservice.DescribeTableInput) (*request.Request, *redshiftdataapiservice.DescribeTableOutput)

	DescribeTablePages(*redshiftdataapiservice.DescribeTableInput, func(*redshiftdataapiservice.DescribeTableOutput, bool) bool) error
	DescribeTablePagesWithContext(aws.Context, *redshiftdataapiservice.DescribeTableInput, func(*redshiftdataapiservice.DescribeTableOutput, bool) bool, ...request.Option) error

	ExecuteStatement(*redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error)
	ExecuteStatementWithContext(aws.Context, *redshiftdataapiservice.ExecuteStatementInput, ...request.Option) (*redshiftdataapiservice.ExecuteStatementOutput, error)
	ExecuteStatementRequest(*redshiftdataapiservice.ExecuteStatementInput) (*request.Request, *redshiftdataapiservice.ExecuteStatementOutput)

	GetStatementResult(*redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error)
	GetStatementResultWithContext(aws.Context, *redshiftdataapiservice.GetStatementResultInput, ...request.Option) (*redshiftdataapiservice.GetStatementResultOutput, error)
	GetStatementResultRequest(*redshiftdataapiservice.GetStatementResultInput) (*request.Request, *redshiftdataapiservice.GetStatementResultOutput)

	GetStatementResultPages(*redshiftdataapiservice.GetStatementResultInput, func(*redshiftdataapiservice.GetStatementResultOutput, bool) bool) error
	GetStatementResultPagesWithContext(aws.Context, *redshiftdataapiservice.GetStatementResultInput, func(*redshiftdataapiservice.GetStatementResultOutput, bool) bool, ...request.Option) error

	ListDatabases(*redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error)
	ListDatabasesWithContext(aws.Context, *redshiftdataapiservice.ListDatabasesInput, ...request.Option) (*redshiftdataapiservice.ListDatabasesOutput, error)
	ListDatabasesRequest(*redshiftdataapiservice.ListDatabasesInput) (*request.Request, *redshiftdataapiservice.ListDatabasesOutput)

	ListDatabasesPages(*redshiftdataapiservice.ListDatabasesInput, func(*redshiftdataapiservice.ListDatabasesOutput, bool) bool) error
	ListDatabasesPagesWithContext(aws.Context, *redshiftdataapiservice.ListDatabasesInput, func(*redshiftdataapiservice.ListDatabasesOutput, bool) bool, ...request.Option) error

	ListSchemas(*redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error)
	ListSchemasWithContext(aws.Context, *redshiftdataapiservice.ListSchemasInput, ...request.Option) (*redshiftdataapiservice.ListSchemasOutput, error)
	ListSchemasRequest(*redshiftdataapiservice.ListSchemasInput) (*request.Request, *redshiftdataapiservice.ListSchemasOutput)

	ListSchemasPages(*redshiftdataapiservice.ListSchemasInput, func(*redshiftdataapiservice.ListSchemasOutput, bool) bool) error
	ListSchemasPagesWithContext(aws.Context, *redshiftdataapiservice.ListSchemasInput, func(*redshiftdataapiservice.ListSchemasOutput, bool) bool, ...request.Option) error

	ListStatements(*redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error)
	ListStatementsWithContext(aws.Context, *redshiftdataapiservice.ListStatementsInput, ...request.Option) (*redshiftdataapiservice.ListStatementsOutput, error)
	ListStatementsRequest(*redshiftdataapiservice.ListStatementsInput) (*request.Request, *redshiftdataapiservice.ListStatementsOutput)

	ListStatementsPages(*redshiftdataapiservice.ListStatementsInput, func(*redshiftdataapiservice.ListStatementsOutput, bool) bool) error
	ListStatementsPagesWithContext(aws.Context, *redshiftdataapiservice.ListStatementsInput, func(*redshiftdataapiservice.ListStatementsOutput, bool) bool, ...request.Option) error

	ListTables(*redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error)
	ListTablesWithContext(aws.Context, *redshiftdataapiservice.ListTablesInput, ...request.Option) (*redshiftdataapiservice.ListTablesOutput, error)
	ListTablesRequest(*redshiftdataapiservice.ListTablesInput) (*request.Request, *redshiftdataapiservice.ListTablesOutput)

	ListTablesPages(*redshiftdataapiservice.ListTablesInput, func(*redshiftdataapiservice.ListTablesOutput, bool) bool) error
	ListTablesPagesWithContext(aws.Context, *redshiftdataapiservice.ListTablesInput, func(*redshiftdataapiservice.ListTablesOutput, bool) bool, ...request.Option) error
}

var _ RedshiftDataAPIServiceAPI = (*redshiftdataapiservice.RedshiftDataAPIService)(nil)