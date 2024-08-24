package abb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/atmassey/abb-lib-rws/structures"
)

func TestControllerActions(t *testing.T) {

	actions := structures.ControllerActionsHTML{}
	actions_struct := structures.ControllerActions{}
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
	mode := structures.OperationMode{}
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
	var EnergyMetricsRaw structures.SystemEnergy
	var EnergyMetricsDecoded structures.SystemEnergyMetrics
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
				axisEnergy := structures.SystemAxisEnergy{Axis: axis.Title, Energy: axis.IntervalEnergy}
				EnergyMetricsDecoded.AxisEnergy = append(EnergyMetricsDecoded.AxisEnergy, axisEnergy)
			}
		}
	}
	fmt.Printf("Accumulated Energy: %s\n", EnergyMetricsDecoded.AccumulatedEnergy)
	for _, axis := range EnergyMetricsDecoded.AxisEnergy {
		fmt.Printf("Axis: %s, Energy: %s\n", axis.Axis, axis.Energy)
	}
}

func TestGetIOSignalsJson(t *testing.T) {
	signals_raw := IOSignalsJson{}
	signals := IOSignals{}
	//sample response from the api documentation
	data := `{
    "_links": {
        "base": {
            "href": "http://localhost:80/rw/iosystem/"
        },
        "next": {
            "href": "signals?start=100&limit=100&json=1"
        }
    },
    "_embedded": {
        "_state": [
            {
                "_links": {
                    "self": {
                        "href": "signals/PROFINET/pn_57380_DO8xDC24V_0_5A/do_Paint_Nozzle_Trig?json=1"
                    }
                },
                "_type": "ios-signal-li",
                "_title": "PROFINET/pn_57380_DO8xDC24V_0_5A/do_Paint_Nozzle_Trig",
                "name": "doTest1",
                "type": "DO",
                "category": "",
                "lvalue": 1,
                "lstate": "not simulated"
            },
            {
                "_links": {
                    "self": {
                        "href": "signals/PROFINET/SLAVE_PLC/soRBT_Safety_OK?json=1"
                    }
                },
                "_type": "ios-signal-li",
                "_title": "PROFINET/SLAVE_PLC/soRBT_Safety_OK",
                "name": "doTest2",
                "type": "DO",
                "category": "ProfiSafe",
                "lvalue": 0,
                "lstate": "not simulated"
            },
            {
                "_links": {
                    "self": {
                        "href": "signals/PROFINET/SLAVE_PLC/soRBT_Sync_OK?json=1"
                    }
                },
                "_type": "ios-signal-li",
                "_title": "PROFINET/SLAVE_PLC/soRBT_Sync_OK",
                "name": "doTest3",
                "type": "DO",
                "category": "ProfiSafe",
                "lvalue": 1,
                "lstate": "not simulated"
            },
            {
                "_links": {
                    "self": {
                        "href": "signals/PROFINET/SLAVE_PLC/soRBT_AutoStop_OK?json=1"
                    }
                },
                "_type": "ios-signal-li",
                "_title": "PROFINET/SLAVE_PLC/soRBT_AutoStop_OK",
                "name": "doTest4",
                "type": "DO",
                "category": "ProfiSafe",
                "lvalue": 0,
                "lstate": "not simulated"
            },
            {
                "_links": {
                    "self": {
                        "href": "signals/PROFINET/SLAVE_PLC/soRBT_GenStop_OK?json=1"
                    }
                },
                "_type": "ios-signal-li",
                "_title": "PROFINET/SLAVE_PLC/soRBT_GenStop_OK",
                "name": "doTest5",
                "type": "DO",
                "category": "ProfiSafe",
                "lvalue": 1,
                "lstate": "not simulated"
            }
    	]
		}
	}`
	err := json.Unmarshal([]byte(data), &signals_raw)
	if err != nil {
		t.Errorf("Error decoding response: %s", err)
	}
	for _, state := range signals_raw.Embedded.State {
		signals.SignalName = append(signals.SignalName, state.Name)
		signals.SignalType = append(signals.SignalType, state.Type)
		signals.SignalValue = append(signals.SignalValue, state.Value)
	}
	for i, signal := range signals.SignalName {
		fmt.Printf("Signal: %s, Type: %s, Value: %v\n", signal, signals.SignalType[i], signals.SignalValue[i])
	}
}

func TestGetMechUnit(t *testing.T) {
	mechUnits := structures.MechUnitsJson{}
	mechUnitsDecoded := structures.MechUnits{}
	//sample response from the api documentation
	data := `{
    "_links": {
        "base": {
            "href": "http://10.40.36.102:80/rw/motionsystem/"
        }
    },
    "_embedded": {
        "_state": [
            {
                "_type": "ms-mechunit-li",
                "_title": "ROB_1",
                "_links": {
                    "self": {
                        "href": "mechunits/ROB_1?json=1"
                    }
                },
                "mode": "Activated",
                "activation-allowed": "True",
                "drive-module": "0"
            }
        ]
    }
}`
	err := json.Unmarshal([]byte(data), &mechUnits)
	if err != nil {
		t.Errorf("Error decoding response: %s", err)
	}
	for _, mechUnit := range mechUnits.Embedded.State {
		mechUnitsDecoded.ActivationAllowed = append(mechUnitsDecoded.ActivationAllowed, mechUnit.ActivationAllowed)
		mechUnitsDecoded.DriveModule = append(mechUnitsDecoded.DriveModule, mechUnit.DriveModule)
		mechUnitsDecoded.Mode = append(mechUnitsDecoded.Mode, mechUnit.Mode)
		mechUnitsDecoded.Title = append(mechUnitsDecoded.Title, mechUnit.Title)
	}
	for i, unit := range mechUnitsDecoded.Title {
		fmt.Printf("Title: %s, Mode: %s, Activation Allowed: %s, Drive Module: %s\n", unit, mechUnitsDecoded.Mode[i], mechUnitsDecoded.ActivationAllowed[i], mechUnitsDecoded.DriveModule[i])
	}
}
