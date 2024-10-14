function get_fmc_openapi() {
    url="https://demo-red.app.us.cdo.cisco.com/api/api-explorer/fmc_oas3.json"
    filename=$1
    curl -s "$url" | \
    # Replace cdFMC URLs with the new Public API proxy URL
    jq '.paths |= with_entries(.key |= gsub("api/"; "v1/cdfmc/api/"))' | \
    jq 'walk(if type == "object" and .operationId == "getDevice" then .operationId = "getDeviceOnCdFmc" else . end)' | \
    jq 'walk(if type == "object" and .operationId == "deleteAccessRule" then .operationId = "deleteCdFmcAccessRule" else . end)' | \
    jq '
      .tags |= map(if .name then .name |= "cdFMC " + . else . end) |
      .paths |= with_entries(
        .value |= with_entries(
          .value |= if .tags then .tags |= map("cdFMC " + .) else . end
        )
      )
    ' | \
    jq 'walk(if type == "object" and ."$ref" == "#/components/schemas/AccessRule" then ."$ref" = "#/components/schemas/FmcAccessRule" else . end)' | \
    jq 'if .components.schemas.AccessRule then .components.schemas.FmcAccessRule = .components.schemas.AccessRule | del(.components.schemas.AccessRule) else . end' | \
    jq 'walk(if type == "object" and ."$ref" == "#/components/schemas/Device" then ."$ref" = "#/components/schemas/FmcDevice" else . end)' | \
    jq 'if .components.schemas.Device then .components.schemas.FmcDevice = .components.schemas.Device | del(.components.schemas.Device) else . end' | \
    # Add servers section
    jq '.servers = [
        {"url": "https://us.manage.security.cisco.com/api/rest", "description": "US"},
        {"url": "https://eu.manage.security.cisco.com/api/rest", "description": "EU"},
        {"url": "https://apj.manage.security.cisco.com/api/rest", "description": "APJ"},
        {"url": "https://staging.manage.security.cisco.com/api/rest", "description": "Staging"},
        {"url": "https://scale.manage.security.cisco.com/api/rest", "description": "Scale"},
        {"url": "https://ci.manage.security.cisco.com/api/rest", "description": "CI"}
      ]' | \
    yq e -P - > "${filename}"
}