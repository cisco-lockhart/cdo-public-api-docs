root_dir=$(dirname $(dirname $(readlink -f "$0")))
filename=openapi.yaml
declare -A urls=( ["public-api"]="https://edge.staging.cdo.cisco.com/api/platform/public-api/v3/api-docs.yaml" ["object-service"]="https://edge.staging.cdo.cisco.com/api/platform/object-service/v3/api-docs.yaml" )

filenames=()
for service in "${!urls[@]}"; do
    filename="${service}-openapi.yaml"
    filenames+=("${root_dir}/${filename}")
    url=${urls[${service}]}
    echo -n "$(yellow Downloading file from) $(yellow_underlined ${url}) to $(blue_bold ${root_dir}/${filename})... "
    curl -X GET --silent --url  "${url}" -o $root_dir/$filename
    echo "✅︎"
done

echo "$(yellow Installing redocly from npm)... "
npm i
echo "Installed ✅︎"

echo -n "$(yellow Combining all OpenAPI YAMLs into one)... "
 ./node_modules/.bin/redocly join ${filenames[*]} -o openapi.yaml
echo "Combined ✅︎"

if [[ -z ${args[--do-not-commit]} ]]; then
    cd ${root_dir}
    git checkout main
    git pull origin main
    git add .
    git commit -am "Update OpenAPI.yaml from staging"
    git push origin main
    cd -
fi