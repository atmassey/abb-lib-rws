package abb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestIOSignals(t *testing.T) {

	signals := IOSignalsHTML{}
	signals_struct := IOSignals{}
	//sample response from the api documentation
	signals_raw := `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>io</title>
    <base href="http://localhost/rw/iosystem/"/>
</head>
<body>
    <div class="state">
    <a href="signals" rel="self"/>
    <a href= "signals?action=show" rel="action"/>
    <ul>
        <li class="ios-signal-li" title="Local/DRV_1/DRV1TESTE2">
            <a href="signals/Local/DRV_1/DRV1TESTE2" rel="self"/>
            <span class="name">DRV1TESTE2</span>
            <span class="type">DO</span>
            <span class="category">safety</span>
            <span class="lvalue">0</span>
            <span class="lstate">simulated</span>
        </li>
        <li class="ios-signal-li" title="Local/DRV_1/DRV1BRAKE">
            <a href="signals/Local/DRV_1/DRV1BRAKE" rel="self"/>
            <span class="name">DRV1BRAKE</span>
            <span class="type">DO</span>
            <span class="category">safety</span>
            <span class="lvalue">0</span>
            <span class="lstate">simulated</span>
        </li>
    </ul>
    </div>
</body>
</html>`
	err := xml.Unmarshal([]byte(signals_raw), &signals)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Signals: %d\n", len(signals.Body.Div.UL.LIs))

	for _, signal := range signals.Body.Div.UL.LIs {
		name, sigType, lvalue := "", "", ""
		for _, span := range signal.Spans {
			switch span.Class {
			case "name":
				name = span.Content
				signals_struct.SignalName = append(signals_struct.SignalName, span.Content)
			case "type":
				sigType = span.Content
				signals_struct.SignalType = append(signals_struct.SignalType, span.Content)
			case "lvalue":
				lvalue = span.Content
				signals_struct.SignalValue = append(signals_struct.SignalValue, span.Content)
			}
		}
		fmt.Printf("Name: %s, Type: %s, Value: %s\n", name, sigType, lvalue)
	}
	fmt.Println("Struct:")
	for i, name := range signals_struct.SignalName {
		fmt.Printf("Name: %s, Type: %s, Value: %s\n", name, signals_struct.SignalType[i], signals_struct.SignalValue[i])
	}
}

func TestControllerActions(t *testing.T) {

	actions := ControllerActionsHTML{}
	actions_struct := ControllerActions{}
	//sample response from the api documentation
	actions_raw := `<?xml version="1.0" encoding="utf-8"?>
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
		<head>
			<title>controller</title>
			<base href="http://localhost/ctrl/" />
		</head>
		<body>
			<div class="state">
				<form name="ctrl-restart" method='post' action="/ctrl">
					<select name="restart-mode">
					<option value="restart"></option>
					<option value="shutdown"></option>
					<option value="xstart"></option>
					<option value="istart"></option>
					<option value="pstart"></option>
					<option value="bstart"></option>
				</select>
				</form>
				<form id="set-ctrl-lang" method="post" action="?action=set-lang">
					<input name="lang" type="text"/>
				</form>
			</div>
		</body>
	</html>`
	err := xml.Unmarshal([]byte(actions_raw), &actions)
	if err != nil {
		t.Error(err)
	}
	for _, option := range actions.Body.Div.Select.Options {
		actions_struct.Actions = append(actions_struct.Actions, option.Value)
	}
	for _, action := range actions_struct.Actions {
		fmt.Printf("Action: %s\n", action)
	}
}

