# Device


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | The address of the device, in &#x60;host:port&#x60; format. Security Cloud Control connects to the device at this address. | [optional] 
**asa_failover_mate** | [**AsaFailoverMate**](AsaFailoverMate.md) |  | [optional] 
**asa_failover_mode** | [**AsaFailoverMode**](AsaFailoverMode.md) |  | [optional] 
**asa_failover_state** | **str** | (High Availability ASAs only) Failover state of this device. | [optional] 
**asa_license_entitlements** | **Dict[str, str]** | (ASAs only) Map of ASA License entitlements. | [optional] 
**asdm_version** | **str** | (ASAs only) Version of the ASDM device manager running on the device. | [optional] 
**cd_fmc_info** | [**CdFmcInfo**](CdFmcInfo.md) |  | [optional] 
**certificate_expiry_date** | **datetime** | (ASAs and FDM-managed FTDs only) Expiration date of the certificate used on the management interface of the device. | [optional] 
**chassis_serial** | **str** | The serial number on the chassis of the device (ASA-only). This is typically used to type up to Cisco SmartNet, and is not the same as the serial number. | [optional] 
**config_state** | [**ConfigState**](ConfigState.md) |  | [optional] 
**conflict_detection_state** | [**ConflictDetectionState**](ConflictDetectionState.md) |  | [optional] 
**connectivity_state** | [**ConnectivityState**](ConnectivityState.md) |  | [optional] 
**connector_type** | [**ConnectorType**](ConnectorType.md) |  | [optional] 
**connector_uid** | **str** | The unique identifier, represented as a UUID, of the Secure Device Connector (SDC) that will be used to communicate with the device. This value is not required if the connector type selected is Cloud Connector (CDG)The name of the Secure Device Connector (SDC) that will be used to communicate with the device. This value is not required if the connector type selected is Cloud Connector (CDG). | [optional] 
**device_maintenance_window** | [**DeviceMaintenanceWindow**](DeviceMaintenanceWindow.md) |  | [optional] 
**device_metadata** | [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**device_record_on_fmc** | [**FmcDeviceRecord**](FmcDeviceRecord.md) |  | [optional] 
**device_role** | [**DeviceRole**](DeviceRole.md) |  | [optional] 
**device_roles** | [**List[DeviceRole]**](DeviceRole.md) | The roles that this device performs on the network. | [optional] 
**device_type** | [**EntityType**](EntityType.md) |  | 
**fmc_access_policy** | [**FmcAccessPolicyReference**](FmcAccessPolicyReference.md) |  | [optional] 
**fmc_domain_uid** | **str** | (FMC device managers only) The unique identifier, represented as a UUID, of the [FMC domain](https://www.cisco.com/c/en/us/td/docs/security/secure-firewall/management-center/admin/740/management-center-admin-74/system-domains.html). | [optional] 
**ftd_cluster_info** | [**FtdClusterInfo**](FtdClusterInfo.md) |  | [optional] 
**ftd_ha_info** | [**FtdHaInfo**](FtdHaInfo.md) |  | [optional] 
**ftd_licenses** | **List[str]** | (FTDs only) List of FTD License entitlements. | [optional] 
**ftd_performance_tier** | **str** | (FTDvs only) The FTDv supports performance-tiered licensing that provides different throughput levels and VPN connection limits based on deployment requirements. This field specifies the performance tier of the FTD. | [optional] 
**hardware_model** | **str** | (ASAs, FDM-managed FTDs, and FMC-managed FTDs only) The hardware model of the device | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**meraki_deployment_mode** | [**MerakiDeploymentMode**](MerakiDeploymentMode.md) |  | [optional] 
**meraki_network** | [**Network**](Network.md) |  | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on (ASA-only). This field can be missing in the case of a partially onboarded device. | [optional] 
**name** | **str** | The name of the device. Device names are unique in Security Cloud Control. | 
**notes** | **str** | Free-form notes on the device. | [optional] 
**on_prem_fmc_info** | [**OnPremFmcInfo**](OnPremFmcInfo.md) |  | [optional] 
**opted_in_to_asa_health_metrics** | **bool** | Indicates whether the device has been opted in to collect ASA health metrics (SDC-managed ASAs only). | [optional] 
**ra_vpn_certificate_expiry_date** | **datetime** | (Remote Access VPN headends — ASA, FDM-managed FTD, and cdFMC-managed FTD — only) Expiration date of the Secure Client certificate installed on the device. | [optional] 
**redundancy_mode** | **str** | The redundancy mode this firewall is running in. Note: for High Availability pairs and clusters, Security Cloud Control represents all of the devices as part of one record. | [optional] 
**serial** | **str** | The serial number of the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**snort_version** | **str** | (FTDs only) List of FTD License entitlements. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**sse_device_data** | [**SseDeviceData**](SseDeviceData.md) |  | [optional] 
**state** | **str** | The device state. | [optional] 
**state_machine_details** | [**StateMachineDetails**](StateMachineDetails.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 
**uid_on_fmc** | **str** | (FMC-managed FTDs only) The unique identifier, represented as a UUID, of the device on a cdFMC. This field is deprecated. Please see &#x60;deviceRecordOnFmc&#x60;. | [optional] 
**universal_ztna_settings** | [**UniversalZtnaSettings**](UniversalZtnaSettings.md) |  | [optional] 
**ztp_onboarding_job_id** | **str** | The unique identifier, represented as a UUID, for an internal job that orchestrates the onboarding of devices through Zero-Touch Provisioning. This applies to devices managed by both on-premises and Cloud-delivered FMC. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device import Device

# TODO update the JSON string below
json = "{}"
# create an instance of Device from a JSON string
device_instance = Device.from_json(json)
# print the JSON string representation of the object
print(Device.to_json())

# convert the object into a dict
device_dict = device_instance.to_dict()
# create an instance of Device from a dict
device_form_dict = device.from_dict(device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


