package structures

type IOSignalsJson struct {
	Links    IOSignalsJsonLinks `json:"_links"`
	Embedded IOSignalsJsonState `json:"_embedded"`
}

type IOSignalsJsonLinks struct {
	Base IOSignalsJsonBase `json:"base"`
}

type IOSignalsJsonBase struct {
	Href string `json:"href"`
}

type IOSignalsJsonState struct {
	State []IOSignalsJsonMeta `json:"_state"`
}

type IOSignalsJsonMeta struct {
	Links           IOSignalsJsonMetaLinks `json:"_links"`
	TypeT           string                 `json:"_type"`
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	Category        string                 `json:"category"`
	Value           int                    `json:"lvalue"`
	SimulationState string                 `json:"lstate"`
}

type IOSignalsJsonMetaLinks struct {
	Self IOSignalsJsonMetaLinksSelf `json:"self"`
}

type IOSignalsJsonMetaLinksSelf struct {
	Href string `json:"href"`
}

type IOSignals struct {
	SignalName  []string
	SignalType  []string
	SignalValue []int
}
