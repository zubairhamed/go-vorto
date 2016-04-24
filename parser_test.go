package govorto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInfoModel(t *testing.T) {
	m := ParseFile("./testresources/TI_SensorTag_CC2650.infomodel")

	assert.NotNil(t, m)

	assert.Equal(t, "examples.informationmodels.sensors", m.Namespace)
	assert.Equal(t, "1.0.0", m.Version)
	assert.Equal(t, "TI SensorTag CC2650", m.Displayname)
	assert.Equal(t, "Information model for the TI SensorTag CC2650.", m.Description)
	assert.Equal(t, "demo", m.Category)
	assert.Equal(t, 4, len(m.Using))
	assert.Equal(t, "examples.functionblockmodels.sensors.Accelerometer ; 1.0.0", m.Using[0])
	assert.Equal(t, "examples.functionblockmodels.sensors.TemperatureSensor ; 1.0.0", m.Using[1])
	assert.Equal(t, "examples.functionblockmodels.sensors.HumiditySensor ; 1.0.0", m.Using[2])
	assert.Equal(t, "examples.functionblockmodels.sensors.PressureSensor ; 1.0.0", m.Using[3])

	assert.Equal(t, "TI_SensorTag_CC2650", m.InfoModel.Name)
	assert.Equal(t, 4, len(m.InfoModel.FunctionBlocks))
	assert.Equal(t, "accelerometer", m.InfoModel.FunctionBlocks[0].Name)
	assert.Equal(t, "Accelerometer", m.InfoModel.FunctionBlocks[0].DataType)
	assert.Equal(t, "Function block representing the accelerometer of the device.", m.InfoModel.FunctionBlocks[0].Description)
	assert.Equal(t, "temperaturesensor", m.InfoModel.FunctionBlocks[1].Name)
	assert.Equal(t, "TemperatureSensor", m.InfoModel.FunctionBlocks[1].DataType)
	assert.Equal(t, "Function block representing the temperature sensor of the device.", m.InfoModel.FunctionBlocks[1].Description)
	assert.Equal(t, "humiditysensor", m.InfoModel.FunctionBlocks[2].Name)
	assert.Equal(t, "HumiditySensor", m.InfoModel.FunctionBlocks[2].DataType)
	assert.Equal(t, "Function block representing the humidity sensor of the device.", m.InfoModel.FunctionBlocks[2].Description)
	assert.Equal(t, "pressuresensor", m.InfoModel.FunctionBlocks[3].Name)
	assert.Equal(t, "PressureSensor", m.InfoModel.FunctionBlocks[3].DataType)
	assert.Equal(t, "Function block representing the pressure sensor of the device.", m.InfoModel.FunctionBlocks[3].Description)
}

func TestParseFunctionBlock(t *testing.T) {
	m := ParseFile("./testresources/Battery.fbmodel")

	assert.NotNil(t, m)

	assert.Equal(t, "examples.functionblockmodels.metering", m.Namespace)
	assert.Equal(t, "1.0.0", m.Version)
	assert.Equal(t, "Battery", m.Displayname)
	assert.Equal(t, "Function block model for describing key functionalities of a battery device.", m.Description)
	assert.Equal(t, "example", m.Category)
	assert.Equal(t, 1, len(m.Using))
	assert.Equal(t, "examples.datatypes.measurement.Percent ; 1.0.0", m.Using[0])

	assert.Equal(t, "Battery", m.FunctionBlock.Name)

	assert.Equal(t, 2, len(m.FunctionBlock.Status))
	assert.Equal(t, "batteryState", m.FunctionBlock.Status[0].Name)
	assert.Equal(t, "Percent", m.FunctionBlock.Status[0].DataType)
	assert.Equal(t, "Indicates the state of the battery.", m.FunctionBlock.Status[0].Description)
	assert.True(t, m.FunctionBlock.Status[0].Mandatory)
	assert.Equal(t, "charging", m.FunctionBlock.Status[1].Name)
	assert.Equal(t, "boolean", m.FunctionBlock.Status[1].DataType)
	assert.Equal(t, "Indicates if the battery is currently charging.", m.FunctionBlock.Status[1].Description)
	assert.True(t, m.FunctionBlock.Status[0].Mandatory)

	assert.Equal(t, 2, len(m.FunctionBlock.Operations))
	assert.Equal(t, "getBatteryState()", m.FunctionBlock.Operations[0].Name)
	assert.Equal(t, "Percent", m.FunctionBlock.Operations[0].DataType)
	assert.Equal(t, "Returns the battery state in percent.", m.FunctionBlock.Operations[0].Description)
	assert.Equal(t, "getCharging()", m.FunctionBlock.Operations[1].Name)
	assert.Equal(t, "boolean", m.FunctionBlock.Operations[1].DataType)
	assert.Equal(t, "Indicates if the battery is currently charging.", m.FunctionBlock.Operations[1].Description)
}

func TestParseDataType(t *testing.T) {
	m := ParseFile("./testresources/DoorState.type")

	assert.NotNil(t, m)

	assert.Equal(t, "examples.datatypes.state", m.Namespace)
	assert.Equal(t, "1.0.0", m.Version)
	assert.Equal(t, "Enum for describing the state of a door.", m.Description)

	assert.Equal(t, "DoorState", m.DataType.Name)
	assert.Equal(t, "open", m.DataType.Fields[0].Name)
	assert.Equal(t, "The door is open.", m.DataType.Fields[0].Description)

	assert.Equal(t, "closed", m.DataType.Fields[1].Name)
	assert.Equal(t, "The door is closed.", m.DataType.Fields[1].Description)

	assert.Equal(t, "undefined", m.DataType.Fields[2].Name)
	assert.Equal(t, "The status is not known.", m.DataType.Fields[2].Description)
}
