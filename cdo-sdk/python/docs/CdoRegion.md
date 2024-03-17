# CdoRegion

The list of items.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | The domain for the region. | [optional] 
**api_domain** | **str** | The domain for API services in the region. | [optional] 
**description** | **str** | Human readable description of the region. | [optional] 

## Example

```python
from cdo_sdk_python.models.cdo_region import CdoRegion

# TODO update the JSON string below
json = "{}"
# create an instance of CdoRegion from a JSON string
cdo_region_instance = CdoRegion.from_json(json)
# print the JSON string representation of the object
print(CdoRegion.to_json())

# convert the object into a dict
cdo_region_dict = cdo_region_instance.to_dict()
# create an instance of CdoRegion from a dict
cdo_region_form_dict = cdo_region.from_dict(cdo_region_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


