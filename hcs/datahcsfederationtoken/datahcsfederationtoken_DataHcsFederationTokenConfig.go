package datahcsfederationtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcsFederationTokenConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the HCS Azure Managed Application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/federation_token#managed_application_name DataHcsFederationToken#managed_application_name}
	ManagedApplicationName *string `field:"required" json:"managedApplicationName" yaml:"managedApplicationName"`
	// The name of the Resource Group in which the HCS Azure Managed Application belongs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/federation_token#resource_group_name DataHcsFederationToken#resource_group_name}
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/federation_token#id DataHcsFederationToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/federation_token#timeouts DataHcsFederationToken#timeouts}
	Timeouts *DataHcsFederationTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

