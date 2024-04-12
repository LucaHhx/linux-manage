import { useUserStore } from '@/pinia/modules/user'
import { starx } from '@/utils/starx-wsclient'
import { protos } from '@/utils/proto'
import { getDownloadInfo } from '@/api/aTorrent'

var onMessage = function(msg) {
  console.log('onMessage', msg)
  // console.log('onMessage', protos.room.userMessage.decode(msg))
}

var join = function(data) {
  // data = protos.common.basicRep.decode(data)
  console.log('join', data)
  // console.log('join', joinResponse.decode(data))
  if (data.code === 0) {
    // v.messages.push({ name: 'system', content: data.result })

  }
}

var onNewUser = function(data) {
  // console.log('onNewUser', data)
  console.log('onNewUser', data)

  // console.log('onNewUser', data)
  // v.messages.push({ name: 'system', content: data.content })
}

var onMembers = function(data) {
  console.log('onMembers', data)
  // console.log('onMembers', protos.room.allMembers.decode(data))
  // v.messages.push({ name: 'system', content: 'members: ' + data.members })
}

export default {
  install: (app) => {
    starx.init({ host: location.hostname, port: 3250, path: '/socket' }, function() {
      starx.on('downloadInfo', getDownloadInfo)
      starx.on('onMessage', onMessage)
      starx.request('socket.create', join)
    })
  }
}

