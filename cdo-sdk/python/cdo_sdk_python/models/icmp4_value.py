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

from pydantic import BaseModel, ConfigDict, Field, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class Icmp4Value(BaseModel):
    """
    Icmp4Value
    """ # noqa: E501
    icmp4_type: Optional[StrictStr] = Field(default=None, alias="icmp4Type")
    icmp4_code: Optional[StrictStr] = Field(default=None, alias="icmp4Code")
    __properties: ClassVar[List[str]] = ["icmp4Type", "icmp4Code"]

    @field_validator('icmp4_type')
    def icmp4_type_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['ANY', 'ECHO_REPLY', 'DESTINATION_UNREACHABLE', 'SOURCE_QUENCH', 'REDIRECT_MESSAGE', 'ALTERNATE_HOST_ADDRESS', 'ECHO_REQUEST', 'ROUTER_ADVERTISEMENT', 'ROUTER_SOLICITATION', 'TIME_EXCEEDED', 'PARAMETER_PROBLEM', 'TIMESTAMP', 'TIMESTAMP_REPLY', 'INFO_REQUEST', 'INFO_REPLY', 'ADDR_MASK_REQUEST', 'ADDR_MASK_REPLY', 'TRACEROUTE', 'DATAGRAM_CONVERSION_ERROR', 'MOBILE_HOST_REDIRECT', 'WHERE_ARE_YOU', 'HERE_I_AM', 'MOBILE_REG_REQUEST', 'MOBILE_REG_REPLY', 'DOMAIN_NAME_REQUEST', 'DOMAIN_NAME_REPLY', 'SKIP_ALGORITHM_DISCOVERY_PROTOCOL', 'PHOTURIS', 'EXPERIMENTAL_MOB_PROTOCOLS']):
            raise ValueError("must be one of enum values ('ANY', 'ECHO_REPLY', 'DESTINATION_UNREACHABLE', 'SOURCE_QUENCH', 'REDIRECT_MESSAGE', 'ALTERNATE_HOST_ADDRESS', 'ECHO_REQUEST', 'ROUTER_ADVERTISEMENT', 'ROUTER_SOLICITATION', 'TIME_EXCEEDED', 'PARAMETER_PROBLEM', 'TIMESTAMP', 'TIMESTAMP_REPLY', 'INFO_REQUEST', 'INFO_REPLY', 'ADDR_MASK_REQUEST', 'ADDR_MASK_REPLY', 'TRACEROUTE', 'DATAGRAM_CONVERSION_ERROR', 'MOBILE_HOST_REDIRECT', 'WHERE_ARE_YOU', 'HERE_I_AM', 'MOBILE_REG_REQUEST', 'MOBILE_REG_REPLY', 'DOMAIN_NAME_REQUEST', 'DOMAIN_NAME_REPLY', 'SKIP_ALGORITHM_DISCOVERY_PROTOCOL', 'PHOTURIS', 'EXPERIMENTAL_MOB_PROTOCOLS')")
        return value

    @field_validator('icmp4_code')
    def icmp4_code_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['NET_UNREACHABLE', 'HOST_UNREACHABLE', 'PROTOCOL_UNREACHABLE', 'PORT_UNREACHABLE', 'FRAGMENTATION_NEEDED', 'SOURCE_ROUTE_FAILED', 'DEST_NETWORK_UNKNOWN', 'DEST_HOST_UNKNOWN', 'SRC_HOST_ISOLATED', 'COMMUNICATION_DEST_NET_PROHIBITED', 'COMMUNICATION_DEST_HOST_PROHIBITED', 'DEST_NET_UNREACHABLE_FOR_TOS', 'DEST_HOST_UNREACHABLE_FOR_TOS', 'COMM_ADMINISTRATIVELY_PROHIBITED', 'HOST_PRECEDENCE_VIOLATION', 'PRECEDENCE_CUTOFF', 'REDIRECT_DATAGRAM_NETWORK', 'REDIRECT_DATAGRAM_HOST', 'REDIRECT_DATAGRAM_SERVICE_NETWORK', 'REDIRECT_DATAGRAM_SERVICE_HOST', 'ALTERNATE_HOST_ADDR', 'DO_NOT_ROUTE_COMMON_TRAFFIC', 'NORMAL_ROUTER_ADV', 'TTL_EXPIRED_TRANSIT', 'FRAG_ASSEMBLY', 'PTR_ERROR', 'MISSING_REQD_OPTION', 'BAD_LENGTH', 'BAD_SPI', 'AUTH_FAILED', 'DECOMPRESSION_FAILED', 'DECRYPTION_FAILED', 'NEED_AUTHENTICATION', 'NEED_AUTHORIZATION']):
            raise ValueError("must be one of enum values ('NET_UNREACHABLE', 'HOST_UNREACHABLE', 'PROTOCOL_UNREACHABLE', 'PORT_UNREACHABLE', 'FRAGMENTATION_NEEDED', 'SOURCE_ROUTE_FAILED', 'DEST_NETWORK_UNKNOWN', 'DEST_HOST_UNKNOWN', 'SRC_HOST_ISOLATED', 'COMMUNICATION_DEST_NET_PROHIBITED', 'COMMUNICATION_DEST_HOST_PROHIBITED', 'DEST_NET_UNREACHABLE_FOR_TOS', 'DEST_HOST_UNREACHABLE_FOR_TOS', 'COMM_ADMINISTRATIVELY_PROHIBITED', 'HOST_PRECEDENCE_VIOLATION', 'PRECEDENCE_CUTOFF', 'REDIRECT_DATAGRAM_NETWORK', 'REDIRECT_DATAGRAM_HOST', 'REDIRECT_DATAGRAM_SERVICE_NETWORK', 'REDIRECT_DATAGRAM_SERVICE_HOST', 'ALTERNATE_HOST_ADDR', 'DO_NOT_ROUTE_COMMON_TRAFFIC', 'NORMAL_ROUTER_ADV', 'TTL_EXPIRED_TRANSIT', 'FRAG_ASSEMBLY', 'PTR_ERROR', 'MISSING_REQD_OPTION', 'BAD_LENGTH', 'BAD_SPI', 'AUTH_FAILED', 'DECOMPRESSION_FAILED', 'DECRYPTION_FAILED', 'NEED_AUTHENTICATION', 'NEED_AUTHORIZATION')")
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
        """Create an instance of Icmp4Value from a JSON string"""
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
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Icmp4Value from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "icmp4Type": obj.get("icmp4Type"),
            "icmp4Code": obj.get("icmp4Code")
        })
        return _obj

