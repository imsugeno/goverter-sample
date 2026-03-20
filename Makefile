.PHONY: generate generate-api generate-converter

generate: generate-api generate-converter

generate-api:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		--config api/oapi-codegen.yaml \
		api/openapi.yaml

generate-converter:
	go run github.com/jmattheis/goverter/cmd/goverter gen ./converter/
