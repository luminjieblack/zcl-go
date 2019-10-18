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
	MultistateInput                ClusterId = 0x0012
	OTA                            ClusterId = 0x0019
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
			IASZone: {
				Name: "IAS Zone",
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
				Name: "IAS Ancillary Control Equipment",
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
				Name: "IAS Warning Device",
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
