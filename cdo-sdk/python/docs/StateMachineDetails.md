# StateMachineDetails

The device state machine details.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**identifier** | **str** |  | [optional] 
**last_error** | [**StateMachineError**](StateMachineError.md) |  | [optional] 

## Example

```python
from cdo_python_sdk.models.state_machine_details import StateMachineDetails

# TODO update the JSON string below
json = "{}"
# create an instance of StateMachineDetails from a JSON string
state_machine_details_instance = StateMachineDetails.from_json(json)
# print the JSON string representation of the object
print(StateMachineDetails.to_json())

# convert the object into a dict
state_machine_details_dict = state_machine_details_instance.to_dict()
# create an instance of StateMachineDetails from a dict
state_machine_details_form_dict = state_machine_details.from_dict(state_machine_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


