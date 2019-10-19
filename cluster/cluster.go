package cluster

type AttributeDescriptor struct {
	Name   string
	Type   ZclDataType
	Access Access
}

type CommandDescriptor struct {
	Name    string
	Command interface{}
}

type CommandDescriptors struct {
	Received  map[uint8]*CommandDescriptor
	Generated map[uint8]*CommandDescriptor
}

type Cluster struct {
	Name                 string
	AttributeDescriptors map[uint16]*AttributeDescriptor
	CommandDescriptors   *CommandDescriptors
}

type ClusterLibrary struct {
	global   map[uint8]*CommandDescriptor
	clusters map[ClusterId]*Cluster
}

type Access uint8

const (
	Read       Access = 0x01
	Write      Access = 0x02
	Reportable Access = 0x04
	Scene      Access = 0x08
)

type ClusterId uint16

const (
	Basic                          ClusterId = 0x0000
	PowerConfiguration             ClusterId = 0x0001
	DeviceTemperatureConfiguration ClusterId = 0x0002
	Identify                       ClusterId = 0x0003
	OnOff                          ClusterId = 0x0006
	LevelControl                   ClusterId = 0x0008
	Time                           ClusterId = 0x000a
	AnalogInputBasic               ClusterId = 0x000c
	BinaryOutputBasic              ClusterId = 0x0010
	MultistateInput                ClusterId = 0x0012
	OTA                            ClusterId = 0x0019
	IlluminanceMeasurement         ClusterId = 0x0400
	IlluminanceLevelSensing        ClusterId = 0x0401
	TemperatureMeasurement         ClusterId = 0x0402
	PressureMeasurement            ClusterId = 0x0403
	FlowMeasurement                ClusterId = 0x0404
	RelativeHumidityMeasurement    ClusterId = 0x0405
	OccupancySensing               ClusterId = 0x0406
	ElectricalMeasurement          ClusterId = 0x0b04
	IASZone                        ClusterId = 0x0500
	IASACE                         ClusterId = 0x0501
	IASWarningDevice               ClusterId = 0x0502
)

