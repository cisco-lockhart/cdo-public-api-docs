# CdFmcResult

Results from the Cloud-delivered FMC devices, objects or policies that match the search term.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices** | [**List[CdFmcObject]**](CdFmcObject.md) | Devices that match the search term. | [optional] 
**objects** | [**List[CdFmcObject]**](CdFmcObject.md) | Objects that match the search term. | [optional] 
**policies** | [**List[CdFmcObject]**](CdFmcObject.md) | Policies that match the search term. | [optional] 

## Example

```python
from cdo_sdk_python.models.cd_fmc_result import CdFmcResult

# TODO update the JSON string below
json = "{}"
# create an instance of CdFmcResult from a JSON string
cd_fmc_result_instance = CdFmcResult.from_json(json)
# print the JSON string representation of the object
print(CdFmcResult.to_json())

# convert the object into a dict
cd_fmc_result_dict = cd_fmc_result_instance.to_dict()
# create an instance of CdFmcResult from a dict
cd_fmc_result_form_dict = cd_fmc_result.from_dict(cd_fmc_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


