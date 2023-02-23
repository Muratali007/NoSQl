package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type Cpu struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Manufacturer   string             `bson:"manufacturer" json:"manufacturer,omitempty"`
	Model          string             `bson:"model" json:"model,omitempty"`
	Cores          int                `bson:"cores" json:"cores,omitempty"`
	Threads        int                `bson:"threads" json:"threads,omitempty"`
	ClockFrequency []float64          `bson:"clock_frequency" json:"clock_frequency,omitempty"`
	Socket         string             `bson:"socket" json:"socket,omitempty"`
	Ram            RamCpu             `bson:"ram"`
	Tdp            int                `bson:"tdp" json:"tdp,omitempty"`
	PciE           int                `bson:"pci-e" json:"pci_e,omitempty"`
	MaxTemperature int                `bson:"max_temperature"`
	Image          string             `bson:"image"`
}

type RamCpu struct {
	Channels     int    `bson:"channels"`
	Type         string `bson:"type"`
	MaxFrequency int    `bson:"max_frequency"`
	MaxCapacity  int    `bson:"max_capacity"`
}

type Motherboard struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	Socket       string             `bson:"socket"`
	Chipset      string             `bson:"chipset"`
	FormFactor   string             `bson:"form_factor"`
	Ram          ramMb              `bson:"ram"`
	Interfaces   interfaces         `bson:"interfaces"`
	PciStandard  int                `bson:"pci_standard"`
	MbPower      int                `bson:"mb_power"`
	CpuPower     int                `bson:"cpu_power"`
	Image        string             `bson:"image"`
}

type ramMb struct {
	Slots        int    `bson:"slots"`
	Type         string `bson:"type"`
	MaxFrequency int    `bson:"max_frequency"`
	MaxCapacity  int    `bson:"max_capacity"`
}

type interfaces struct {
	Sata3   int `bson:"SATA3"`
	M2      int `bson:"M2"`
	PciE1x  int `bson:"PCI_E_1x"`
	PciE16x int `bson:"PCI_E_16x"`
}

type Ram struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	Capacity     int                `bson:"capacity"`
	Number       int                `bson:"number"`
	FormFactor   string             `bson:"form_factor"`
	Rank         int                `bson:"rank"`
	Type         string             `bson:"type"`
	Frequency    int                `bson:"frequency"`
	Bandwidth    int                `bson:"bandwidth"`
	CasLatency   string             `bson:"cas_latency"`
	TimingScheme []int              `bson:"timing_scheme"`
	Voltage      float64            `bson:"voltage"`
	Cooling      string             `bson:"cooling"`
	Height       int                `bson:"height"`
}

type Ssd struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	Type         string             `bson:"type"`
	Capacity     int                `bson:"capacity"`
	Interface    string             `bson:"interface"`
	MemoryType   string             `bson:"memory_type"`
	Read         int                `bson:"read"`
	Write        int                `bson:"write"`
	FormFactor   string             `bson:"form_factor"`
	Mftb         float64            `bson:"mftb"`
	Size         []int              `bson:"size"`
	Weight       int                `bson:"weight"`
}

type Hdd struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	Type         string             `bson:"type"`
	Capacity     int                `bson:"capacity"`
	Interface    string             `bson:"interface"`
	WriteMethod  string             `bson:"write_method"`
	TransferRate int                `bson:"transfer_rate"`
	SpindleSpeed int                `bson:"spindle_speed"`
	FormFactor   string             `bson:"form_factor"`
	Mftb         int                `bson:"mftb"`
	Size         []int              `bson:"size"`
	Weight       int                `bson:"weight"`
}

type Gpu struct {
	ID            primitive.ObjectID `bson:"_id"`
	Manufacturer  string             `bson:"manufacturer"`
	Model         string             `bson:"model"`
	Architecture  string             `bson:"architecture"`
	Memory        memoryGpu          `bson:"memory"`
	GpuFrequency  int                `bson:"gpu_frequency"`
	ProcessSize   int                `bson:"process_size"`
	MaxResolution string             `bson:"max_resolution"`
	Interfaces    []interfacesGpu    `bson:"interfaces"`
	MaxMonitors   int                `bson:"max_monitors"`
	Cooling       coolingGpu         `bson:"cooling"`
	Tdp           int                `bson:"tdp"`
	TdpR          int                `bson:"tdp_r"`
	PowerSupply   []int              `bson:"power_supply"`
	Slots         float64            `bson:"slots"`
	Size          []int              `bson:"size"`
	Image         string             `bson:"image"`
}

type memoryGpu struct {
	Capacity       int    `bson:"capacity"`
	Type           string `bson:"type"`
	InterfaceWidth int    `bson:"interface_width"`
	Frequency      int    `bson:"frequency"`
}

type interfacesGpu struct {
	Type   string `bson:"type"`
	Number int    `bson:"number"`
}

type coolingGpu struct {
	Type      string `bson:"type"`
	FanNumber int    `bson:"fan_number"`
}

type Cooling struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	Type         string             `bson:"type"`
	Sockets      []string           `bson:"sockets"`
	Fans         []int              `bson:"fans"`
	Rpm          []int              `bson:"rpm"`
	Tdp          int                `bson:"tdp"`
	NoiseLevel   int                `bson:"noise_level"`
	MountType    string             `bson:"mount_type"`
	Power        int                `bson:"power"`
	Height       int                `bson:"height"`
	Image        string             `bson:"image"`
}

type Housing struct {
	ID              primitive.ObjectID `bson:"_id"`
	Manufacturer    string             `bson:"manufacturer"`
	Model           string             `bson:"model"`
	FormFactor      string             `bson:"form_factor"`
	DriveBays       driveBays          `bson:"drive_bays"`
	MbFormFactor    string             `bson:"mb_form_factor"`
	PsFormFactor    string             `bson:"ps_form_factor"`
	ExpansionSlots  int                `bson:"expansion_slots"`
	GraphicCardSize int                `bson:"graphic_card_size"`
	CoolerHeight    int                `bson:"cooler_height"`
	Size            []int              `bson:"size"`
	Weight          float64            `bson:"weight"`
	Image           string             `bson:"image"`
}

type driveBays struct {
	D35 int `bson:"3_5"`
	D25 int `bson:"2_5"`
}

type PowerSupply struct {
	ID           primitive.ObjectID `bson:"_id"`
	Manufacturer string             `bson:"manufacturer"`
	Model        string             `bson:"model"`
	FormFactor   string             `bson:"form_factor"`
	OutputPower  int                `bson:"output_power"`
	Connectors   connectors         `bson:"connectors"`
	Modules      bool               `bson:"modules"`
	MbPower      int                `bson:"mb_power"`
	CpuPower     bsoncore.Array     `bson:"cpu_power"`
	Image        string             `bson:"image"`
}

type connectors struct {
	Sata  int   `bson:"SATA"`
	Molex int   `bson:"MOLEX"`
	PciE  []int `bson:"PCI_E"`
}
