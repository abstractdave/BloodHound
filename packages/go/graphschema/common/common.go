// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by Cuelang code gen. DO NOT EDIT!
// Cuelang source: github.com/specterops/bloodhound/-/tree/main/packages/cue/schemas/

package common

import (
	"errors"
	graph "github.com/specterops/bloodhound/dawgs/graph"
)

var (
	MigrationData = graph.StringKind("MigrationData")
	AllADAttacks  = graph.StringKind("ALL_AD_ATTACKS")
	AllAZAttacks  = graph.StringKind("ALL_AZ_ATTACKS")
)

type Property string

const (
	ObjectID        Property = "objectid"
	Name            Property = "name"
	DisplayName     Property = "displayname"
	Description     Property = "description"
	OwnerObjectID   Property = "owner_objectid"
	Collected       Property = "collected"
	OperatingSystem Property = "operatingsystem"
	SystemTags      Property = "system_tags"
	UserTags        Property = "user_tags"
	LastSeen        Property = "lastseen"
	WhenCreated     Property = "whencreated"
	Enabled         Property = "enabled"
	PasswordLastSet Property = "pwdlastset"
	Title           Property = "title"
	Email           Property = "email"
	IsInherited     Property = "isinherited"
)

func AllProperties() []Property {
	return []Property{ObjectID, Name, DisplayName, Description, OwnerObjectID, Collected, OperatingSystem, SystemTags, UserTags, LastSeen, WhenCreated, Enabled, PasswordLastSet, Title, Email, IsInherited}
}
func ParseProperty(source string) (Property, error) {
	switch source {
	case "objectid":
		return ObjectID, nil
	case "name":
		return Name, nil
	case "displayname":
		return DisplayName, nil
	case "description":
		return Description, nil
	case "owner_objectid":
		return OwnerObjectID, nil
	case "collected":
		return Collected, nil
	case "operatingsystem":
		return OperatingSystem, nil
	case "system_tags":
		return SystemTags, nil
	case "user_tags":
		return UserTags, nil
	case "lastseen":
		return LastSeen, nil
	case "whencreated":
		return WhenCreated, nil
	case "enabled":
		return Enabled, nil
	case "pwdlastset":
		return PasswordLastSet, nil
	case "title":
		return Title, nil
	case "email":
		return Email, nil
	case "isinherited":
		return IsInherited, nil
	default:
		return "", errors.New("Invalid enumeration value: " + source)
	}
}
func (s Property) String() string {
	switch s {
	case ObjectID:
		return string(ObjectID)
	case Name:
		return string(Name)
	case DisplayName:
		return string(DisplayName)
	case Description:
		return string(Description)
	case OwnerObjectID:
		return string(OwnerObjectID)
	case Collected:
		return string(Collected)
	case OperatingSystem:
		return string(OperatingSystem)
	case SystemTags:
		return string(SystemTags)
	case UserTags:
		return string(UserTags)
	case LastSeen:
		return string(LastSeen)
	case WhenCreated:
		return string(WhenCreated)
	case Enabled:
		return string(Enabled)
	case PasswordLastSet:
		return string(PasswordLastSet)
	case Title:
		return string(Title)
	case Email:
		return string(Email)
	case IsInherited:
		return string(IsInherited)
	default:
		return "Invalid enumeration case: " + string(s)
	}
}
func (s Property) Name() string {
	switch s {
	case ObjectID:
		return "Object ID"
	case Name:
		return "Name"
	case DisplayName:
		return "Display Name"
	case Description:
		return "Description"
	case OwnerObjectID:
		return "Owner Object ID"
	case Collected:
		return "Collected"
	case OperatingSystem:
		return "Operating System"
	case SystemTags:
		return "Node System Tags"
	case UserTags:
		return "Node User Tags"
	case LastSeen:
		return "Last Collected by BloodHound"
	case WhenCreated:
		return "Created"
	case Enabled:
		return "Enabled"
	case PasswordLastSet:
		return "Password Last Set"
	case Title:
		return "Title"
	case Email:
		return "Email"
	case IsInherited:
		return "Is Inherited"
	default:
		return "Invalid enumeration case: " + string(s)
	}
}
func (s Property) Is(others ...graph.Kind) bool {
	for _, other := range others {
		if value, err := ParseProperty(other.String()); err == nil && value == s {
			return true
		}
	}
	return false
}
func Nodes() []graph.Kind {
	return []graph.Kind{MigrationData}
}
func Relationships() []graph.Kind {
	return []graph.Kind{AllADAttacks, AllAZAttacks}
}
func NodeKinds() []graph.Kind {
	return []graph.Kind{MigrationData}
}
