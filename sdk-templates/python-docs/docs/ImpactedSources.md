# ImpactedSources


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices_count** | **int** |  | [optional] 
**templates_count** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.impacted_sources import ImpactedSources

# TODO update the JSON string below
json = "{}"
# create an instance of ImpactedSources from a JSON string
impacted_sources_instance = ImpactedSources.from_json(json)
# print the JSON string representation of the object
print(ImpactedSources.to_json())

# convert the object into a dict
impacted_sources_dict = impacted_sources_instance.to_dict()
# create an instance of ImpactedSources from a dict
impacted_sources_form_dict = impacted_sources.from_dict(impacted_sources_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


