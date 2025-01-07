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
  Type RecordType
  Host string
  Content string
  TTL uint16
}

type Actions interface {
  AddRecord(Record) response 
  DeleteRecord() response
  UpdateRecord(Record) response
  ReadRecords() (error, *[]Record)
}

type Manager interface {
  New() *Actions
}

type response struct {
  error string
  message string
}

func New(manager_config Manager) *Actions {
  return manager_config.New() 
}
