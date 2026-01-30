# CreateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | The human-readable description of the object | [optional] 
**labels** | **List[str]** | The labels for the object | [optional] 
**name** | **str** | The name of the object | 
**tags** | **Dict[str, List[str]]** | The tags for the object | [optional] 
**target_ids** | **List[str]** | Set of IDs for targets that contain the object. A target can be, for example, a device, service, or a shared policy (Ruleset). | [optional] 
**value** | [**SharedObjectValue**](SharedObjectValue.md) |  | 

## Example

```python
from scc_firewall_manager_sdk.models.create_request import CreateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateRequest from a JSON string
create_request_instance = CreateRequest.from_json(json)
# print the JSON string representation of the object
print(CreateRequest.to_json())

# convert the object into a dict
create_request_dict = create_request_instance.to_dict()
# create an instance of CreateRequest from a dict
create_request_form_dict = create_request.from_dict(create_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


