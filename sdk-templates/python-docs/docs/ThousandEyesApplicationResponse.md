# ThousandEyesApplicationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application** | **str** |  | [optional] 
**available** | **bool** |  | [optional] 
**exist_in_cd_fmc_policy** | **bool** |  | [optional] 
**subscribed_to_notifications** | **bool** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_application_response import ThousandEyesApplicationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesApplicationResponse from a JSON string
thousand_eyes_application_response_instance = ThousandEyesApplicationResponse.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesApplicationResponse.to_json())

# convert the object into a dict
thousand_eyes_application_response_dict = thousand_eyes_application_response_instance.to_dict()
# create an instance of ThousandEyesApplicationResponse from a dict
thousand_eyes_application_response_form_dict = thousand_eyes_application_response.from_dict(thousand_eyes_application_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


