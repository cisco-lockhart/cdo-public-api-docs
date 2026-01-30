# AsaCompatibleVersion


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_image_url** | **str** |  | [optional] 
**asdm_version** | **str** |  | [optional] 
**software_image_url** | **str** |  | [optional] 
**software_version** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_compatible_version import AsaCompatibleVersion

# TODO update the JSON string below
json = "{}"
# create an instance of AsaCompatibleVersion from a JSON string
asa_compatible_version_instance = AsaCompatibleVersion.from_json(json)
# print the JSON string representation of the object
print(AsaCompatibleVersion.to_json())

# convert the object into a dict
asa_compatible_version_dict = asa_compatible_version_instance.to_dict()
# create an instance of AsaCompatibleVersion from a dict
asa_compatible_version_form_dict = asa_compatible_version.from_dict(asa_compatible_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


