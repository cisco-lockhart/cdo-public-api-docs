# MspSalUsagePeriodMonthly

Event usage for one UTC calendar month. Dates are used because SAL aggregates usage by calendar month rather than by time of day.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_date** | **date** | The first day of the following calendar month in UTC. This date is not included in the period. | [optional] 
**firewall_bytes** | **int** | Firewall event data ingested during this month, in bytes. | [optional] 
**sdwan_bytes** | **int** | SD-WAN event data ingested during this month, in bytes. | [optional] 
**start_date** | **date** | The first day of the calendar month in UTC. This date is included in the period. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_period_monthly import MspSalUsagePeriodMonthly

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsagePeriodMonthly from a JSON string
msp_sal_usage_period_monthly_instance = MspSalUsagePeriodMonthly.from_json(json)
# print the JSON string representation of the object
print(MspSalUsagePeriodMonthly.to_json())

# convert the object into a dict
msp_sal_usage_period_monthly_dict = msp_sal_usage_period_monthly_instance.to_dict()
# create an instance of MspSalUsagePeriodMonthly from a dict
msp_sal_usage_period_monthly_form_dict = msp_sal_usage_period_monthly.from_dict(msp_sal_usage_period_monthly_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


