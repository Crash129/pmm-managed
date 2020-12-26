// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package dbaas contains all APIs related to DBaaS.
package dbaas

import (
	"context"

	controllerv1beta1 "github.com/percona-platform/dbaas-api/gen/controller"
	"google.golang.org/grpc"
)

//go:generate mockery -name=dbaasClient -case=snake -inpkg -testonly

type dbaasClient interface {
	// CheckKubernetesClusterConnection checks connection to Kubernetes cluster and returns statuses of the cluster and operators.
	CheckKubernetesClusterConnection(ctx context.Context, kubeConfig string) (*controllerv1beta1.CheckKubernetesClusterConnectionResponse, error)
	// ListXtraDBClusters returns a list of XtraDB clusters.
	ListXtraDBClusters(ctx context.Context, in *controllerv1beta1.ListXtraDBClustersRequest, opts ...grpc.CallOption) (*controllerv1beta1.ListXtraDBClustersResponse, error)
	// CreateXtraDBCluster creates a new XtraDB cluster.
	CreateXtraDBCluster(ctx context.Context, in *controllerv1beta1.CreateXtraDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.CreateXtraDBClusterResponse, error)
	// UpdateXtraDBCluster updates existing XtraDB cluster.
	UpdateXtraDBCluster(ctx context.Context, in *controllerv1beta1.UpdateXtraDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.UpdateXtraDBClusterResponse, error)
	// DeleteXtraDBCluster deletes XtraDB cluster.
	DeleteXtraDBCluster(ctx context.Context, in *controllerv1beta1.DeleteXtraDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.DeleteXtraDBClusterResponse, error)
	// RestartXtraDBCluster restarts XtraDB cluster.
	RestartXtraDBCluster(ctx context.Context, in *controllerv1beta1.RestartXtraDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.RestartXtraDBClusterResponse, error)
	// GetXtraDBCluster returns an XtraDB cluster credentials.
	GetXtraDBCluster(ctx context.Context, in *controllerv1beta1.GetXtraDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetXtraDBClusterResponse, error)
	// ListPSMDBClusters returns a list of PSMDB clusters.
	ListPSMDBClusters(ctx context.Context, in *controllerv1beta1.ListPSMDBClustersRequest, opts ...grpc.CallOption) (*controllerv1beta1.ListPSMDBClustersResponse, error)
	// CreatePSMDBCluster creates a new PSMDB cluster.
	CreatePSMDBCluster(ctx context.Context, in *controllerv1beta1.CreatePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.CreatePSMDBClusterResponse, error)
	// UpdatePSMDBCluster updates existing PSMDB cluster.
	UpdatePSMDBCluster(ctx context.Context, in *controllerv1beta1.UpdatePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.UpdatePSMDBClusterResponse, error)
	// DeletePSMDBCluster deletes PSMDB cluster.
	DeletePSMDBCluster(ctx context.Context, in *controllerv1beta1.DeletePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.DeletePSMDBClusterResponse, error)
	// RestartPSMDBCluster restarts PSMDB cluster.
	RestartPSMDBCluster(ctx context.Context, in *controllerv1beta1.RestartPSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.RestartPSMDBClusterResponse, error)
	// GetPSMDBCluster gets a PSMDB cluster.
	GetPSMDBCluster(ctx context.Context, in *controllerv1beta1.GetPSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetPSMDBClusterResponse, error)
}