/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "room";

export interface userMessage {
  name: string;
  content: string;
}

export interface newUser {
  content: string;
}

export interface allMembers {
  members: Long[];
}

export interface joinResponse {
  code: number;
  result: string;
}

function createBaseuserMessage(): userMessage {
  return { name: "", content: "" };
}

export const userMessage = {
  encode(message: userMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): userMessage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseuserMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.content = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): userMessage {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      content: isSet(object.content) ? globalThis.String(object.content) : "",
    };
  },

  toJSON(message: userMessage): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.content !== "") {
      obj.content = message.content;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<userMessage>, I>>(base?: I): userMessage {
    return userMessage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<userMessage>, I>>(object: I): userMessage {
    const message = createBaseuserMessage();
    message.name = object.name ?? "";
    message.content = object.content ?? "";
    return message;
  },
};

function createBasenewUser(): newUser {
  return { content: "" };
}

export const newUser = {
  encode(message: newUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content !== "") {
      writer.uint32(10).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): newUser {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasenewUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.content = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): newUser {
    return { content: isSet(object.content) ? globalThis.String(object.content) : "" };
  },

  toJSON(message: newUser): unknown {
    const obj: any = {};
    if (message.content !== "") {
      obj.content = message.content;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<newUser>, I>>(base?: I): newUser {
    return newUser.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<newUser>, I>>(object: I): newUser {
    const message = createBasenewUser();
    message.content = object.content ?? "";
    return message;
  },
};

function createBaseallMembers(): allMembers {
  return { members: [] };
}

export const allMembers = {
  encode(message: allMembers, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.members) {
      writer.int64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): allMembers {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseallMembers();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag === 8) {
            message.members.push(reader.int64() as Long);

            continue;
          }

          if (tag === 10) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.members.push(reader.int64() as Long);
            }

            continue;
          }

          break;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): allMembers {
    return {
      members: globalThis.Array.isArray(object?.members) ? object.members.map((e: any) => Long.fromValue(e)) : [],
    };
  },

  toJSON(message: allMembers): unknown {
    const obj: any = {};
    if (message.members?.length) {
      obj.members = message.members.map((e) => (e || Long.ZERO).toString());
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<allMembers>, I>>(base?: I): allMembers {
    return allMembers.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<allMembers>, I>>(object: I): allMembers {
    const message = createBaseallMembers();
    message.members = object.members?.map((e) => Long.fromValue(e)) || [];
    return message;
  },
};

function createBasejoinResponse(): joinResponse {
  return { code: 0, result: "" };
}

export const joinResponse = {
  encode(message: joinResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.result !== "") {
      writer.uint32(18).string(message.result);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): joinResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasejoinResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.result = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): joinResponse {
    return {
      code: isSet(object.code) ? globalThis.Number(object.code) : 0,
      result: isSet(object.result) ? globalThis.String(object.result) : "",
    };
  },

  toJSON(message: joinResponse): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = Math.round(message.code);
    }
    if (message.result !== "") {
      obj.result = message.result;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<joinResponse>, I>>(base?: I): joinResponse {
    return joinResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<joinResponse>, I>>(object: I): joinResponse {
    const message = createBasejoinResponse();
    message.code = object.code ?? 0;
    message.result = object.result ?? "";
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
