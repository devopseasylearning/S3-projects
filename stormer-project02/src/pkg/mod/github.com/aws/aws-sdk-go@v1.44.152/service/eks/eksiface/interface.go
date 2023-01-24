// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package eksiface provides an interface to enable mocking the Amazon Elastic Kubernetes Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package eksiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/eks"
)

// EKSAPI provides an interface to enable mocking the
// eks.EKS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Elastic Kubernetes Service.
//	func myFunc(svc eksiface.EKSAPI) bool {
//	    // Make svc.AssociateEncryptionConfig request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := eks.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockEKSClient struct {
//	    eksiface.EKSAPI
//	}
//	func (m *mockEKSClient) AssociateEncryptionConfig(input *eks.AssociateEncryptionConfigInput) (*eks.AssociateEncryptionConfigOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockEKSClient{}
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
type EKSAPI interface {
	AssociateEncryptionConfig(*eks.AssociateEncryptionConfigInput) (*eks.AssociateEncryptionConfigOutput, error)
	AssociateEncryptionConfigWithContext(aws.Context, *eks.AssociateEncryptionConfigInput, ...request.Option) (*eks.AssociateEncryptionConfigOutput, error)
	AssociateEncryptionConfigRequest(*eks.AssociateEncryptionConfigInput) (*request.Request, *eks.AssociateEncryptionConfigOutput)

	AssociateIdentityProviderConfig(*eks.AssociateIdentityProviderConfigInput) (*eks.AssociateIdentityProviderConfigOutput, error)
	AssociateIdentityProviderConfigWithContext(aws.Context, *eks.AssociateIdentityProviderConfigInput, ...request.Option) (*eks.AssociateIdentityProviderConfigOutput, error)
	AssociateIdentityProviderConfigRequest(*eks.AssociateIdentityProviderConfigInput) (*request.Request, *eks.AssociateIdentityProviderConfigOutput)

	CreateAddon(*eks.CreateAddonInput) (*eks.CreateAddonOutput, error)
	CreateAddonWithContext(aws.Context, *eks.CreateAddonInput, ...request.Option) (*eks.CreateAddonOutput, error)
	CreateAddonRequest(*eks.CreateAddonInput) (*request.Request, *eks.CreateAddonOutput)

	CreateCluster(*eks.CreateClusterInput) (*eks.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *eks.CreateClusterInput, ...request.Option) (*eks.CreateClusterOutput, error)
	CreateClusterRequest(*eks.CreateClusterInput) (*request.Request, *eks.CreateClusterOutput)

	CreateFargateProfile(*eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error)
	CreateFargateProfileWithContext(aws.Context, *eks.CreateFargateProfileInput, ...request.Option) (*eks.CreateFargateProfileOutput, error)
	CreateFargateProfileRequest(*eks.CreateFargateProfileInput) (*request.Request, *eks.CreateFargateProfileOutput)

	CreateNodegroup(*eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error)
	CreateNodegroupWithContext(aws.Context, *eks.CreateNodegroupInput, ...request.Option) (*eks.CreateNodegroupOutput, error)
	CreateNodegroupRequest(*eks.CreateNodegroupInput) (*request.Request, *eks.CreateNodegroupOutput)

	DeleteAddon(*eks.DeleteAddonInput) (*eks.DeleteAddonOutput, error)
	DeleteAddonWithContext(aws.Context, *eks.DeleteAddonInput, ...request.Option) (*eks.DeleteAddonOutput, error)
	DeleteAddonRequest(*eks.DeleteAddonInput) (*request.Request, *eks.DeleteAddonOutput)

	DeleteCluster(*eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *eks.DeleteClusterInput, ...request.Option) (*eks.DeleteClusterOutput, error)
	DeleteClusterRequest(*eks.DeleteClusterInput) (*request.Request, *eks.DeleteClusterOutput)

	DeleteFargateProfile(*eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error)
	DeleteFargateProfileWithContext(aws.Context, *eks.DeleteFargateProfileInput, ...request.Option) (*eks.DeleteFargateProfileOutput, error)
	DeleteFargateProfileRequest(*eks.DeleteFargateProfileInput) (*request.Request, *eks.DeleteFargateProfileOutput)

	DeleteNodegroup(*eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error)
	DeleteNodegroupWithContext(aws.Context, *eks.DeleteNodegroupInput, ...request.Option) (*eks.DeleteNodegroupOutput, error)
	DeleteNodegroupRequest(*eks.DeleteNodegroupInput) (*request.Request, *eks.DeleteNodegroupOutput)

	DeregisterCluster(*eks.DeregisterClusterInput) (*eks.DeregisterClusterOutput, error)
	DeregisterClusterWithContext(aws.Context, *eks.DeregisterClusterInput, ...request.Option) (*eks.DeregisterClusterOutput, error)
	DeregisterClusterRequest(*eks.DeregisterClusterInput) (*request.Request, *eks.DeregisterClusterOutput)

	DescribeAddon(*eks.DescribeAddonInput) (*eks.DescribeAddonOutput, error)
	DescribeAddonWithContext(aws.Context, *eks.DescribeAddonInput, ...request.Option) (*eks.DescribeAddonOutput, error)
	DescribeAddonRequest(*eks.DescribeAddonInput) (*request.Request, *eks.DescribeAddonOutput)

	DescribeAddonVersions(*eks.DescribeAddonVersionsInput) (*eks.DescribeAddonVersionsOutput, error)
	DescribeAddonVersionsWithContext(aws.Context, *eks.DescribeAddonVersionsInput, ...request.Option) (*eks.DescribeAddonVersionsOutput, error)
	DescribeAddonVersionsRequest(*eks.DescribeAddonVersionsInput) (*request.Request, *eks.DescribeAddonVersionsOutput)

	DescribeAddonVersionsPages(*eks.DescribeAddonVersionsInput, func(*eks.DescribeAddonVersionsOutput, bool) bool) error
	DescribeAddonVersionsPagesWithContext(aws.Context, *eks.DescribeAddonVersionsInput, func(*eks.DescribeAddonVersionsOutput, bool) bool, ...request.Option) error

	DescribeCluster(*eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *eks.DescribeClusterInput, ...request.Option) (*eks.DescribeClusterOutput, error)
	DescribeClusterRequest(*eks.DescribeClusterInput) (*request.Request, *eks.DescribeClusterOutput)

	DescribeFargateProfile(*eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error)
	DescribeFargateProfileWithContext(aws.Context, *eks.DescribeFargateProfileInput, ...request.Option) (*eks.DescribeFargateProfileOutput, error)
	DescribeFargateProfileRequest(*eks.DescribeFargateProfileInput) (*request.Request, *eks.DescribeFargateProfileOutput)

	DescribeIdentityProviderConfig(*eks.DescribeIdentityProviderConfigInput) (*eks.DescribeIdentityProviderConfigOutput, error)
	DescribeIdentityProviderConfigWithContext(aws.Context, *eks.DescribeIdentityProviderConfigInput, ...request.Option) (*eks.DescribeIdentityProviderConfigOutput, error)
	DescribeIdentityProviderConfigRequest(*eks.DescribeIdentityProviderConfigInput) (*request.Request, *eks.DescribeIdentityProviderConfigOutput)

	DescribeNodegroup(*eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error)
	DescribeNodegroupWithContext(aws.Context, *eks.DescribeNodegroupInput, ...request.Option) (*eks.DescribeNodegroupOutput, error)
	DescribeNodegroupRequest(*eks.DescribeNodegroupInput) (*request.Request, *eks.DescribeNodegroupOutput)

	DescribeUpdate(*eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error)
	DescribeUpdateWithContext(aws.Context, *eks.DescribeUpdateInput, ...request.Option) (*eks.DescribeUpdateOutput, error)
	DescribeUpdateRequest(*eks.DescribeUpdateInput) (*request.Request, *eks.DescribeUpdateOutput)

	DisassociateIdentityProviderConfig(*eks.DisassociateIdentityProviderConfigInput) (*eks.DisassociateIdentityProviderConfigOutput, error)
	DisassociateIdentityProviderConfigWithContext(aws.Context, *eks.DisassociateIdentityProviderConfigInput, ...request.Option) (*eks.DisassociateIdentityProviderConfigOutput, error)
	DisassociateIdentityProviderConfigRequest(*eks.DisassociateIdentityProviderConfigInput) (*request.Request, *eks.DisassociateIdentityProviderConfigOutput)

	ListAddons(*eks.ListAddonsInput) (*eks.ListAddonsOutput, error)
	ListAddonsWithContext(aws.Context, *eks.ListAddonsInput, ...request.Option) (*eks.ListAddonsOutput, error)
	ListAddonsRequest(*eks.ListAddonsInput) (*request.Request, *eks.ListAddonsOutput)

	ListAddonsPages(*eks.ListAddonsInput, func(*eks.ListAddonsOutput, bool) bool) error
	ListAddonsPagesWithContext(aws.Context, *eks.ListAddonsInput, func(*eks.ListAddonsOutput, bool) bool, ...request.Option) error

	ListClusters(*eks.ListClustersInput) (*eks.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *eks.ListClustersInput, ...request.Option) (*eks.ListClustersOutput, error)
	ListClustersRequest(*eks.ListClustersInput) (*request.Request, *eks.ListClustersOutput)

	ListClustersPages(*eks.ListClustersInput, func(*eks.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *eks.ListClustersInput, func(*eks.ListClustersOutput, bool) bool, ...request.Option) error

	ListFargateProfiles(*eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error)
	ListFargateProfilesWithContext(aws.Context, *eks.ListFargateProfilesInput, ...request.Option) (*eks.ListFargateProfilesOutput, error)
	ListFargateProfilesRequest(*eks.ListFargateProfilesInput) (*request.Request, *eks.ListFargateProfilesOutput)

	ListFargateProfilesPages(*eks.ListFargateProfilesInput, func(*eks.ListFargateProfilesOutput, bool) bool) error
	ListFargateProfilesPagesWithContext(aws.Context, *eks.ListFargateProfilesInput, func(*eks.ListFargateProfilesOutput, bool) bool, ...request.Option) error

	ListIdentityProviderConfigs(*eks.ListIdentityProviderConfigsInput) (*eks.ListIdentityProviderConfigsOutput, error)
	ListIdentityProviderConfigsWithContext(aws.Context, *eks.ListIdentityProviderConfigsInput, ...request.Option) (*eks.ListIdentityProviderConfigsOutput, error)
	ListIdentityProviderConfigsRequest(*eks.ListIdentityProviderConfigsInput) (*request.Request, *eks.ListIdentityProviderConfigsOutput)

	ListIdentityProviderConfigsPages(*eks.ListIdentityProviderConfigsInput, func(*eks.ListIdentityProviderConfigsOutput, bool) bool) error
	ListIdentityProviderConfigsPagesWithContext(aws.Context, *eks.ListIdentityProviderConfigsInput, func(*eks.ListIdentityProviderConfigsOutput, bool) bool, ...request.Option) error

	ListNodegroups(*eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error)
	ListNodegroupsWithContext(aws.Context, *eks.ListNodegroupsInput, ...request.Option) (*eks.ListNodegroupsOutput, error)
	ListNodegroupsRequest(*eks.ListNodegroupsInput) (*request.Request, *eks.ListNodegroupsOutput)

	ListNodegroupsPages(*eks.ListNodegroupsInput, func(*eks.ListNodegroupsOutput, bool) bool) error
	ListNodegroupsPagesWithContext(aws.Context, *eks.ListNodegroupsInput, func(*eks.ListNodegroupsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *eks.ListTagsForResourceInput, ...request.Option) (*eks.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*eks.ListTagsForResourceInput) (*request.Request, *eks.ListTagsForResourceOutput)

	ListUpdates(*eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error)
	ListUpdatesWithContext(aws.Context, *eks.ListUpdatesInput, ...request.Option) (*eks.ListUpdatesOutput, error)
	ListUpdatesRequest(*eks.ListUpdatesInput) (*request.Request, *eks.ListUpdatesOutput)

	ListUpdatesPages(*eks.ListUpdatesInput, func(*eks.ListUpdatesOutput, bool) bool) error
	ListUpdatesPagesWithContext(aws.Context, *eks.ListUpdatesInput, func(*eks.ListUpdatesOutput, bool) bool, ...request.Option) error

	RegisterCluster(*eks.RegisterClusterInput) (*eks.RegisterClusterOutput, error)
	RegisterClusterWithContext(aws.Context, *eks.RegisterClusterInput, ...request.Option) (*eks.RegisterClusterOutput, error)
	RegisterClusterRequest(*eks.RegisterClusterInput) (*request.Request, *eks.RegisterClusterOutput)

	TagResource(*eks.TagResourceInput) (*eks.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *eks.TagResourceInput, ...request.Option) (*eks.TagResourceOutput, error)
	TagResourceRequest(*eks.TagResourceInput) (*request.Request, *eks.TagResourceOutput)

	UntagResource(*eks.UntagResourceInput) (*eks.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *eks.UntagResourceInput, ...request.Option) (*eks.UntagResourceOutput, error)
	UntagResourceRequest(*eks.UntagResourceInput) (*request.Request, *eks.UntagResourceOutput)

	UpdateAddon(*eks.UpdateAddonInput) (*eks.UpdateAddonOutput, error)
	UpdateAddonWithContext(aws.Context, *eks.UpdateAddonInput, ...request.Option) (*eks.UpdateAddonOutput, error)
	UpdateAddonRequest(*eks.UpdateAddonInput) (*request.Request, *eks.UpdateAddonOutput)

	UpdateClusterConfig(*eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error)
	UpdateClusterConfigWithContext(aws.Context, *eks.UpdateClusterConfigInput, ...request.Option) (*eks.UpdateClusterConfigOutput, error)
	UpdateClusterConfigRequest(*eks.UpdateClusterConfigInput) (*request.Request, *eks.UpdateClusterConfigOutput)

	UpdateClusterVersion(*eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error)
	UpdateClusterVersionWithContext(aws.Context, *eks.UpdateClusterVersionInput, ...request.Option) (*eks.UpdateClusterVersionOutput, error)
	UpdateClusterVersionRequest(*eks.UpdateClusterVersionInput) (*request.Request, *eks.UpdateClusterVersionOutput)

	UpdateNodegroupConfig(*eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error)
	UpdateNodegroupConfigWithContext(aws.Context, *eks.UpdateNodegroupConfigInput, ...request.Option) (*eks.UpdateNodegroupConfigOutput, error)
	UpdateNodegroupConfigRequest(*eks.UpdateNodegroupConfigInput) (*request.Request, *eks.UpdateNodegroupConfigOutput)

	UpdateNodegroupVersion(*eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error)
	UpdateNodegroupVersionWithContext(aws.Context, *eks.UpdateNodegroupVersionInput, ...request.Option) (*eks.UpdateNodegroupVersionOutput, error)
	UpdateNodegroupVersionRequest(*eks.UpdateNodegroupVersionInput) (*request.Request, *eks.UpdateNodegroupVersionOutput)

	WaitUntilAddonActive(*eks.DescribeAddonInput) error
	WaitUntilAddonActiveWithContext(aws.Context, *eks.DescribeAddonInput, ...request.WaiterOption) error

	WaitUntilAddonDeleted(*eks.DescribeAddonInput) error
	WaitUntilAddonDeletedWithContext(aws.Context, *eks.DescribeAddonInput, ...request.WaiterOption) error

	WaitUntilClusterActive(*eks.DescribeClusterInput) error
	WaitUntilClusterActiveWithContext(aws.Context, *eks.DescribeClusterInput, ...request.WaiterOption) error

	WaitUntilClusterDeleted(*eks.DescribeClusterInput) error
	WaitUntilClusterDeletedWithContext(aws.Context, *eks.DescribeClusterInput, ...request.WaiterOption) error

	WaitUntilFargateProfileActive(*eks.DescribeFargateProfileInput) error
	WaitUntilFargateProfileActiveWithContext(aws.Context, *eks.DescribeFargateProfileInput, ...request.WaiterOption) error

	WaitUntilFargateProfileDeleted(*eks.DescribeFargateProfileInput) error
	WaitUntilFargateProfileDeletedWithContext(aws.Context, *eks.DescribeFargateProfileInput, ...request.WaiterOption) error

	WaitUntilNodegroupActive(*eks.DescribeNodegroupInput) error
	WaitUntilNodegroupActiveWithContext(aws.Context, *eks.DescribeNodegroupInput, ...request.WaiterOption) error

	WaitUntilNodegroupDeleted(*eks.DescribeNodegroupInput) error
	WaitUntilNodegroupDeletedWithContext(aws.Context, *eks.DescribeNodegroupInput, ...request.WaiterOption) error
}

var _ EKSAPI = (*eks.EKS)(nil)