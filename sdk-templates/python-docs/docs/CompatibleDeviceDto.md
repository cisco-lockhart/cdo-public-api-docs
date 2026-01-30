# CompatibleDeviceDto

Information about a device that is compatible with a specific upgrade version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**current_software_version** | **str** | The software version currently running on the device | [optional] 
**hardware_model** | **str** | The hardware model of the device | 
**managed_tenant_display_name** | **str** | The display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant in CDO. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this device belongs to. | 
**name** | **str** | The name of the device | 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_device_dto import CompatibleDeviceDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleDeviceDto from a JSON string
compatible_device_dto_instance = CompatibleDeviceDto.from_json(json)
# print the JSON string representation of the object
print(CompatibleDeviceDto.to_json())

# convert the object into a dict
compatible_device_dto_dict = compatible_device_dto_instance.to_dict()
# create an instance of CompatibleDeviceDto from a dict
compatible_device_dto_form_dict = compatible_device_dto.from_dict(compatible_device_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


