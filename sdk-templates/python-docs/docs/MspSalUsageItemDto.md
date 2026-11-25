# MspSalUsageItemDto

SAL event usage for one managed tenant.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**managed_tenant_display_name** | **str** | The human-readable display name of the managed tenant. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant. | [optional] 
**monthly_usage** | [**List[MspSalUsagePeriodMonthly]**](MspSalUsagePeriodMonthly.md) | Event usage for each UTC calendar month in the requested time range, ordered from oldest to newest. | [optional] 
**pay_type** | **str** | The billing classification of the managed tenant. This value is null when no pay type is stored. | 
**sal_usage_ingest_date** | **datetime** | When SAL usage was last successfully ingested for this managed tenant. This value is null when SAL usage has not yet been ingested. | 
**summary** | [**MspSalUsageSummary**](MspSalUsageSummary.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_item_dto import MspSalUsageItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsageItemDto from a JSON string
msp_sal_usage_item_dto_instance = MspSalUsageItemDto.from_json(json)
# print the JSON string representation of the object
print(MspSalUsageItemDto.to_json())

# convert the object into a dict
msp_sal_usage_item_dto_dict = msp_sal_usage_item_dto_instance.to_dict()
# create an instance of MspSalUsageItemDto from a dict
msp_sal_usage_item_dto_form_dict = msp_sal_usage_item_dto.from_dict(msp_sal_usage_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


