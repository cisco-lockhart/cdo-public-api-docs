# FtdHaInfo

(High Availability Devices managed by FMC only) High-Available information information. Note: Security Cloud Control represents all of the nodes on an FTD cluster in a single device record with the UID of the cluster control node.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**primary_node** | [**HaNode**](HaNode.md) |  | [optional] 
**secondary_node** | [**HaNode**](HaNode.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.ftd_ha_info import FtdHaInfo

# TODO update the JSON string below
json = "{}"
# create an instance of FtdHaInfo from a JSON string
ftd_ha_info_instance = FtdHaInfo.from_json(json)
# print the JSON string representation of the object
print(FtdHaInfo.to_json())

# convert the object into a dict
ftd_ha_info_dict = ftd_ha_info_instance.to_dict()
# create an instance of FtdHaInfo from a dict
ftd_ha_info_form_dict = ftd_ha_info.from_dict(ftd_ha_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


