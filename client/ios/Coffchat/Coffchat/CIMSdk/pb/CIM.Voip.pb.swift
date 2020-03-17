// DO NOT EDIT.
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: CIM.Voip.proto
//
// For information on using the generated types, please see the documenation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that your are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

/// 音视频呼叫邀请
struct CIM_Voip_CIMVoipInviteReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// cmd id:		0x0401
  var creatorUserID: UInt64 {
    get {return _storage._creatorUserID}
    set {_uniqueStorage()._creatorUserID = newValue}
  }

  /// 被邀请方列表，如果是一对一，无需设置
  var inviteUserList: [UInt64] {
    get {return _storage._inviteUserList}
    set {_uniqueStorage()._inviteUserList = newValue}
  }

  /// 通话类型
  var inviteType: CIM_Def_CIMVoipInviteType {
    get {return _storage._inviteType}
    set {_uniqueStorage()._inviteType = newValue}
  }

  /// 频道信息，由服务端分配
  var channelInfo: CIM_Def_CIMChannelInfo {
    get {return _storage._channelInfo ?? CIM_Def_CIMChannelInfo()}
    set {_uniqueStorage()._channelInfo = newValue}
  }
  /// Returns true if `channelInfo` has been explicitly set.
  var hasChannelInfo: Bool {return _storage._channelInfo != nil}
  /// Clears the value of `channelInfo`. Subsequent reads from it will return its default value.
  mutating func clearChannelInfo() {_uniqueStorage()._channelInfo = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

/// 音视频呼叫应答状态
struct CIM_Voip_CIMVoipInviteReply {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// cmd id:		0x0402
  var userID: UInt64 {
    get {return _storage._userID}
    set {_uniqueStorage()._userID = newValue}
  }

  /// 当前状态
  var rspCode: CIM_Def_CIMInviteRspCode {
    get {return _storage._rspCode}
    set {_uniqueStorage()._rspCode = newValue}
  }

  /// 频道信息，由服务端分配
  var channelInfo: CIM_Def_CIMChannelInfo {
    get {return _storage._channelInfo ?? CIM_Def_CIMChannelInfo()}
    set {_uniqueStorage()._channelInfo = newValue}
  }
  /// Returns true if `channelInfo` has been explicitly set.
  var hasChannelInfo: Bool {return _storage._channelInfo != nil}
  /// Clears the value of `channelInfo`. Subsequent reads from it will return its default value.
  mutating func clearChannelInfo() {_uniqueStorage()._channelInfo = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

/// 音视频呼叫ACK
/// 100 Trying->180 Ringing->200 Ok->ACK(this message)
struct CIM_Voip_CIMVoipInviteReplyAck {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 频道信息，由服务端分配
  var channelInfo: CIM_Def_CIMChannelInfo {
    get {return _storage._channelInfo ?? CIM_Def_CIMChannelInfo()}
    set {_uniqueStorage()._channelInfo = newValue}
  }
  /// Returns true if `channelInfo` has been explicitly set.
  var hasChannelInfo: Bool {return _storage._channelInfo != nil}
  /// Clears the value of `channelInfo`. Subsequent reads from it will return its default value.
  mutating func clearChannelInfo() {_uniqueStorage()._channelInfo = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

/// 心跳
struct CIM_Voip_CIMVoipHeartbeat {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

/// 挂断请求
struct CIM_Voip_CIMVoipByeReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 本地通话时长计数
  var localCallTimeLen: UInt64 {
    get {return _storage._localCallTimeLen}
    set {_uniqueStorage()._localCallTimeLen = newValue}
  }

  /// 挂断方用户ID
  var userID: UInt64 {
    get {return _storage._userID}
    set {_uniqueStorage()._userID = newValue}
  }

  /// 频道信息
  var channelInfo: CIM_Def_CIMChannelInfo {
    get {return _storage._channelInfo ?? CIM_Def_CIMChannelInfo()}
    set {_uniqueStorage()._channelInfo = newValue}
  }
  /// Returns true if `channelInfo` has been explicitly set.
  var hasChannelInfo: Bool {return _storage._channelInfo != nil}
  /// Clears the value of `channelInfo`. Subsequent reads from it will return its default value.
  mutating func clearChannelInfo() {_uniqueStorage()._channelInfo = nil}

  /// 挂断原因
  var byeReason: CIM_Def_CIMVoipByeReason {
    get {return _storage._byeReason}
    set {_uniqueStorage()._byeReason = newValue}
  }

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

/// 挂断响应
struct CIM_Voip_CIMVoipByeRsp {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 结果
  var errorCode: CIM_Def_CIMErrorCode = .kCimErrUnknown

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

/// 挂断通知
struct CIM_Voip_CIMVoipByeNotify {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 挂断方用户ID
  var userID: UInt64 {
    get {return _storage._userID}
    set {_uniqueStorage()._userID = newValue}
  }

  /// 频道信息
  var channelInfo: CIM_Def_CIMChannelInfo {
    get {return _storage._channelInfo ?? CIM_Def_CIMChannelInfo()}
    set {_uniqueStorage()._channelInfo = newValue}
  }
  /// Returns true if `channelInfo` has been explicitly set.
  var hasChannelInfo: Bool {return _storage._channelInfo != nil}
  /// Clears the value of `channelInfo`. Subsequent reads from it will return its default value.
  mutating func clearChannelInfo() {_uniqueStorage()._channelInfo = nil}

  /// 挂断原因
  var byeReason: CIM_Def_CIMVoipByeReason {
    get {return _storage._byeReason}
    set {_uniqueStorage()._byeReason = newValue}
  }

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "CIM.Voip"

extension CIM_Voip_CIMVoipInviteReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipInviteReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "creator_user_id"),
    2: .standard(proto: "invite_user_list"),
    3: .standard(proto: "invite_type"),
    4: .standard(proto: "channel_info"),
  ]

  fileprivate class _StorageClass {
    var _creatorUserID: UInt64 = 0
    var _inviteUserList: [UInt64] = []
    var _inviteType: CIM_Def_CIMVoipInviteType = .kCimVoipInviteTypeUnknown
    var _channelInfo: CIM_Def_CIMChannelInfo? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _creatorUserID = source._creatorUserID
      _inviteUserList = source._inviteUserList
      _inviteType = source._inviteType
      _channelInfo = source._channelInfo
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularUInt64Field(value: &_storage._creatorUserID)
        case 2: try decoder.decodeRepeatedUInt64Field(value: &_storage._inviteUserList)
        case 3: try decoder.decodeSingularEnumField(value: &_storage._inviteType)
        case 4: try decoder.decodeSingularMessageField(value: &_storage._channelInfo)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if _storage._creatorUserID != 0 {
        try visitor.visitSingularUInt64Field(value: _storage._creatorUserID, fieldNumber: 1)
      }
      if !_storage._inviteUserList.isEmpty {
        try visitor.visitPackedUInt64Field(value: _storage._inviteUserList, fieldNumber: 2)
      }
      if _storage._inviteType != .kCimVoipInviteTypeUnknown {
        try visitor.visitSingularEnumField(value: _storage._inviteType, fieldNumber: 3)
      }
      if let v = _storage._channelInfo {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipInviteReq, rhs: CIM_Voip_CIMVoipInviteReq) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._creatorUserID != rhs_storage._creatorUserID {return false}
        if _storage._inviteUserList != rhs_storage._inviteUserList {return false}
        if _storage._inviteType != rhs_storage._inviteType {return false}
        if _storage._channelInfo != rhs_storage._channelInfo {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipInviteReply: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipInviteReply"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "user_id"),
    2: .standard(proto: "rsp_code"),
    3: .standard(proto: "channel_info"),
  ]

  fileprivate class _StorageClass {
    var _userID: UInt64 = 0
    var _rspCode: CIM_Def_CIMInviteRspCode = .kCimVoipInviteCodeUnknown
    var _channelInfo: CIM_Def_CIMChannelInfo? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _userID = source._userID
      _rspCode = source._rspCode
      _channelInfo = source._channelInfo
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularUInt64Field(value: &_storage._userID)
        case 2: try decoder.decodeSingularEnumField(value: &_storage._rspCode)
        case 3: try decoder.decodeSingularMessageField(value: &_storage._channelInfo)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if _storage._userID != 0 {
        try visitor.visitSingularUInt64Field(value: _storage._userID, fieldNumber: 1)
      }
      if _storage._rspCode != .kCimVoipInviteCodeUnknown {
        try visitor.visitSingularEnumField(value: _storage._rspCode, fieldNumber: 2)
      }
      if let v = _storage._channelInfo {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipInviteReply, rhs: CIM_Voip_CIMVoipInviteReply) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._userID != rhs_storage._userID {return false}
        if _storage._rspCode != rhs_storage._rspCode {return false}
        if _storage._channelInfo != rhs_storage._channelInfo {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipInviteReplyAck: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipInviteReplyAck"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "channel_info"),
  ]

  fileprivate class _StorageClass {
    var _channelInfo: CIM_Def_CIMChannelInfo? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _channelInfo = source._channelInfo
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularMessageField(value: &_storage._channelInfo)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if let v = _storage._channelInfo {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipInviteReplyAck, rhs: CIM_Voip_CIMVoipInviteReplyAck) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._channelInfo != rhs_storage._channelInfo {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipHeartbeat: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipHeartbeat"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipHeartbeat, rhs: CIM_Voip_CIMVoipHeartbeat) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipByeReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipByeReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "local_call_time_len"),
    2: .standard(proto: "user_id"),
    3: .standard(proto: "channel_info"),
    4: .standard(proto: "bye_reason"),
  ]

  fileprivate class _StorageClass {
    var _localCallTimeLen: UInt64 = 0
    var _userID: UInt64 = 0
    var _channelInfo: CIM_Def_CIMChannelInfo? = nil
    var _byeReason: CIM_Def_CIMVoipByeReason = .kCimVoipByeReasonUnknown

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _localCallTimeLen = source._localCallTimeLen
      _userID = source._userID
      _channelInfo = source._channelInfo
      _byeReason = source._byeReason
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularUInt64Field(value: &_storage._localCallTimeLen)
        case 2: try decoder.decodeSingularUInt64Field(value: &_storage._userID)
        case 3: try decoder.decodeSingularMessageField(value: &_storage._channelInfo)
        case 4: try decoder.decodeSingularEnumField(value: &_storage._byeReason)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if _storage._localCallTimeLen != 0 {
        try visitor.visitSingularUInt64Field(value: _storage._localCallTimeLen, fieldNumber: 1)
      }
      if _storage._userID != 0 {
        try visitor.visitSingularUInt64Field(value: _storage._userID, fieldNumber: 2)
      }
      if let v = _storage._channelInfo {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
      }
      if _storage._byeReason != .kCimVoipByeReasonUnknown {
        try visitor.visitSingularEnumField(value: _storage._byeReason, fieldNumber: 4)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipByeReq, rhs: CIM_Voip_CIMVoipByeReq) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._localCallTimeLen != rhs_storage._localCallTimeLen {return false}
        if _storage._userID != rhs_storage._userID {return false}
        if _storage._channelInfo != rhs_storage._channelInfo {return false}
        if _storage._byeReason != rhs_storage._byeReason {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipByeRsp: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipByeRsp"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "error_code"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeSingularEnumField(value: &self.errorCode)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.errorCode != .kCimErrUnknown {
      try visitor.visitSingularEnumField(value: self.errorCode, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipByeRsp, rhs: CIM_Voip_CIMVoipByeRsp) -> Bool {
    if lhs.errorCode != rhs.errorCode {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension CIM_Voip_CIMVoipByeNotify: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".CIMVoipByeNotify"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "user_id"),
    2: .standard(proto: "channel_info"),
    3: .same(proto: "byeReason"),
  ]

  fileprivate class _StorageClass {
    var _userID: UInt64 = 0
    var _channelInfo: CIM_Def_CIMChannelInfo? = nil
    var _byeReason: CIM_Def_CIMVoipByeReason = .kCimVoipByeReasonUnknown

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _userID = source._userID
      _channelInfo = source._channelInfo
      _byeReason = source._byeReason
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularUInt64Field(value: &_storage._userID)
        case 2: try decoder.decodeSingularMessageField(value: &_storage._channelInfo)
        case 3: try decoder.decodeSingularEnumField(value: &_storage._byeReason)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if _storage._userID != 0 {
        try visitor.visitSingularUInt64Field(value: _storage._userID, fieldNumber: 1)
      }
      if let v = _storage._channelInfo {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
      }
      if _storage._byeReason != .kCimVoipByeReasonUnknown {
        try visitor.visitSingularEnumField(value: _storage._byeReason, fieldNumber: 3)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: CIM_Voip_CIMVoipByeNotify, rhs: CIM_Voip_CIMVoipByeNotify) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._userID != rhs_storage._userID {return false}
        if _storage._channelInfo != rhs_storage._channelInfo {return false}
        if _storage._byeReason != rhs_storage._byeReason {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}