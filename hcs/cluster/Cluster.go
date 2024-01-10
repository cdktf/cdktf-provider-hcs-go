// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcs-go/hcs/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcs-go/hcs/v8/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs/resources/cluster hcs_cluster}.
type Cluster interface {
	cdktf.TerraformResource
	AuditLoggingEnabled() interface{}
	SetAuditLoggingEnabled(val interface{})
	AuditLoggingEnabledInput() interface{}
	AuditLogStorageContainerUrl() *string
	SetAuditLogStorageContainerUrl(val *string)
	AuditLogStorageContainerUrlInput() *string
	BlobContainerName() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterMode() *string
	SetClusterMode(val *string)
	ClusterModeInput() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsulAutomaticUpgrades() cdktf.IResolvable
	ConsulCaFile() *string
	ConsulClusterId() *string
	ConsulConfigFile() *string
	ConsulConnect() cdktf.IResolvable
	ConsulDatacenter() *string
	SetConsulDatacenter(val *string)
	ConsulDatacenterInput() *string
	ConsulExternalEndpoint() interface{}
	SetConsulExternalEndpoint(val interface{})
	ConsulExternalEndpointInput() interface{}
	ConsulExternalEndpointUrl() *string
	ConsulFederationToken() *string
	SetConsulFederationToken(val *string)
	ConsulFederationTokenInput() *string
	ConsulPrivateEndpointUrl() *string
	ConsulRootTokenAccessorId() *string
	ConsulRootTokenSecretId() *string
	ConsulSnapshotInterval() *string
	ConsulSnapshotRetention() *string
	ConsulVersion() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedApplicationId() *string
	ManagedApplicationName() *string
	SetManagedApplicationName(val *string)
	ManagedApplicationNameInput() *string
	ManagedIdentityName() *string
	ManagedResourceGroupName() *string
	SetManagedResourceGroupName(val *string)
	ManagedResourceGroupNameInput() *string
	MinConsulVersion() *string
	SetMinConsulVersion(val *string)
	MinConsulVersionInput() *string
	// The tree node.
	Node() constructs.Node
	PlanName() *string
	SetPlanName(val *string)
	PlanNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	State() *string
	StorageAccountName() *string
	StorageAccountResourceGroup() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VnetCidr() *string
	SetVnetCidr(val *string)
	VnetCidrInput() *string
	VnetId() *string
	VnetName() *string
	VnetResourceGroupName() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ClusterTimeouts)
	ResetAuditLoggingEnabled()
	ResetAuditLogStorageContainerUrl()
	ResetClusterName()
	ResetConsulDatacenter()
	ResetConsulExternalEndpoint()
	ResetConsulFederationToken()
	ResetId()
	ResetLocation()
	ResetManagedResourceGroupName()
	ResetMinConsulVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlanName()
	ResetTags()
	ResetTimeouts()
	ResetVnetCidr()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Cluster) AuditLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AuditLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AuditLogStorageContainerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogStorageContainerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AuditLogStorageContainerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogStorageContainerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BlobContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulAutomaticUpgrades() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"consulAutomaticUpgrades",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulConfigFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulConfigFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulConnect() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"consulConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulDatacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulDatacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulDatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulDatacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulExternalEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consulExternalEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulExternalEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consulExternalEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulExternalEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulExternalEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulFederationToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulFederationToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulFederationTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulFederationTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulPrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulPrivateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulRootTokenAccessorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulRootTokenAccessorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulRootTokenSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulRootTokenSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulSnapshotInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulSnapshotInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulSnapshotRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulSnapshotRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConsulVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ManagedResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MinConsulVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minConsulVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MinConsulVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minConsulVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PlanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) StorageAccountResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Timeouts() ClusterTimeoutsOutputReference {
	var returns ClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VnetResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetResourceGroupName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs/resources/cluster hcs_cluster} Resource.
func NewCluster(scope constructs.Construct, id *string, config *ClusterConfig) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"@cdktf/provider-hcs.cluster.Cluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs/resources/cluster hcs_cluster} Resource.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, config *ClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcs.cluster.Cluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Cluster)SetAuditLoggingEnabled(val interface{}) {
	if err := j.validateSetAuditLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetAuditLogStorageContainerUrl(val *string) {
	if err := j.validateSetAuditLogStorageContainerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogStorageContainerUrl",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetClusterMode(val *string) {
	if err := j.validateSetClusterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMode",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConsulDatacenter(val *string) {
	if err := j.validateSetConsulDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consulDatacenter",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConsulExternalEndpoint(val interface{}) {
	if err := j.validateSetConsulExternalEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consulExternalEndpoint",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConsulFederationToken(val *string) {
	if err := j.validateSetConsulFederationTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consulFederationToken",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetManagedApplicationName(val *string) {
	if err := j.validateSetManagedApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedApplicationName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetManagedResourceGroupName(val *string) {
	if err := j.validateSetManagedResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetMinConsulVersion(val *string) {
	if err := j.validateSetMinConsulVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minConsulVersion",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetPlanName(val *string) {
	if err := j.validateSetPlanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetVnetCidr(val *string) {
	if err := j.validateSetVnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetCidr",
		val,
	)
}

// Generates CDKTF code for importing a Cluster resource upon running "cdktf plan <stack-name>".
func Cluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.cluster.Cluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.cluster.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.cluster.Cluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.cluster.Cluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Cluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcs.cluster.Cluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Cluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Cluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Cluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Cluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Cluster) PutTimeouts(value *ClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) ResetAuditLoggingEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAuditLoggingEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAuditLogStorageContainerUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetAuditLogStorageContainerUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClusterName() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetConsulDatacenter() {
	_jsii_.InvokeVoid(
		c,
		"resetConsulDatacenter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetConsulExternalEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetConsulExternalEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetConsulFederationToken() {
	_jsii_.InvokeVoid(
		c,
		"resetConsulFederationToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetManagedResourceGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedResourceGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetMinConsulVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinConsulVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetPlanName() {
	_jsii_.InvokeVoid(
		c,
		"resetPlanName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetVnetCidr() {
	_jsii_.InvokeVoid(
		c,
		"resetVnetCidr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