func New() *ClusterLibrary {
	return &ClusterLibrary{
		global: map[uint8]*CommandDescriptor{
			0x00: {"ReadAttributes", &ReadAttributesCommand{}},
			0x01: {"ReadAttributesResponse", &ReadAttributesResponse{}},
			0x02: {"WriteAttributes", &WriteAttributesCommand{}},
			0x03: {"WriteAttributesUndivided", &WriteAttributesUndividedCommand{}},
			0x04: {"WriteAttributesResponse", &WriteAttributesResponse{}},
			0x05: {"WriteAttributesNoResponse", &WriteAttributesNoResponseCommand{}},
			0x06: {"ConfigureReporting", &ConfigureReportingCommand{}},
			0x07: {"ConfigureReportingResponse", &ConfigureReportingResponse{}},
			0x08: {"ReadReportingConfiguration", &ReadReportingConfigurationCommand{}},
			0x09: {"ReadReportingConfigurationResponse", &ReadReportingConfigurationResponse{}},
			0x0a: {"ReportAttributes", &ReportAttributesCommand{}},
			0x0b: {"DefaultResponse", &DefaultResponseCommand{}},
			0x0c: {"DiscoverAttributes", &DiscoverAttributesCommand{}},
			0x0d: {"DiscoverAttributesResponse", &DiscoverAttributesResponse{}},
			0x0e: {"ReadAttributesStructured", &ReadAttributesStructuredCommand{}},
			0x0f: {"WriteAttributesStructured", &WriteAttributesStructuredCommand{}},
			0x10: {"WriteAttributesStructuredResponse", &WriteAttributesStructuredResponse{}},
			0x11: {"DiscoverCommandsReceived", &DiscoverCommandsReceivedCommand{}},
			0x12: {"DiscoverCommandsReceivedResponse", &DiscoverCommandsReceivedResponse{}},
			0x13: {"DiscoverCommandsGenerated", &DiscoverCommandsGeneratedCommand{}},
			0x14: {"DiscoverCommandsGeneratedResponse", &DiscoverCommandsGeneratedResponse{}},
			0x15: {"DiscoverAttributesExtended", &DiscoverAttributesExtendedCommand{}},
			0x16: {"DiscoverAttributesExtendedResponse", &DiscoverAttributesExtendedResponse{}},
		},
		clusters: map[ClusterId]*Cluster{
			Basic: {
				Name: "Basic",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"ZLibraryVersion", ZclDataTypeUint8, Read},
					0x0001: {"ApplicationVersion", ZclDataTypeUint8, Read},
					0x0002: {"StackVersion", ZclDataTypeUint8, Read},
					0x0003: {"HWVersion", ZclDataTypeUint8, Read},
					0x0004: {"ManufacturerName", ZclDataTypeCharStr, Read},
					0x0005: {"ModelIdentifier", ZclDataTypeCharStr, Read},
					0x0006: {"DateCode", ZclDataTypeCharStr, Read},
					0x0007: {"PowerSource", ZclDataTypeEnum8, Read},
					0x0010: {"LocationDescription", ZclDataTypeCharStr, Read | Write},
					0x0011: {"PhysicalEnvironment", ZclDataTypeEnum8, Read | Write},
					0x0012: {"DeviceEnabled", ZclDataTypeBoolean, Read | Write},
					0x0013: {"AlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0014: {"DisableLocalConfig", ZclDataTypeBitmap8, Read | Write},
					0x4000: {"SWBuildID", ZclDataTypeCharStr, Read},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"ResetToFactoryDefaults", &ResetToFactoryDefaultsCommand{}},
					},
				},
			},
			PowerConfiguration: {
				Name: "PowerConfiguration",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MainsVoltage", ZclDataTypeUint16, Read},
					0x0001: {"MainsFrequency", ZclDataTypeUint8, Read},
					0x0010: {"MainsAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0011: {"MainsVoltageMinThreshold", ZclDataTypeUint16, Read | Write},
					0x0012: {"MainsVoltageMaxThreshold", ZclDataTypeUint16, Read | Write},
					0x0013: {"MainsVoltageDwellTripPoint", ZclDataTypeUint16, Read | Write},
					0x0020: {"BatteryVoltage", ZclDataTypeUint8, Read},
					0x0021: {"BatteryPercentageRemaining", ZclDataTypeUint8, Read | Reportable},
					0x0030: {"BatteryManufacturer", ZclDataTypeCharStr, Read | Write},
					0x0031: {"BatterySize", ZclDataTypeEnum8, Read | Write},
					0x0032: {"BatteryAHrRating", ZclDataTypeUint16, Read | Write},
					0x0033: {"BatteryQuantity", ZclDataTypeUint8, Read | Write},
					0x0034: {"BatteryRatedVoltage", ZclDataTypeUint8, Read | Write},
					0x0035: {"BatteryAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0036: {"BatteryVoltageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x0037: {"BatteryVoltageThreshold1", ZclDataTypeUint8, Read | Write},
					0x0038: {"BatteryVoltageThreshold2", ZclDataTypeUint8, Read | Write},
					0x0039: {"BatteryVoltageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003a: {"BatteryPercentageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x003b: {"BatteryPercentageThreshold1", ZclDataTypeUint8, Read | Write},
					0x003c: {"BatteryPercentageThreshold2", ZclDataTypeUint8, Read | Write},
					0x003d: {"BatteryPercentageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003e: {"BatteryAlarmState", ZclDataTypeBitmap32, Read},
				},
			},
			DeviceTemperatureConfiguration: {
				Name: "DeviceTemperatureConfiguration",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"CurrentTemperature", ZclDataTypeInt16, Read},
					0x0001: {"MinTempExperienced", ZclDataTypeInt16, Read},
					0x0002: {"MaxTempExperienced", ZclDataTypeInt16, Read},
					0x0003: {"OverTempTotalDwell", ZclDataTypeInt16, Read},
					0x0010: {"DeviceTempAlarmMask", ZclDataTypeBitmap16, Read | Write},
					0x0011: {"LowTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0012: {"HighTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0013: {"LowTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
					0x0014: {"HighTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
				},
			},
			Identify: {
				Name: "Identify",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"IdentifyTime", ZclDataTypeInt16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"Identify", &IdentifyCommand{}},
						0x01: {"IdentifyQuery", &IdentifyQueryCommand{}},
						0x40: {"TriggerEffect ", &TriggerEffectCommand{}},
					},
					Generated: map[uint8]*CommandDescriptor{
						0x00: {"IdentifyQueryResponse ", &IdentifyQueryResponse{}},
					},
				},
			},
			OnOff: {
				Name: "OnOff",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"OnOff", ZclDataTypeBoolean, Read | Reportable | Scene},
					0x4000: {"GlobalSceneControl", ZclDataTypeBoolean, Read},
					0x4001: {"OnTime", ZclDataTypeUint16, Read | Write},
					0x4002: {"OffWaitTime", ZclDataTypeUint16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"Off", &OffCommand{}},
						0x01: {"On", &OnCommand{}},
						0x02: {"Toggle ", &ToggleCommand{}},
						0x40: {"OffWithEffect ", &OffWithEffectCommand{}},
						0x41: {"OnWithRecallGlobalScene ", &OnWithRecallGlobalSceneCommand{}},
						0x42: {"OnWithTimedOff ", &OnWithTimedOffCommand{}},
					},
				},
			},
			LevelControl: {
				Name: "LevelControl",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"CurrentLevel", ZclDataTypeUint8, Read | Reportable},
					0x0001: {"RemainingTime", ZclDataTypeUint16, Read},
					0x0010: {"OnOffTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0011: {"OnLevel", ZclDataTypeUint8, Read | Write},
					0x0012: {"OnTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0013: {"OffTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0014: {"DefaultMoveRate", ZclDataTypeUint16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"MoveToLevel ", &MoveToLevelCommand{}},
						0x01: {"Move", &MoveCommand{}},
						0x02: {"Step ", &StepCommand{}},
						0x03: {"Stop ", &StopCommand{}},
						0x04: {"MoveToLevel/OnOff", &MoveToLevelOnOffCommand{}},
						0x05: {"Move/OnOff", &MoveOnOffCommand{}},
						0x06: {"Step/OnOff", &StepOnOffCommand{}},
						0x07: {"Stop/OnOff", &StopOnOffCommand{}},
					},
				},
			},
			Time: {
				Name: "Time",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"Time", ZclDataTypeUtc, Read | Write},
					0x0001: {"TimeStatus", ZclDataTypeBitmap8, Read | Write},
					0x0002: {"TimeZone", ZclDataTypeInt32, Read | Write},
					0x0003: {"DstStart", ZclDataTypeUint32, Read | Write},
					0x0004: {"DstEnd", ZclDataTypeUint32, Read | Write},
					0x0005: {"DstShift", ZclDataTypeInt32, Read | Write},
					0x0006: {"StandardTime", ZclDataTypeUint32, Read},
					0x0007: {"LocalTime", ZclDataTypeUint32, Read},
					0x0008: {"LastSetTime", ZclDataTypeUtc, Read},
					0x0009: {"ValidUntilTime", ZclDataTypeUtc, Read | Write},
				},
			},
			AnalogInputBasic: {
				Name: "AnalogInputBasic",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x001c: {"Description", ZclDataTypeCharStr, Read | Write},
					0x0041: {"MaxPresentValue", ZclDataTypeSinglePrec, Read | Write},
					0x0045: {"MinPresentValue", ZclDataTypeSinglePrec, Read | Write},
					0x0050: {"OutOfService", ZclDataTypeBoolean, Read | Write},
					0x0055: {"PresentValue", ZclDataTypeSinglePrec, Read | Write},
					0x0067: {"Reliability", ZclDataTypeBitmap8, Read | Write},
					0x006a: {"Resolution", ZclDataTypeSinglePrec, Read | Write},
					0x006f: {"StatusFlags", ZclDataTypeBitmap8, Read},
					0x0075: {"EngineeringUnits", ZclDataTypeEnum16, Read | Write},
					0x0100: {"ApplicationType", ZclDataTypeUint32, Read},
				},
			},
			BinaryOutputBasic: {
				Name: "BinaryOutputBasic",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0004: {"ActiveText", ZclDataTypeCharStr, Read | Write},
					0x001c: {"Description", ZclDataTypeCharStr, Read | Write},
					0x002e: {"InactiveText", ZclDataTypeCharStr, Read | Write},
					0x0042: {"MinimumOffTime", ZclDataTypeUint32, Read | Write},
					0x0043: {"MaximumOffTime", ZclDataTypeUint32, Read | Write},
					0x0051: {"OutOfService", ZclDataTypeBoolean, Read | Write},
					0x0050: {"Polarity", ZclDataTypeEnum8, Read},
					0x0055: {"PresentValue", ZclDataTypeBoolean, Read | Write},
					0x0067: {"Reliability", ZclDataTypeBitmap8, Read | Write},
					0x0068: {"RelinquishDefault", ZclDataTypeBoolean, Read | Write},
					0x006f: {"StatusFlags", ZclDataTypeBitmap8, Read},
					0x0100: {"ApplicationType", ZclDataTypeUint32, Read},
				},
			},
			MultistateInput: {
				Name: "MultistateInput",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x000E: {"StateText", ZclDataTypeArray, Read | Write},
					0x001C: {"Description", ZclDataTypeCharStr, Read | Write},
					0x004A: {"NumberOfStates", ZclDataTypeUint16, Read | Write},
					0x0051: {"OutOfService", ZclDataTypeBoolean, Read | Write},
					0x0055: {"PresentValue", ZclDataTypeUint16, Read | Write},
					0x0067: {"Reliability", ZclDataTypeEnum8, Read | Write},
					0x006F: {"StatusFlags", ZclDataTypeBitmap8, Read},
					0x0100: {"ApplicationType", ZclDataTypeUint32, Read},
				},
			},
			OTA: {
				Name: "OTA",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"UpgradeServerID", ZclDataTypeIeeeAddr, Read},
					0x0001: {"FileOffset", ZclDataTypeUint32, Read},
					0x0002: {"CurrentFileVersion", ZclDataTypeUint32, Read},
					0x0003: {"CurrentZigBeeStackVersion", ZclDataTypeUint16, Read},
					0x0004: {"DownloadedFileVersion", ZclDataTypeUint32, Read},
					0x0005: {"DownloadedZigBeeStackVersion", ZclDataTypeUint16, Read},
					0x0006: {"ImageUpgradeStatus", ZclDataTypeEnum8, Read},
					0x0007: {"ManufacturerID", ZclDataTypeUint16, Read},
					0x0008: {"ImageTypeID ", ZclDataTypeUint16, Read},
					0x0009: {"MinimumBlockPeriod ", ZclDataTypeUint16, Read},
					0x000a: {"ImageStamp ", ZclDataTypeUint32, Read},
				},
			},
			IlluminanceMeasurement: {
				Name: "IlluminanceMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasuredValue", ZclDataTypeUint16, Read},
					0x0001: {"MinMeasuredValue", ZclDataTypeUint16, Read},
					0x0002: {"MaxMeasuredValue", ZclDataTypeUint16, Read},
					0x0003: {"Tolerance", ZclDataTypeUint16, Read},
					0x0004: {"LightSensorType", ZclDataTypeEnum8, Read},
				},
			},
			IlluminanceLevelSensing: {
				Name: "IlluminanceLevelSensing",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"LevelStatus", ZclDataTypeEnum8, Read},
					0x0001: {"LightSensorType", ZclDataTypeEnum8, Read},
					0x0010: {"IlluminanceTarget", ZclDataTypeUint16, Read | Write},
				},
			},
			TemperatureMeasurement: {
				Name: "TemperatureMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasuredValue", ZclDataTypeInt16, Read},
					0x0001: {"MinMeasuredValue", ZclDataTypeInt16, Read},
					0x0002: {"MaxMeasuredValue", ZclDataTypeInt16, Read},
					0x0003: {"Tolerance", ZclDataTypeUint16, Read},
				},
			},
			PressureMeasurement: {
				Name: "PressureMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasuredValue", ZclDataTypeInt16, Read},
					0x0001: {"MinMeasuredValue", ZclDataTypeInt16, Read},
					0x0002: {"MaxMeasuredValue", ZclDataTypeInt16, Read},
					0x0003: {"Tolerance", ZclDataTypeUint16, Read},
					0x0010: {"ScaledValue", ZclDataTypeInt16, Read},
					0x0011: {"MinScaledValue", ZclDataTypeInt16, Read},
					0x0012: {"MaxScaledValue", ZclDataTypeInt16, Read},
					0x0013: {"ScaledTolerance", ZclDataTypeUint16, Read},
					0x0014: {"Scale", ZclDataTypeInt8, Read},
				},
			},
			FlowMeasurement: {
				Name: "FlowMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasuredValue", ZclDataTypeUint16, Read},
					0x0001: {"MinMeasuredValue", ZclDataTypeUint16, Read},
					0x0002: {"MaxMeasuredValue", ZclDataTypeUint16, Read},
					0x0003: {"Tolerance", ZclDataTypeUint16, Read},
				},
			},
			RelativeHumidityMeasurement: {
				Name: "RelativeHumidityMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasuredValue", ZclDataTypeUint16, Read},
					0x0001: {"MinMeasuredValue", ZclDataTypeUint16, Read},
					0x0002: {"MaxMeasuredValue", ZclDataTypeUint16, Read},
					0x0003: {"Tolerance", ZclDataTypeUint16, Read},
				},
			},
			OccupancySensing: {
				Name: "OccupancySensing",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"Occupancy", ZclDataTypeBitmap8, Read},
					0x0001: {"OccupancySensorType", ZclDataTypeEnum8, Read},
					0x0010: {"PIROccupiedToUnoccupiedDelay", ZclDataTypeUint16, Read | Write},
					0x0011: {"PIRUnoccupiedToOccupiedDelay", ZclDataTypeUint16, Read | Write},
					0x0012: {"PIRUnoccupiedToOccupiedThreshold", ZclDataTypeUint8, Read | Write},
					0x0020: {"UltrasonicOccupiedToUnoccupiedDelay", ZclDataTypeUint16, Read | Write},
					0x0021: {"UltrasonicUnoccupiedToOccupiedDelay", ZclDataTypeUint16, Read | Write},
					0x0022: {"UltrasonicUnoccupiedToOccupiedThreshold", ZclDataTypeUint8, Read | Write},
				},
			},
			ElectricalMeasurement: {
				Name: "ElectricalMeasurement",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MeasurementType", ZclDataTypeBitmap32, Read},

					0x0100: {"DCVoltage", ZclDataTypeInt16, Read},
					0x0101: {"DCVoltageMin", ZclDataTypeInt16, Read},
					0x0102: {"DCVoltageMax", ZclDataTypeInt16, Read},
					0x0103: {"DCCurrent", ZclDataTypeInt16, Read},
					0x0104: {"DCCurrentMin", ZclDataTypeInt16, Read},
					0x0105: {"DCCurrentMax", ZclDataTypeInt16, Read},
					0x0106: {"DCPower", ZclDataTypeInt16, Read},
					0x0107: {"DCPowerMin", ZclDataTypeInt16, Read},
					0x0108: {"DCPowerMax", ZclDataTypeInt16, Read},

					0x0200: {"DCVoltageMultiplier", ZclDataTypeUint16, Read},
					0x0201: {"DCVoltageDivisor", ZclDataTypeUint16, Read},
					0x0202: {"DCCurrentMultiplier", ZclDataTypeUint16, Read},
					0x0203: {"DCCurrentDivisor", ZclDataTypeUint16, Read},
					0x0204: {"DCPowerMultiplier", ZclDataTypeUint16, Read},
					0x0205: {"DCPowerDivisor", ZclDataTypeUint16, Read},

					0x0300: {"ACFrequency", ZclDataTypeUint16, Read},
					0x0301: {"ACFrequencyMin", ZclDataTypeUint16, Read},
					0x0302: {"ACFrequencyMax", ZclDataTypeUint16, Read},
					0x0303: {"NeutralCurrent", ZclDataTypeUint16, Read},
					0x0304: {"TotalActivePower", ZclDataTypeInt32, Read},
					0x0305: {"TotalReactivePower", ZclDataTypeInt32, Read},
					0x0306: {"ApparentPower", ZclDataTypeUint32, Read},
					0x0307: {"Measured1stHarmonicCurrent", ZclDataTypeInt16, Read},
					0x0308: {"Measured3rdHarmonicCurrent", ZclDataTypeInt16, Read},
					0x0309: {"Measured5thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030a: {"Measured7thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030b: {"Measured9thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030c: {"Measured11thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030d: {"MeasuredPhase1stHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030e: {"MeasuredPhase3rdHarmonicCurrent", ZclDataTypeInt16, Read},
					0x030f: {"MeasuredPhase5thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x0310: {"MeasuredPhase7thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x0311: {"MeasuredPhase9thHarmonicCurrent", ZclDataTypeInt16, Read},
					0x0312: {"MeasuredPhase11thHarmonicCurrent", ZclDataTypeInt16, Read},

					0x0400: {"ACFrequencyMultiplier", ZclDataTypeUint16, Read},
					0x0401: {"ACFrequencyDivisor", ZclDataTypeUint16, Read},
					0x0402: {"PowerMultiplier", ZclDataTypeUint32, Read},
					0x0403: {"PowerDivisor", ZclDataTypeUint32, Read},
					0x0404: {"HarmonicCurrentMultiplier", ZclDataTypeInt8, Read},
					0x0405: {"PhaseHarmonicCurrentMultiplier", ZclDataTypeInt8, Read},

					0x0500: {"Reserved", ZclDataTypeInt16, Read},
					0x0501: {"LineCurrent", ZclDataTypeUint16, Read},
					0x0502: {"ActiveCurrent", ZclDataTypeInt16, Read},
					0x0503: {"ReactiveCurrent", ZclDataTypeInt16, Read},
					0x0504: {"Reserved", ZclDataTypeInt8, Read},
					0x0505: {"RMSVoltage", ZclDataTypeUint16, Read},
					0x0506: {"RMSVoltageMin", ZclDataTypeUint16, Read},
					0x0507: {"RMSVoltageMax", ZclDataTypeUint16, Read},
					0x0508: {"RMSCurrent", ZclDataTypeUint16, Read},
					0x0509: {"RMSCurrentMin", ZclDataTypeUint16, Read},
					0x050a: {"RMSCurrentMax", ZclDataTypeUint16, Read},
					0x050b: {"ActivePower", ZclDataTypeInt16, Read},
					0x050c: {"ActivePowerMin", ZclDataTypeInt16, Read},
					0x050d: {"ActivePowerMax", ZclDataTypeInt16, Read},
					0x050e: {"ReactivePower", ZclDataTypeInt16, Read},
					0x050f: {"ApparentPower", ZclDataTypeUint16, Read},
					0x0510: {"PowerFactor", ZclDataTypeInt8, Read},
					0x0511: {"AverageRMSVoltageMeasurementPeriod", ZclDataTypeUint16, Read | Write},
					0x0512: {"AverageRMSOverVoltageCounter", ZclDataTypeUint16, Read | Write},
					0x0513: {"AverageRMSUnderVoltageCounter", ZclDataTypeUint16, Read | Write},
					0x0514: {"RMSExtremeOverVoltagePeriod", ZclDataTypeUint16, Read | Write},
					0x0515: {"RMSExtremeUnderVoltagePeriod", ZclDataTypeUint16, Read | Write},
					0x0516: {"RMSVoltageSagPeriod", ZclDataTypeUint16, Read | Write},
					0x0517: {"RMSVoltageSwellPeriod", ZclDataTypeUint16, Read | Write},

					0x0600: {"ACVoltageMultiplier", ZclDataTypeUint16, Read},
					0x0601: {"ACVoltageDivisor", ZclDataTypeUint16, Read},
					0x0602: {"ACCurrentMultiplier", ZclDataTypeUint16, Read},
					0x0603: {"ACCurrentDivisor", ZclDataTypeUint16, Read},
					0x0604: {"ACPowerMultiplier", ZclDataTypeUint16, Read},
					0x0605: {"ACPowerDivisor", ZclDataTypeUint16, Read},

					0x0700: {"DCOverloadAlarmsMask", ZclDataTypeBitmap8, Read | Write},
					0x0701: {"DCVoltageOverload", ZclDataTypeInt16, Read},
					0x0702: {"DCVoltageOverload", ZclDataTypeInt16, Read},

					0x0800: {"ACAlarmsMask", ZclDataTypeBitmap16, Read | Write},
					0x0801: {"ACVoltageOverload", ZclDataTypeInt16, Read},
					0x0802: {"ACCurrentOverload", ZclDataTypeInt16, Read},
					0x0803: {"ACActivePowerOverload", ZclDataTypeInt16, Read},
					0x0804: {"ACReactivePowerOverload", ZclDataTypeInt16, Read},
					0x0805: {"AverageRMSOverVoltage", ZclDataTypeInt16, Read},
					0x0806: {"AverageRMSUnderVoltage", ZclDataTypeInt16, Read},
					0x0807: {"RMSExtremeOverVoltage", ZclDataTypeInt16, Read | Write},
					0x0808: {"RMSExtremeUnderVoltage", ZclDataTypeInt16, Read | Write},
					0x0809: {"RMSVoltageSag", ZclDataTypeInt16, Read | Write},
					0x080a: {"RMSVoltageSwell", ZclDataTypeInt16, Read | Write},

					0x0900: {"ReservedPhB", ZclDataTypeInt16, Read},
					0x0901: {"LineCurrentPhB", ZclDataTypeUint16, Read},
					0x0902: {"ActiveCurrentPhB", ZclDataTypeInt16, Read},
					0x0903: {"ReactiveCurrentPhB", ZclDataTypeInt16, Read},
					0x0904: {"ReservedPhB", ZclDataTypeInt8, Read},
					0x0905: {"RMSVoltagePhB", ZclDataTypeUint16, Read},
					0x0906: {"RMSVoltageMinPhB", ZclDataTypeUint16, Read},
					0x0907: {"RMSVoltageMaxPhB", ZclDataTypeUint16, Read},
					0x0908: {"RMSCurrentPhB", ZclDataTypeUint16, Read},
					0x0909: {"RMSCurrentMinPhB", ZclDataTypeUint16, Read},
					0x090a: {"RMSCurrentMaxPhB", ZclDataTypeUint16, Read},
					0x090b: {"ActivePowerPhB", ZclDataTypeInt16, Read},
					0x090c: {"ActivePowerMinPhB", ZclDataTypeInt16, Read},
					0x090d: {"ActivePowerMaxPhB", ZclDataTypeInt16, Read},
					0x090e: {"ReactivePowerPhB", ZclDataTypeInt16, Read},
					0x090f: {"ApparentPowerPhB", ZclDataTypeUint16, Read},
					0x0910: {"PowerFactorPhB", ZclDataTypeInt8, Read},
					0x0911: {"AverageRMSVoltageMeasurementPeriodPhB", ZclDataTypeUint16, Read | Write},
					0x0912: {"AverageRMSOverVoltageCounterPhB", ZclDataTypeUint16, Read | Write},
					0x0913: {"AverageRMSUnderVoltageCounterPhB", ZclDataTypeUint16, Read | Write},
					0x0914: {"RMSExtremeOverVoltagePeriodPhB", ZclDataTypeUint16, Read | Write},
					0x0915: {"RMSExtremeUnderVoltagePeriodPhB", ZclDataTypeUint16, Read | Write},
					0x0916: {"RMSVoltageSagPeriodPhB", ZclDataTypeUint16, Read | Write},
					0x0917: {"RMSVoltageSwellPeriodPhB", ZclDataTypeUint16, Read | Write},

					0x0a00: {"ReservedPhC", ZclDataTypeInt16, Read},
					0x0a01: {"LineCurrentPhC", ZclDataTypeUint16, Read},
					0x0a02: {"ActiveCurrentPhC", ZclDataTypeInt16, Read},
					0x0a03: {"ReactiveCurrentPhC", ZclDataTypeInt16, Read},
					0x0a04: {"ReservedPhC", ZclDataTypeInt8, Read},
					0x0a05: {"RMSVoltagePhC", ZclDataTypeUint16, Read},
					0x0a06: {"RMSVoltageMinPhC", ZclDataTypeUint16, Read},
					0x0a07: {"RMSVoltageMaxPhC", ZclDataTypeUint16, Read},
					0x0a08: {"RMSCurrentPhC", ZclDataTypeUint16, Read},
					0x0a09: {"RMSCurrentMinPhC", ZclDataTypeUint16, Read},
					0x0a0a: {"RMSCurrentMaxPhC", ZclDataTypeUint16, Read},
					0x0a0b: {"ActivePowerPhC", ZclDataTypeInt16, Read},
					0x0a0c: {"ActivePowerMinPhC", ZclDataTypeInt16, Read},
					0x0a0d: {"ActivePowerMaxPhC", ZclDataTypeInt16, Read},
					0x0a0e: {"ReactivePowerPhC", ZclDataTypeInt16, Read},
					0x0a0f: {"ApparentPowerPhC", ZclDataTypeUint16, Read},
					0x0a10: {"PowerFactorPhC", ZclDataTypeInt8, Read},
					0x0a11: {"AverageRMSVoltageMeasurementPeriodPhC", ZclDataTypeUint16, Read | Write},
					0x0a12: {"AverageRMSOverVoltageCounterPhC", ZclDataTypeUint16, Read | Write},
					0x0a13: {"AverageRMSUnderVoltageCounterPhC", ZclDataTypeUint16, Read | Write},
					0x0a14: {"RMSExtremeOverVoltagePeriodPhC", ZclDataTypeUint16, Read | Write},
					0x0a15: {"RMSExtremeUnderVoltagePeriodPhC", ZclDataTypeUint16, Read | Write},
					0x0a16: {"RMSVoltageSagPeriodPhC", ZclDataTypeUint16, Read | Write},
					0x0a17: {"RMSVoltageSwellPeriodPhC", ZclDataTypeUint16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"GetProfileInfoResponse", &GetProfileInfoResponse{}},
						0x01: {"GetMeasurementProfileResponse", &GetMeasurementProfileResponse{}},
					},
					Generated: map[uint8]*CommandDescriptor{
						0x00: {"GetProfileInfoCommand", &GetProfileInfoCommand{}},
						0x01: {"GetMeasurementProfileCommand", &GetMeasurementProfileCommand{}},
					},
				},
			},
			IASZone: {
				Name: "IASZone",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"ZoneState", ZclDataTypeEnum8, Read},
					0x0001: {"ZoneType", ZclDataTypeEnum16, Read},
					0x0002: {"ZoneStatus", ZclDataTypeBitmap16, Read},
					0x0010: {"IAS_CIE_Address", ZclDataTypeIeeeAddr, Read | Write},
					0x0011: {"ZoneID", ZclDataTypeUint8, Read},
					0x0012: {"NumberOfZoneSensitivityLevelsSupported", ZclDataTypeUint8, Read},
					0x0013: {"CurrentZoneSensitivityLevel", ZclDataTypeUint8, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"ZoneEnrollResponse", &ZoneEnrollResponse{}},
						0x01: {"InitiateNormalOperationMode", &InitiateNormalOperationModeCommand{}},
						0x02: {"InitiateTestMode", &InitiateTestModeCommand{}},
					},
					Generated: map[uint8]*CommandDescriptor{
						0x00: {"ZoneStatusChangeNotification", &ZoneStatusChangeNotificationCommand{}},
						0x01: {"ZoneEnrollRequest", &ZoneEnrollCommand{}},
					},
				},
			},
			IASACE: {
				Name: "IASAncillaryControlEquipment",
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"Arm", &ArmCommand{}},
						0x01: {"Bypass", &BypassCommand{}},
						0x02: {"Emergency", &EmergencyCommand{}},
						0x03: {"Fire", &FireCommand{}},
						0x04: {"Panic", &PanicCommand{}},
						0x05: {"GetZoneIDMap", &GetZoneIDMapCommand{}},
						0x06: {"GetZoneInformation", &GetZoneInformationCommand{}},
						0x07: {"GetPanelStatus", &GetPanelStatusCommand{}},
						0x08: {"GetBypassedZoneList", &GetBypassedZoneListCommand{}},
						0x09: {"GetZoneStatus", &GetZoneStatus{}},
					},
					Generated: map[uint8]*CommandDescriptor{
						0x00: {"ArmResponse", &ArmResponse{}},
						0x01: {"GetZoneIDMapResponse", &GetZoneIDMapResponse{}},
						0x02: {"GetZoneInformationResponse", &GetZoneInformationResponse{}},
						0x03: {"ZoneStatusChanged", &ZoneStatusChanged{}},
						0x04: {"PanelStatusChanged", &PanelStatusChanged{}},
						0x05: {"GetPanelStatusResponse", &PanelStatusChanged{}},
						0x06: {"SetBypassedZoneList", &BypassedZoneList{}},
						0x07: {"BypassResponse", &BypassedZoneList{}},
						0x08: {"GetZoneStatusResponse", &GetZoneStatusResponse{}},
					},
				},
			},
			IASWarningDevice: {
				Name: "IASWarningDevice",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: {"MaxDuration", ZclDataTypeUint16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: {"StartWarning", &StartWarning{}},
						0x01: {"Squawk", &Squark{}},
					},
				},
			},
		},
	}
}

func (cl *ClusterLibrary) Clusters() map[ClusterId]*Cluster {
	return cl.clusters
}

func (cl *ClusterLibrary) Global() map[uint8]*CommandDescriptor {
	return cl.global
}
