# MspMslaDeviceEntitlementDto

A single MSLA entitlement assignment for one device across an MSP-managed tenant.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_name** | **str** | The name of the device. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 
**entitlement** | **str** | MSLA entitlement identifier. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant that owns this device. | [optional] 
**model_number** | **str** | The hardware model number of the device. | [optional] 
**serial** | **str** | The serial number of the device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the MSLA entitlement assignment. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_msla_device_entitlement_dto import MspMslaDeviceEntitlementDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspMslaDeviceEntitlementDto from a JSON string
msp_msla_device_entitlement_dto_instance = MspMslaDeviceEntitlementDto.from_json(json)
# print the JSON string representation of the object
print(MspMslaDeviceEntitlementDto.to_json())

# convert the object into a dict
msp_msla_device_entitlement_dto_dict = msp_msla_device_entitlement_dto_instance.to_dict()
# create an instance of MspMslaDeviceEntitlementDto from a dict
msp_msla_device_entitlement_dto_form_dict = msp_msla_device_entitlement_dto.from_dict(msp_msla_device_entitlement_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


