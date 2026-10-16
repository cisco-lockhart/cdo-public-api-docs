# LogicalDevice

The logical devices on this chassis that are managed on the same Security Cloud Control tenant as the chassis. A chassis can host logical devices managed on other tenants; those are not listed.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the logical device in Security Cloud Control. Use it to retrieve the logical device from the inventory. | [optional] 
**uid_on_fmc** | **str** | The unique identifier, represented as a UUID, of the logical device&#39;s device record on the cdFMC. This field is deprecated: retrieve the logical device by &#x60;uid&#x60; and use its &#x60;deviceRecordOnFmc&#x60; instead. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.logical_device import LogicalDevice

# TODO update the JSON string below
json = "{}"
# create an instance of LogicalDevice from a JSON string
logical_device_instance = LogicalDevice.from_json(json)
# print the JSON string representation of the object
print(LogicalDevice.to_json())

# convert the object into a dict
logical_device_dict = logical_device_instance.to_dict()
# create an instance of LogicalDevice from a dict
logical_device_form_dict = logical_device.from_dict(logical_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


