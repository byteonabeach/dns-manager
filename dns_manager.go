package dns_manager

type RecordType int

const (
	None RecordType = iota
	A
	AAAA
	MX
	CNAME
	TXT
)

type Record struct {
	Type    RecordType
	Host    string
	Content string
	TTL     uint16
}

type Actions interface {
	AddRecord(Record) Response
	DeleteRecord() Response
	UpdateRecord(Record) Response
	ReadRecords() (error, *[]Record)
}

type manager interface {
	New() Actions
}

type Response struct {
	error   string
	message string
}

func New(manager_config manager) Actions {
	return manager_config.New()
}
