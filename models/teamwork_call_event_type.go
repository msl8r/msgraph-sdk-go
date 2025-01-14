package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type TeamworkCallEventType int

const (
    CALL_TEAMWORKCALLEVENTTYPE TeamworkCallEventType = iota
    MEETING_TEAMWORKCALLEVENTTYPE
    SCREENSHARE_TEAMWORKCALLEVENTTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKCALLEVENTTYPE
)

func (i TeamworkCallEventType) String() string {
    return []string{"call", "meeting", "screenShare", "unknownFutureValue"}[i]
}
func ParseTeamworkCallEventType(v string) (interface{}, error) {
    result := CALL_TEAMWORKCALLEVENTTYPE
    switch v {
        case "call":
            result = CALL_TEAMWORKCALLEVENTTYPE
        case "meeting":
            result = MEETING_TEAMWORKCALLEVENTTYPE
        case "screenShare":
            result = SCREENSHARE_TEAMWORKCALLEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKCALLEVENTTYPE
        default:
            return 0, errors.New("Unknown TeamworkCallEventType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkCallEventType(values []TeamworkCallEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
