# MspSalRateDailyItemDto

SAL event rate for one managed tenant.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**daily_rates** | [**List[MspSalRatePeriodDaily]**](MspSalRatePeriodDaily.md) | Event rate for each stored, completed UTC day in the requested time range, ordered from oldest to newest. Days without data and the current day are not listed. | [optional] 
**managed_tenant_display_name** | **str** | The human-readable display name of the managed tenant. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant. | [optional] 
**pay_type** | **str** | The billing classification of the managed tenant. This value is null when no pay type is stored. | 
**sal_rate_ingest_date** | **datetime** | When the SAL event rate was last successfully ingested for this managed tenant. This value is null when the rate has not yet been ingested. | 
**summary** | [**MspSalRateSummary**](MspSalRateSummary.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_rate_daily_item_dto import MspSalRateDailyItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalRateDailyItemDto from a JSON string
msp_sal_rate_daily_item_dto_instance = MspSalRateDailyItemDto.from_json(json)
# print the JSON string representation of the object
print(MspSalRateDailyItemDto.to_json())

# convert the object into a dict
msp_sal_rate_daily_item_dto_dict = msp_sal_rate_daily_item_dto_instance.to_dict()
# create an instance of MspSalRateDailyItemDto from a dict
msp_sal_rate_daily_item_dto_form_dict = msp_sal_rate_daily_item_dto.from_dict(msp_sal_rate_daily_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


