# MslaUsageItemDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_count** | **int** | Number of devices in the tenant that have this entitlement assigned in cdFMC. Zero if no devices have consumed this entitlement. | [optional] 
**entitlement** | **str** | The MSLA entitlement identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msla_usage_item_dto import MslaUsageItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MslaUsageItemDto from a JSON string
msla_usage_item_dto_instance = MslaUsageItemDto.from_json(json)
# print the JSON string representation of the object
print(MslaUsageItemDto.to_json())

# convert the object into a dict
msla_usage_item_dto_dict = msla_usage_item_dto_instance.to_dict()
# create an instance of MslaUsageItemDto from a dict
msla_usage_item_dto_form_dict = msla_usage_item_dto.from_dict(msla_usage_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


