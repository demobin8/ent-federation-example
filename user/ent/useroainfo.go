// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"user/ent/useroainfo"

	"entgo.io/ent/dialect/sql"
)

// UserOaInfo is the model entity for the UserOaInfo schema.
type UserOaInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailPwd holds the value of the "email_pwd" field.
	EmailPwd          string `json:"email_pwd,omitempty"`
	user_user_oa_info *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserOaInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case useroainfo.FieldID:
			values[i] = new(sql.NullInt64)
		case useroainfo.FieldEmail, useroainfo.FieldEmailPwd:
			values[i] = new(sql.NullString)
		case useroainfo.ForeignKeys[0]: // user_user_oa_info
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserOaInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserOaInfo fields.
func (uoi *UserOaInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case useroainfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uoi.ID = int(value.Int64)
		case useroainfo.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				uoi.Email = value.String
			}
		case useroainfo.FieldEmailPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_pwd", values[i])
			} else if value.Valid {
				uoi.EmailPwd = value.String
			}
		case useroainfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_user_oa_info", value)
			} else if value.Valid {
				uoi.user_user_oa_info = new(int)
				*uoi.user_user_oa_info = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserOaInfo.
// Note that you need to call UserOaInfo.Unwrap() before calling this method if this UserOaInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (uoi *UserOaInfo) Update() *UserOaInfoUpdateOne {
	return (&UserOaInfoClient{config: uoi.config}).UpdateOne(uoi)
}

// Unwrap unwraps the UserOaInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uoi *UserOaInfo) Unwrap() *UserOaInfo {
	tx, ok := uoi.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserOaInfo is not a transactional entity")
	}
	uoi.config.driver = tx.drv
	return uoi
}

// String implements the fmt.Stringer.
func (uoi *UserOaInfo) String() string {
	var builder strings.Builder
	builder.WriteString("UserOaInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", uoi.ID))
	builder.WriteString(", email=")
	builder.WriteString(uoi.Email)
	builder.WriteString(", email_pwd=")
	builder.WriteString(uoi.EmailPwd)
	builder.WriteByte(')')
	return builder.String()
}

// UserOaInfos is a parsable slice of UserOaInfo.
type UserOaInfos []*UserOaInfo

func (uoi UserOaInfos) config(cfg config) {
	for _i := range uoi {
		uoi[_i].config = cfg
	}
}