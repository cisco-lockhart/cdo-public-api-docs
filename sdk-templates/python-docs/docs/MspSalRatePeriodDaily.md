# MspSalRatePeriodDaily

SAL event rate for one UTC day.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average_events_per_second** | **float** | The average number of events per second sent by the managed tenant during this day. | [optional] 
**end_date** | **date** | The following UTC day. This date is not included in the period. | [optional] 
**start_date** | **date** | The UTC day. This date is included in the period. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_rate_period_daily import MspSalRatePeriodDaily

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalRatePeriodDaily from a JSON string
msp_sal_rate_period_daily_instance = MspSalRatePeriodDaily.from_json(json)
# print the JSON string representation of the object
print(MspSalRatePeriodDaily.to_json())

# convert the object into a dict
msp_sal_rate_period_daily_dict = msp_sal_rate_period_daily_instance.to_dict()
# create an instance of MspSalRatePeriodDaily from a dict
msp_sal_rate_period_daily_form_dict = msp_sal_rate_period_daily.from_dict(msp_sal_rate_period_daily_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


