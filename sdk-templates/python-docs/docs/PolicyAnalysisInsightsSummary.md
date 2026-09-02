# PolicyAnalysisInsightsSummary


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**anomalous_rules** | **int** | The total number of anomalous rules | [optional] 
**data_source_uid** | **str** | The data source UID for the summarized NAT policies | [optional] 
**disabled_rules** | **int** | The total number of disabled rules | [optional] 
**policies** | **int** | The number of policies included in the summary | [optional] 
**policy_type** | **str** | The policy type for the summarized NAT policies | [optional] 
**redundant_rules** | **int** | The total number of redundant rules using v1 semantics | [optional] 
**shadowed_rules** | **int** | The total number of shadowed rules using v1 semantics | [optional] 
**total_rules** | **int** | The total number of rules | [optional] 
**unhealthy_rules** | **int** | The number of unique unhealthy rules across shadowed and redundant sets | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.policy_analysis_insights_summary import PolicyAnalysisInsightsSummary

# TODO update the JSON string below
json = "{}"
# create an instance of PolicyAnalysisInsightsSummary from a JSON string
policy_analysis_insights_summary_instance = PolicyAnalysisInsightsSummary.from_json(json)
# print the JSON string representation of the object
print(PolicyAnalysisInsightsSummary.to_json())

# convert the object into a dict
policy_analysis_insights_summary_dict = policy_analysis_insights_summary_instance.to_dict()
# create an instance of PolicyAnalysisInsightsSummary from a dict
policy_analysis_insights_summary_form_dict = policy_analysis_insights_summary.from_dict(policy_analysis_insights_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


