package mustekala

type memoryDB struct {
	nodes []node
}

type node struct {
	id       string
	ip       string
	tcp      string
	udp      string
	statuses []status
}

type status struct {
	timeStamp string
	Code      string
	Msg       string
	Info      string
}

func NewMemoryDB() *memoryDB {
	return &memoryDB{
		nodes: make([]node, 0),
	}
}

func (m *memoryDB) UpdateNodeStatus() {

}

func (m *memoryDB) DumpDB() {

}
