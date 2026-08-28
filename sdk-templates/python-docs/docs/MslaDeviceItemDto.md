# MslaDeviceItemDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_name** | **str** | The name of the device. | [optional] 
**device_uid** | **str** | The UID of the device in Security Cloud Control. | [optional] 
**entitlement** | **str** | The MSLA entitlement identifier. | [optional] 
**model_number** | **str** | The hardware model number. | [optional] 
**serial** | **str** | The serial number of the device. | [optional] 
**uid** | **str** | The UID of the entitlement assignment row. | [optional] 

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


