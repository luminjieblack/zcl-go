package cluster

type ResetToFactoryDefaultsCommand struct {
}

type IdentifyCommand struct {
	IdentifyTime uint16
}

type IdentifyQueryCommand struct{}

type TriggerEffectCommand struct {
	EffectIdentifier uint8
	EffectVariant    uint8
}

type IdentifyQueryResponse struct {
	Timeout uint16
}

type OffCommand struct{}

type OnCommand struct{}

type ToggleCommand struct{}

type OffWithEffectCommand struct {
	EffectIdentifier uint8
	EffectVariant    uint8
}

type OnWithRecallGlobalSceneCommand struct{}

type OnWithTimedOffCommand struct {
	OnOffControl uint8
	OnTime       uint16
	OffWaitTime  uint16
}

type MoveToLevelCommand struct {
	Level          uint8
	TransitionTime uint16
}

type MoveCommand struct {
	MoveMode uint8
	Rate     uint8
}

type StepCommand struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type StopCommand struct{}

type MoveToLevelOnOffCommand struct {
	Level          uint8
	TransitionTime uint16
}

type MoveOnOffCommand struct {
	MoveMode uint8
	Rate     uint8
}

type StepOnOffCommand struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type StopOnOffCommand struct{}

type StartWarning struct {
	WarningControl  uint8
	WarningDuration uint16
	StrobeDutyCycle uint8
	StrobeLevel     uint8
}

type Squark struct {
	SquarkControl uint8
}

type ZoneEnrollResponse struct {
	ResponseCode uint8
	ZoneID       uint8
}

type InitiateNormalOperationModeCommand struct {
}

type InitiateTestModeCommand struct {
	TestModeDuration            uint8
	CurrentZoneSensitivityLevel uint8
}

type ZoneStatusChangeNotificationCommand struct {
	ZoneStatus     uint16
	ExtendedStatus uint8
	ZoneID         uint8
	Delay          uint16
}

type ZoneEnrollCommand struct {
	ZoneType         uint16
	ManufacturerCode uint16
}

type ArmCommand struct {
	ArmMode       uint8
	ArmDisarmCode string
	ZoneID        uint8
}

type BypassCommand struct {
	NumberOfZones uint8
	ZoneID        []uint8
	ArmDisarmCode string
}

type EmergencyCommand struct{}

type FireCommand struct{}

type PanicCommand struct{}

type GetZoneIDMapCommand struct{}

type GetZoneInformationCommand struct {
	ZoneID uint8
}
type GetPanelStatusCommand struct{}

type GetBypassedZoneListCommand struct{}

type GetZoneStatus struct {
	StartingZoneID     uint8
	MaxNumberZoneIDs   uint8
	ZoneStatusMaskFlag bool
	ZoneStatusMask     uint16
}

type ArmResponse struct {
	ArmNotification uint8
}

type GetZoneIDMapResponse struct {
	ZoneIDMapSection0  uint16
	ZoneIDMapSection1  uint16
	ZoneIDMapSection2  uint16
	ZoneIDMapSection3  uint16
	ZoneIDMapSection4  uint16
	ZoneIDMapSection5  uint16
	ZoneIDMapSection6  uint16
	ZoneIDMapSection7  uint16
	ZoneIDMapSection8  uint16
	ZoneIDMapSection9  uint16
	ZoneIDMapSection10 uint16
	ZoneIDMapSection11 uint16
	ZoneIDMapSection12 uint16
	ZoneIDMapSection13 uint16
	ZoneIDMapSection14 uint16
	ZoneIDMapSection15 uint16
}

type GetZoneInformationResponse struct {
	ZoneId      uint8
	ZoneType    uint16
	IEEEAddress [6]byte
	ZoneLabel   string
}

type ZoneStatusChanged struct {
	ZoneId              uint8
	ZoneStatus          uint16
	AudibleNotification uint8
	ZoneLabel           string
}

type PanelStatusChanged struct {
	PanelStatus         uint8
	SecondsRemaining    uint8
	AudibleNotification uint8
	AlarmStatus         uint8
}

type BypassedZoneList struct {
	NumberOfZones uint8
	ZoneID        []uint8
}

type GetZoneStatusResponse struct {
	ZoneStatusComplete bool
	NumberOfZones      uint8
	ZoneID             []uint16
	ZoneStatus         []uint16
}

type GetProfileInfoResponse struct {
	ProfileCount          uint8
	ProfileIntervalPeriod uint8
	MaxNumberOfIntervals  uint8
	ListOfAttributes      []uint16
}

type GetProfileInfoCommand struct {
}

type GetMeasurementProfileResponse struct {
	StartTime                  uint32
	Status                     uint8
	ProfileIntervalPeriod      uint8
	NumberOfIntervalsDelivered uint8
	AttributeId                uint8
	AttributeValues            []uint16
}

type GetMeasurementProfileCommand struct {
	AttributeID       uint16
	StartTime         uint32
	NumberOfIntervals uint8
}

type ResetAlarmCommand struct {
	AlarmCode         uint8
	ClusterIdentifier uint16
}

type ResetAllAlarmsCommand struct{}

type GetAlarmCommand struct{}

type ResetAlarmLogCommand struct{}

type AlarmCommand struct {
	AlarmCode         uint8
	ClusterIdentifier uint16
}

type GetAlarmResponse struct {
	Status            uint8
	AlarmCode         uint8
	ClusterIdentifier uint16
	TimeStamp         uint32
}
