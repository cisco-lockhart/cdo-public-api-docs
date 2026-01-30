# InterfaceMetricError


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | [optional] 
**name** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_metric_error import InterfaceMetricError

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceMetricError from a JSON string
interface_metric_error_instance = InterfaceMetricError.from_json(json)
# print the JSON string representation of the object
print(InterfaceMetricError.to_json())

# convert the object into a dict
interface_metric_error_dict = interface_metric_error_instance.to_dict()
# create an instance of InterfaceMetricError from a dict
interface_metric_error_form_dict = interface_metric_error.from_dict(interface_metric_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


