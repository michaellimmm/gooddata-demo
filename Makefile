gen-proto:
	buf generate proto

gen-open-api:
	pnpm openapi-zod-client "./generated/openapi.yaml" -o "./views/apiclient.ts"