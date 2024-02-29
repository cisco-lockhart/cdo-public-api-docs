root_dir=$(dirname $(dirname $(readlink -f "$0")))
filename=openapi.yaml
declare -A urls=(
    # Currently dependent on the CDO_cisco-ai-team-cdo tenant cdFMC - need to update
    ["cdFmc-service"]="https://demo-red.app.us.cdo.cisco.com/api/api-explorer/fmc_oas3.json"
)

filenames=()
for service in "${!urls[@]}"; do
    url=${urls[${service}]}
    filename="${service}-openapi.yaml"
    echo -n "$(yellow Downloading file from) $(yellow_underlined ${url}) to $(blue_bold ${root_dir}/${filename})... "
    curl -s "$url" | \
    # Replace cdFMC URLs with the new Public API proxy URL
    jq '.paths |= with_entries(.key |= gsub("api/fmc_config/"; "v1/cdfmc/api/fmc_config/"))' | \
    # Add servers section
    jq '.servers = [
        {"url": "https://edge.us.cdo.cisco.com/api/rest", "description": "US"},
        {"url": "https://edge.eu.cdo.cisco.com/api/rest", "description": "EU"},
        {"url": "https://edge.apj.cdo.cisco.com/api/rest", "description": "APJ"},
        {"url": "https://edge.staging.cdo.cisco.com/api/rest", "description": "Staging"},
        {"url": "https://edge.scale.cdo.cisco.com/api/rest", "description": "Scale"},
        {"url": "https://edge.ci.cdo.cisco.com/api/rest", "description": "CI"}
      ]' | \
    yq e -P - > "${root_dir}/${filename}"
    echo "✅︎"
done