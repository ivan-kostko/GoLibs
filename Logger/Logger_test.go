//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package Logger

import (
	"reflect"
	"testing"
)

func TestLevelStringer(t *testing.T) {

	testCases := []struct {
		Name     string
		Actual   string
		Expected string
	}{
		{
			"Nane",
			None.String(),
			"None",
		},
		{
			"Emergency",
			Emergency.String(),
			"Emergency",
		},
		{
			"Critical",
			Critical.String(),
			"Critical",
		},
		{
			"Alert",
			Alert.String(),
			"Alert",
		},
		{
			"Error",
			Error.String(),
			"Error",
		},
		{
			"Warning",
			Warning.String(),
			"Warning",
		},
		{
			"Notice",
			Notice.String(),
			"Notice",
		},
		{
			"Info",
			Info.String(),
			"Info",
		},
		{
			"Debug",
			Debug.String(),
			"Debug",
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		expected := testCase.Expected
		actual := testCase.Actual
		if actual != expected {
			t.Errorf("Testing %v.String: returned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}

}

func TestLoggerContainerPrintLikeFunctions(t *testing.T) {
	// Just Print like functions

	originalArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	expectedArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	var actualArgs []interface{}

	var actualLevel Level

	reInitActualLevel := func() { actualLevel = 0 }
	ilf := func(level Level, args ...interface{}) { actualLevel = level; actualArgs = args }

	l := GetNewLogAdapter(ilf)

	reInitActualLevel()

	testCases := []struct {
		Name          string
		Func          func()
		ExpectedLevel Level
	}{
		{
			"Log",
			func() { l.Log(None, originalArgs...) },
			None,
		},
		{
			"Emergency",
			func() { l.Emergency(originalArgs...) },
			Emergency,
		},
		{
			"Alert",
			func() { l.Alert(originalArgs...) },
			Alert,
		},
		{
			"Critical",
			func() { l.Critical(originalArgs...) },
			Critical,
		},
		{
			"Error",
			func() { l.Error(originalArgs...) },
			Error,
		},
		{
			"Warning",
			func() { l.Warning(originalArgs...) },
			Warning,
		},
		{
			"Notice",
			func() { l.Notice(originalArgs...) },
			Notice,
		},
		{
			"Info",
			func() { l.Info(originalArgs...) },
			Info,
		},
		{
			"Debug",
			func() { l.Debug(originalArgs...) },
			Debug,
		},
	}

	for _, testCase := range testCases {
		// Before any test - reinitialize actual to avoid influnce of previouse tests
		reInitActualLevel()

		name := testCase.Name
		fn := testCase.Func
		expectedLevel := testCase.ExpectedLevel

		// invoce test function
		fn()

		if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
			t.Errorf("l.%s passed %s %v when expected %s %v", name, actualLevel, actualArgs, expectedLevel, expectedArgs)
		}
	}

}

func TestLoggerContainerPrintFLikeFunctions(t *testing.T) {

	// Just PrintF like functions

	originalArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	formatString := " %s %s %s %v "
	expectedArgs := []interface{}{" Arg1 Arg2 Arg3 4 "}
	var actualArgs []interface{}

	var actualLevel Level

	reInitActualLevel := func() { actualLevel = 0 }
	ilf := func(level Level, args ...interface{}) { actualLevel = level; actualArgs = args }

	l := GetNewLogAdapter(ilf)

	reInitActualLevel()

	testCases := []struct {
		Name          string
		Func          func()
		ExpectedLevel Level
	}{
		{
			"Log",
			func() { l.Logf(None, formatString, originalArgs...) },
			None,
		},
		{
			"Emergency",
			func() { l.Emergencyf(formatString, originalArgs...) },
			Emergency,
		},
		{
			"Alert",
			func() { l.Alertf(formatString, originalArgs...) },
			Alert,
		},
		{
			"Critical",
			func() { l.Criticalf(formatString, originalArgs...) },
			Critical,
		},
		{
			"Error",
			func() { l.Errorf(formatString, originalArgs...) },
			Error,
		},
		{
			"Warning",
			func() { l.Warningf(formatString, originalArgs...) },
			Warning,
		},
		{
			"Notice",
			func() { l.Noticef(formatString, originalArgs...) },
			Notice,
		},
		{
			"Info",
			func() { l.Infof(formatString, originalArgs...) },
			Info,
		},
		{
			"Debug",
			func() { l.Debugf(formatString, originalArgs...) },
			Debug,
		},
	}

	for _, testCase := range testCases {
		// Before any test - reinitialize actual to avoid influnce of previouse tests
		reInitActualLevel()

		name := testCase.Name
		fn := testCase.Func
		expectedLevel := testCase.ExpectedLevel

		// invoce test function
		fn()

		if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
			t.Errorf("l.%s passed %s %v when expected %s %v", name, actualLevel, actualArgs, expectedLevel, expectedArgs)
		}
	}
}
