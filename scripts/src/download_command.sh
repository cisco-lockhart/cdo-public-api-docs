root_dir=$(dirname $(dirname $(readlink -f "$0")))
filename=openapi.yaml
echo -n "$(yellow Triggering transformation)... "
transformation_id=`curl -X POST --silent --url "https://api.apimatic.io/transformations/transform-via-url" -H "Authorization: X-Auth-Key: ${args[--apimatic-api-key]}" -H "Accept: application/json" -H "Content-Type: application/vnd.apimatic.urlTransformDto.v1+json" --data-raw "{\"url\": \"https://edge.staging.cdo.cisco.com/api/platform/public-api/v3/api-docs\", \"export_format\": \"OpenApi3Yaml\"}" | jq -r ".id"`
echo "✅︎"
transformation_url="https://api.apimatic.io/transformations/${transformation_id}/converted-file"
echo -n "$(yellow Downloading transformed file from) $(yellow_underlined ${transformation_url}) to $(blue_bold ${root_dir}/${filename})... "
curl -X GET --silent --url  "${transformation_url}" -H "Authorization: X-Auth-Key: ${args[--apimatic-api-key]}" -H 'Accept: application/json' -H 'Content-Type: application/vnd.apimatic.urlTransformDto.v1+json' -o $root_dir/openapi.yaml
echo "✅︎"

if [[ -z ${args[--do-not-commit]} ]]; then
    git checkout main
    git pull origin main
    git add .
    git commit -am "Update OpenAPI.yaml from staging"
    git push origin main
fi