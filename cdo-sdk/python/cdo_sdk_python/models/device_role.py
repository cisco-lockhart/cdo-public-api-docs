# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.5.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
from enum import Enum
from typing_extensions import Self


class DeviceRole(str, Enum):
    """
    The role this device performs on the network.
    """

    """
    allowed enum values
    """
    RA_VPN_HEADEND = 'RA_VPN_HEADEND'
    MFA_PROVIDER = 'MFA_PROVIDER'
    ZTNA_PROVIDER = 'ZTNA_PROVIDER'
    ANYCONNECT_VPN_HEAD_END = 'ANYCONNECT_VPN_HEAD_END'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of DeviceRole from a JSON string"""
        return cls(json.loads(json_str))


