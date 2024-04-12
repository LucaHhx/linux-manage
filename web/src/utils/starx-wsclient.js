import Protocol from './protocol'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { protos } from '@/utils/proto'
class Emitter {
  constructor(obj) {
    if (obj) return mixin(obj)
  }

  on(event, fn) {
    this._callbacks = this._callbacks || {};
    (this._callbacks[event] = this._callbacks[event] || []).push(fn)
    return this
  }

  once(event, fn) {
    function on() {
      this.off(event, on)
      fn.apply(this, arguments)
    }

    on.fn = fn
    this.on(event, on)
    return this
  }

  off(event, fn) {
    this._callbacks = this._callbacks || {}

    if (arguments.length === 0) {
      this._callbacks = {}
      return this
    }

    const callbacks = this._callbacks[event]
    if (!callbacks) return this

    if (arguments.length === 1) {
      delete this._callbacks[event]
      return this
    }

    for (let i = 0; i < callbacks.length; i++) {
      const cb = callbacks[i]
      if (cb === fn || cb.fn === fn) {
        callbacks.splice(i, 1)
        break
      }
    }
    return this
  }

  emit(event) {
    this._callbacks = this._callbacks || {}
    var args = [].slice.call(arguments, 1)
    var callbacks = this._callbacks[event]
    const message = converProto(args[0])
    if (callbacks && message) {
      callbacks = callbacks.slice(0)
      for (var i = 0, len = callbacks.length; i < len; ++i) {
        callbacks[i].apply(this, [message])
      }
    }
    return this
  }

  listeners(event) {
    this._callbacks = this._callbacks || {}
    return this._callbacks[event] || []
  }

  hasListeners(event) {
    return !!this.listeners(event).length
  }
}

// Helper function to mixin Emitter properties
function mixin(obj) {
  for (const key in Emitter.prototype) {
    obj[key] = Emitter.prototype[key]
  }
  return obj
}

var JS_WS_CLIENT_TYPE = 'js-websocket'
var JS_WS_CLIENT_VERSION = '0.0.1'

var decodeIO_encoder = null
var decodeIO_decoder = null
var isJson = false
var Package = Protocol.Package
var Message = Protocol.Message
var EventEmitter = Emitter
var rsa = window.rsa

if (typeof (window) !== 'undefined' && typeof (sys) !== 'undefined' && sys.localStorage) {
  window.localStorage = sys.localStorage
}

var RES_OK = 200
var RES_FAIL = 500
var RES_OLD_CLIENT = 501

if (typeof Object.create !== 'function') {
  Object.create = function(o) {
    function F() { }
    F.prototype = o
    return new F()
  }
}

class Starx extends Emitter {
  constructor() {
    super()
    // Extend Emitter with starx specific properties
    mixin(this)
  }

  init(params, cb, closeCb) {
    initCallback = cb
    var host = params.host
    var port = params.port
    var path = params.path

    encode = params.encode || defaultEncode
    decode = params.decode || defaultDecode

    var url = 'ws://' + host
    if (port) {
      url += ':' + port
    }

    if (path) {
      url += path
    }

    handshakeBuffer.user = params.user
    if (params.encrypt) {
      useCrypto = true
      rsa.generate(1024, '10001')
      var data = {
        rsa_n: rsa.n.toString(16),
        rsa_e: rsa.e
      }
      handshakeBuffer.sys.rsa = data
    }
    handshakeCallback = params.handshakeCallback
    connect(params, url, cb, closeCb)
  }

  decode(data) {
    var msg = Message.decode(data)

    if (msg.id > 0) {
      msg.route = routeMap[msg.id]
      delete routeMap[msg.id]
      if (!msg.route) {
        return
      }
    }

    msg.body = deCompose(msg)
    return msg
  }

  encode(reqId, route, msg) {
    var type = reqId ? Message.TYPE_REQUEST : Message.TYPE_NOTIFY

    // if (isJson) {
    //   msg = Protocol.strencode(JSON.stringify(msg))
    // }

    var compressRoute = 0
    if (dict && dict[route]) {
      route = dict[route]
      compressRoute = 1
    }

    return Message.encode(reqId, type, compressRoute, route, msg)
  }

  disconnect() {
    if (socket) {
      if (socket.disconnect) socket.disconnect()
      if (socket.close) socket.close()
      // console.log('disconnect')
      socket = null
    }

    if (heartbeatId) {
      clearTimeout(heartbeatId)
      heartbeatId = null
    }
    if (heartbeatTimeoutId) {
      clearTimeout(heartbeatTimeoutId)
      heartbeatTimeoutId = null
    }
  }

