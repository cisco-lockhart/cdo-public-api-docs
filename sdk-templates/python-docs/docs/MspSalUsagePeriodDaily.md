# MspSalUsagePeriodDaily

Event usage for one completed UTC day.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_date** | **date** | The following UTC day. This date is not included in the period. | [optional] 
**firewall_bytes** | **int** | Firewall event data ingested during this day, in bytes. | [optional] 
**sdwan_bytes** | **int** | SD-WAN event data ingested during this day, in bytes. | [optional] 
**start_date** | **date** | The UTC day. This date is included in the period. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_period_daily import MspSalUsagePeriodDaily

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsagePeriodDaily from a JSON string
msp_sal_usage_period_daily_instance = MspSalUsagePeriodDaily.from_json(json)
# print the JSON string representation of the object
print(MspSalUsagePeriodDaily.to_json())

# convert the object into a dict
msp_sal_usage_period_daily_dict = msp_sal_usage_period_daily_instance.to_dict()
# create an instance of MspSalUsagePeriodDaily from a dict
msp_sal_usage_period_daily_form_dict = msp_sal_usage_period_daily.from_dict(msp_sal_usage_period_daily_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


