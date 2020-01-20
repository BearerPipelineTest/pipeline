// Code generated by mga tool. DO NOT EDIT.
package clusterdriver

import (
	"github.com/banzaicloud/pipeline/internal/cluster"
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// Endpoint name constants
const (
	CreateNodePoolNodePoolEndpoint = "cluster.CreateNodePool"
	DeleteNodePoolNodePoolEndpoint = "cluster.DeleteNodePool"
)

// NodePoolEndpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type NodePoolEndpoints struct {
	CreateNodePool endpoint.Endpoint
	DeleteNodePool endpoint.Endpoint
}

// MakeNodePoolEndpoints returns a(n) NodePoolEndpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeNodePoolEndpoints(service cluster.NodePoolService, middleware ...endpoint.Middleware) NodePoolEndpoints {
	mw := kitxendpoint.Combine(middleware...)

	return NodePoolEndpoints{
		CreateNodePool: kitxendpoint.OperationNameMiddleware(CreateNodePoolNodePoolEndpoint)(mw(MakeCreateNodePoolEndpoint(service))),
		DeleteNodePool: kitxendpoint.OperationNameMiddleware(DeleteNodePoolNodePoolEndpoint)(mw(MakeDeleteNodePoolEndpoint(service))),
	}
}

// TraceNodePoolEndpoints returns a(n) NodePoolEndpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceNodePoolEndpoints(endpoints NodePoolEndpoints) NodePoolEndpoints {
	return NodePoolEndpoints{
		CreateNodePool: kitoc.TraceEndpoint("cluster.CreateNodePool")(endpoints.CreateNodePool),
		DeleteNodePool: kitoc.TraceEndpoint("cluster.DeleteNodePool")(endpoints.DeleteNodePool),
	}
}
