# AllowListSource

The five named CDO sources.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**available** | **bool** | Whether the source was successfully resolved for this request. | [optional] 
**fqdns** | **List[str]** | Authoritative hostnames for the source, if any. | [optional] 
**ipv4_addresses** | **List[str]** | Resolved IPv4 addresses. | [optional] 
**ipv4_cidrs** | **List[str]** | IPv4 CIDRs (always empty for CDO sources). | [optional] 
**ipv6_addresses** | **List[str]** | Resolved IPv6 addresses. | [optional] 
**ipv6_cidrs** | **List[str]** | IPv6 CIDRs (always empty for CDO sources). | [optional] 
**name** | **str** | The named CDO source. | [optional] 
**scope** | **str** | Scope at which the source&#39;s addresses apply. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.allow_list_source import AllowListSource

# TODO update the JSON string below
json = "{}"
# create an instance of AllowListSource from a JSON string
allow_list_source_instance = AllowListSource.from_json(json)
# print the JSON string representation of the object
print(AllowListSource.to_json())

# convert the object into a dict
allow_list_source_dict = allow_list_source_instance.to_dict()
# create an instance of AllowListSource from a dict
allow_list_source_form_dict = allow_list_source.from_dict(allow_list_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


