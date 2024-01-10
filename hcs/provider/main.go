// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-hcs.provider.HcsProvider",
		reflect.TypeOf((*HcsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientCertificatePassword", GoGetter: "AzureClientCertificatePassword"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientCertificatePasswordInput", GoGetter: "AzureClientCertificatePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientCertificatePath", GoGetter: "AzureClientCertificatePath"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientCertificatePathInput", GoGetter: "AzureClientCertificatePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientId", GoGetter: "AzureClientId"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientIdInput", GoGetter: "AzureClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientSecret", GoGetter: "AzureClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "azureClientSecretInput", GoGetter: "AzureClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureEnvironment", GoGetter: "AzureEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "azureEnvironmentInput", GoGetter: "AzureEnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureMetadataHost", GoGetter: "AzureMetadataHost"},
			_jsii_.MemberProperty{JsiiProperty: "azureMetadataHostInput", GoGetter: "AzureMetadataHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureMsiEndpoint", GoGetter: "AzureMsiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "azureMsiEndpointInput", GoGetter: "AzureMsiEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureSubscriptionId", GoGetter: "AzureSubscriptionId"},
			_jsii_.MemberProperty{JsiiProperty: "azureSubscriptionIdInput", GoGetter: "AzureSubscriptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantId", GoGetter: "AzureTenantId"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantIdInput", GoGetter: "AzureTenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureUseMsi", GoGetter: "AzureUseMsi"},
			_jsii_.MemberProperty{JsiiProperty: "azureUseMsiInput", GoGetter: "AzureUseMsiInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "hcpApiDomain", GoGetter: "HcpApiDomain"},
			_jsii_.MemberProperty{JsiiProperty: "hcpApiDomainInput", GoGetter: "HcpApiDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "hcsMarketplaceProductName", GoGetter: "HcsMarketplaceProductName"},
			_jsii_.MemberProperty{JsiiProperty: "hcsMarketplaceProductNameInput", GoGetter: "HcsMarketplaceProductNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureClientCertificatePassword", GoMethod: "ResetAzureClientCertificatePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureClientCertificatePath", GoMethod: "ResetAzureClientCertificatePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureClientId", GoMethod: "ResetAzureClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureClientSecret", GoMethod: "ResetAzureClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureEnvironment", GoMethod: "ResetAzureEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureMetadataHost", GoMethod: "ResetAzureMetadataHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureMsiEndpoint", GoMethod: "ResetAzureMsiEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureSubscriptionId", GoMethod: "ResetAzureSubscriptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureTenantId", GoMethod: "ResetAzureTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureUseMsi", GoMethod: "ResetAzureUseMsi"},
			_jsii_.MemberMethod{JsiiMethod: "resetHcpApiDomain", GoMethod: "ResetHcpApiDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetHcsMarketplaceProductName", GoMethod: "ResetHcsMarketplaceProductName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_HcsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcs.provider.HcsProviderConfig",
		reflect.TypeOf((*HcsProviderConfig)(nil)).Elem(),
	)
}
