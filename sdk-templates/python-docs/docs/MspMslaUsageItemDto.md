# MspMslaUsageItemDto

MSLA entitlement consumption across all managed tenants of an MSP portal.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_count** | **int** | The total number of devices across all managed tenants consuming this entitlement. | [optional] 
**entitlement** | **str** | MSLA entitlement identifier. | [optional] 
**managed_tenants** | [**List[MspMslaUsageTenantDto]**](MspMslaUsageTenantDto.md) | The managed tenants with at least one device consuming this entitlement. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_msla_usage_item_dto import MspMslaUsageItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspMslaUsageItemDto from a JSON string
msp_msla_usage_item_dto_instance = MspMslaUsageItemDto.from_json(json)
# print the JSON string representation of the object
print(MspMslaUsageItemDto.to_json())

# convert the object into a dict
msp_msla_usage_item_dto_dict = msp_msla_usage_item_dto_instance.to_dict()
# create an instance of MspMslaUsageItemDto from a dict
msp_msla_usage_item_dto_form_dict = msp_msla_usage_item_dto.from_dict(msp_msla_usage_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


