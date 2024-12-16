# FtdHaInfo

(High Availability Devices managed by FMC only) High-Available information information. Note: Security Cloud Control represents all of the nodes on an FTD cluster in a single device record with the UID of the cluster control node.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ha_pair_uid** | **str** | The unique identifier, represented as a UUID, of the HA Pair, on the FMC | [optional] 
**primary_node** | [**HaNode**](HaNode.md) |  | [optional] 
**secondary_node** | [**HaNode**](HaNode.md) |  | [optional] 
**ha_node_type** | **str** | (on-prem FMC-managed FTDs only) Information on the type of this node in the HA Pair. Note: Each node in an on-prem-FMC-managed FTD HA Pair is represented as a separate device entry in the API. | [optional] 
**current_role** | **str** | (on-prem FMC-managed FTDs only) Information on the current role of the node in the HA Pair. Note: Each node in an on-prem-FMC-managed FTD HA Pair is represented as a separate device entry in the API. | [optional] 

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


