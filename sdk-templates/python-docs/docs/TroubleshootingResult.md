# TroubleshootingResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**download_urls** | **List[str]** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.troubleshooting_result import TroubleshootingResult

# TODO update the JSON string below
json = "{}"
# create an instance of TroubleshootingResult from a JSON string
troubleshooting_result_instance = TroubleshootingResult.from_json(json)
# print the JSON string representation of the object
print(TroubleshootingResult.to_json())

# convert the object into a dict
troubleshooting_result_dict = troubleshooting_result_instance.to_dict()
# create an instance of TroubleshootingResult from a dict
troubleshooting_result_form_dict = troubleshooting_result.from_dict(troubleshooting_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


