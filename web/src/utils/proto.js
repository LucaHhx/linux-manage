import * as room from '@/protobuf/proto/room'
import * as common from '@/protobuf/proto/common'
import * as googleAny from '@/google/protobuf/any'
import * as resources from '@/protobuf/proto/resources'

export var protos = {
  room,
  common,
  googleAny,
  resources
}

for (const namespace in protos) {
  for (const key in protos[namespace]) {
    if (protos[namespace][key].encode) {
      protos[namespace][key].getFullName = () => {
        return `${protos[namespace].protobufPackage}.${key}`
      }
    }
  }
}

// console.log('Protos', protos)

