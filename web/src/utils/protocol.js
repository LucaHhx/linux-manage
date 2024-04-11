// protocol.js
const PKG_HEAD_BYTES = 4
const MSG_FLAG_BYTES = 1
const MSG_ROUTE_CODE_BYTES = 2
const MSG_ID_MAX_BYTES = 5
const MSG_ROUTE_LEN_BYTES = 1

const MSG_ROUTE_CODE_MAX = 0xffff

const MSG_COMPRESS_ROUTE_MASK = 0x1
const MSG_TYPE_MASK = 0x7

export class Protocol {
  static strencode(str) {
    const byteArray = new Uint8Array(str.length * 3)
    let offset = 0
    for (let i = 0; i < str.length; i++) {
      const charCode = str.charCodeAt(i)
      let codes = null
      if (charCode <= 0x7f) {
        codes = [charCode]
      } else if (charCode <= 0x7ff) {
        codes = [0xc0 | (charCode >> 6), 0x80 | (charCode & 0x3f)]
      } else {
        codes = [0xe0 | (charCode >> 12), 0x80 | ((charCode & 0xfc0) >> 6), 0x80 | (charCode & 0x3f)]
      }
      for (let j = 0; j < codes.length; j++) {
        byteArray[offset] = codes[j]
        ++offset
      }
    }
    const buffer = byteArray.subarray(0, offset)
    return buffer
  }

  static strdecode(buffer) {
    const bytes = new Uint8Array(buffer)
    const array = []
    let offset = 0
    let charCode = 0
    const end = bytes.length
    while (offset < end) {
      if (bytes[offset] < 128) {
        charCode = bytes[offset]
        offset += 1
      } else if (bytes[offset] < 224) {
        charCode = ((bytes[offset] & 0x3f) << 6) + (bytes[offset + 1] & 0x3f)
        offset += 2
      } else {
        charCode = ((bytes[offset] & 0x0f) << 12) + ((bytes[offset + 1] & 0x3f) << 6) + (bytes[offset + 2] & 0x3f)
        offset += 3
      }
      array.push(charCode)
    }
    return String.fromCharCode.apply(null, array)
  }
}

// Defining the Package and Message as static properties on the Protocol class
Protocol.Package = {
  TYPE_HANDSHAKE: 1,
  TYPE_HANDSHAKE_ACK: 2,
  TYPE_HEARTBEAT: 3,
  TYPE_DATA: 4,
  TYPE_KICK: 5,
  // ... Package's encode and decode methods go here
  encode: function(type, body) {
    const length = body ? body.length : 0
    const buffer = new Uint8Array(PKG_HEAD_BYTES + length)
    let index = 0
    buffer[index++] = type & 0xff
    buffer[index++] = (length >> 16) & 0xff
    buffer[index++] = (length >> 8) & 0xff
    buffer[index++] = length & 0xff
    if (body) {
      copyArray(buffer, index, body, 0, length)
    }
    return buffer
  },
  decode: function(buffer) {
    let offset = 0
    const bytes = new Uint8Array(buffer)
    let length = 0
    const rs = []
    while (offset < bytes.length) {
      const type = bytes[offset++]
      length = ((bytes[offset++]) << 16 | (bytes[offset++]) << 8 | bytes[offset++]) >>> 0
      const body = length ? new Uint8Array(length) : null
      copyArray(body, 0, bytes, offset, length)
      offset += length
      rs.push({ 'type': type, 'body': body })
    }
    return rs.length === 1 ? rs[0] : rs
  }
}

