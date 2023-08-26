// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/tomato3713/nullwiki/pkg/ent/schema"
	"github.com/tomato3713/nullwiki/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
