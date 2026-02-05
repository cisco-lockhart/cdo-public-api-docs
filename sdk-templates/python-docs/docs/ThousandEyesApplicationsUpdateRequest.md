# ThousandEyesApplicationsUpdateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applications** | [**List[ThousandEyesApplicationUpdateRequest]**](ThousandEyesApplicationUpdateRequest.md) |  | 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_applications_update_request import ThousandEyesApplicationsUpdateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesApplicationsUpdateRequest from a JSON string
thousand_eyes_applications_update_request_instance = ThousandEyesApplicationsUpdateRequest.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesApplicationsUpdateRequest.to_json())

# convert the object into a dict
thousand_eyes_applications_update_request_dict = thousand_eyes_applications_update_request_instance.to_dict()
# create an instance of ThousandEyesApplicationsUpdateRequest from a dict
thousand_eyes_applications_update_request_form_dict = thousand_eyes_applications_update_request.from_dict(thousand_eyes_applications_update_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


