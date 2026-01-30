# MspManagedDevice


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | The address of the device, in &lt;code&gt;host:port&lt;/code&gt; format. Security Cloud Control connects to the device at this address. | [optional] 
**asa_failover_mate** | [**AsaFailoverMate**](AsaFailoverMate.md) |  | [optional] 
**asa_failover_mode** | [**AsaFailoverMode**](AsaFailoverMode.md) |  | [optional] 
**asa_failover_state** | **str** | (High Availability ASAs only) Failover state of this device. | [optional] 
**certificate_expiry_date** | **datetime** | (ASAs and FDM-managed FTDs only) Expiration date of the certificate used on the management interface of the device. | [optional] 
**chassis_serial** | **str** | The serial number on the chassis of the device (ASA-only). This is typically used to type up to Cisco SmartNet, and is not the same as the serial number. | [optional] 
**config_state** | [**ConfigState**](ConfigState.md) |  | [optional] 
**conflict_detection_state** | [**ConflictDetectionState**](ConflictDetectionState.md) |  | [optional] 
**connectivity_state** | [**ConnectivityState**](ConnectivityState.md) |  | [optional] 
**device_maintenance_window** | [**DeviceMaintenanceWindow**](DeviceMaintenanceWindow.md) |  | [optional] 
**device_record_on_fmc** | [**FmcDeviceRecord**](FmcDeviceRecord.md) |  | [optional] 
**device_type** | [**EntityType**](EntityType.md) |  | [optional] 
**ftd_cluster_info** | [**FtdClusterInfo**](FtdClusterInfo.md) |  | [optional] 
**ftd_ha_info** | [**FtdHaInfo**](FtdHaInfo.md) |  | [optional] 
**hardware_model** | **str** | (ASAs, FDM-managed FTDs, and FMC-managed FTDs only) The hardware model of the device | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**managed_tenant_display_name** | **str** | The display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant in CDO. | [optional] 
**managed_tenant_region** | **str** | The region of the managed tenant in CDO. This is the region where the device is located. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this device belongs to. | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on (ASA-only). This field can be missing in the case of a partially onboarded device. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in Security Cloud Control. | 
**ra_vpn_certificate_expiry_date** | **datetime** | (Remote Access VPN headends — ASA, FDM-managed FTD, and cdFMC-managed FTD — only) Expiration date of the Secure Client certificate installed on the device. | [optional] 
**redundancy_mode** | **str** | The redundancy mode this firewall is running in. Note: for High Availability pairs and clusters, Security Cloud Control represents all of the devices as part of one record. | [optional] 
**serial** | **str** | The serial number of the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**tenant_uid** | **str** |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_device import MspManagedDevice

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedDevice from a JSON string
msp_managed_device_instance = MspManagedDevice.from_json(json)
# print the JSON string representation of the object
print(MspManagedDevice.to_json())

# convert the object into a dict
msp_managed_device_dict = msp_managed_device_instance.to_dict()
# create an instance of MspManagedDevice from a dict
msp_managed_device_form_dict = msp_managed_device.from_dict(msp_managed_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


