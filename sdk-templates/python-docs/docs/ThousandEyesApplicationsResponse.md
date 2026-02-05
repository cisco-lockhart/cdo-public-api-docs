# ThousandEyesApplicationsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applications** | [**List[ThousandEyesApplicationResponse]**](ThousandEyesApplicationResponse.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_applications_response import ThousandEyesApplicationsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesApplicationsResponse from a JSON string
thousand_eyes_applications_response_instance = ThousandEyesApplicationsResponse.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesApplicationsResponse.to_json())

# convert the object into a dict
thousand_eyes_applications_response_dict = thousand_eyes_applications_response_instance.to_dict()
# create an instance of ThousandEyesApplicationsResponse from a dict
thousand_eyes_applications_response_form_dict = thousand_eyes_applications_response.from_dict(thousand_eyes_applications_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


