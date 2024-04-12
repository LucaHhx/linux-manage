/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "resources";

export interface downloadInfo {
  id: number;
  progress: number;
  finishBytes: string;
  totalBytes: string;
  speed: string;
  isFinished: boolean;
}

function createBasedownloadInfo(): downloadInfo {
  return { id: 0, progress: 0, finishBytes: "", totalBytes: "", speed: "", isFinished: false };
}

export const downloadInfo = {
  encode(message: downloadInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.progress !== 0) {
      writer.uint32(16).int32(message.progress);
    }
    if (message.finishBytes !== "") {
      writer.uint32(26).string(message.finishBytes);
    }
    if (message.totalBytes !== "") {
      writer.uint32(34).string(message.totalBytes);
    }
    if (message.speed !== "") {
      writer.uint32(42).string(message.speed);
    }
    if (message.isFinished !== false) {
      writer.uint32(48).bool(message.isFinished);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): downloadInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasedownloadInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.progress = reader.int32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.finishBytes = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.totalBytes = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.speed = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.isFinished = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): downloadInfo {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      progress: isSet(object.progress) ? globalThis.Number(object.progress) : 0,
      finishBytes: isSet(object.finishBytes) ? globalThis.String(object.finishBytes) : "",
      totalBytes: isSet(object.totalBytes) ? globalThis.String(object.totalBytes) : "",
      speed: isSet(object.speed) ? globalThis.String(object.speed) : "",
      isFinished: isSet(object.isFinished) ? globalThis.Boolean(object.isFinished) : false,
    };
  },

  toJSON(message: downloadInfo): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.progress !== 0) {
      obj.progress = Math.round(message.progress);
    }
    if (message.finishBytes !== "") {
      obj.finishBytes = message.finishBytes;
    }
    if (message.totalBytes !== "") {
      obj.totalBytes = message.totalBytes;
    }
    if (message.speed !== "") {
      obj.speed = message.speed;
    }
    if (message.isFinished !== false) {
      obj.isFinished = message.isFinished;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<downloadInfo>, I>>(base?: I): downloadInfo {
    return downloadInfo.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<downloadInfo>, I>>(object: I): downloadInfo {
    const message = createBasedownloadInfo();
    message.id = object.id ?? 0;
    message.progress = object.progress ?? 0;
    message.finishBytes = object.finishBytes ?? "";
    message.totalBytes = object.totalBytes ?? "";
    message.speed = object.speed ?? "";
    message.isFinished = object.isFinished ?? false;
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
