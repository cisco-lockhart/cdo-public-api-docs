# ThousandEyesApplicationUpdateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application** | **str** |  | 
**subscribed_to_notifications** | **bool** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_application_update_request import ThousandEyesApplicationUpdateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesApplicationUpdateRequest from a JSON string
thousand_eyes_application_update_request_instance = ThousandEyesApplicationUpdateRequest.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesApplicationUpdateRequest.to_json())

# convert the object into a dict
thousand_eyes_application_update_request_dict = thousand_eyes_application_update_request_instance.to_dict()
# create an instance of ThousandEyesApplicationUpdateRequest from a dict
thousand_eyes_application_update_request_form_dict = thousand_eyes_application_update_request.from_dict(thousand_eyes_application_update_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


