# PagePolicyAnalysisData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[PolicyAnalysisData]**](PolicyAnalysisData.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.page_policy_analysis_data import PagePolicyAnalysisData

# TODO update the JSON string below
json = "{}"
# create an instance of PagePolicyAnalysisData from a JSON string
page_policy_analysis_data_instance = PagePolicyAnalysisData.from_json(json)
# print the JSON string representation of the object
print(PagePolicyAnalysisData.to_json())

# convert the object into a dict
page_policy_analysis_data_dict = page_policy_analysis_data_instance.to_dict()
# create an instance of PagePolicyAnalysisData from a dict
page_policy_analysis_data_form_dict = page_policy_analysis_data.from_dict(page_policy_analysis_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


