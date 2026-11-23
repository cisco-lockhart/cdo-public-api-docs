# UsageBytes

Average event data ingested per returned calendar month or day. The current month uses its month-to-date value.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**firewall_bytes** | **int** | Firewall event data ingested, in bytes. | [optional] 
**sdwan_bytes** | **int** | SD-WAN event data ingested, in bytes. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.usage_bytes import UsageBytes

# TODO update the JSON string below
json = "{}"
# create an instance of UsageBytes from a JSON string
usage_bytes_instance = UsageBytes.from_json(json)
# print the JSON string representation of the object
print(UsageBytes.to_json())

# convert the object into a dict
usage_bytes_dict = usage_bytes_instance.to_dict()
# create an instance of UsageBytes from a dict
usage_bytes_form_dict = usage_bytes.from_dict(usage_bytes_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


