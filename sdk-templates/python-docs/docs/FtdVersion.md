# FtdVersion


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**filename** | **str** | The name of the file | [optional] 
**is_suggested_version** | **bool** | A boolean value, indicating whether this version is a suggested version to upgrade to. | [optional] 
**software_version** | **str** | The version to which this package will upgrade the FTD. | [optional] 
**upgrade_package_uid** | **str** | The unique identifier, represented as a UUID, of the version in Security Cloud Control | [optional] 
**upgrade_type** | **str** | The type of the upgrade | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_version import FtdVersion

# TODO update the JSON string below
json = "{}"
# create an instance of FtdVersion from a JSON string
ftd_version_instance = FtdVersion.from_json(json)
# print the JSON string representation of the object
print(FtdVersion.to_json())

# convert the object into a dict
ftd_version_dict = ftd_version_instance.to_dict()
# create an instance of FtdVersion from a dict
ftd_version_form_dict = ftd_version.from_dict(ftd_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


