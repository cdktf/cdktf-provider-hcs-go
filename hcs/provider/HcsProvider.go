// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcs-go/hcs/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcs-go/hcs/v6/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs hcs}.
type HcsProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AzureClientCertificatePassword() *string
	SetAzureClientCertificatePassword(val *string)
	AzureClientCertificatePasswordInput() *string
	AzureClientCertificatePath() *string
	SetAzureClientCertificatePath(val *string)
	AzureClientCertificatePathInput() *string
	AzureClientId() *string
	SetAzureClientId(val *string)
	AzureClientIdInput() *string
	AzureClientSecret() *string
	SetAzureClientSecret(val *string)
	AzureClientSecretInput() *string
	AzureEnvironment() *string
	SetAzureEnvironment(val *string)
	AzureEnvironmentInput() *string
	AzureMetadataHost() *string
	SetAzureMetadataHost(val *string)
	AzureMetadataHostInput() *string
	AzureMsiEndpoint() *string
	SetAzureMsiEndpoint(val *string)
	AzureMsiEndpointInput() *string
	AzureSubscriptionId() *string
	SetAzureSubscriptionId(val *string)
	AzureSubscriptionIdInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
	AzureUseMsi() interface{}
	SetAzureUseMsi(val interface{})
	AzureUseMsiInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HcpApiDomain() *string
	SetHcpApiDomain(val *string)
	HcpApiDomainInput() *string
	HcsMarketplaceProductName() *string
	SetHcsMarketplaceProductName(val *string)
	HcsMarketplaceProductNameInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAzureClientCertificatePassword()
	ResetAzureClientCertificatePath()
	ResetAzureClientId()
	ResetAzureClientSecret()
	ResetAzureEnvironment()
	ResetAzureMetadataHost()
	ResetAzureMsiEndpoint()
	ResetAzureSubscriptionId()
	ResetAzureTenantId()
	ResetAzureUseMsi()
	ResetHcpApiDomain()
	ResetHcsMarketplaceProductName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HcsProvider
type jsiiProxy_HcsProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_HcsProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientCertificatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientCertificatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientCertificatePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientCertificatePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureMetadataHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMetadataHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureMetadataHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMetadataHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureMsiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMsiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureMsiEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMsiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureSubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSubscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureUseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureUseMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) AzureUseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureUseMsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) HcpApiDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcpApiDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) HcpApiDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcpApiDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) HcsMarketplaceProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcsMarketplaceProductName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) HcsMarketplaceProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcsMarketplaceProductNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcsProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs hcs} Resource.
func NewHcsProvider(scope constructs.Construct, id *string, config *HcsProviderConfig) HcsProvider {
	_init_.Initialize()

	if err := validateNewHcsProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HcsProvider{}

	_jsii_.Create(
		"@cdktf/provider-hcs.provider.HcsProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs hcs} Resource.
func NewHcsProvider_Override(h HcsProvider, scope constructs.Construct, id *string, config *HcsProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcs.provider.HcsProvider",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HcsProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"azureClientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"azureClientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureClientId(val *string) {
	_jsii_.Set(
		j,
		"azureClientId",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureClientSecret(val *string) {
	_jsii_.Set(
		j,
		"azureClientSecret",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureEnvironment(val *string) {
	_jsii_.Set(
		j,
		"azureEnvironment",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureMetadataHost(val *string) {
	_jsii_.Set(
		j,
		"azureMetadataHost",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureMsiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"azureMsiEndpoint",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"azureSubscriptionId",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureTenantId(val *string) {
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetAzureUseMsi(val interface{}) {
	if err := j.validateSetAzureUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureUseMsi",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetHcpApiDomain(val *string) {
	_jsii_.Set(
		j,
		"hcpApiDomain",
		val,
	)
}

func (j *jsiiProxy_HcsProvider)SetHcsMarketplaceProductName(val *string) {
	_jsii_.Set(
		j,
		"hcsMarketplaceProductName",
		val,
	)
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
func HcsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.provider.HcsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HcsProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcsProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.provider.HcsProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HcsProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcsProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcs.provider.HcsProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HcsProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcs.provider.HcsProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HcsProvider) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HcsProvider) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HcsProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		h,
		"resetAlias",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureClientCertificatePassword() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureClientCertificatePassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureClientCertificatePath() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureClientCertificatePath",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureClientId() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureClientId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureClientSecret() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureClientSecret",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureEnvironment() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureEnvironment",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureMetadataHost() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureMetadataHost",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureMsiEndpoint() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureMsiEndpoint",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureSubscriptionId() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureSubscriptionId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetAzureUseMsi() {
	_jsii_.InvokeVoid(
		h,
		"resetAzureUseMsi",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetHcpApiDomain() {
	_jsii_.InvokeVoid(
		h,
		"resetHcpApiDomain",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetHcsMarketplaceProductName() {
	_jsii_.InvokeVoid(
		h,
		"resetHcsMarketplaceProductName",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcsProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcsProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcsProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

