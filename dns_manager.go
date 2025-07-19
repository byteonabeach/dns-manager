package dns_manager

type RecordType string

// Defining supported record types
const (
	None  RecordType = "None essential"
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	MX    RecordType = "MX"
	CNAME RecordType = "CNAME"
	TXT   RecordType = "TXT"
)

// This is the DNS record structure.
// Ready to be marshaled into JSON
type Record struct {
	Type    RecordType `json:"type,omitempty"`
	Name    string     `json:"name,omitempty"`
	Content string     `json:"content,omitempty"`
	TTL     uint       `json:"ttl,omitempty"`
}

// This is representation of simple API response
// to driver query
type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// The core of library<->drivers architecture
// All drivers must implement those methods
type Actions interface {
	AddRecord(*Record) (*Response, error)
	DeleteRecord(*Record) (*Response, error)
	UpdateRecord(*Record) (*Response, error)
	GetRecords() ([]*Record, error)
}

// This is local scope interface that defines
// method New() which must be implemented by all
// libraries. The New() represents fabric method.
type manager interface {
	New() (Actions, error)
}

// Naive method that checks if record type is
// supported by this library.
func (rt RecordType) Check() bool {
	switch rt {
	case None, A, AAAA, MX, CNAME, TXT:
		return true
	default:
		return false
	}
}

func New(manager_config manager) (Actions, error) {
	return manager_config.New()
}
