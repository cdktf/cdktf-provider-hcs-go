package datahcsagenthelmconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcsAgentHelmConfigConfig struct {
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
	// The name of the AKS cluster that will consume the Helm config.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#aks_cluster_name DataHcsAgentHelmConfig#aks_cluster_name}
	AksClusterName *string `field:"required" json:"aksClusterName" yaml:"aksClusterName"`
	// The name of the HCS Azure Managed Application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#managed_application_name DataHcsAgentHelmConfig#managed_application_name}
	ManagedApplicationName *string `field:"required" json:"managedApplicationName" yaml:"managedApplicationName"`
	// The name of the Resource Group in which the HCS Azure Managed Application belongs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#resource_group_name DataHcsAgentHelmConfig#resource_group_name}
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The resource group name of the AKS cluster that will consume the Helm config.
	//
	// If not specified, it is defaulted to the value of `resource_group_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#aks_resource_group DataHcsAgentHelmConfig#aks_resource_group}
	AksResourceGroup *string `field:"optional" json:"aksResourceGroup" yaml:"aksResourceGroup"`
	// Denotes that the gossip ports should be exposed. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#expose_gossip_ports DataHcsAgentHelmConfig#expose_gossip_ports}
	ExposeGossipPorts interface{} `field:"optional" json:"exposeGossipPorts" yaml:"exposeGossipPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#id DataHcsAgentHelmConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/d/agent_helm_config#timeouts DataHcsAgentHelmConfig#timeouts}
	Timeouts *DataHcsAgentHelmConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

