// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-hcs.cluster.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "auditLoggingEnabled", GoGetter: "AuditLoggingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "auditLoggingEnabledInput", GoGetter: "AuditLoggingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogStorageContainerUrl", GoGetter: "AuditLogStorageContainerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogStorageContainerUrlInput", GoGetter: "AuditLogStorageContainerUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "blobContainerName", GoGetter: "BlobContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterMode", GoGetter: "ClusterMode"},
			_jsii_.MemberProperty{JsiiProperty: "clusterModeInput", GoGetter: "ClusterModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "consulAutomaticUpgrades", GoGetter: "ConsulAutomaticUpgrades"},
			_jsii_.MemberProperty{JsiiProperty: "consulCaFile", GoGetter: "ConsulCaFile"},
			_jsii_.MemberProperty{JsiiProperty: "consulClusterId", GoGetter: "ConsulClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "consulConfigFile", GoGetter: "ConsulConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "consulConnect", GoGetter: "ConsulConnect"},
			_jsii_.MemberProperty{JsiiProperty: "consulDatacenter", GoGetter: "ConsulDatacenter"},
			_jsii_.MemberProperty{JsiiProperty: "consulDatacenterInput", GoGetter: "ConsulDatacenterInput"},
			_jsii_.MemberProperty{JsiiProperty: "consulExternalEndpoint", GoGetter: "ConsulExternalEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "consulExternalEndpointInput", GoGetter: "ConsulExternalEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "consulExternalEndpointUrl", GoGetter: "ConsulExternalEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "consulFederationToken", GoGetter: "ConsulFederationToken"},
			_jsii_.MemberProperty{JsiiProperty: "consulFederationTokenInput", GoGetter: "ConsulFederationTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "consulPrivateEndpointUrl", GoGetter: "ConsulPrivateEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "consulRootTokenAccessorId", GoGetter: "ConsulRootTokenAccessorId"},
			_jsii_.MemberProperty{JsiiProperty: "consulRootTokenSecretId", GoGetter: "ConsulRootTokenSecretId"},
			_jsii_.MemberProperty{JsiiProperty: "consulSnapshotInterval", GoGetter: "ConsulSnapshotInterval"},
			_jsii_.MemberProperty{JsiiProperty: "consulSnapshotRetention", GoGetter: "ConsulSnapshotRetention"},
			_jsii_.MemberProperty{JsiiProperty: "consulVersion", GoGetter: "ConsulVersion"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedApplicationId", GoGetter: "ManagedApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "managedApplicationName", GoGetter: "ManagedApplicationName"},
			_jsii_.MemberProperty{JsiiProperty: "managedApplicationNameInput", GoGetter: "ManagedApplicationNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedIdentityName", GoGetter: "ManagedIdentityName"},
			_jsii_.MemberProperty{JsiiProperty: "managedResourceGroupName", GoGetter: "ManagedResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "managedResourceGroupNameInput", GoGetter: "ManagedResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "minConsulVersion", GoGetter: "MinConsulVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minConsulVersionInput", GoGetter: "MinConsulVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "planName", GoGetter: "PlanName"},
			_jsii_.MemberProperty{JsiiProperty: "planNameInput", GoGetter: "PlanNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLoggingEnabled", GoMethod: "ResetAuditLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogStorageContainerUrl", GoMethod: "ResetAuditLogStorageContainerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterName", GoMethod: "ResetClusterName"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsulDatacenter", GoMethod: "ResetConsulDatacenter"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsulExternalEndpoint", GoMethod: "ResetConsulExternalEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsulFederationToken", GoMethod: "ResetConsulFederationToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedResourceGroupName", GoMethod: "ResetManagedResourceGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinConsulVersion", GoMethod: "ResetMinConsulVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlanName", GoMethod: "ResetPlanName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetVnetCidr", GoMethod: "ResetVnetCidr"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountName", GoGetter: "StorageAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountResourceGroup", GoGetter: "StorageAccountResourceGroup"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vnetCidr", GoGetter: "VnetCidr"},
			_jsii_.MemberProperty{JsiiProperty: "vnetCidrInput", GoGetter: "VnetCidrInput"},
			_jsii_.MemberProperty{JsiiProperty: "vnetId", GoGetter: "VnetId"},
			_jsii_.MemberProperty{JsiiProperty: "vnetName", GoGetter: "VnetName"},
			_jsii_.MemberProperty{JsiiProperty: "vnetResourceGroupName", GoGetter: "VnetResourceGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcs.cluster.ClusterConfig",
		reflect.TypeOf((*ClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcs.cluster.ClusterTimeouts",
		reflect.TypeOf((*ClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-hcs.cluster.ClusterTimeoutsOutputReference",
		reflect.TypeOf((*ClusterTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInput", GoGetter: "DefaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefault", GoMethod: "ResetDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
