package schema

// MicrosoftForms .
type MicrosoftForms struct {
	FormsUserTypes     []FormsUserTypes `json:"FormsUserTypes"`
	SourceApp          *string          `json:"SourceApp"`
	FormName           *string          `json:"FormName,omitempty"`
	FormID             *string          `json:"FormId,omitempty"`
	FormTypes          []FormTypes      `json:"FormTypes,omitempty"`
	ActivityParameters *string          `json:"ActivityParameters,omitempty"`
}

// FormsUserTypes .
type FormsUserTypes int

// FormsUserTypes enum.
const (
	AdminUT FormsUserTypes = iota
	OwnerUT
	ResponderUT
	CoauthorUT
)

func (t FormsUserTypes) String() string {
	literals := map[FormsUserTypes]string{
		AdminUT:     "Admin",
		OwnerUT:     "Owner",
		ResponderUT: "Responder",
		CoauthorUT:  "Coauthor",
	}
	return literals[t]
}

// FormTypes .
type FormTypes int

// FormTypes enum.
const (
	Form FormTypes = iota
	Quiz
	Survey
)

func (t FormTypes) String() string {
	literals := map[FormTypes]string{
		Form:   "Form",
		Quiz:   "Quiz",
		Survey: "Survey",
	}
	return literals[t]
}
