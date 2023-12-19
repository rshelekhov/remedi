package models

import (
	"github.com/google/uuid"
	"time"
)

type MedicalReport struct {
	ID              uuid.UUID
	Diagnosis       string
	Recommendations string
	AppointmentID   uuid.UUID
	Attachments     []Attachment
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Attachment struct {
	ID              uuid.UUID
	FileName        string
	URL             string
	AttachmentSize  string
	MedicalReportID uuid.UUID
	AttachedByID    uuid.UUID
	AttachedAt      time.Time
	UpdatedAt       time.Time
}
