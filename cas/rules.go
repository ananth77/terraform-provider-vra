package cas

import (
	"github.com/vmware/cas-sdk-go/pkg/models"

	"github.com/hashicorp/terraform/helper/schema"
)

// nicsSDKSchema returns the schema to use for the nics property
func rulesSDKSchema(isRequired bool) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: isRequired,
		Optional: !isRequired,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"access": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"ip_range_cidr": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"ports": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"protocol": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func expandSDKRules(configRules []interface{}) []*models.Rule {
	rules := make([]*models.Rule, 0, len(configRules))

	for _, configRule := range configRules {
		ruleMap := configRule.(map[string]interface{})

		rule := models.Rule{
			Access:      withString(ruleMap["access"].(string)),
			IPRangeCidr: withString(ruleMap["ip_range_cidr"].(string)),
			Ports:       withString(ruleMap["ports"].(string)),
			Protocol:    withString(ruleMap["protocol"].(string)),
		}

		if v, ok := ruleMap["name"].(string); ok && v != "" {
			rule.Name = v
		}

		rules = append(rules, &rule)
	}

	return rules
}