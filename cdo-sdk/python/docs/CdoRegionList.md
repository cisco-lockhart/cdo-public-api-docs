# CdoRegionList


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[CdoRegion]**](CdoRegion.md) | The list of items. | [optional] 

## Example

```python
from openapi_client.models.cdo_region_list import CdoRegionList

# TODO update the JSON string below
json = "{}"
# create an instance of CdoRegionList from a JSON string
cdo_region_list_instance = CdoRegionList.from_json(json)
# print the JSON string representation of the object
print(CdoRegionList.to_json())

# convert the object into a dict
cdo_region_list_dict = cdo_region_list_instance.to_dict()
# create an instance of CdoRegionList from a dict
cdo_region_list_form_dict = cdo_region_list.from_dict(cdo_region_list_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


