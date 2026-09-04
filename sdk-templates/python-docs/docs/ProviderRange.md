# ProviderRange

Ranges grouped by service and region.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ipv4_cidrs** | **List[str]** | IPv4 CIDRs for this service/region. | [optional] 
**ipv6_cidrs** | **List[str]** | IPv6 CIDRs for this service/region. | [optional] 
**region** | **str** | AWS region, or GLOBAL. | [optional] 
**service** | **str** | AWS service the ranges belong to. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.provider_range import ProviderRange

# TODO update the JSON string below
json = "{}"
# create an instance of ProviderRange from a JSON string
provider_range_instance = ProviderRange.from_json(json)
# print the JSON string representation of the object
print(ProviderRange.to_json())

# convert the object into a dict
provider_range_dict = provider_range_instance.to_dict()
# create an instance of ProviderRange from a dict
provider_range_form_dict = provider_range.from_dict(provider_range_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


