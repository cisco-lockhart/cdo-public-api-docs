# StateMachineError


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action_identifier** | **str** |  | [optional] 
**error_code** | **str** |  | [optional] 
**error_message** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.state_machine_error import StateMachineError

# TODO update the JSON string below
json = "{}"
# create an instance of StateMachineError from a JSON string
state_machine_error_instance = StateMachineError.from_json(json)
# print the JSON string representation of the object
print(StateMachineError.to_json())

# convert the object into a dict
state_machine_error_dict = state_machine_error_instance.to_dict()
# create an instance of StateMachineError from a dict
state_machine_error_form_dict = state_machine_error.from_dict(state_machine_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


