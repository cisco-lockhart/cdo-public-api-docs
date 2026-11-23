# UsagePeriod

Event usage for one UTC calendar month or day. Dates are used because SAL aggregates usage by day rather than by time of day.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end** | **date** | The first day after the period in UTC. This date is not included in the period. | [optional] 
**firewall_bytes** | **int** | Firewall event data ingested during this period, in bytes. | [optional] 
**sdwan_bytes** | **int** | SD-WAN event data ingested during this period, in bytes. | [optional] 
**start** | **date** | The first day of the period in UTC. This date is included in the period. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.usage_period import UsagePeriod

# TODO update the JSON string below
json = "{}"
# create an instance of UsagePeriod from a JSON string
usage_period_instance = UsagePeriod.from_json(json)
# print the JSON string representation of the object
print(UsagePeriod.to_json())

# convert the object into a dict
usage_period_dict = usage_period_instance.to_dict()
# create an instance of UsagePeriod from a dict
usage_period_form_dict = usage_period.from_dict(usage_period_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