  request(route, type, msg, cb) {
    if (arguments.length === 2 && typeof type === 'function') {
      cb = type
      msg = {}
      type = null
    } else {
      msg = msg || {}
    }
    route = route || msg.route
    if (!route) {
      return
    }

    reqId++
    sendMessage(reqId, route, type, msg)

    callbacks[reqId] = cb
    routeMap[reqId] = route
  }

  notify(route, type, msg) {
    msg = msg || {}
    sendMessage(0, route, type, msg)
  }
  // ... (Starx-specific methods like init, request, notify, etc.) ...
}

var root = window

export const starx = new Starx()
window.starx = starx

var socket = null
var reqId = 0
var callbacks = {}
var handlers = {}
// Map from request id to route
var routeMap = {}
var dict = {} // route string to code
var abbrs = {} // code to route string

var heartbeatInterval = 0
var heartbeatTimeout = 0
var nextHeartbeatTimeout = 0
var gapThreshold = 100 // heartbeat gap threashold
var heartbeatId = null
var heartbeatTimeoutId = null
var handshakeCallback = null

var decode = null
var encode = null

var reconnect = false
var reconncetTimer = null
var reconnectUrl = null
var reconnectAttempts = 0
var reconnectionDelay = 5000
var DEFAULT_MAX_RECONNECT_ATTEMPTS = 10

var useCrypto

var handshakeBuffer = {
  'sys': {
    type: JS_WS_CLIENT_TYPE,
    version: JS_WS_CLIENT_VERSION,
    rsa: {}
  },
  'user': {
  }
}

var initCallback = null

var defaultDecode = starx.decode

var defaultEncode = starx.encode

var connect = function(params, url, cb, closeCb) {
  // console.log('connect to ' + url)
  console.log(`%c connect to %c ` + url + ` %c`,
    'background:#0081ff; padding: 1px; border-radius: 3px 0 0 3px; color: #fff',
    'background:#354855; padding: 1px 5px; border-radius: 0 3px 3px 0; color: #fff; font-weight: bold;',
    'background:transparent')
  var params = params || {}
  var maxReconnectAttempts = params.maxReconnectAttempts || DEFAULT_MAX_RECONNECT_ATTEMPTS
  reconnectUrl = url

  var onopen = function(event) {
    if (reconnect) {
      starx.emit('reconnect')
    }
    reset()
    var obj = Package.encode(Package.TYPE_HANDSHAKE, Protocol.strencode(JSON.stringify(handshakeBuffer)))
    send(obj)
  }

  var onmessage = function(event) {
    processPackage(Package.decode(event.data), cb)
    // new package arrived, update the heartbeat timeout
    if (heartbeatTimeout) {
      nextHeartbeatTimeout = Date.now() + heartbeatTimeout
    }
  }
  var onerror = function(event) {
    starx.emit('io-error', event)
    ElMessage.error('socket error: ', event)
  }

  var onclose = function(event) {
    starx.emit('close', event)
    starx.emit('disconnect', event)
    console.log('socket close: ', event)
    // 断线回调
    if (typeof closeCb === 'function') {
      closeCb(event)
    }

    if (!!params.reconnect && reconnectAttempts < maxReconnectAttempts) {
      reconnect = true
      reconnectAttempts++
      reconncetTimer = setTimeout(function() {
        connect(params, reconnectUrl, cb, closeCb)
      }, reconnectionDelay)
      reconnectionDelay *= 2
    }
  }
  socket = new WebSocket(url)
  socket.binaryType = 'arraybuffer'
  socket.onopen = onopen
  socket.onmessage = onmessage
  socket.onerror = onerror
  socket.onclose = onclose
}

var reset = function() {
  reconnect = false
  reconnectionDelay = 1000 * 5
  reconnectAttempts = 0
  clearTimeout(reconncetTimer)
}

var sendMessage = function(reqId, route, type, msg) {
  if (useCrypto) {
    msg = JSON.stringify(msg)
    var sig = rsa.signString(msg, 'sha256')
    msg = JSON.parse(msg)
    msg['__crypto__'] = sig
  }
  var any = protos.googleAny.Any.fromPartial({})
  if (type) {
    any = protos.googleAny.Any.fromPartial({ typeUrl: type.getFullName(), value: type.encode(type.fromPartial(msg)).finish() })
  }
  const userStore = useUserStore()
  const data = protos.common.basicReq.fromPartial({ head: {
    token: userStore.token,
    userId: userStore.userInfo.ID,
  },
  data: any
  })
  msg = protos.common.basicReq.encode(data).finish()
  if (encode) {
    msg = encode(reqId, route, msg)
  }

  var packet = Package.encode(Package.TYPE_DATA, msg)
  send(packet)
}

var send = function(packet) {
  socket.send(packet.buffer)
}

var handler = {}

