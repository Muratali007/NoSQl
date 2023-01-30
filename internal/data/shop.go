package data

type Cpu struct {
	ID              string    `bson:"_id"`
	Manufacturer    string    `bson:"manufacturer"`
	Model           string    `bson:"model"`
	Cores           int       `bson:"cores"`
	Threads         int       `bson:"threads"`
	ClockFrequency  []float64 `bson:"clock_frequency"`
	Socket          string    `bson:"socket"`
	RamChannels     int       `bson:"ram_channels"`
	RamTypes        []string  `bson:"ram_types"`
	MaxRamFrequency []int     `bson:"max_ram_frequency"`
	MaxRamCapacity  int       `bson:"max_ram_capacity"`
	Tdp             int       `bson:"tdp"`
	PciE            int       `bson:"pci-e"`
}

type Motherboard struct {
	ID              string     `bson:"_id"`
	Manufacturer    string     `bson:"manufacturer"`
	Model           string     `bson:"model"`
	Socket          string     `bson:"socket"`
	Chipset         string     `bson:"chipset"`
	FormFactor      string     `bson:"form_factor"`
	RamSlots        int        `bson:"ram_slots"`
	RamTypes        string     `bson:"ram_types"`
	MaxRamFrequency int        `bson:"max_ram_frequency"`
	MaxRam          int        `bson:"max_ram"`
	Interfaces      interfaces `bson:"interfaces"`
	PciStandard     int        `bson:"pci_standard"`
	MbPower         int        `bson:"mb_power"`
	CpuPower        int        `bson:"cpu_power"`
}

type interfaces struct {
	Sata3   int `bson:"SATA3"`
	M2      int `bson:"M2"`
	PciE1x  int `bson:"PCI_E_1x"`
	PciE16x int `bson:"PCI_E_16x"`
}

type Ram struct {
	ID           string  `bson:"_id"`
	Manufacturer string  `bson:"manufacturer"`
	Model        string  `bson:"model"`
	Capacity     int     `bson:"capacity"`
	Number       int     `bson:"number"`
	FormFactor   string  `bson:"form_factor"`
	Rank         int     `bson:"rank"`
	RamType      string  `bson:"ram_type"`
	Bandwidth    int     `bson:"bandwidth"`
	CasLatency   string  `bson:"cas_latency"`
	TimingScheme []int   `bson:"timing_scheme"`
	Voltage      float64 `bson:"voltage"`
	Cooling      string  `bson:"cooling"`
	Height       int     `bson:"height"`
}

type Ssd struct {
	ID           string  `bson:"_id"`
	Manufacturer string  `bson:"manufacturer"`
	Model        string  `bson:"model"`
	Type         string  `bson:"type"`
	Capacity     int     `bson:"capacity"`
	Interface    int     `bson:"interface"`
	MemoryType   string  `bson:"memory_type"`
	Read         int     `bson:"read"`
	Write        int     `bson:"write"`
	FormFactor   string  `bson:"form_factor"`
	Mftb         float64 `bson:"mftb"`
	Size         []int   `bson:"size"`
	Weight       int     `bson:"weight"`
}

type Hdd struct {
	ID           string `bson:"_id"`
	Manufacturer string `bson:"manufacturer"`
	Model        string `bson:"model"`
	Type         string `bson:"type"`
	Capacity     int    `bson:"capacity"`
	Interface    string `bson:"interface"`
	WriteMethod  string `bson:"write_method"`
	TransferRate int    `bson:"transfer_rate"`
	SpindleSpeed int    `bson:"spindle_speed"`
	FormFactor   string `bson:"form_factor"`
	Mftb         int    `bson:"mftb"`
	Size         []int  `bson:"size"`
	Weight       int    `bson:"weight"`
}

type GraphicCard struct {
}

type Cooling struct {
	ID           string   `bson:"_id"`
	Manufacturer string   `bson:"manufacturer"`
	Model        string   `bson:"model"`
	Type         string   `bson:"type"`
	Sockets      []string `bson:"sockets"`
	Fans         []int    `bson:"fans"`
	Rpm          []int    `bson:"rpm"`
	Tdp          int      `bson:"tdp"`
	NoiseLevel   int      `bson:"noise_level"`
	MountType    string   `bson:"mount_type"`
	Power        int      `bson:"power"`
	Height       float32  `bson:"height"`
}

type Housing struct {
	ID              string    `bson:"_id"`
	Manufacturer    string    `bson:"manufacturer"`
	Model           string    `bson:"model"`
	FormFactor      string    `bson:"form_factor"`
	DriveBays       driveBays `bson:"drive_bays"`
	MbFormFactor    string    `bson:"mb_form_factor"`
	PsFormFactor    string    `bson:"ps_form_factor"`
	ExpansionSlots  int       `bson:"expansion_slots"`
	GraphicCardSize int       `bson:"graphic_card_size"`
	CoolerHeight    int       `bson:"cooler_height"`
	Size            []int     `bson:"size"`
	Weight          float64   `bson:"weight"`
}

type driveBays struct {
	D25 int `bson:"25"`
	D35 int `bson:"35"`
}

type PowerSupply struct {
	ID           string     `bson:"_id"`
	Manufacturer string     `bson:"manufacturer"`
	Model        string     `bson:"model"`
	FormFactor   string     `bson:"form_factor"`
	OutputPower  int        `bson:"output_power"`
	Connectors   connectors `bson:"connectors"`
}

type connectors []sata

type sata struct {
	Sata int `bson:"SATA"`
}

type molex struct {
	Molex int `bson:"MOLEX"`
}

type pciE struct {
	PciE   string `bson:"PCI_E"`
	Number int    `bson:"number"`
}
