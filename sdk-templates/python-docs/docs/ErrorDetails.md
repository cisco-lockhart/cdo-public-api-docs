# ErrorDetails


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**details** | **Dict[str, str]** |  | [optional] 
**severity** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.error_details import ErrorDetails

# TODO update the JSON string below
json = "{}"
# create an instance of ErrorDetails from a JSON string
error_details_instance = ErrorDetails.from_json(json)
# print the JSON string representation of the object
print(ErrorDetails.to_json())

# convert the object into a dict
error_details_dict = error_details_instance.to_dict()
# create an instance of ErrorDetails from a dict
error_details_form_dict = error_details.from_dict(error_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


