# ThousandEyesAiOpsStatusUpdateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**service_enabled** | **bool** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_ai_ops_status_update_request import ThousandEyesAiOpsStatusUpdateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesAiOpsStatusUpdateRequest from a JSON string
thousand_eyes_ai_ops_status_update_request_instance = ThousandEyesAiOpsStatusUpdateRequest.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesAiOpsStatusUpdateRequest.to_json())

# convert the object into a dict
thousand_eyes_ai_ops_status_update_request_dict = thousand_eyes_ai_ops_status_update_request_instance.to_dict()
# create an instance of ThousandEyesAiOpsStatusUpdateRequest from a dict
thousand_eyes_ai_ops_status_update_request_form_dict = thousand_eyes_ai_ops_status_update_request.from_dict(thousand_eyes_ai_ops_status_update_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


