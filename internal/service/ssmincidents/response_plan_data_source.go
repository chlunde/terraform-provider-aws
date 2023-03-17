package ssmincidents

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_ssmincidents_response_plan")
func DataSourceResponsePlan() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceResponsePlanRead,

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssm_automation": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"document_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role_arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"document_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_account": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"values": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"dynamic_parameters": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"chat_channel": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engagements": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"incident_template": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"impact": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dedupe_string": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"incident_tags": tftags.TagsSchemaComputed(),
						"notification_target": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sns_topic_arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"integration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pagerduty": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"secret_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
		},
	}
}

const (
	DSNameResponsePlan = "Response Plan Data Source"
)

func dataSourceResponsePlanRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*conns.AWSClient).SSMIncidentsClient()

	d.SetId(d.Get("arn").(string))

	responsePlan, err := FindResponsePlanByID(context, client, d.Id())

	if err != nil {
		return create.DiagError(names.SSMIncidents, create.ErrActionReading, DSNameResponsePlan, d.Id(), err)
	}

	if err := setResponsePlanResourceData(d, responsePlan); err != nil {
		return create.DiagError(names.SSMIncidents, create.ErrActionReading, DSNameResponsePlan, d.Id(), err)
	}

	tags, err := listResourceTags(context, client, d.Id())
	if err != nil {
		return create.DiagError(names.SSMIncidents, create.ErrActionReading, DSNameResponsePlan, d.Id(), err)
	}

	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	//lintignore:AWSR002
	if err := d.Set("tags", tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return create.DiagError(names.SSMIncidents, create.ErrActionSetting, DSNameResponsePlan, d.Id(), err)
	}

	return nil
}
