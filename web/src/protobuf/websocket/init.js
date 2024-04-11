import { useUserStore } from '@/pinia/modules/user'
// import { starx } from '@/utils/starx-wsclient'
import { websocket } from '@/protobuf/websocket/register'
import { Protos } from '@/utils/proto'

websocket.add('ss')
var onMessage = function(msg) {
  // console.log('onMessage', msg)
  console.log('onMessage', userMessage.decode(msg))
}

var join = function(data) {
  data = joinResponse.decode(data)
  // console.log('join', joinResponse.decode(data))
  if (data.code === 0) {
    // v.messages.push({ name: 'system', content: data.result })
    starx.on('onMessage', onMessage)
  }
}

var onNewUser = function(data) {
  console.log('onNewUser', data)
  // v.messages.push({ name: 'system', content: data.content })
}

var onMembers = function(data) {
  console.log('onMembers', Protos.room.allMembers.decode(data))
  // v.messages.push({ name: 'system', content: 'members: ' + data.members })
}

export default {
  install: (app) => {
    const userStore = useUserStore()
    // var userMessage = Protos['room.userMessage']
    // var BasicReq = Protos['common.basicReq']
    // var Any = Protos['google.protobuf.Any']
    console.log(typeof (Protos.room.userMessage).fullName)
    const userMessageInfo = Protos.room.userMessage.encode({ name: 'Huang', content: 'hello' })
    console.log('userStore.userInfo.ID', userStore.userInfo.ID)
    const aa = Protos.common.basicReq.encode({ head: {
      token: userStore.token,
      userId: userStore.userInfo.ID,
    },
    data: Protos.googleAny.Any.fromPartial({ typeUrl: Protos.room.userMessage.$type.fullName, value: userMessageInfo.finish() })
    })

    starx.init({ host: location.hostname, port: 3250, path: '/nano' }, function() {
      // console.log('initialized')
      // starx.on('onMessage', onMessage)
      starx.on('onNewUser', onNewUser)
      starx.on('onMembers', onMembers)
      starx.request('room.join', aa.finish(), join)
    })
  }
}

