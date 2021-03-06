package sqlapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2018-06-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

// DatabaseSecurityAlertPoliciesClientAPI contains the set of methods on the DatabaseSecurityAlertPoliciesClient type.
type DatabaseSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseSecurityAlertPolicy) (result sql.DatabaseSecurityAlertPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseSecurityAlertPolicy, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseSecurityAlertListResultPage, err error)
}

var _ DatabaseSecurityAlertPoliciesClientAPI = (*sql.DatabaseSecurityAlertPoliciesClient)(nil)

// ManagedDatabaseSensitivityLabelsClientAPI contains the set of methods on the ManagedDatabaseSensitivityLabelsClient type.
type ManagedDatabaseSensitivityLabelsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters sql.SensitivityLabel) (result sql.SensitivityLabel, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	DisableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	EnableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource sql.SensitivityLabelSource) (result sql.SensitivityLabel, err error)
	ListCurrentByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string) (result sql.SensitivityLabelListResultPage, err error)
	ListRecommendedByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result sql.SensitivityLabelListResultPage, err error)
}

var _ ManagedDatabaseSensitivityLabelsClientAPI = (*sql.ManagedDatabaseSensitivityLabelsClient)(nil)

// ManagedInstanceVulnerabilityAssessmentsClientAPI contains the set of methods on the ManagedInstanceVulnerabilityAssessmentsClient type.
type ManagedInstanceVulnerabilityAssessmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstanceVulnerabilityAssessment) (result sql.ManagedInstanceVulnerabilityAssessment, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceVulnerabilityAssessment, err error)
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceVulnerabilityAssessmentListResultPage, err error)
}

var _ ManagedInstanceVulnerabilityAssessmentsClientAPI = (*sql.ManagedInstanceVulnerabilityAssessmentsClient)(nil)

// ManagedInstanceOperationsClientAPI contains the set of methods on the ManagedInstanceOperationsClient type.
type ManagedInstanceOperationsClientAPI interface {
	ListByManagedInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstanceOperationListResultPage, err error)
}

var _ ManagedInstanceOperationsClientAPI = (*sql.ManagedInstanceOperationsClient)(nil)

// ServerVulnerabilityAssessmentsClientAPI contains the set of methods on the ServerVulnerabilityAssessmentsClient type.
type ServerVulnerabilityAssessmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerVulnerabilityAssessment) (result sql.ServerVulnerabilityAssessment, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerVulnerabilityAssessment, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerVulnerabilityAssessmentListResultPage, err error)
}

var _ ServerVulnerabilityAssessmentsClientAPI = (*sql.ServerVulnerabilityAssessmentsClient)(nil)

// InstancePoolsClientAPI contains the set of methods on the InstancePoolsClient type.
type InstancePoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, instancePoolName string, parameters sql.InstancePool) (result sql.InstancePoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, instancePoolName string) (result sql.InstancePoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, instancePoolName string) (result sql.InstancePool, err error)
	List(ctx context.Context) (result sql.InstancePoolListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.InstancePoolListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, instancePoolName string, parameters sql.InstancePoolUpdate) (result sql.InstancePoolsUpdateFuture, err error)
}

var _ InstancePoolsClientAPI = (*sql.InstancePoolsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByInstancePool(ctx context.Context, resourceGroupName string, instancePoolName string, expandChildren *bool) (result sql.UsageListResultPage, err error)
}

var _ UsagesClientAPI = (*sql.UsagesClient)(nil)

// ManagedInstancesClientAPI contains the set of methods on the ManagedInstancesClient type.
type ManagedInstancesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstance) (result sql.ManagedInstancesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstancesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstance, err error)
	List(ctx context.Context) (result sql.ManagedInstanceListResultPage, err error)
	ListByInstancePool(ctx context.Context, resourceGroupName string, instancePoolName string) (result sql.ManagedInstanceListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.ManagedInstanceListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstanceUpdate) (result sql.ManagedInstancesUpdateFuture, err error)
}

var _ ManagedInstancesClientAPI = (*sql.ManagedInstancesClient)(nil)

// ManagedDatabaseRestoreDetailsClientAPI contains the set of methods on the ManagedDatabaseRestoreDetailsClient type.
type ManagedDatabaseRestoreDetailsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.ManagedDatabaseRestoreDetailsResult, err error)
}

var _ ManagedDatabaseRestoreDetailsClientAPI = (*sql.ManagedDatabaseRestoreDetailsClient)(nil)

// ManagedDatabasesClientAPI contains the set of methods on the ManagedDatabasesClient type.
type ManagedDatabasesClientAPI interface {
	CompleteRestore(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters sql.CompleteDatabaseRestoreDefinition) (result sql.ManagedDatabasesCompleteRestoreFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters sql.ManagedDatabase) (result sql.ManagedDatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.ManagedDatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result sql.ManagedDatabase, err error)
	ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedDatabaseListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters sql.ManagedDatabaseUpdate) (result sql.ManagedDatabasesUpdateFuture, err error)
}

var _ ManagedDatabasesClientAPI = (*sql.ManagedDatabasesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	Failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, replicaType sql.ReplicaType) (result sql.DatabasesFailoverFuture, err error)
}

var _ DatabasesClientAPI = (*sql.DatabasesClient)(nil)

// ElasticPoolsClientAPI contains the set of methods on the ElasticPoolsClient type.
type ElasticPoolsClientAPI interface {
	Failover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolsFailoverFuture, err error)
}

var _ ElasticPoolsClientAPI = (*sql.ElasticPoolsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string, parameters sql.PrivateEndpointConnection) (result sql.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string) (result sql.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string) (result sql.PrivateEndpointConnection, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.PrivateEndpointConnectionListResultPage, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*sql.PrivateEndpointConnectionsClient)(nil)

// ServerAzureADAdministratorsClientAPI contains the set of methods on the ServerAzureADAdministratorsClient type.
type ServerAzureADAdministratorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerAzureADAdministrator) (result sql.ServerAzureADAdministratorsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerAzureADAdministratorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerAzureADAdministrator, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdministratorListResultPage, err error)
}

var _ ServerAzureADAdministratorsClientAPI = (*sql.ServerAzureADAdministratorsClient)(nil)
