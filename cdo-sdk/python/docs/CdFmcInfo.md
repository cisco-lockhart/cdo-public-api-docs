# CdFmcInfo

(FTDs managed by cdFMC only) Information on the cloud-delivered FMC managing this FTD. This information is not available for FTDs managed using FDM or on-prem FMCs.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cli_key** | **str** | The CLI key to paste into the FTD CLI to register the FTD with a cdFMC. You need to paste this value in only once, when the FTD is being onboarded. Refer to the [CDO Documentation](https://www.cisco.com/c/en/us/td/docs/security/cdo/cloud-delivered-firewall-management-center-in-cdo/managing-firewall-threat-defense-services-with-cisco-defense-orchestrator/m-onboard-for-ftd-management.html) for details. | [optional] 
**reg_key** | **str** | The Network Address Translation (NAT) ID of this FTD. Refer to the [CDO Documentation](https://www.cisco.com/c/en/us/td/docs/security/cdo/cloud-delivered-firewall-management-center-in-cdo/managing-firewall-threat-defense-services-with-cisco-defense-orchestrator/m-onboard-for-ftd-management.html) for details. | [optional] 
**nat_id** | **str** | The Registration Key of this FTD. Refer to the [CDO Documentation](https://www.cisco.com/c/en/us/td/docs/security/cdo/cloud-delivered-firewall-management-center-in-cdo/managing-firewall-threat-defense-services-with-cisco-defense-orchestrator/m-onboard-for-ftd-management.html) for details. | [optional] 

## Example

```python
from openapi_client.models.cd_fmc_info import CdFmcInfo

# TODO update the JSON string below
json = "{}"
# create an instance of CdFmcInfo from a JSON string
cd_fmc_info_instance = CdFmcInfo.from_json(json)
# print the JSON string representation of the object
print(CdFmcInfo.to_json())

# convert the object into a dict
cd_fmc_info_dict = cd_fmc_info_instance.to_dict()
# create an instance of CdFmcInfo from a dict
cd_fmc_info_form_dict = cd_fmc_info.from_dict(cd_fmc_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


