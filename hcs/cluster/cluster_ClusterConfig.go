package cluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterConfig struct {
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
	// The mode of the cluster ('Development' or 'Production').
	//
	// Development clusters only have a single Consul server. Production clusters are fully supported, full featured, and deploy with a minimum of three hosts.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#cluster_mode Cluster#cluster_mode}
	ClusterMode *string `field:"required" json:"clusterMode" yaml:"clusterMode"`
	// The contact email for the primary owner of the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#email Cluster#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The name of the HCS Azure Managed Application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#managed_application_name Cluster#managed_application_name}
	ManagedApplicationName *string `field:"required" json:"managedApplicationName" yaml:"managedApplicationName"`
	// The name of the Resource Group in which the HCS Azure Managed Application belongs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#resource_group_name Cluster#resource_group_name}
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Enables Consul audit logging for the cluster resource. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#audit_logging_enabled Cluster#audit_logging_enabled}
	AuditLoggingEnabled interface{} `field:"optional" json:"auditLoggingEnabled" yaml:"auditLoggingEnabled"`
	// The url of the Azure blob storage container to write audit logs to if `audit_logging_enabled` is `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#audit_log_storage_container_url Cluster#audit_log_storage_container_url}
	AuditLogStorageContainerUrl *string `field:"optional" json:"auditLogStorageContainerUrl" yaml:"auditLogStorageContainerUrl"`
	// The name of the cluster Managed Resource. If not specified, it is defaulted to the value of `managed_application_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#cluster_name Cluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The Consul data center name of the cluster. If not specified, it is defaulted to the value of `managed_application_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#consul_datacenter Cluster#consul_datacenter}
	ConsulDatacenter *string `field:"optional" json:"consulDatacenter" yaml:"consulDatacenter"`
	// Denotes that the cluster has an external endpoint for the Consul UI. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#consul_external_endpoint Cluster#consul_external_endpoint}
	ConsulExternalEndpoint interface{} `field:"optional" json:"consulExternalEndpoint" yaml:"consulExternalEndpoint"`
	// The token used to join a federation of Consul clusters.
	//
	// If the cluster is not part of a federation, this field will be empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#consul_federation_token Cluster#consul_federation_token}
	ConsulFederationToken *string `field:"optional" json:"consulFederationToken" yaml:"consulFederationToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#id Cluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The Azure region that the cluster is deployed to.
	//
	// If not specified, it is defaulted to the region of the Resource Group the Managed Application belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#location Cluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the Managed Resource Group in which the cluster resources belong.
	//
	// If not specified, it is defaulted to the value of `managed_application_name` with 'mrg-' prepended.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#managed_resource_group_name Cluster#managed_resource_group_name}
	ManagedResourceGroupName *string `field:"optional" json:"managedResourceGroupName" yaml:"managedResourceGroupName"`
	// The minimum Consul version of the cluster.
	//
	// If not specified, it is defaulted to the version that is currently recommended by HCS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#min_consul_version Cluster#min_consul_version}
	MinConsulVersion *string `field:"optional" json:"minConsulVersion" yaml:"minConsulVersion"`
	// The name of the Azure Marketplace HCS plan for the cluster.
	//
	// If not specified, it will default to the current HCS default plan (see the `hcs_plan_defaults` data source).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#plan_name Cluster#plan_name}
	PlanName *string `field:"optional" json:"planName" yaml:"planName"`
	// A mapping of tags to assign to the HCS Azure Managed Application resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#tags Cluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#timeouts Cluster#timeouts}
	Timeouts *ClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The VNET CIDR range of the Consul cluster. Defaults to `172.25.16.0/24`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/cluster#vnet_cidr Cluster#vnet_cidr}
	VnetCidr *string `field:"optional" json:"vnetCidr" yaml:"vnetCidr"`
}

