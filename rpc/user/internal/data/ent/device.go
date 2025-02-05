// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"user/internal/data/ent/device"

	"entgo.io/ent/dialect/sql"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 创建时间
	Created time.Time `json:"created,omitempty"`
	// 更新时间
	Updated time.Time `json:"updated,omitempty"`
	// DeviceID holds the value of the "device_id" field.
	DeviceID string `json:"device_id,omitempty"`
	// AppVersion holds the value of the "app_version" field.
	AppVersion int32 `json:"app_version,omitempty"`
	// OsVersion holds the value of the "os_version" field.
	OsVersion string `json:"os_version,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldID, device.FieldAppVersion:
			values[i] = new(sql.NullInt64)
		case device.FieldDeviceID, device.FieldOsVersion:
			values[i] = new(sql.NullString)
		case device.FieldCreated, device.FieldUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Device", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int32(value.Int64)
		case device.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				d.Created = value.Time
			}
		case device.FieldUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated", values[i])
			} else if value.Valid {
				d.Updated = value.Time
			}
		case device.FieldDeviceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				d.DeviceID = value.String
			}
		case device.FieldAppVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_version", values[i])
			} else if value.Valid {
				d.AppVersion = int32(value.Int64)
			}
		case device.FieldOsVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os_version", values[i])
			} else if value.Valid {
				d.OsVersion = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created=")
	builder.WriteString(d.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated=")
	builder.WriteString(d.Updated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(d.DeviceID)
	builder.WriteString(", ")
	builder.WriteString("app_version=")
	builder.WriteString(fmt.Sprintf("%v", d.AppVersion))
	builder.WriteString(", ")
	builder.WriteString("os_version=")
	builder.WriteString(d.OsVersion)
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
