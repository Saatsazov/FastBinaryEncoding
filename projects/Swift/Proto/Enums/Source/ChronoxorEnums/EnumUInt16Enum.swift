// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

import Foundation

public enum EnumUInt16Enum {
    typealias RawValue = UInt16
    case ENUM_VALUE_0
    case ENUM_VALUE_1
    case ENUM_VALUE_2
    case ENUM_VALUE_3
    case ENUM_VALUE_4
    case ENUM_VALUE_5

    var rawValue: RawValue {
        switch self {
        case .ENUM_VALUE_0: return 0 + 0
        case .ENUM_VALUE_1: return 0 + 0
        case .ENUM_VALUE_2: return 0 + 1
        case .ENUM_VALUE_3: return 65534 + 0
        case .ENUM_VALUE_4: return 65534 + 1
        case .ENUM_VALUE_5: return Self.ENUM_VALUE_3.rawValue
        }
    }

    init(value: UInt8) { self = EnumUInt16Enum(rawValue: NSNumber(value: value).uint16Value) }
    init(value: UInt16) { self = EnumUInt16Enum(rawValue: NSNumber(value: value).uint16Value) }
    init(value: UInt32) { self = EnumUInt16Enum(rawValue: NSNumber(value: value).uint16Value) }
    init(value: UInt64) { self = EnumUInt16Enum(rawValue: NSNumber(value: value).uint16Value) }
    init(value: EnumUInt16Enum) { self = EnumUInt16Enum(rawValue: value.rawValue) }
    init(rawValue: UInt16) { self = Self.mapValue(value: rawValue)! }

    var description: String {
        switch self {
        case .ENUM_VALUE_0: return "ENUM_VALUE_0"
        case .ENUM_VALUE_1: return "ENUM_VALUE_1"
        case .ENUM_VALUE_2: return "ENUM_VALUE_2"
        case .ENUM_VALUE_3: return "ENUM_VALUE_3"
        case .ENUM_VALUE_4: return "ENUM_VALUE_4"
        case .ENUM_VALUE_5: return "ENUM_VALUE_5"
        }
    }

    static let rawValuesMap: [RawValue: EnumUInt16Enum] = {
        var value = [RawValue: EnumUInt16Enum]()
        value[EnumUInt16Enum.ENUM_VALUE_0.rawValue] = .ENUM_VALUE_0
        value[EnumUInt16Enum.ENUM_VALUE_1.rawValue] = .ENUM_VALUE_1
        value[EnumUInt16Enum.ENUM_VALUE_2.rawValue] = .ENUM_VALUE_2
        value[EnumUInt16Enum.ENUM_VALUE_3.rawValue] = .ENUM_VALUE_3
        value[EnumUInt16Enum.ENUM_VALUE_4.rawValue] = .ENUM_VALUE_4
        value[EnumUInt16Enum.ENUM_VALUE_5.rawValue] = .ENUM_VALUE_5
        return value
    }()

    static func mapValue(value: UInt16) -> EnumUInt16Enum? {
        return rawValuesMap[value]
    }
}
