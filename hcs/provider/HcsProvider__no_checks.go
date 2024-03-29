// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HcsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HcsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHcsProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateHcsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHcsProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHcsProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HcsProvider) validateSetAzureUseMsiParameters(val interface{}) error {
	return nil
}

func validateNewHcsProviderParameters(scope constructs.Construct, id *string, config *HcsProviderConfig) error {
	return nil
}

