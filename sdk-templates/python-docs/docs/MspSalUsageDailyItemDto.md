# MspSalUsageDailyItemDto

Daily SAL event usage for one managed tenant.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**daily_usage** | [**List[MspSalUsagePeriodDaily]**](MspSalUsagePeriodDaily.md) | Event usage for each completed UTC day in the requested time range, ordered from oldest to newest. Days run from the 1st of the first calendar month in the range up to and including yesterday (UTC); days without stored usage are listed with 0. | [optional] 
**managed_tenant_display_name** | **str** | The human-readable display name of the managed tenant. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant. | [optional] 
**pay_type** | **str** | The billing classification of the managed tenant. This value is null when no pay type is stored. | 
**sal_usage_ingest_date** | **datetime** | When daily SAL usage was last successfully ingested for this managed tenant. Daily usage is ingested separately from monthly usage. This value is null when none of the ingested daily usage falls within the past 12 calendar months: before the first successful ingest, or when ingests have failed for so long that all ingested usage is older. | 
**summary** | [**MspSalUsageSummary**](MspSalUsageSummary.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_daily_item_dto import MspSalUsageDailyItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsageDailyItemDto from a JSON string
msp_sal_usage_daily_item_dto_instance = MspSalUsageDailyItemDto.from_json(json)
# print the JSON string representation of the object
print(MspSalUsageDailyItemDto.to_json())

# convert the object into a dict
msp_sal_usage_daily_item_dto_dict = msp_sal_usage_daily_item_dto_instance.to_dict()
# create an instance of MspSalUsageDailyItemDto from a dict
msp_sal_usage_daily_item_dto_form_dict = msp_sal_usage_daily_item_dto.from_dict(msp_sal_usage_daily_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


