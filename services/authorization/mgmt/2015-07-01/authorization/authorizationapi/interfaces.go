package authorizationapi

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
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
)

// PermissionsClientAPI contains the set of methods on the PermissionsClient type.
type PermissionsClientAPI interface {
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result authorization.PermissionGetResultPage, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string) (result authorization.PermissionGetResultPage, err error)
}

var _ PermissionsClientAPI = (*authorization.PermissionsClient)(nil)

// ProviderOperationsMetadataClientAPI contains the set of methods on the ProviderOperationsMetadataClient type.
type ProviderOperationsMetadataClientAPI interface {
	Get(ctx context.Context, resourceProviderNamespace string, expand string) (result authorization.ProviderOperationsMetadata, err error)
	List(ctx context.Context, expand string) (result authorization.ProviderOperationsMetadataListResultPage, err error)
}

var _ ProviderOperationsMetadataClientAPI = (*authorization.ProviderOperationsMetadataClient)(nil)

// RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
type RoleAssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (result authorization.RoleAssignment, err error)
	CreateByID(ctx context.Context, roleAssignmentID string, parameters authorization.RoleAssignmentCreateParameters) (result authorization.RoleAssignment, err error)
	Delete(ctx context.Context, scope string, roleAssignmentName string) (result authorization.RoleAssignment, err error)
	DeleteByID(ctx context.Context, roleAssignmentID string) (result authorization.RoleAssignment, err error)
	Get(ctx context.Context, scope string, roleAssignmentName string) (result authorization.RoleAssignment, err error)
	GetByID(ctx context.Context, roleAssignmentID string) (result authorization.RoleAssignment, err error)
	List(ctx context.Context, filter string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForScope(ctx context.Context, scope string, filter string) (result authorization.RoleAssignmentListResultPage, err error)
}

var _ RoleAssignmentsClientAPI = (*authorization.RoleAssignmentsClient)(nil)

// RoleDefinitionsClientAPI contains the set of methods on the RoleDefinitionsClient type.
type RoleDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, roleDefinitionID string, roleDefinition authorization.RoleDefinition) (result authorization.RoleDefinition, err error)
	Delete(ctx context.Context, scope string, roleDefinitionID string) (result authorization.RoleDefinition, err error)
	Get(ctx context.Context, scope string, roleDefinitionID string) (result authorization.RoleDefinition, err error)
	GetByID(ctx context.Context, roleDefinitionID string) (result authorization.RoleDefinition, err error)
	List(ctx context.Context, scope string, filter string) (result authorization.RoleDefinitionListResultPage, err error)
}

var _ RoleDefinitionsClientAPI = (*authorization.RoleDefinitionsClient)(nil)

// ClassicAdministratorsClientAPI contains the set of methods on the ClassicAdministratorsClient type.
type ClassicAdministratorsClientAPI interface {
	List(ctx context.Context) (result authorization.ClassicAdministratorListResultPage, err error)
}

var _ ClassicAdministratorsClientAPI = (*authorization.ClassicAdministratorsClient)(nil)