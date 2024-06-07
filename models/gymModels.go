package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleMember Role = "Member"
	RoleAdmin  Role = "Admin"
)

// User model for admins and members
type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Role         Role   `gorm:"not null"`
	Membership   Membership
	MembershipID uint
}

// Trainer model
type Trainer struct {
	gorm.Model
	FirstName string   `gorm:"not null"`
	LastName  string   `gorm:"not null"`
	ImageURL  string   `gorm:"not null"`
	Email     string   `gorm:"unique;not null"`
	Bio       string   `gorm:"not null"`
	Phone     string   `gorm:"not null"`
	Classes   []*Class `gorm:"many2many:trainer_classes"`
}

// Class to attend model
type Class struct {
	gorm.Model
	Name        string `gorm:"not null"`
	ImageURL    string
	Description string
	StartTime   string     `gorm:"not null"`
	EndTime     string     `gorm:"not null"`
	Trainers    []*Trainer `gorm:"many2many:trainer_classes"`
	Attendees   []*User    `gorm:"many2many:class_attendees"`
}

type MembershipType string

const (
	Basic   MembershipType = "Basic"
	Premium MembershipType = "Premium"
	VIP     MembershipType = "VIP"
)

type Membership struct {
	gorm.Model
	Type      MembershipType `gorm:"not null"`
	ImageURL  string         `gorm:"not null"`
	StartDate string         `gorm:"not null"`
	EndDate   string         `gorm:"not null"`
	UserID    uint
}

// Attendace model
type Attendance struct {
	gorm.Model
	Date    time.Time `gorm:"not null"`
	UserID  uint      `gorm:"not null"`
	ClassID uint      `gorm:"not null"`
	Present bool      `gorm:"not null"`
}
