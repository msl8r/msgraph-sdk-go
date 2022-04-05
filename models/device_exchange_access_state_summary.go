package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceExchangeAccessStateSummary device Exchange Access State summary
type DeviceExchangeAccessStateSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total count of devices with Exchange Access State: Allowed.
    allowedDeviceCount *int32;
    // Total count of devices with Exchange Access State: Blocked.
    blockedDeviceCount *int32;
    // Total count of devices with Exchange Access State: Quarantined.
    quarantinedDeviceCount *int32;
    // Total count of devices for which no Exchange Access State could be found.
    unavailableDeviceCount *int32;
    // Total count of devices with Exchange Access State: Unknown.
    unknownDeviceCount *int32;
}
// NewDeviceExchangeAccessStateSummary instantiates a new deviceExchangeAccessStateSummary and sets the default values.
func NewDeviceExchangeAccessStateSummary()(*DeviceExchangeAccessStateSummary) {
    m := &DeviceExchangeAccessStateSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceExchangeAccessStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceExchangeAccessStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceExchangeAccessStateSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceExchangeAccessStateSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedDeviceCount gets the allowedDeviceCount property value. Total count of devices with Exchange Access State: Allowed.
func (m *DeviceExchangeAccessStateSummary) GetAllowedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.allowedDeviceCount
    }
}
// GetBlockedDeviceCount gets the blockedDeviceCount property value. Total count of devices with Exchange Access State: Blocked.
func (m *DeviceExchangeAccessStateSummary) GetBlockedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blockedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceExchangeAccessStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedDeviceCount(val)
        }
        return nil
    }
    res["blockedDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockedDeviceCount(val)
        }
        return nil
    }
    res["quarantinedDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantinedDeviceCount(val)
        }
        return nil
    }
    res["unavailableDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnavailableDeviceCount(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetQuarantinedDeviceCount gets the quarantinedDeviceCount property value. Total count of devices with Exchange Access State: Quarantined.
func (m *DeviceExchangeAccessStateSummary) GetQuarantinedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.quarantinedDeviceCount
    }
}
// GetUnavailableDeviceCount gets the unavailableDeviceCount property value. Total count of devices for which no Exchange Access State could be found.
func (m *DeviceExchangeAccessStateSummary) GetUnavailableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unavailableDeviceCount
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Total count of devices with Exchange Access State: Unknown.
func (m *DeviceExchangeAccessStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// Serialize serializes information the current object
func (m *DeviceExchangeAccessStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("allowedDeviceCount", m.GetAllowedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("blockedDeviceCount", m.GetBlockedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("quarantinedDeviceCount", m.GetQuarantinedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unavailableDeviceCount", m.GetUnavailableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceExchangeAccessStateSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedDeviceCount sets the allowedDeviceCount property value. Total count of devices with Exchange Access State: Allowed.
func (m *DeviceExchangeAccessStateSummary) SetAllowedDeviceCount(value *int32)() {
    if m != nil {
        m.allowedDeviceCount = value
    }
}
// SetBlockedDeviceCount sets the blockedDeviceCount property value. Total count of devices with Exchange Access State: Blocked.
func (m *DeviceExchangeAccessStateSummary) SetBlockedDeviceCount(value *int32)() {
    if m != nil {
        m.blockedDeviceCount = value
    }
}
// SetQuarantinedDeviceCount sets the quarantinedDeviceCount property value. Total count of devices with Exchange Access State: Quarantined.
func (m *DeviceExchangeAccessStateSummary) SetQuarantinedDeviceCount(value *int32)() {
    if m != nil {
        m.quarantinedDeviceCount = value
    }
}
// SetUnavailableDeviceCount sets the unavailableDeviceCount property value. Total count of devices for which no Exchange Access State could be found.
func (m *DeviceExchangeAccessStateSummary) SetUnavailableDeviceCount(value *int32)() {
    if m != nil {
        m.unavailableDeviceCount = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Total count of devices with Exchange Access State: Unknown.
func (m *DeviceExchangeAccessStateSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}