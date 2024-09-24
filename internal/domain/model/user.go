package model

import (
	"time"

	"github.com/sofisoft-tech/go-common/crypto"
)

type User struct {
	ID                     string `json:"_key,omitempty"`
	WorkspaceId            string `json:"workspaceId,omitempty"`
	UserName               string `json:"userName,omitempty"`
	Email                  string `json:"email,omitempty"`
	NormalizedUserName     string `json:"normalizedUserName,omitempty"`
	NormalizedEmail        string `json:"normalizedEmail,omitempty"`
	PasswordHash           string `json:"passwordHash,omitempty"`
	FirstName              string `json:"firstName,omitempty"`
	LastName               string `json:"lastName,omitempty"`
	Alias                  string `json:"alias,omitempty"`
	ImageUri               string `json:"imageUri,omitempty"`
	AccessFailedCount      int16  `json:"accessFailedCount,omitempty"`
	LockoutEnabled         bool   `json:"lockoutEnabled,omitempty"`
	LockoutEnd             *int64 `json:"lockoutEnd,omitempty"`
	PasswordExpiresEnabled bool   `json:"passwordExpiresEnabled,omitempty"`
	PasswordExpires        *int64 `json:"passwordExpires,omitempty"`
	RequestPasswordChange  bool   `json:"requestPasswordChange,omitempty"`
	SecurityStamp          string `json:"securityStamp,omitempty"`
	Active                 bool   `json:"active,omitempty"`
	CreatedAt              int64  `json:"createdAt,omitempty"`
	CreatedBy              string `json:"createdBy,omitempty"`
	ModifiedAt             *int64 `json:"modifiedAt,omitempty"`
	ModifiedBy             string `json:"modifiedBy,omitempty"`
}

func (User) GetCollectionName() string {
	return "Users"
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) IsLockedOut() bool {
	if !u.LockoutEnabled || u.LockoutEnd == nil {
		return false
	}

	return time.Now().Unix()-*u.LockoutEnd <= 0
}

func (u *User) SetPasswordHash(password string) (err error) {
	u.PasswordHash, err = crypto.HashPassword(password)
	if err != nil {
		return err
	}

	if u.PasswordExpiresEnabled {
		passwordExpires := time.Now().Add(time.Hour * 2190).UTC().Unix()
		u.PasswordExpires = &passwordExpires
	}

	return nil
}

func (u User) Validate() bool {
	fieldsPass := u.Email != "" && u.UserName != "" &&
		u.FirstName != "" && u.LastName != "" &&
		u.PasswordHash != "" && u.WorkspaceId != ""

	if fieldsPass {
		return true
	} else {
		return false
	}
}
