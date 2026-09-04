# AllowListResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_ranges** | [**ProviderRanges**](ProviderRanges.md) |  | [optional] 
**sources** | [**List[AllowListSource]**](AllowListSource.md) | The five named CDO sources. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.allow_list_response import AllowListResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AllowListResponse from a JSON string
allow_list_response_instance = AllowListResponse.from_json(json)
# print the JSON string representation of the object
print(AllowListResponse.to_json())

# convert the object into a dict
allow_list_response_dict = allow_list_response_instance.to_dict()
# create an instance of AllowListResponse from a dict
allow_list_response_form_dict = allow_list_response.from_dict(allow_list_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


