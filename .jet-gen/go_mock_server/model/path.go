//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Path struct {
	ID          int32 `sql:"primary_key"`
	Path        string
	DefaultHost *string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
