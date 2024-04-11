import * as room from '@/protobuf/proto/room'
import * as common from '@/protobuf/proto/common'
import * as googleAny from '@/google/protobuf/any'

export var Protos = {
  room,
  common,
  googleAny
}

// function mergeNamespacesIntoProtos(namespace, namespaceName) {
//   for (const key in namespace) {
//     const fullName = `${namespaceName}.${key}`
//     Protos[fullName] = {
//       ...namespace[key],
//       getFullName: function() { return fullName } // 修改了方法名
//     }
//   }
// }

// mergeNamespacesIntoProtos(room, room.protobufPackage)
// mergeNamespacesIntoProtos(common, common.protobufPackage)
// mergeNamespacesIntoProtos(googleAny, googleAny.protobufPackage)

console.log('Protos', Protos)
