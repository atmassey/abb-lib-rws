package structures

type MotionErrorStateJson struct {
	Links    MotionErrorStateJsonLinks `json:"_links"`
	Embedded MotionErrorStateJsonState `json:"_embedded"`
}

type MotionErrorStateJsonLinks struct {
	Base MotionErrorStateJsonBase `json:"base"`
}

type MotionErrorStateJsonBase struct {
	Href string `json:"href"`
}

type MotionErrorStateJsonState struct {
	State []MotionErrorStateJsonMeta `json:"_state"`
}

type MotionErrorStateJsonMeta struct {
	Type  string `json:"_type"`
	Title string `json:"_title"`
	State string `json:"err-state"`
	Count string `json:"err-count"`
}

type MotionErrorState struct {
	State []string
	Count []string
}

type MechUnits struct {
	Title             []string
	Mode              []string
	DriveModule       []string
	ActivationAllowed []string
}

type MechUnitsJson struct {
	Links    MechUnitsJsonLinks `json:"_links"`
	Embedded MechUnitsJsonState `json:"_embedded"`
}

type MechUnitsJsonLinks struct {
	Base MechUnitsJsonBase `json:"base"`
}

type MechUnitsJsonBase struct {
	Href string `json:"href"`
}

type MechUnitsJsonState struct {
	State []MechUnitsJsonMeta `json:"_state"`
}

type MechUnitsJsonMeta struct {
	Type              string                 `json:"_type"`
	Title             string                 `json:"_title"`
	Links             MechUnitsJsonMetaLinks `json:"_links"`
	Mode              string                 `json:"mode"`
	ActivationAllowed string                 `json:"activation-allowed"`
	DriveModule       string                 `json:"drive-module"`
}

type MechUnitsJsonMetaLinks struct {
	Self MechUnitsJsonMetaLinksSelf `json:"self"`
}

type MechUnitsJsonMetaLinksSelf struct {
	Href string `json:"href"`
}

type SystemEnergy struct {
	Links    SystemEneryLinks     `json:"_links"`
	Embedded SystemEnergyEmbedded `json:"_embedded"`
}

type SystemEneryLinks struct {
	Base SystemEneryBase `json:"base"`
}

type SystemEneryBase struct {
	Href string `json:"href"`
}

type SystemEnergyEmbedded struct {
	State []StateListLinks `json:"_state"`
}

type StateListLinks struct {
	Links             StateListLinksSelf `json:"_links"`
	Type              string             `json:"_type"`
	Title             string             `json:"_title"`
	State             string             `json:"state"`
	EnergyState       string             `json:"energy-state"`
	ChangeCount       string             `json:"change-count"`
	TimeStamp         string             `json:"time-stamp"`
	ResetTime         string             `json:"reset-time"`
	IntervalLength    string             `json:"interval-length"`
	IntervalEnergy    string             `json:"interval-energy"`
	AccumulatedEnergy string             `json:"accumulated-energy"`
	MechUnits         []SystemEnergyAxes `json:"mechunits"`
}

type StateListLinksSelf struct {
	Self StateListLinksSelfHref `json:"self"`
}

type StateListLinksSelfHref struct {
	Href string `json:"href"`
}

type SystemEnergyAxes struct {
	Type  string             `json:"_type"`
	Title string             `json:"_title"`
	Axis  []SystemEnergyAxis `json:"axes"`
}

type SystemEnergyAxis struct {
	Type           string `json:"_type"`
	Title          string `json:"_title"`
	IntervalEnergy string `json:"interval-energy"`
}

type SystemEnergyMetrics struct {
	AccumulatedEnergy string
	AxisEnergy        []SystemAxisEnergy
}

type SystemAxisEnergy struct {
	Axis   string
	Energy string
}

type InstalledSystemProducts struct {
	Title   []string
	Version []string
}

type InstalledProducts struct {
	Links InstalledProductsLinks   `json:"_links"`
	State []InstalledProductsState `json:"state"`
}

type InstalledProductsLinks struct {
	Base InstalledProductsBase `json:"base"`
	Self InstalledProductsSelf `json:"self"`
}

type InstalledProductsBase struct {
	Href string `json:"href"`
}

type InstalledProductsSelf struct {
	Href string `json:"href"`
}

type InstalledProductsState struct {
	Type        string `json:"_type"`
	Title       string `json:"_title"`
	VersionName string `json:"version-name"`
}
