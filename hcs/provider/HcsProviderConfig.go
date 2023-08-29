// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HcsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#alias HcsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The password associated with the Azure Client Certificate.
	//
	// For use when authenticating as a Service Principal using a Client Certificate
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_client_certificate_password HcsProvider#azure_client_certificate_password}
	AzureClientCertificatePassword *string `field:"optional" json:"azureClientCertificatePassword" yaml:"azureClientCertificatePassword"`
	// The path to the Azure Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_client_certificate_path HcsProvider#azure_client_certificate_path}
	AzureClientCertificatePath *string `field:"optional" json:"azureClientCertificatePath" yaml:"azureClientCertificatePath"`
	// The Azure Client ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_client_id HcsProvider#azure_client_id}
	AzureClientId *string `field:"optional" json:"azureClientId" yaml:"azureClientId"`
	// The Azure Client Secret which should be used.
	//
	// For use when authenticating as a Service Principal using a Client Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_client_secret HcsProvider#azure_client_secret}
	AzureClientSecret *string `field:"optional" json:"azureClientSecret" yaml:"azureClientSecret"`
	// The Azure Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_environment HcsProvider#azure_environment}
	AzureEnvironment *string `field:"optional" json:"azureEnvironment" yaml:"azureEnvironment"`
	// The hostname which should be used for the Azure Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_metadata_host HcsProvider#azure_metadata_host}
	AzureMetadataHost *string `field:"optional" json:"azureMetadataHost" yaml:"azureMetadataHost"`
	// The path to a custom endpoint for Azure Managed Service Identity - in most circumstances this should be detected automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_msi_endpoint HcsProvider#azure_msi_endpoint}
	AzureMsiEndpoint *string `field:"optional" json:"azureMsiEndpoint" yaml:"azureMsiEndpoint"`
	// The Azure Subscription ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_subscription_id HcsProvider#azure_subscription_id}
	AzureSubscriptionId *string `field:"optional" json:"azureSubscriptionId" yaml:"azureSubscriptionId"`
	// The Azure Tenant ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_tenant_id HcsProvider#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Allowed Azure Managed Service Identity be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#azure_use_msi HcsProvider#azure_use_msi}
	AzureUseMsi interface{} `field:"optional" json:"azureUseMsi" yaml:"azureUseMsi"`
	// The HashiCorp Cloud Platform API domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#hcp_api_domain HcsProvider#hcp_api_domain}
	HcpApiDomain *string `field:"optional" json:"hcpApiDomain" yaml:"hcpApiDomain"`
	// The HashiCorp Consul Service product name on the Azure marketplace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcs/0.5.1/docs#hcs_marketplace_product_name HcsProvider#hcs_marketplace_product_name}
	HcsMarketplaceProductName *string `field:"optional" json:"hcsMarketplaceProductName" yaml:"hcsMarketplaceProductName"`
}

