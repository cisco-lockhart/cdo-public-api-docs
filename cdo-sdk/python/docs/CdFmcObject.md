# CdFmcObject

Policies that match the search term.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the entity in Cloud-delivered FMC. | [optional] 
**name** | **str** | The name of the entity in Cloud-delivered FMC. | [optional] 
**link** | **str** | A URL to access the entity in Cloud-delivered FMC. | [optional] 

## Example

```python
from cdo_python_sdk.models.cd_fmc_object import CdFmcObject

# TODO update the JSON string below
json = "{}"
# create an instance of CdFmcObject from a JSON string
cd_fmc_object_instance = CdFmcObject.from_json(json)
# print the JSON string representation of the object
print(CdFmcObject.to_json())

# convert the object into a dict
cd_fmc_object_dict = cd_fmc_object_instance.to_dict()
# create an instance of CdFmcObject from a dict
cd_fmc_object_form_dict = cd_fmc_object.from_dict(cd_fmc_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


