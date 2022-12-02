// Code generated by ent, DO NOT EDIT.

package project

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRepository holds the string denoting the repository field in the database.
	FieldRepository = "repository"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldStoreLocation holds the string denoting the storelocation field in the database.
	FieldStoreLocation = "store_location"
	// Table holds the table name of the project in the database.
	Table = "projects"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
	FieldRepository,
	FieldURL,
	FieldDomain,
	FieldStoreLocation,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
