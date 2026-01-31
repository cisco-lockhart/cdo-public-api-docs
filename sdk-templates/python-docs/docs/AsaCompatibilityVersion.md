# AsaCompatibilityVersion


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_asa_versions** | [**List[AsaCompatibleVersion]**](AsaCompatibleVersion.md) |  | [optional] 
**device_uids** | **List[str]** |  | [optional] 
**hardware_model** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_compatibility_version import AsaCompatibilityVersion

# TODO update the JSON string below
json = "{}"
# create an instance of AsaCompatibilityVersion from a JSON string
asa_compatibility_version_instance = AsaCompatibilityVersion.from_json(json)
# print the JSON string representation of the object
print(AsaCompatibilityVersion.to_json())

# convert the object into a dict
asa_compatibility_version_dict = asa_compatibility_version_instance.to_dict()
# create an instance of AsaCompatibilityVersion from a dict
asa_compatibility_version_form_dict = asa_compatibility_version.from_dict(asa_compatibility_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


