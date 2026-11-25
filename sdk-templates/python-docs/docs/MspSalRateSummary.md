# MspSalRateSummary

Average and peak event rate for the requested time range.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average_events_per_second** | **float** | The mean of the stored average events per second in the requested time range. A null monthly rate is not counted; this value is 0 when there are no stored rates. | [optional] 
**peak_events_per_second** | **float** | The highest stored average events per second in the requested time range. This value is 0 when there are no stored rates. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_rate_summary import MspSalRateSummary

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalRateSummary from a JSON string
msp_sal_rate_summary_instance = MspSalRateSummary.from_json(json)
# print the JSON string representation of the object
print(MspSalRateSummary.to_json())

# convert the object into a dict
msp_sal_rate_summary_dict = msp_sal_rate_summary_instance.to_dict()
# create an instance of MspSalRateSummary from a dict
msp_sal_rate_summary_form_dict = msp_sal_rate_summary.from_dict(msp_sal_rate_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


