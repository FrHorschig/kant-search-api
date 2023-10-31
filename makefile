.PHONY: build

validate:
	docker run --rm -v ${PWD}:/work:ro dshanley/vacuum lint -d src/openapi/openapi.yaml