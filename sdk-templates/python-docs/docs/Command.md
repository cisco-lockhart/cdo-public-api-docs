# Command


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**body** | **str** | The body of the command (optional). | [optional] 
**link** | **str** | The FMC endpoint to invoke for this command. | 
**method** | **str** | The HTTP method for the command. | 

## Example

```python
from scc_firewall_manager_sdk.models.command import Command

# TODO update the JSON string below
json = "{}"
# create an instance of Command from a JSON string
command_instance = Command.from_json(json)
# print the JSON string representation of the object
print(Command.to_json())

# convert the object into a dict
command_dict = command_instance.to_dict()
# create an instance of Command from a dict
command_form_dict = command.from_dict(command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


