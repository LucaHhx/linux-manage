/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Any } from "../../google/protobuf/any";

export const protobufPackage = "common";

/** 通用code */
export enum Code {
  /** Unknown - 未知 */
  Unknown = 0,
  /** Ok - 成功 */
  Ok = 1,
  /** TokenInvalid - token失效 */
  TokenInvalid = 401,
  /** ParameterError - 参数错误 */
  ParameterError = 402,
  /** Exists - 已存在 */
  Exists = 403,
  /** NotExists - 不存在 */
  NotExists = 404,
  /** SystemError - 系统错误 */
  SystemError = 405,
  /** DbError - 数据库错误 */
  DbError = 406,
  /** StatusError - 状态错误 */
  StatusError = 407,
  /** AmountInsufficientError - 余额不足 */
  AmountInsufficientError = 408,
  UNRECOGNIZED = -1,
}

export function codeFromJSON(object: any): Code {
  switch (object) {
    case 0:
    case "Unknown":
      return Code.Unknown;
    case 1:
    case "Ok":
      return Code.Ok;
    case 401:
    case "TokenInvalid":
      return Code.TokenInvalid;
    case 402:
    case "ParameterError":
      return Code.ParameterError;
    case 403:
    case "Exists":
      return Code.Exists;
    case 404:
    case "NotExists":
      return Code.NotExists;
    case 405:
    case "SystemError":
      return Code.SystemError;
    case 406:
    case "DbError":
      return Code.DbError;
    case 407:
    case "StatusError":
      return Code.StatusError;
    case 408:
    case "AmountInsufficientError":
      return Code.AmountInsufficientError;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Code.UNRECOGNIZED;
  }
}

export function codeToJSON(object: Code): string {
  switch (object) {
    case Code.Unknown:
      return "Unknown";
    case Code.Ok:
      return "Ok";
    case Code.TokenInvalid:
      return "TokenInvalid";
    case Code.ParameterError:
      return "ParameterError";
    case Code.Exists:
      return "Exists";
    case Code.NotExists:
      return "NotExists";
    case Code.SystemError:
      return "SystemError";
    case Code.DbError:
      return "DbError";
    case Code.StatusError:
      return "StatusError";
    case Code.AmountInsufficientError:
      return "AmountInsufficientError";
    case Code.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface reqHead {
  token: string;
  userId: number;
}

export interface repHead {
  code: Code;
  msg: string;
}

export interface basicReq {
  head: reqHead | undefined;
  data: Any | undefined;
}

function createBasereqHead(): reqHead {
  return { token: "", userId: 0 };
}

export const reqHead = {
  encode(message: reqHead, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    if (message.userId !== 0) {
      writer.uint32(24).int32(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): reqHead {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasereqHead();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.token = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.userId = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): reqHead {
    return {
      token: isSet(object.token) ? globalThis.String(object.token) : "",
      userId: isSet(object.userId) ? globalThis.Number(object.userId) : 0,
    };
  },

  toJSON(message: reqHead): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<reqHead>, I>>(base?: I): reqHead {
    return reqHead.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<reqHead>, I>>(object: I): reqHead {
    const message = createBasereqHead();
    message.token = object.token ?? "";
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaserepHead(): repHead {
  return { code: 0, msg: "" };
}

export const repHead = {
  encode(message: repHead, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.msg !== "") {
      writer.uint32(18).string(message.msg);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): repHead {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaserepHead();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.msg = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): repHead {
    return {
      code: isSet(object.code) ? codeFromJSON(object.code) : 0,
      msg: isSet(object.msg) ? globalThis.String(object.msg) : "",
    };
  },

  toJSON(message: repHead): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = codeToJSON(message.code);
    }
    if (message.msg !== "") {
      obj.msg = message.msg;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<repHead>, I>>(base?: I): repHead {
    return repHead.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<repHead>, I>>(object: I): repHead {
    const message = createBaserepHead();
    message.code = object.code ?? 0;
    message.msg = object.msg ?? "";
    return message;
  },
};

function createBasebasicReq(): basicReq {
  return { head: undefined, data: undefined };
}

export const basicReq = {
  encode(message: basicReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.head !== undefined) {
      reqHead.encode(message.head, writer.uint32(10).fork()).ldelim();
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): basicReq {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasebasicReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.head = reqHead.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.data = Any.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): basicReq {
    return {
      head: isSet(object.head) ? reqHead.fromJSON(object.head) : undefined,
      data: isSet(object.data) ? Any.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: basicReq): unknown {
    const obj: any = {};
    if (message.head !== undefined) {
      obj.head = reqHead.toJSON(message.head);
    }
    if (message.data !== undefined) {
      obj.data = Any.toJSON(message.data);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<basicReq>, I>>(base?: I): basicReq {
    return basicReq.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<basicReq>, I>>(object: I): basicReq {
    const message = createBasebasicReq();
    message.head = (object.head !== undefined && object.head !== null) ? reqHead.fromPartial(object.head) : undefined;
    message.data = (object.data !== undefined && object.data !== null) ? Any.fromPartial(object.data) : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
