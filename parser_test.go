package govorto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInfoModel(t *testing.T) {
	m := ParseFile("./testresources/TI_SensorTag_CC2650.infomodel")

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

	// assert.Equal(t, "TI_SensorTag_CC2650", m.InfoModel.name)
	// assert.Equal(t, 4, len(m.InfoModel.functionBlocks))

	//accelerometer
	//m.InfoModel.functionBlocks[0].name
	//
	//Accelerometer
	//m.InfoModel.functionBlocks[0].fbType
	//
	//Function block representing the accelerometer of the device.
	//m.InfoModel.functionBlocks[0].description
	//
	//temperaturesensor
	//m.InfoModel.functionBlocks[0].name
	//
	//TemperatureSensor
	//m.InfoModel.functionBlocks[0].fbType
	//
	//Function block representing the temperature sensor of the device.
	//m.InfoModel.functionBlocks[0].description
	//
	//humiditysensor
	//m.InfoModel.functionBlocks[0].name
	//
	//HumiditySensor
	//m.InfoModel.functionBlocks[0].fbType
	//
	//Function block representing the humidity sensor of the device.
	//m.InfoModel.functionBlocks[0].description
	//
	//pressuresensor
	//m.InfoModel.functionBlocks[0].name
	//
	//PressureSensor
	//m.InfoModel.functionBlocks[0].fbType
	//
	//Function block representing the pressure sensor of the device.
	//m.InfoModel.functionBlocks[0].description
}

func TestParseFunctionBlock(t *testing.T) {
	m := ParseFile("./testresources/Battery.fbmodel")

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

}
