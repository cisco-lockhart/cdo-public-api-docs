# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from cdo_sdk_python.models.labels import Labels
from typing import Optional, Set
from typing_extensions import Self

class FtdCreateOrUpdateInput(BaseModel):
    """
    FtdCreateOrUpdateInput
    """ # noqa: E501
    name: StrictStr = Field(description="Specify a human-readable name for the device.")
    licenses: List[StrictStr] = Field(description="Specify a set of licenses to apply to the device.")
    virtual: Optional[StrictBool] = Field(default=None, description="Indicate whether the FTD is a virtual or a physical device.")
    fmc_access_policy_uid: StrictStr = Field(description="Specify the unique identifier of the FMC access policy to apply to this device.", alias="fmcAccessPolicyUid")
    performance_tier: Optional[StrictStr] = Field(default=None, description="Specify the performance tier of the FTDv (required only if isVirtual is set to true)", alias="performanceTier")
    labels: Optional[Labels] = None
    device_type: StrictStr = Field(description="Specify the type of the FTD. The only supported type of FTD is CDFMC_MANAGED_FTD", alias="deviceType")
    __properties: ClassVar[List[str]] = ["name", "licenses", "virtual", "fmcAccessPolicyUid", "performanceTier", "labels", "deviceType"]

    @field_validator('licenses')
    def licenses_validate_enum(cls, value):
        """Validates the enum"""
        for i in value:
            if i not in set(['BASE', 'CARRIER', 'THREAT', 'MALWARE', 'URLFilter']):
                raise ValueError("each list item must be one of ('BASE', 'CARRIER', 'THREAT', 'MALWARE', 'URLFilter')")
        return value

    @field_validator('performance_tier')
    def performance_tier_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['FTDv5', 'FTDv10', 'FTDv20', 'FTDv30', 'FTDv50', 'FTDv100', 'FTDv']):
            raise ValueError("must be one of enum values ('FTDv5', 'FTDv10', 'FTDv20', 'FTDv30', 'FTDv50', 'FTDv100', 'FTDv')")
        return value

    @field_validator('device_type')
    def device_type_validate_enum(cls, value):
        """Validates the enum"""
        if value not in set(['CDFMC_MANAGED_FTD']):
            raise ValueError("must be one of enum values ('CDFMC_MANAGED_FTD')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of FtdCreateOrUpdateInput from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of labels
        if self.labels:
            _dict['labels'] = self.labels.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of FtdCreateOrUpdateInput from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "name": obj.get("name"),
            "licenses": obj.get("licenses"),
            "virtual": obj.get("virtual"),
            "fmcAccessPolicyUid": obj.get("fmcAccessPolicyUid"),
            "performanceTier": obj.get("performanceTier"),
            "labels": Labels.from_dict(obj["labels"]) if obj.get("labels") is not None else None,
            "deviceType": obj.get("deviceType")
        })
        return _obj

