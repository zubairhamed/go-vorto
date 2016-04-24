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
	Status     []*FunctionBlockStatus
	Operations []*FunctionBlockOperation
}

type FunctionBlockConfiguration struct {
/*
	mandatory blinking as boolean "if the lamp is currently blinking or not"
	mandatory on as boolean "if the lamp is currently switched on"
	mandatory powerConsumption as int "the amount of power the lamp is consuming"
 */
}

type FunctionBlockFault struct {
// mandatory bulbDefect as boolean "true if the light bulb of the lamp is defect"
}

type FunctionBlockStatus struct {
	Name        string
	DataType    string
	Description string
	Mandatory	bool
}

type FunctionBlockOperation struct {
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
