# MspManagedCloudServiceDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | The address of the cloud service (if applicable), in &lt;code&gt;host:port&lt;/code&gt; format. Security Cloud Control connects to the cloud service at this address. | [optional] 
**cloud_service_type** | [**EntityType**](EntityType.md) |  | [optional] 
**config_state** | [**ConfigState**](ConfigState.md) |  | [optional] 
**conflict_detection_state** | [**ConflictDetectionState**](ConflictDetectionState.md) |  | [optional] 
**connectivity_state** | [**ConnectivityState**](ConnectivityState.md) |  | [optional] 
**managed_scc_organization_uid** | **str** | The unique identifier, represented as a UUID, of the managed organization on Security Cloud Control that this cloud service belongs to. | [optional] 
**managed_tenant_display_name** | **str** | The display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant in CDO. | [optional] 
**managed_tenant_region** | **str** | The region of the managed tenant in CDO. This is the region where the cloud service is located. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control Firewall Manager that this cloud service belongs to. | [optional] 
**name** | **str** | The name of the cloud service in CDO. Cloud service names are unique in Security Cloud Control. | 
**tenant_uid** | **str** | The unique identifier, represented as a UUID, of the MSP portal that manages this cloud service. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_cloud_service_dto import MspManagedCloudServiceDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedCloudServiceDto from a JSON string
msp_managed_cloud_service_dto_instance = MspManagedCloudServiceDto.from_json(json)
# print the JSON string representation of the object
print(MspManagedCloudServiceDto.to_json())

# convert the object into a dict
msp_managed_cloud_service_dto_dict = msp_managed_cloud_service_dto_instance.to_dict()
# create an instance of MspManagedCloudServiceDto from a dict
msp_managed_cloud_service_dto_form_dict = msp_managed_cloud_service_dto.from_dict(msp_managed_cloud_service_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


