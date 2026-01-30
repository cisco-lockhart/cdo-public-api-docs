# CommandResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | **object** | The data associated with the command response. | [optional] 
**message** | **str** | The response message, if applicable. | [optional] 
**request_id** | **str** | The unique identifier of the request, represented as a UUID. | [optional] 
**response_origin** | **str** | The origin system of the response. | [optional] 
**status** | **str** | The status of the command. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.command_response import CommandResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CommandResponse from a JSON string
command_response_instance = CommandResponse.from_json(json)
# print the JSON string representation of the object
print(CommandResponse.to_json())

# convert the object into a dict
command_response_dict = command_response_instance.to_dict()
# create an instance of CommandResponse from a dict
command_response_form_dict = command_response.from_dict(command_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


