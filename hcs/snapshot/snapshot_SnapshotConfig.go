package snapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnapshotConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/snapshot#managed_application_name Snapshot#managed_application_name}
	ManagedApplicationName *string `field:"required" json:"managedApplicationName" yaml:"managedApplicationName"`
	// The name of the Resource Group in which the HCS Azure Managed Application belongs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/snapshot#resource_group_name Snapshot#resource_group_name}
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The name of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/snapshot#snapshot_name Snapshot#snapshot_name}
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/snapshot#id Snapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcs/r/snapshot#timeouts Snapshot#timeouts}
	Timeouts *SnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

