package ssmincidents_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

// only one replication set resource can be active at once, so we must have serialised tests
func TestAccSSMIncidentsResponsePlan_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Response Plan Resource Tests": {
			"basic":                  testResponsePlan_basic,
			"update":                 testResponsePlan_updateRequiredFields,
			"updateTags":             testResponsePlan_updateTags,
			"updateEmptyTags":        testResponsePlan_updateEmptyTags,
			"disappears":             testResponsePlan_disappears,
			"incidentTemplateFields": testResponsePlan_incidentTemplateOptionalFields,
			"displayName":            testResponsePlan_displayName,
			"chatChannel":            testResponsePlan_chatChannel,
			"engagement":             testResponsePlan_engagement,
			"action":                 testResponsePlan_action,
			/*
				Comment out integration test as the configured PagerDuty secretId is invalid and the test will fail,
				as we do not want to expose credentials to public repository.

				Tested locally and PagerDuty integration work with response plan.
			*/
			//"integration": testResponsePlan_integration,
		},
		"Response Plan Data Source Tests": {
			"basic": testResponsePlanDataSource_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
