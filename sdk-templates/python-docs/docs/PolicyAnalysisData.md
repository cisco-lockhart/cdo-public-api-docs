# PolicyAnalysisData

Latest NAT policy analysis data.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**analysis_details** | [**AnalysisDetails**](AnalysisDetails.md) |  | [optional] 
**data_source** | [**DataSource**](DataSource.md) |  | [optional] 
**impacted_sources** | [**ImpactedSources**](ImpactedSources.md) |  | [optional] 
**link** | **str** |  | [optional] 
**policy** | [**Policy**](Policy.md) |  | [optional] 
**remediation_status** | **str** |  | [optional] 
**uid** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.policy_analysis_data import PolicyAnalysisData

# TODO update the JSON string below
json = "{}"
# create an instance of PolicyAnalysisData from a JSON string
policy_analysis_data_instance = PolicyAnalysisData.from_json(json)
# print the JSON string representation of the object
print(PolicyAnalysisData.to_json())

# convert the object into a dict
policy_analysis_data_dict = policy_analysis_data_instance.to_dict()
# create an instance of PolicyAnalysisData from a dict
policy_analysis_data_form_dict = policy_analysis_data.from_dict(policy_analysis_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


