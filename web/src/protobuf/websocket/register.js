import { starx } from '@/utils/starx-wsclient'
export var websocket = {
  notify: function(route, data) {
    starx.notify(route, data)
  }
}

websocket.add = function(name) {
  console.log('add', name)
}
