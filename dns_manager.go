package dns_manager

type RecordType string

const (
	None  RecordType = "None essential"
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	MX    RecordType = "MX"
	CNAME RecordType = "CNAME"
	TXT   RecordType = "TXT"
)

type Record struct {
	Type    RecordType `json:"type,omitempty"`
	Name    string     `json:"name,omitempty"`
	Content string     `json:"content,omitempty"`
	TTL     uint       `json:"ttl,omitempty"`
}

type Actions interface {
	AddRecord(*Record) (*Response, error)
	DeleteRecord(*Record) (*Response, error)
	UpdateRecord(*Record) (*Response, error)
	GetRecords() ([]*Record, error)
}

type manager interface {
	New() (Actions, error)
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

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
