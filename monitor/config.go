package monitor

type config []serverInfo

type serverInfo struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    uint16 `json:"port"`
	Network string `json:"network"`
	Repeat  uint   `json:"repeat"`
	Timeout uint  `json:"timeout"`
	Peroid  uint  `json:"peroid"`
}
