package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func agentDeploymentCustomizationOverrideResourceRequirementFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cpu_limit": {
			Type:        schema.TypeString,
			Description: "The maximum CPU limit for agent",
			Optional:    true,
		},
		"cpu_request": {
			Type:        schema.TypeString,
			Description: "The minimum CPU required for agent",
			Optional:    true,
		},
		"memory_limit": {
			Type:        schema.TypeString,
			Description: "The maximum memory limit for agent",
			Optional:    true,
		},
		"memory_request": {
			Type:        schema.TypeString,
			Description: "The minimum memory required for agent",
			Optional:    true,
		},
	}
	return s
}

func agentDeploymentCustomizationFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"append_tolerations": {
			Type:        schema.TypeString,
			Description: "User defined tolerations to append to agent",
			Optional:    true,
		},
		"override_affinity": {
			Type:        schema.TypeString,
			Description: "User defined affinity to override default agent affinity",
			Optional:    true,
		},
		"override_resource_requirements": {
			Type:        schema.TypeList,
			Description: "User defined resource requirements to set on the agent",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: agentDeploymentCustomizationOverrideResourceRequirementFields(),
			},
		},
	}
	return s
}
