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

// "testing"

func ExampleGetStdTerminalLogger() {

	l := GetStdTerminalLogger()
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
	// Alert [TestAlert]
	// Alert [TestAlertf Extra]
	// Emergency [TestEmergency]
	// Emergency [TestEmergencyf Extra]
	// Critical [TestCritical]
	// Critical [TestCriticalf Extra]
	// Error [TestError]
	// Error [TestErrorf Extra]
	// Warning [TestWarning]
	// Warning [TestWarningf Extra]
	// Notice [TestNotice]
	// Notice [TestNoticef Extra]
	// Info [TestInfo]
	// Info [TestInfof Extra]
	// Debug [TestDebug]
	// Debug [TestDebugf Extra]
	// None [TestLog]
	// None [TestLogf Extra]

}
