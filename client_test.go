package abb

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestIOSignals(t *testing.T) {

	signals := IOSignals{}
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
			case "type":
				sigType = span.Content
			case "lvalue":
				lvalue = span.Content
			}
		}
		fmt.Printf("Name: %s, Type: %s, Value: %s\n", name, sigType, lvalue)
	}
}

func TestControllerActions(t *testing.T) {

	actions := ControllerActions{}
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
		fmt.Println("Option value:", option.Value)
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
