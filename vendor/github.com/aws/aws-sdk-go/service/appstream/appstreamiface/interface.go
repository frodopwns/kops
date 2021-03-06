// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package appstreamiface provides an interface to enable mocking the Amazon AppStream service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appstreamiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appstream"
)

// AppStreamAPI provides an interface to enable mocking the
// appstream.AppStream service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon AppStream.
//    func myFunc(svc appstreamiface.AppStreamAPI) bool {
//        // Make svc.AssociateFleet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appstream.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppStreamClient struct {
//        appstreamiface.AppStreamAPI
//    }
//    func (m *mockAppStreamClient) AssociateFleet(input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppStreamClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AppStreamAPI interface {
	AssociateFleetRequest(*appstream.AssociateFleetInput) (*request.Request, *appstream.AssociateFleetOutput)

	AssociateFleet(*appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error)

	CreateFleetRequest(*appstream.CreateFleetInput) (*request.Request, *appstream.CreateFleetOutput)

	CreateFleet(*appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error)

	CreateStackRequest(*appstream.CreateStackInput) (*request.Request, *appstream.CreateStackOutput)

	CreateStack(*appstream.CreateStackInput) (*appstream.CreateStackOutput, error)

	CreateStreamingURLRequest(*appstream.CreateStreamingURLInput) (*request.Request, *appstream.CreateStreamingURLOutput)

	CreateStreamingURL(*appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error)

	DeleteFleetRequest(*appstream.DeleteFleetInput) (*request.Request, *appstream.DeleteFleetOutput)

	DeleteFleet(*appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error)

	DeleteStackRequest(*appstream.DeleteStackInput) (*request.Request, *appstream.DeleteStackOutput)

	DeleteStack(*appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error)

	DescribeFleetsRequest(*appstream.DescribeFleetsInput) (*request.Request, *appstream.DescribeFleetsOutput)

	DescribeFleets(*appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error)

	DescribeImagesRequest(*appstream.DescribeImagesInput) (*request.Request, *appstream.DescribeImagesOutput)

	DescribeImages(*appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error)

	DescribeSessionsRequest(*appstream.DescribeSessionsInput) (*request.Request, *appstream.DescribeSessionsOutput)

	DescribeSessions(*appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error)

	DescribeStacksRequest(*appstream.DescribeStacksInput) (*request.Request, *appstream.DescribeStacksOutput)

	DescribeStacks(*appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error)

	DisassociateFleetRequest(*appstream.DisassociateFleetInput) (*request.Request, *appstream.DisassociateFleetOutput)

	DisassociateFleet(*appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error)

	ExpireSessionRequest(*appstream.ExpireSessionInput) (*request.Request, *appstream.ExpireSessionOutput)

	ExpireSession(*appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error)

	ListAssociatedFleetsRequest(*appstream.ListAssociatedFleetsInput) (*request.Request, *appstream.ListAssociatedFleetsOutput)

	ListAssociatedFleets(*appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error)

	ListAssociatedStacksRequest(*appstream.ListAssociatedStacksInput) (*request.Request, *appstream.ListAssociatedStacksOutput)

	ListAssociatedStacks(*appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error)

	StartFleetRequest(*appstream.StartFleetInput) (*request.Request, *appstream.StartFleetOutput)

	StartFleet(*appstream.StartFleetInput) (*appstream.StartFleetOutput, error)

	StopFleetRequest(*appstream.StopFleetInput) (*request.Request, *appstream.StopFleetOutput)

	StopFleet(*appstream.StopFleetInput) (*appstream.StopFleetOutput, error)

	UpdateFleetRequest(*appstream.UpdateFleetInput) (*request.Request, *appstream.UpdateFleetOutput)

	UpdateFleet(*appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error)

	UpdateStackRequest(*appstream.UpdateStackInput) (*request.Request, *appstream.UpdateStackOutput)

	UpdateStack(*appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error)

	WaitUntilFleetStarted(*appstream.DescribeFleetsInput) error

	WaitUntilFleetStopped(*appstream.DescribeFleetsInput) error
}

var _ AppStreamAPI = (*appstream.AppStream)(nil)
