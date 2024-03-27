# MfaEvent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the MFA event. | 
**username** | **str** | The name of the user associated with the MFA event. | 
**timestamp** | **datetime** | The time (in UTC) at which the user logged in to the MFA event, represented using the RFC-3339 standard. | [optional] 
**application** | **str** | The name of the application associated with the MFA event. | [optional] 
**result** | **str** | The result of the MFA event. | [optional] 
**reason** | **str** | The reason for the result of the MFA event. When the &#x60;result&#x60; is &#x60;DENIED&#x60;, this field contains information on why the MFA event failed. | [optional] 
**second_factor** | **str** | The second factor used for the MFA event. | [optional] 
**client_device** | [**ClientDevice**](ClientDevice.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.mfa_event import MfaEvent

# TODO update the JSON string below
json = "{}"
# create an instance of MfaEvent from a JSON string
mfa_event_instance = MfaEvent.from_json(json)
# print the JSON string representation of the object
print(MfaEvent.to_json())

# convert the object into a dict
mfa_event_dict = mfa_event_instance.to_dict()
# create an instance of MfaEvent from a dict
mfa_event_form_dict = mfa_event.from_dict(mfa_event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


