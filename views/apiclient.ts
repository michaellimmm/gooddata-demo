import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const LoginRequest = z
  .object({ email: z.string(), password: z.string() })
  .partial()
  .passthrough();
const LoginResponse = z
  .object({
    email: z.string(),
    name: z.string(),
    tenantId: z.string(),
    accessToken: z.string(),
  })
  .partial()
  .passthrough();
const GoogleProtobufAny = z
  .object({ "@type": z.string() })
  .partial()
  .passthrough();
const Status = z
  .object({
    code: z.number().int(),
    message: z.string(),
    details: z.array(GoogleProtobufAny),
  })
  .partial()
  .passthrough();
const RegisterAccountRequest = z
  .object({
    email: z.string(),
    password: z.string(),
    name: z.string(),
    tenantId: z.string(),
  })
  .partial()
  .passthrough();
const RequestAccountResponse = z
  .object({
    email: z.string(),
    name: z.string(),
    tenantId: z.string(),
    accessToken: z.string(),
  })
  .partial()
  .passthrough();
const GetTokenRequest = z
  .object({ tenantId: z.string() })
  .partial()
  .passthrough();
const GetTokenResponse = z
  .object({ accessToken: z.string() })
  .partial()
  .passthrough();

export const schemas = {
  LoginRequest,
  LoginResponse,
  GoogleProtobufAny,
  Status,
  RegisterAccountRequest,
  RequestAccountResponse,
  GetTokenRequest,
  GetTokenResponse,
};

const endpoints = makeApi([
  {
    method: "post",
    path: "/v1/login",
    alias: "AnalyticService_Login",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: LoginRequest,
      },
    ],
    response: LoginResponse,
  },
  {
    method: "post",
    path: "/v1/register",
    alias: "AnalyticService_RegisterAccount",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: RegisterAccountRequest,
      },
    ],
    response: RequestAccountResponse,
  },
  {
    method: "post",
    path: "/v1/token",
    alias: "AnalyticService_GetToken",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z.object({ tenantId: z.string() }).partial().passthrough(),
      },
    ],
    response: z.object({ accessToken: z.string() }).partial().passthrough(),
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
