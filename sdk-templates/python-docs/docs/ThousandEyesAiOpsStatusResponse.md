# ThousandEyesAiOpsStatusResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**aiops_onboarded** | **bool** |  | [optional] 
**service_enabled** | **bool** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_ai_ops_status_response import ThousandEyesAiOpsStatusResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesAiOpsStatusResponse from a JSON string
thousand_eyes_ai_ops_status_response_instance = ThousandEyesAiOpsStatusResponse.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesAiOpsStatusResponse.to_json())

# convert the object into a dict
thousand_eyes_ai_ops_status_response_dict = thousand_eyes_ai_ops_status_response_instance.to_dict()
# create an instance of ThousandEyesAiOpsStatusResponse from a dict
thousand_eyes_ai_ops_status_response_form_dict = thousand_eyes_ai_ops_status_response.from_dict(thousand_eyes_ai_ops_status_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