var heartbeat = function(data) {
  if (!heartbeatInterval) {
    // no heartbeat
    return
  }

  var obj = Package.encode(Package.TYPE_HEARTBEAT)
  if (heartbeatTimeoutId) {
    clearTimeout(heartbeatTimeoutId)
    heartbeatTimeoutId = null
  }

  if (heartbeatId) {
    // already in a heartbeat interval
    return
  }
  heartbeatId = setTimeout(function() {
    heartbeatId = null
    send(obj)

    nextHeartbeatTimeout = Date.now() + heartbeatTimeout
    heartbeatTimeoutId = setTimeout(heartbeatTimeoutCb, heartbeatTimeout)
  }, heartbeatInterval)
}

var heartbeatTimeoutCb = function() {
  var gap = nextHeartbeatTimeout - Date.now()
  if (gap > gapThreshold) {
    heartbeatTimeoutId = setTimeout(heartbeatTimeoutCb, gap)
  } else {
    ElMessage.error('server heartbeat timeout')
    starx.emit('heartbeat timeout')
    starx.disconnect()
  }
}

var handshake = function(data) {
  data = JSON.parse(Protocol.strdecode(data))
  if (data.code === RES_OLD_CLIENT) {
    starx.emit('error', 'client version not fullfill')
    return
  }

  if (data.code !== RES_OK) {
    starx.emit('error', 'handshake fail')
    return
  }

  handshakeInit(data)

  var obj = Package.encode(Package.TYPE_HANDSHAKE_ACK)
  send(obj)
  if (initCallback) {
    initCallback(socket)
  }
}

var onData = function(data) {
  var msg = data
  if (decode) {
    msg = decode(msg)
  }
  processMessage(starx, msg)
}

var onKick = function(data) {
  data = JSON.parse(Protocol.strdecode(data))
  starx.emit('onKick', data)
}

handlers[Package.TYPE_HANDSHAKE] = handshake
handlers[Package.TYPE_HEARTBEAT] = heartbeat
handlers[Package.TYPE_DATA] = onData
handlers[Package.TYPE_KICK] = onKick

var processPackage = function(msgs) {
  if (Array.isArray(msgs)) {
    for (var i = 0; i < msgs.length; i++) {
      var msg = msgs[i]
      handlers[msg.type](msg.body)
    }
  } else {
    handlers[msgs.type](msgs.body)
  }
}

var processMessage = function(starx, msg) {
  // console.log('processMessage', msg)
  if (!msg.id) {
    // server push message
    // console.log('server push message', msg.route, msg.body)
    starx.emit(msg.route, msg.body)
    return
  }

  // if have a id then find the callback function with the request
  var cb = callbacks[msg.id]

  delete callbacks[msg.id]
  if (typeof cb !== 'function') {
    return
  }
  // console.log('processMessage', msg)
  const message = converProto(msg.body)
  if (message) {
    cb(message)
  }
}

var processMessageBatch = function(starx, msgs) {
  // console.log('processMessageBatch', msgs)
  for (var i = 0, l = msgs.length; i < l; i++) {
    processMessage(starx, msgs[i])
  }
}

var deCompose = function(msg) {
  var route = msg.route

  // Decompose route from dict
  if (msg.compressRoute) {
    if (!abbrs[route]) {
      return {}
    }

    route = msg.route = abbrs[route]
  }

  if (isJson) {
    return JSON.parse(Protocol.strdecode(msg.body))
  }

  return msg.body
}

var handshakeInit = function(data) {
  if (data.sys && data.sys.heartbeat) {
    heartbeatInterval = data.sys.heartbeat * 1000 // heartbeat interval
    heartbeatTimeout = heartbeatInterval * 2 // max heartbeat timeout
  } else {
    heartbeatInterval = 0
    heartbeatTimeout = 0
  }

  initData(data)

  if (typeof handshakeCallback === 'function') {
    handshakeCallback(data.user)
  }
}

// Initilize data used in starx client
var initData = function(data) {
  if (!data || !data.sys) {
    return
  }
  dict = data.sys.dict

  // Init compress dict
  if (dict) {
    dict = dict
    abbrs = {}

    for (var route in dict) {
      abbrs[dict[route]] = route
    }
  }

  window.starx = starx
}

var converProto = (arg) => {
  if (arg.type) {
    ElMessage.error(arg.type)
    return null
  }
  // console.log(arg)
  var proto = protos.common.basicRep.decode(arg)
  if (proto.head) {
    const head = proto.head
    if (head.code !== 0) {
      ElMessage.error(head.msg)
      return null
    }
    if (proto.data) {
      const data = proto.data
      const path = data.typeUrl.split('/')[1].split('.')
      const type = protos[path[0]][path[1]]
      return type.decode(data.value)
    }
  }
  return null
}
