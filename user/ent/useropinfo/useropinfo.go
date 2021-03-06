// Code generated by entc, DO NOT EDIT.

package useropinfo

const (
	// Label holds the string label denoting the useropinfo type in the database.
	Label = "user_op_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// Table holds the table name of the useropinfo in the database.
	Table = "user_op_infos"
)

// Columns holds all SQL columns for useropinfo fields.
var Columns = []string{
	FieldID,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_op_infos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_user_op_info",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
