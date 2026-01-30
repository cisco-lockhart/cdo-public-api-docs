# InterfaceErrorSummary


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error** | **int** |  | [optional] 
**stable** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_error_summary import InterfaceErrorSummary

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceErrorSummary from a JSON string
interface_error_summary_instance = InterfaceErrorSummary.from_json(json)
# print the JSON string representation of the object
print(InterfaceErrorSummary.to_json())

# convert the object into a dict
interface_error_summary_dict = interface_error_summary_instance.to_dict()
# create an instance of InterfaceErrorSummary from a dict
interface_error_summary_form_dict = interface_error_summary.from_dict(interface_error_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


