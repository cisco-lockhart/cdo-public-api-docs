# MslaDeviceItemDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_name** | **str** | The name of the device. Device names are unique in Security Cloud Control. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 
**entitlement** | **str** | The unique identifier of the MSLA entitlement assigned to the device. | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on (ASA-only). This field can be missing in the case of a partially onboarded device. | [optional] 
**serial** | **str** | The serial number of the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the entitlement assignment row. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msla_device_item_dto import MslaDeviceItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MslaDeviceItemDto from a JSON string
msla_device_item_dto_instance = MslaDeviceItemDto.from_json(json)
# print the JSON string representation of the object
print(MslaDeviceItemDto.to_json())

# convert the object into a dict
msla_device_item_dto_dict = msla_device_item_dto_instance.to_dict()
# create an instance of MslaDeviceItemDto from a dict
msla_device_item_dto_form_dict = msla_device_item_dto.from_dict(msla_device_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


