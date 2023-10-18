package schema

// Quarantine .
type Quarantine struct {
	AuditRecord
	RequestType      *RequestType   `json:"RequestType,omitempty"`
	RequestSource    *RequestSource `json:"RequestSource,omitempty"`
	NetworkMessageID *string        `json:"NetworkMessageId,omitempty"`
	ReleaseTo        *string        `json:"ReleaseTo,omitempty"`
}

// RequestType .
type RequestType int

// RequestType enum.
const (
	Preview RequestType = iota
	Delete
	Release
	Export
	ViewHeader
)

func (t RequestType) String() string {
	literals := map[RequestType]string{
		Preview:    "Preview",
		Delete:     "Delete",
		Release:    "Release",
		Export:     "Export",
		ViewHeader: "ViewHeader",
	}
	return literals[t]
}

// RequestSource .
type RequestSource int

// RequestSource enum.
const (
	SCC RequestSource = iota
	Cmdlet
	URLlink
)

func (t RequestSource) String() string {
	literals := map[RequestSource]string{
		SCC:     "SCC",
		Cmdlet:  "Cmdlet",
		URLlink: "URLlink",
	}
	return literals[t]
}
