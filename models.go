package govorto

func NewVortoModel() *VortoModel {
	return &VortoModel{
		Using:         []string{},
		InfoModel:     &InformationModel{},
		DataType:      &DataType{},
		FunctionBlock: &FunctionBlockModel{},
	}
}

type VortoModel struct {
	Namespace   string
	Version     string
	Displayname string
	Description string
	Category    string
	Using       []string

	InfoModel     *InformationModel
	DataType      *DataType
	FunctionBlock *FunctionBlockModel
}

type DataTypeField struct {
	Name        string
	Description string
}
type DataType struct {
	Name   string
	Fields []*DataTypeField
}

type FunctionBlockModel struct {
	Name       string
	Status     []*Status
	Operations []*Operation
}

type Status struct {
	Name        string
	DataType    string
	Description string
	Mandatory	bool
}

type Operation struct {
	Name        string
	DataType    string
	Description string
}

type InformationModel struct {
	Name           string
	FunctionBlocks []*FunctionBlock
}

type FunctionBlock struct {
	Name        string
	DataType      string
	Description string
}
