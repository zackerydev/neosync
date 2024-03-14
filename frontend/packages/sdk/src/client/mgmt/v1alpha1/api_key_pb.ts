// @generated by protoc-gen-es v1.8.0 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/api_key.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message mgmt.v1alpha1.CreateAccountApiKeyRequest
 */
export class CreateAccountApiKeyRequest extends Message<CreateAccountApiKeyRequest> {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * Validate between now and one year: now < x < 365 days
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 3;
   */
  expiresAt?: Timestamp;

  constructor(data?: PartialMessage<CreateAccountApiKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CreateAccountApiKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "expires_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountApiKeyRequest {
    return new CreateAccountApiKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountApiKeyRequest {
    return new CreateAccountApiKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountApiKeyRequest {
    return new CreateAccountApiKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateAccountApiKeyRequest | PlainMessage<CreateAccountApiKeyRequest> | undefined, b: CreateAccountApiKeyRequest | PlainMessage<CreateAccountApiKeyRequest> | undefined): boolean {
    return proto3.util.equals(CreateAccountApiKeyRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.CreateAccountApiKeyResponse
 */
export class CreateAccountApiKeyResponse extends Message<CreateAccountApiKeyResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.AccountApiKey api_key = 1;
   */
  apiKey?: AccountApiKey;

  constructor(data?: PartialMessage<CreateAccountApiKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CreateAccountApiKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "api_key", kind: "message", T: AccountApiKey },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountApiKeyResponse {
    return new CreateAccountApiKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountApiKeyResponse {
    return new CreateAccountApiKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountApiKeyResponse {
    return new CreateAccountApiKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateAccountApiKeyResponse | PlainMessage<CreateAccountApiKeyResponse> | undefined, b: CreateAccountApiKeyResponse | PlainMessage<CreateAccountApiKeyResponse> | undefined): boolean {
    return proto3.util.equals(CreateAccountApiKeyResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.AccountApiKey
 */
export class AccountApiKey extends Message<AccountApiKey> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * The friendly name of the API Key
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string account_id = 3;
   */
  accountId = "";

  /**
   * @generated from field: string created_by_id = 4;
   */
  createdById = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string updated_by_id = 6;
   */
  updatedById = "";

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  /**
   * key_value is only returned on initial creation or when it is regenerated
   *
   * @generated from field: optional string key_value = 8;
   */
  keyValue?: string;

  /**
   * @generated from field: string user_id = 9;
   */
  userId = "";

  /**
   * The timestamp of what the API key expires and will not longer be usable.
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 10;
   */
  expiresAt?: Timestamp;

  constructor(data?: PartialMessage<AccountApiKey>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.AccountApiKey";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "created_by_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "updated_by_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
    { no: 8, name: "key_value", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 9, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "expires_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AccountApiKey {
    return new AccountApiKey().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AccountApiKey {
    return new AccountApiKey().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AccountApiKey {
    return new AccountApiKey().fromJsonString(jsonString, options);
  }

  static equals(a: AccountApiKey | PlainMessage<AccountApiKey> | undefined, b: AccountApiKey | PlainMessage<AccountApiKey> | undefined): boolean {
    return proto3.util.equals(AccountApiKey, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetAccountApiKeysRequest
 */
export class GetAccountApiKeysRequest extends Message<GetAccountApiKeysRequest> {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId = "";

  constructor(data?: PartialMessage<GetAccountApiKeysRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetAccountApiKeysRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountApiKeysRequest {
    return new GetAccountApiKeysRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountApiKeysRequest {
    return new GetAccountApiKeysRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountApiKeysRequest {
    return new GetAccountApiKeysRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountApiKeysRequest | PlainMessage<GetAccountApiKeysRequest> | undefined, b: GetAccountApiKeysRequest | PlainMessage<GetAccountApiKeysRequest> | undefined): boolean {
    return proto3.util.equals(GetAccountApiKeysRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetAccountApiKeysResponse
 */
export class GetAccountApiKeysResponse extends Message<GetAccountApiKeysResponse> {
  /**
   * @generated from field: repeated mgmt.v1alpha1.AccountApiKey api_keys = 1;
   */
  apiKeys: AccountApiKey[] = [];

  constructor(data?: PartialMessage<GetAccountApiKeysResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetAccountApiKeysResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "api_keys", kind: "message", T: AccountApiKey, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountApiKeysResponse {
    return new GetAccountApiKeysResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountApiKeysResponse {
    return new GetAccountApiKeysResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountApiKeysResponse {
    return new GetAccountApiKeysResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountApiKeysResponse | PlainMessage<GetAccountApiKeysResponse> | undefined, b: GetAccountApiKeysResponse | PlainMessage<GetAccountApiKeysResponse> | undefined): boolean {
    return proto3.util.equals(GetAccountApiKeysResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetAccountApiKeyRequest
 */
export class GetAccountApiKeyRequest extends Message<GetAccountApiKeyRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetAccountApiKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetAccountApiKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountApiKeyRequest {
    return new GetAccountApiKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountApiKeyRequest {
    return new GetAccountApiKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountApiKeyRequest {
    return new GetAccountApiKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountApiKeyRequest | PlainMessage<GetAccountApiKeyRequest> | undefined, b: GetAccountApiKeyRequest | PlainMessage<GetAccountApiKeyRequest> | undefined): boolean {
    return proto3.util.equals(GetAccountApiKeyRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetAccountApiKeyResponse
 */
export class GetAccountApiKeyResponse extends Message<GetAccountApiKeyResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.AccountApiKey api_key = 1;
   */
  apiKey?: AccountApiKey;

  constructor(data?: PartialMessage<GetAccountApiKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetAccountApiKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "api_key", kind: "message", T: AccountApiKey },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountApiKeyResponse {
    return new GetAccountApiKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountApiKeyResponse {
    return new GetAccountApiKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountApiKeyResponse {
    return new GetAccountApiKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountApiKeyResponse | PlainMessage<GetAccountApiKeyResponse> | undefined, b: GetAccountApiKeyResponse | PlainMessage<GetAccountApiKeyResponse> | undefined): boolean {
    return proto3.util.equals(GetAccountApiKeyResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.RegenerateAccountApiKeyRequest
 */
export class RegenerateAccountApiKeyRequest extends Message<RegenerateAccountApiKeyRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Validate between now and one year: now < x < 365 days
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 2;
   */
  expiresAt?: Timestamp;

  constructor(data?: PartialMessage<RegenerateAccountApiKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.RegenerateAccountApiKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegenerateAccountApiKeyRequest {
    return new RegenerateAccountApiKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegenerateAccountApiKeyRequest {
    return new RegenerateAccountApiKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegenerateAccountApiKeyRequest {
    return new RegenerateAccountApiKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RegenerateAccountApiKeyRequest | PlainMessage<RegenerateAccountApiKeyRequest> | undefined, b: RegenerateAccountApiKeyRequest | PlainMessage<RegenerateAccountApiKeyRequest> | undefined): boolean {
    return proto3.util.equals(RegenerateAccountApiKeyRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.RegenerateAccountApiKeyResponse
 */
export class RegenerateAccountApiKeyResponse extends Message<RegenerateAccountApiKeyResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.AccountApiKey api_key = 1;
   */
  apiKey?: AccountApiKey;

  constructor(data?: PartialMessage<RegenerateAccountApiKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.RegenerateAccountApiKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "api_key", kind: "message", T: AccountApiKey },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegenerateAccountApiKeyResponse {
    return new RegenerateAccountApiKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegenerateAccountApiKeyResponse {
    return new RegenerateAccountApiKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegenerateAccountApiKeyResponse {
    return new RegenerateAccountApiKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RegenerateAccountApiKeyResponse | PlainMessage<RegenerateAccountApiKeyResponse> | undefined, b: RegenerateAccountApiKeyResponse | PlainMessage<RegenerateAccountApiKeyResponse> | undefined): boolean {
    return proto3.util.equals(RegenerateAccountApiKeyResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.DeleteAccountApiKeyRequest
 */
export class DeleteAccountApiKeyRequest extends Message<DeleteAccountApiKeyRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteAccountApiKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.DeleteAccountApiKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteAccountApiKeyRequest {
    return new DeleteAccountApiKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteAccountApiKeyRequest {
    return new DeleteAccountApiKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteAccountApiKeyRequest {
    return new DeleteAccountApiKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteAccountApiKeyRequest | PlainMessage<DeleteAccountApiKeyRequest> | undefined, b: DeleteAccountApiKeyRequest | PlainMessage<DeleteAccountApiKeyRequest> | undefined): boolean {
    return proto3.util.equals(DeleteAccountApiKeyRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.DeleteAccountApiKeyResponse
 */
export class DeleteAccountApiKeyResponse extends Message<DeleteAccountApiKeyResponse> {
  constructor(data?: PartialMessage<DeleteAccountApiKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.DeleteAccountApiKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteAccountApiKeyResponse {
    return new DeleteAccountApiKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteAccountApiKeyResponse {
    return new DeleteAccountApiKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteAccountApiKeyResponse {
    return new DeleteAccountApiKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteAccountApiKeyResponse | PlainMessage<DeleteAccountApiKeyResponse> | undefined, b: DeleteAccountApiKeyResponse | PlainMessage<DeleteAccountApiKeyResponse> | undefined): boolean {
    return proto3.util.equals(DeleteAccountApiKeyResponse, a, b);
  }
}

