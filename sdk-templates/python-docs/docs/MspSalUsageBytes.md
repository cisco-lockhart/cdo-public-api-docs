# MspSalUsageBytes

Average event data ingested per returned period: per calendar month for monthly usage, where the current month uses its month-to-date value, or per completed UTC day for daily usage.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**firewall_bytes** | **int** | Firewall event data ingested, in bytes. | [optional] 
**sdwan_bytes** | **int** | SD-WAN event data ingested, in bytes. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_sal_usage_bytes import MspSalUsageBytes

# TODO update the JSON string below
json = "{}"
# create an instance of MspSalUsageBytes from a JSON string
msp_sal_usage_bytes_instance = MspSalUsageBytes.from_json(json)
# print the JSON string representation of the object
print(MspSalUsageBytes.to_json())

# convert the object into a dict
msp_sal_usage_bytes_dict = msp_sal_usage_bytes_instance.to_dict()
# create an instance of MspSalUsageBytes from a dict
msp_sal_usage_bytes_form_dict = msp_sal_usage_bytes.from_dict(msp_sal_usage_bytes_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


