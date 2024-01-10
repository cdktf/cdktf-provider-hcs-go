// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutTimeoutsParameters(value *ClusterTimeouts) error {
	return nil
}

func validateCluster_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAuditLoggingEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAuditLogStorageContainerUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetClusterModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetClusterNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConsulDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConsulExternalEndpointParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConsulFederationTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetEmailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetManagedApplicationNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetManagedResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetMinConsulVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetPlanNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetVnetCidrParameters(val *string) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, config *ClusterConfig) error {
	return nil
}

