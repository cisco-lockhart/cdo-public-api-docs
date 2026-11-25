# MspSalRatePeriodMonthly

SAL event rate for one UTC calendar month.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average_events_per_second** | **float** | The average number of events per second sent by the managed tenant during this month. This value is null for the current month until its first UTC day is over. | 
**end_date** | **date** | The first day of the following calendar month in UTC. This date is not included in the period. | [optional] 
**start_date** | **date** | The first day of the calendar month in UTC. This date is included in the period. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_rate_period_monthly import MspSalRatePeriodMonthly

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalRatePeriodMonthly from a JSON string
msp_sal_rate_period_monthly_instance = MspSalRatePeriodMonthly.from_json(json)
# print the JSON string representation of the object
print(MspSalRatePeriodMonthly.to_json())

# convert the object into a dict
msp_sal_rate_period_monthly_dict = msp_sal_rate_period_monthly_instance.to_dict()
# create an instance of MspSalRatePeriodMonthly from a dict
msp_sal_rate_period_monthly_form_dict = msp_sal_rate_period_monthly.from_dict(msp_sal_rate_period_monthly_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


