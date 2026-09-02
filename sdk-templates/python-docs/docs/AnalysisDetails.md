# AnalysisDetails


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**analysis_start_time** | **datetime** |  | [optional] 
**status** | **str** |  | [optional] 
**summary** | [**Summary**](Summary.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.analysis_details import AnalysisDetails

# TODO update the JSON string below
json = "{}"
# create an instance of AnalysisDetails from a JSON string
analysis_details_instance = AnalysisDetails.from_json(json)
# print the JSON string representation of the object
print(AnalysisDetails.to_json())

# convert the object into a dict
analysis_details_dict = analysis_details_instance.to_dict()
# create an instance of AnalysisDetails from a dict
analysis_details_form_dict = analysis_details.from_dict(analysis_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


