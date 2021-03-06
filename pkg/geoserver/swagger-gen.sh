rm -rf openapi

#workspace
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i https://docs.geoserver.org/latest/en/api/1.0.0/workspaces.yaml \
  -g go \
  -o /local/openapi

#layer
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i https://docs.geoserver.org/latest/en/api/1.0.0/layers.yaml \
  -g go \
  -o /local/openapi \
	--skip-validate-spec

# #datastores
# docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
#   -i https://docs.geoserver.org/latest/en/api/1.0.0/datastores.yaml \
#   -g go \
#   -o /local/datastores