Protocol.Message = {
  TYPE_REQUEST: 0,
  TYPE_NOTIFY: 1,
  TYPE_RESPONSE: 2,
  TYPE_PUSH: 3,
  // ... Message's encode and decode methods go here
  encode: function(id, type, compressRoute, route, msg) {
    // caculate message max length
    const idBytes = msgHasId(type) ? caculateMsgIdBytes(id) : 0
    let msgLen = MSG_FLAG_BYTES + idBytes

    if (msgHasRoute(type)) {
      if (compressRoute) {
        if (typeof route !== 'number') {
          throw new Error('error flag for number route!')
        }
        msgLen += MSG_ROUTE_CODE_BYTES
      } else {
        msgLen += MSG_ROUTE_LEN_BYTES
        if (route) {
          route = Protocol.strencode(route)
          if (route.length > 255) {
            throw new Error('route maxlength is overflow')
          }
          msgLen += route.length
        }
      }
    }

    if (msg) {
      msgLen += msg.length
    }

    const buffer = new Uint8Array(msgLen)
    let offset = 0

    // add flag
    offset = encodeMsgFlag(type, compressRoute, buffer, offset)

    // add message id
    if (msgHasId(type)) {
      offset = encodeMsgId(id, buffer, offset)
    }

    // add route
    if (msgHasRoute(type)) {
      offset = encodeMsgRoute(compressRoute, route, buffer, offset)
    }

    // add body
    if (msg) {
      offset = encodeMsgBody(msg, buffer, offset)
    }

    return buffer
  },

  /**
   * Message protocol decode.
   *
   * @param  {Buffer|Uint8Array} buffer message bytes
   * @return {Object}            message object
   */
  decode: function(buffer) {
    const bytes = new Uint8Array(buffer)
    const bytesLen = bytes.length || bytes.byteLength
    let offset = 0
    let id = 0
    let route = null

    // parse flag
    const flag = bytes[offset++]
    const compressRoute = flag & MSG_COMPRESS_ROUTE_MASK
    const type = (flag >> 1) & MSG_TYPE_MASK

    // parse id
    if (msgHasId(type)) {
      let m = parseInt(bytes[offset])
      let i = 0
      do {
        m = parseInt(bytes[offset])
        id = id + ((m & 0x7f) * Math.pow(2, (7 * i)))
        offset++
        i++
      } while (m >= 128)
    }

    // parse route
    if (msgHasRoute(type)) {
      if (compressRoute) {
        route = (bytes[offset++]) << 8 | bytes[offset++]
      } else {
        const routeLen = bytes[offset++]
        if (routeLen) {
          route = new Uint8Array(routeLen)
          copyArray(route, 0, bytes, offset, routeLen)
          route = Protocol.strdecode(route)
        } else {
          route = ''
        }
        offset += routeLen
      }
    }

    // parse body
    const bodyLen = bytesLen - offset
    const body = new Uint8Array(bodyLen)

    copyArray(body, 0, bytes, offset, bodyLen)

    return {
      'id': id, 'type': type, 'compressRoute': compressRoute,
      'route': route, 'body': body
    }
  },
}

// Expose the Protocol to the global window object if it's not being required as a module
if (typeof window !== 'undefined') {
  window.Protocol = Protocol
}

const copyArray = (dest, doffset, src, soffset, length) => {
  if (typeof src.copy === 'function') {
    // Buffer
    src.copy(dest, doffset, soffset, soffset + length)
  } else {
    // Uint8Array
    for (let index = 0; index < length; index++) {
      dest[doffset++] = src[soffset++]
    }
  }
}

const msgHasId = (type) => {
  return type === Protocol.Message.TYPE_REQUEST || type === Protocol.Message.TYPE_RESPONSE
}

const msgHasRoute = (type) => {
  return type === Protocol.Message.TYPE_REQUEST || type === Protocol.Message.TYPE_NOTIFY ||
    type === Protocol.Message.TYPE_PUSH
}

const caculateMsgIdBytes = (id) => {
  let len = 0
  do {
    len += 1
    id >>= 7
  } while (id > 0)
  return len
}

const encodeMsgFlag = (type, compressRoute, buffer, offset) => {
  if (type !== Protocol.Message.TYPE_REQUEST && type !== Protocol.Message.TYPE_NOTIFY &&
    type !== Protocol.Message.TYPE_RESPONSE && type !== Protocol.Message.TYPE_PUSH) {
    throw new Error('unkonw message type: ' + type)
  }

  buffer[offset] = (type << 1) | (compressRoute ? 1 : 0)

  return offset + MSG_FLAG_BYTES
}

const encodeMsgId = (id, buffer, offset) => {
  do {
    let tmp = id % 128
    const next = Math.floor(id / 128)

    if (next !== 0) {
      tmp = tmp + 128
    }
    buffer[offset++] = tmp

    id = next
  } while (id !== 0)

  return offset
}

const encodeMsgRoute = (compressRoute, route, buffer, offset) => {
  if (compressRoute) {
    if (route > MSG_ROUTE_CODE_MAX) {
      throw new Error('route number is overflow')
    }

    buffer[offset++] = (route >> 8) & 0xff
    buffer[offset++] = route & 0xff
  } else {
    if (route) {
      buffer[offset++] = route.length & 0xff
      copyArray(buffer, offset, route, 0, route.length)
      offset += route.length
    } else {
      buffer[offset++] = 0
    }
  }

  return offset
}

const encodeMsgBody = (msg, buffer, offset) => {
  copyArray(buffer, offset, msg, 0, msg.length)
  return offset + msg.length
}

// Export the Protocol for module use
export default Protocol
