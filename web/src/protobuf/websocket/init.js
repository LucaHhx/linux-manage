
import { starx } from '@/utils/starx-wsclient'
import { joinResponse, allMembers, userMessage } from '@/protobuf/proto/room'

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
  console.log('onMembers', allMembers.decode(data))
  // v.messages.push({ name: 'system', content: 'members: ' + data.members })
}

starx.init({ host: location.hostname, port: 3250, path: '/nano' }, function() {
  // console.log('initialized')
  // starx.on('onMessage', onMessage)
  starx.on('onNewUser', onNewUser)
  starx.on('onMembers', onMembers)
  starx.request('room.join', {}, join)
})
