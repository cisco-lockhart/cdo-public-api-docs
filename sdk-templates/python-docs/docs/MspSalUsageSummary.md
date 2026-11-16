# MspSalUsageSummary

Aggregated event usage for the requested time range.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average** | [**MspSalUsageBytes**](MspSalUsageBytes.md) |  | [optional] 
**total** | [**MspSalUsageBytes**](MspSalUsageBytes.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_summary import MspSalUsageSummary

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsageSummary from a JSON string
msp_sal_usage_summary_instance = MspSalUsageSummary.from_json(json)
# print the JSON string representation of the object
print(MspSalUsageSummary.to_json())

# convert the object into a dict
msp_sal_usage_summary_dict = msp_sal_usage_summary_instance.to_dict()
# create an instance of MspSalUsageSummary from a dict
msp_sal_usage_summary_form_dict = msp_sal_usage_summary.from_dict(msp_sal_usage_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