func TestRobotType(t *testing.T) {

	robotType := RobotType{}
	//sample response from the api documentation
	robotType_raw := `<?xml version="1.0" encoding="UTF-8"?>
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head>
    <title>system</title>
    <base href="http://localhost/rw/system/robottype/"/>
    </head>
    <body>
    <div class="state">
    <a href="" rel="self"/>
    <ul>
    <li class="sys-robottype" title="1">
    <span class="robot-type">IRB 120-3/0.6</span>
    </li>
    <li class="sys-robottype" title="2">
    <span class="robot-type">IRB 140T-5/0.8 Type C</span>
    </li>
    </ul>
    </div>
    </body>
    </html>`
	err := xml.Unmarshal([]byte(robotType_raw), &robotType)
	if err != nil {
		t.Error(err)
	}
	for _, robot := range robotType.Body.State.Robots {
		fmt.Printf("Robot Type: %s, Title: %s\n", robot.RobotType, robot.Title)
	}
}

func TestControllerMode(t *testing.T) {
	mode := OperationMode{}
	//sample response from the api documentation
	mode_raw := `{
    "_links": {
        "base": {
            "href": "http://10.40.36.102:80/rw/panel/opmode/"
        }
    },
    "_embedded": {
        "_state": [
            {
                "_type": "pnl-opmode",
                "_title": "opmode",
                "opmode": "AUTO"
            }
        ]
    }
}
`
	err := json.Unmarshal([]byte(mode_raw), &mode)
	if err != nil {
		t.Error(err)
	}
	for _, state := range mode.Embedded.State {
		fmt.Printf("Operation Mode: %s\n", state.Opmode)
	}
}

func TestSystemEnergyMetrics(t *testing.T) {
	var EnergyMetricsRaw SystemEnergy
	var EnergyMetricsDecoded SystemEnergyMetrics
	data := `{
		"_links": {
			"base": {
				"href": "http://10.40.36.103:80/rw/system/energy/"
			}
		},
		"_embedded": {
			"_state": [
				{
					"_links": {
						"self": {
							"href": "?json=1"
						}
					},
					"_type": "sys-energy-state-li",
					"_title": "energy-state",
					"state": "0",
					"energy-state": "blocked",
					"change-count": "14640",
					"time-stamp": "2024-07-16 T 12:00:00",
					"reset-time": "2021-08-12 T 05:53:05",
					"interval-length": "3600",
					"interval-energy": "0",
					"accumulated-energy": "458726901.0453996",
					"mechunits": [
						{
							"_type": "sys-energy-mec-li",
							"_title": "ROB_1",
							"axes": [
								{
									"_type": "sys-energy-axis-li",
									"_title": "1",
									"interval-energy": "0"
								},
								{
									"_type": "sys-energy-axis-li",
									"_title": "2",
									"interval-energy": "0"
								},
								{
									"_type": "sys-energy-axis-li",
									"_title": "3",
									"interval-energy": "0"
								},
								{
									"_type": "sys-energy-axis-li",
									"_title": "4",
									"interval-energy": "0"
								},
								{
									"_type": "sys-energy-axis-li",
									"_title": "5",
									"interval-energy": "0"
								},
								{
									"_type": "sys-energy-axis-li",
									"_title": "6",
									"interval-energy": "0"
								}
							]
						}
					]
				}
			]
		}
	}`
	err := json.Unmarshal([]byte(data), &EnergyMetricsRaw)
	if err != nil {
		t.Errorf("Error decoding response: %s", err)
	}
	EnergyMetricsDecoded.AccumulatedEnergy = EnergyMetricsRaw.Embedded.State[0].AccumulatedEnergy
	for _, state := range EnergyMetricsRaw.Embedded.State {
		for _, MechUnits := range state.MechUnits {
			for _, axis := range MechUnits.Axis {
				axisEnergy := SystemAxisEnergy{Axis: axis.Title, Energy: axis.IntervalEnergy}
				EnergyMetricsDecoded.AxisEnergy = append(EnergyMetricsDecoded.AxisEnergy, axisEnergy)
			}
		}
	}
	fmt.Printf("Accumulated Energy: %s\n", EnergyMetricsDecoded.AccumulatedEnergy)
	for _, axis := range EnergyMetricsDecoded.AxisEnergy {
		fmt.Printf("Axis: %s, Energy: %s\n", axis.Axis, axis.Energy)
	}
}
