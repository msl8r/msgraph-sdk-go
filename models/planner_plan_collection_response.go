package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanCollectionResponse
type PlannerPlanCollectionResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	// The OdataNextLink property
	odataNextLink *string
	// The value property
	value []PlannerPlanable
}

// NewPlannerPlanCollectionResponse instantiates a new PlannerPlanCollectionResponse and sets the default values.
func NewPlannerPlanCollectionResponse() *PlannerPlanCollectionResponse {
	m := &PlannerPlanCollectionResponse{}
	m.SetAdditionalData(make(map[string]interface{}))
	return m
}

// CreatePlannerPlanCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPlannerPlanCollectionResponse(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanCollectionResponse) GetAdditionalData() map[string]interface{} {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["@odata.nextLink"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOdataNextLink(val)
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]PlannerPlanable, len(val))
			for i, v := range val {
				res[i] = v.(PlannerPlanable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *PlannerPlanCollectionResponse) GetOdataNextLink() *string {
	return m.odataNextLink
}

// GetValue gets the value property value. The value property
func (m *PlannerPlanCollectionResponse) GetValue() []PlannerPlanable {
	return m.value
}

// Serialize serializes information the current object
func (m *PlannerPlanCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
		if err != nil {
			return err
		}
	}
	if m.GetValue() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
		for i, v := range m.GetValue() {
			cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
		}
		err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *PlannerPlanCollectionResponse) SetAdditionalData(value map[string]interface{}) {
	m.additionalData = value
}

// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *PlannerPlanCollectionResponse) SetOdataNextLink(value *string) {
	m.odataNextLink = value
}

// SetValue sets the value property value. The value property
func (m *PlannerPlanCollectionResponse) SetValue(value []PlannerPlanable) {
	m.value = value
}