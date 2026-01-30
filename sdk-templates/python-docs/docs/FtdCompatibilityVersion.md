# FtdCompatibilityVersion


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_ftd_versions** | [**List[FtdVersion]**](FtdVersion.md) |  | [optional] 
**device_uids** | **List[str]** |  | [optional] 
**hardware_model** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_compatibility_version import FtdCompatibilityVersion

# TODO update the JSON string below
json = "{}"
# create an instance of FtdCompatibilityVersion from a JSON string
ftd_compatibility_version_instance = FtdCompatibilityVersion.from_json(json)
# print the JSON string representation of the object
print(FtdCompatibilityVersion.to_json())

# convert the object into a dict
ftd_compatibility_version_dict = ftd_compatibility_version_instance.to_dict()
# create an instance of FtdCompatibilityVersion from a dict
ftd_compatibility_version_form_dict = ftd_compatibility_version.from_dict(ftd_compatibility_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


