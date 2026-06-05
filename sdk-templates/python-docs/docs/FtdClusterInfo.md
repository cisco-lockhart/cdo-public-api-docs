# FtdClusterInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cluster_name** | **str** | The name of the cluster on the FMC. | [optional] 
**cluster_node_status** | **str** | (on-prem FMC-managed FTDs only) Information on the type of this node in the FTD cluster. Note: Each node in an on-prem-FMC-managed FTD cluster is represented as a separate device entry in the API. | [optional] 
**cluster_node_type** | **str** | (on-prem FMC-managed FTDs only) Information on the type of this node in the FTD cluster. Note: Each node in an on-prem-FMC-managed FTD cluster is represented as a separate device entry in the API. | [optional] 
**cluster_uid** | **str** | The unique identifier, represented as a UUID, of the cluster, on the FMC | [optional] 
**control_node** | [**ClusterNode**](ClusterNode.md) |  | [optional] 
**data_nodes** | [**List[ClusterNode]**](ClusterNode.md) | (cdFMC-managed FTDs only) Information on the data nodes, which are individual units within a cluster that process and forward network traffic based on policies and configurations managed by the control node. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_cluster_info import FtdClusterInfo

# TODO update the JSON string below
json = "{}"
# create an instance of FtdClusterInfo from a JSON string
ftd_cluster_info_instance = FtdClusterInfo.from_json(json)
# print the JSON string representation of the object
print(FtdClusterInfo.to_json())

# convert the object into a dict
ftd_cluster_info_dict = ftd_cluster_info_instance.to_dict()
# create an instance of FtdClusterInfo from a dict
ftd_cluster_info_form_dict = ftd_cluster_info.from_dict(ftd_cluster_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


