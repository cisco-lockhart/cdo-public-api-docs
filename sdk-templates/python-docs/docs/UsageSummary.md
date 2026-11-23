# UsageSummary

Totals and per-period averages for the requested time range.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average** | [**UsageBytes**](UsageBytes.md) |  | [optional] 
**total** | [**UsageBytes**](UsageBytes.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.usage_summary import UsageSummary

# TODO update the JSON string below
json = "{}"
# create an instance of UsageSummary from a JSON string
usage_summary_instance = UsageSummary.from_json(json)
# print the JSON string representation of the object
print(UsageSummary.to_json())

# convert the object into a dict
usage_summary_dict = usage_summary_instance.to_dict()
# create an instance of UsageSummary from a dict
usage_summary_form_dict = usage_summary.from_dict(usage_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


