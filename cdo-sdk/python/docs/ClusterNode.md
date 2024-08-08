# ClusterNode

Information on the data nodes, which are individual units within a cluster that process and forward network traffic based on policies and configurations managed by the control node.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**serial** | **str** | The serial number of the node on the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**uid_on_fmc** | **str** | The unique identifier of the device on a cdFMC. | [optional] 
**status** | **str** | The status of the cluster node. | [optional] 

## Example

```python
from cdo_sdk_python.models.cluster_node import ClusterNode

# TODO update the JSON string below
json = "{}"
# create an instance of ClusterNode from a JSON string
cluster_node_instance = ClusterNode.from_json(json)
# print the JSON string representation of the object
print(ClusterNode.to_json())

# convert the object into a dict
cluster_node_dict = cluster_node_instance.to_dict()
# create an instance of ClusterNode from a dict
cluster_node_form_dict = cluster_node.from_dict(cluster_node_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


