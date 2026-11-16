# EventUsageResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**summary** | [**UsageSummary**](UsageSummary.md) |  | [optional] 
**usage** | [**List[UsagePeriod]**](UsagePeriod.md) | UTC calendar-month usage buckets, ordered from oldest to newest. Each bucket has an inclusive start date and an exclusive end date. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.event_usage_response import EventUsageResponse

# TODO update the JSON string below
json = "{}"
# create an instance of EventUsageResponse from a JSON string
event_usage_response_instance = EventUsageResponse.from_json(json)
# print the JSON string representation of the object
print(EventUsageResponse.to_json())

# convert the object into a dict
event_usage_response_dict = event_usage_response_instance.to_dict()
# create an instance of EventUsageResponse from a dict
event_usage_response_form_dict = event_usage_response.from_dict(event_usage_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


