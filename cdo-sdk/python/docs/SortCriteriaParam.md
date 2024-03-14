# SortCriteriaParam


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sort** | **List[str]** |  | [optional] 
**sort_fields** | **List[str]** |  | [optional] 

## Example

```python
from openapi_client.models.sort_criteria_param import SortCriteriaParam

# TODO update the JSON string below
json = "{}"
# create an instance of SortCriteriaParam from a JSON string
sort_criteria_param_instance = SortCriteriaParam.from_json(json)
# print the JSON string representation of the object
print(SortCriteriaParam.to_json())

# convert the object into a dict
sort_criteria_param_dict = sort_criteria_param_instance.to_dict()
# create an instance of SortCriteriaParam from a dict
sort_criteria_param_form_dict = sort_criteria_param.from_dict(sort_criteria_param_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


