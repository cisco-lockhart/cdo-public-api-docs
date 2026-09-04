# ProviderRanges

AWS-managed shared CIDR ranges.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ranges** | [**List[ProviderRange]**](ProviderRange.md) | Ranges grouped by service and region. | [optional] 
**sync_token** | **str** | Version identifier of the underlying AWS ip-ranges document. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.provider_ranges import ProviderRanges

# TODO update the JSON string below
json = "{}"
# create an instance of ProviderRanges from a JSON string
provider_ranges_instance = ProviderRanges.from_json(json)
# print the JSON string representation of the object
print(ProviderRanges.to_json())

# convert the object into a dict
provider_ranges_dict = provider_ranges_instance.to_dict()
# create an instance of ProviderRanges from a dict
provider_ranges_form_dict = provider_ranges.from_dict(provider_ranges_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


