# UpgradedDeviceDto

The devices included in this upgrade run.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error_msg** | **str** | Error message if the device encountered a failure during the upgrade. | [optional] 
**managed_tenant_display_name** | **str** | Display name of the managed tenant in Security Cloud Control. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this device belongs to. | [optional] 
**name** | **str** | The name of the device. | [optional] 
**status** | **str** | The overall status of the device in this upgrade run. One of: PENDING, IN_PROGRESS, SUCCESS, FAILED, CLEANUP_FAILED. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.upgraded_device_dto import UpgradedDeviceDto

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradedDeviceDto from a JSON string
upgraded_device_dto_instance = UpgradedDeviceDto.from_json(json)
# print the JSON string representation of the object
print(UpgradedDeviceDto.to_json())

# convert the object into a dict
upgraded_device_dto_dict = upgraded_device_dto_instance.to_dict()
# create an instance of UpgradedDeviceDto from a dict
upgraded_device_dto_form_dict = upgraded_device_dto.from_dict(upgraded_device_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


