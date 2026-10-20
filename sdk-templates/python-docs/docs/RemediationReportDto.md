# RemediationReportDto

NAT policy remediation report data.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completed_at** | **datetime** | The time the remediation finished, in UTC (RFC 3339). | [optional] 
**remediation_job_uid** | **str** | The unique identifier of the NAT policy remediation job. | [optional] 
**report_url** | **str** | Presigned download URL for the NAT policy remediation report, valid for 30 minutes after generation. Get the report again to obtain a new URL. Null if no PDF report is available. | [optional] 
**started_at** | **datetime** | The time the remediation started, in UTC (RFC 3339). Omitted if unavailable. | [optional] 
**status** | **str** | The final status of the NAT policy remediation. | [optional] 
**triggered_by** | **str** | The username of the user who started the NAT policy remediation. | [optional] 
**version** | **int** | The version of the remediation report for the NAT policy analysis. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.remediation_report_dto import RemediationReportDto

# TODO update the JSON string below
json = "{}"
# create an instance of RemediationReportDto from a JSON string
remediation_report_dto_instance = RemediationReportDto.from_json(json)
# print the JSON string representation of the object
print(RemediationReportDto.to_json())

# convert the object into a dict
remediation_report_dto_dict = remediation_report_dto_instance.to_dict()
# create an instance of RemediationReportDto from a dict
remediation_report_dto_form_dict = remediation_report_dto.from_dict(remediation_report_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


