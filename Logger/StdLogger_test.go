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

func ExampleGetStdTerminalLogger() {

	// mock time for testing
	now = func() string { return "2017-03-11 08:55:34.043958533 +0100 CET" }

	// ....

	// Get a logger to stdout
	l := GetStdTerminalLogger()

	// Lst's log
	l.Alert("TestAlert")
	l.Alertf("TestAlertf %v", "Extra")
	l.Emergency("TestEmergency")
	l.Emergencyf("TestEmergencyf %v", "Extra")
	l.Critical("TestCritical")
	l.Criticalf("TestCriticalf %v", "Extra")
	l.Error("TestError")
	l.Errorf("TestErrorf %v", "Extra")
	l.Warning("TestWarning")
	l.Warningf("TestWarningf %v", "Extra")
	l.Notice("TestNotice")
	l.Noticef("TestNoticef %v", "Extra")
	l.Info("TestInfo")
	l.Infof("TestInfof %v", "Extra")
	l.Debug("TestDebug")
	l.Debugf("TestDebugf %v", "Extra")
	l.Log(None, "TestLog")
	l.Logf(None, "TestLogf %v", "Extra")

	// Output:
	// 2017-03-11 08:55:34.043958533 +0100 CET [Alert] TestAlert
	// 2017-03-11 08:55:34.043958533 +0100 CET [Alert] TestAlertf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Emergency] TestEmergency
	// 2017-03-11 08:55:34.043958533 +0100 CET [Emergency] TestEmergencyf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Critical] TestCritical
	// 2017-03-11 08:55:34.043958533 +0100 CET [Critical] TestCriticalf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Error] TestError
	// 2017-03-11 08:55:34.043958533 +0100 CET [Error] TestErrorf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Warning] TestWarning
	// 2017-03-11 08:55:34.043958533 +0100 CET [Warning] TestWarningf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Notice] TestNotice
	// 2017-03-11 08:55:34.043958533 +0100 CET [Notice] TestNoticef Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Info] TestInfo
	// 2017-03-11 08:55:34.043958533 +0100 CET [Info] TestInfof Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [Debug] TestDebug
	// 2017-03-11 08:55:34.043958533 +0100 CET [Debug] TestDebugf Extra
	// 2017-03-11 08:55:34.043958533 +0100 CET [None] TestLog
	// 2017-03-11 08:55:34.043958533 +0100 CET [None] TestLogf Extra

}
