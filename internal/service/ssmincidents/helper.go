package ssmincidents

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssmincidents"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
)

func getReplicationSetARN(context context.Context, client *ssmincidents.Client) (string, error) {
	replicationSets, err := client.ListReplicationSets(context, &ssmincidents.ListReplicationSetsInput{})

	if err != nil {
		return "", err
	}

	if len(replicationSets.ReplicationSetArns) == 0 {
		return "", fmt.Errorf("replication set could not be found")
	}

	// currently only one replication set is supported
	return replicationSets.ReplicationSetArns[0], nil
}

func setResponsePlanResourceData(
	d *schema.ResourceData,
	getResponsePlanOutput *ssmincidents.GetResponsePlanOutput,
) error {
	if err := d.Set("action", flattenAction(getResponsePlanOutput.Actions)); err != nil {
		return err
	}
	if err := d.Set("arn", getResponsePlanOutput.Arn); err != nil {
		return err
	}
	if err := d.Set("chat_channel", flattenChatChannel(getResponsePlanOutput.ChatChannel)); err != nil {
		return err
	}
	if err := d.Set("display_name", getResponsePlanOutput.DisplayName); err != nil {
		return err
	}
	if err := d.Set("engagements", flex.FlattenStringValueSet(getResponsePlanOutput.Engagements)); err != nil {
		return err
	}
	if err := d.Set("incident_template", flattenIncidentTemplate(getResponsePlanOutput.IncidentTemplate)); err != nil {
		return err
	}
	if err := d.Set("integration", flattenIntegration(getResponsePlanOutput.Integrations)); err != nil {
		return err
	}
	if err := d.Set("name", getResponsePlanOutput.Name); err != nil {
		return err
	}
	return nil
}
