# Event

The events recorded in this Change Log.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | A human-readable description of the Change Log event. | [optional] 
**diff** | **str** | A textual Diff representation of the changes made to the configuration. | [optional] 
**username** | **str** | The username of the user that triggered the Change Log event. | [optional] 
**var_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the Change Log Event occurred. | [optional] 
**action** | **str** | The action performed. | [optional] 
**change_request_uid** | **str** | The unique Identifier of the Change Request associated with the Change Log event. | [optional] 
**change_request_name** | **str** | The name of the Change Request associated with the Change Log event. | [optional] 

## Example

```python
from cdo_python_sdk.models.event import Event

# TODO update the JSON string below
json = "{}"
# create an instance of Event from a JSON string
event_instance = Event.from_json(json)
# print the JSON string representation of the object
print(Event.to_json())

# convert the object into a dict
event_dict = event_instance.to_dict()
# create an instance of Event from a dict
event_form_dict = event.from_dict(event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


