package WorkerPool

import (
	"errors"
	"testing"
	"time"
)

func Test_NewWorkerPool(t *testing.T) {
	testCases := []struct {
		TestAlias            string
		InitWorkerNumber     int
		ExpectedWorkerNumber int
		ExpectedChanCapacity int
	}{
		{
			TestAlias:            "Simple 0 workers",
			InitWorkerNumber:     0,
			ExpectedWorkerNumber: 0,
			ExpectedChanCapacity: 0,
		},
		{
			TestAlias:            "Simple 10 workers",
			InitWorkerNumber:     10,
			ExpectedWorkerNumber: 10,
			ExpectedChanCapacity: 10,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		expectedWorkerNumber := testCase.ExpectedWorkerNumber
		expectedChanCapacity := testCase.ExpectedChanCapacity

		fn := func(t *testing.T) {
			id := NewWorkerPool(initWorkerNumber)

			d, ok := id.(*workerPool)

			if !ok {
				t.Skipf("\r\n For TestAlias: '%s' the NewWorkerPool(%v)  returned unknown implementation of WorkerPool inteface\r\n", initWorkerNumber, testAlias)
			}

			actualWorkerNumber := len(d.workersChan)

			if actualWorkerNumber != expectedWorkerNumber {
				t.Errorf("\r\n For TestAlias: '%s' the NewWorkerPool(%v)  returned department{} \r\n with numberOfWorkers = %v \r\n while expected %v \r\n", testAlias, initWorkerNumber, actualWorkerNumber, expectedWorkerNumber)
			}

			actualChanCapacity := cap(d.workersChan)

			if actualChanCapacity != expectedChanCapacity {
				t.Errorf("\r\n For TestAlias: '%s' the NewWorkerPool(%v)  returned department{} \r\n with workersPool cap = %v \r\n while expected %v \r\n", testAlias, initWorkerNumber, actualChanCapacity, expectedChanCapacity)
			}

		}
		t.Run(testAlias, fn)

	}
}

func Test_WorkerPoolDoWorkersLimit(t *testing.T) {

	testCases := []struct {
		TestAlias              string
		InitWorkerNumber       int
		StartWorkerNumber      int
		ExpectedStartedWorkers int
	}{
		{
			TestAlias:              "20 work items for 0 workers",
			InitWorkerNumber:       0,
			StartWorkerNumber:      20,
			ExpectedStartedWorkers: 0,
		},
		{
			TestAlias:              "1 work item for 1 worker",
			InitWorkerNumber:       1,
			StartWorkerNumber:      1,
			ExpectedStartedWorkers: 1,
		},
		{
			TestAlias:              "1 work item for 10 workers",
			InitWorkerNumber:       10,
			StartWorkerNumber:      1,
			ExpectedStartedWorkers: 1,
		},
		{
			TestAlias:              "20 work items for 16 workers",
			InitWorkerNumber:       16,
			StartWorkerNumber:      20,
			ExpectedStartedWorkers: 16,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedStartedWorkers := testCase.ExpectedStartedWorkers

		fn := func(t *testing.T) {

			// the chanel where workers report about start
			started := make(chan struct{})

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerNumber := 0

			// Aggregate number of started workers
			go func() {
				for {
					_, more := <-started
					if more {
						actualWorkerNumber++
					} else {
						break
					}
				}
			}()

			timeOut := time.Duration(10000000)

			for i := 0; i < startWorkerNumber; i++ {
				dep.Do(func() { started <- struct{}{}; <-block }, timeOut)
			}

			// Sleep for all timeouts
			time.Sleep(timeOut * time.Duration(initWorkerNumber))

			actualStartedWorkers := actualWorkerNumber

			// cancel workers
			close(block)

			if actualStartedWorkers != expectedStartedWorkers {
				t.Errorf("For TestAlias '%s' WorkerPool.Do()  \r\n started %v workers \r\n while expected %v \r\n", testAlias, actualStartedWorkers, expectedStartedWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_WorkerPoolDoProcessAllWorkers(t *testing.T) {

	testCases := []struct {
		TestAlias            string
		InitWorkerNumber     int
		StartWorkerNumber    int
		ExpectedDonedWorkers int
	}{
		{
			TestAlias:            "20 work items for 0 workers",
			InitWorkerNumber:     0,
			StartWorkerNumber:    20,
			ExpectedDonedWorkers: 0,
		},
		{
			TestAlias:            "1 work item for 1 worker",
			InitWorkerNumber:     1,
			StartWorkerNumber:    1,
			ExpectedDonedWorkers: 1,
		},
		{
			TestAlias:            "1 work item for 10 workers",
			InitWorkerNumber:     10,
			StartWorkerNumber:    1,
			ExpectedDonedWorkers: 1,
		},
		{
			TestAlias:            "20 work items for 16 workers",
			InitWorkerNumber:     16,
			StartWorkerNumber:    20,
			ExpectedDonedWorkers: 20,
		},
		{
			TestAlias:            "2000 work items for 16 workers",
			InitWorkerNumber:     16,
			StartWorkerNumber:    2000,
			ExpectedDonedWorkers: 2000,
		},
		{
			TestAlias:            "2000 work items for 1 workers",
			InitWorkerNumber:     1,
			StartWorkerNumber:    2000,
			ExpectedDonedWorkers: 2000,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedDoneWorkers := testCase.ExpectedDonedWorkers

		fn := func(t *testing.T) {

			// the chanel where workers report about done
			done := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerDone := 0

			// Aggregate number of started workers
			go func() {
				for {

					_, more := <-done
					if more {
						actualWorkerDone++

					} else {
						break
					}
				}
			}()

			timeOut := time.Duration(10000000)

			for i := 0; i < startWorkerNumber; i++ {
				dep.Do(func() {
					done <- struct{}{}
				}, timeOut)
			}

			// Sleep for all timeouts
			time.Sleep(timeOut * time.Duration(initWorkerNumber))

			// wait while all left assignments are done
			actualDoneWorkers := actualWorkerDone

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' WorkerPool.Do()  \r\n has done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_WorkerPoolReadyForGCAfterClose(t *testing.T) {

	testCases := []struct {
		TestAlias                       string
		InitWorkerNumber                int
		StartWorkerNumber               int
		ExpectedWorkerChanLen           int
		ExpectedIsCancelationChanClosed bool
		ExpectedIsWorkersChanClosed     bool
	}{
		{
			TestAlias:                       "20 work items for 0 workers",
			InitWorkerNumber:                0,
			StartWorkerNumber:               20,
			ExpectedWorkerChanLen:           0,
			ExpectedIsCancelationChanClosed: true,
			ExpectedIsWorkersChanClosed:     true,
		},
		{
			TestAlias:                       "1 work item for 1 worker",
			InitWorkerNumber:                1,
			StartWorkerNumber:               1,
			ExpectedWorkerChanLen:           0,
			ExpectedIsCancelationChanClosed: true,
			ExpectedIsWorkersChanClosed:     true,
		},
		{
			TestAlias:                       "1 work item for 10 workers",
			InitWorkerNumber:                10,
			StartWorkerNumber:               1,
			ExpectedWorkerChanLen:           0,
			ExpectedIsCancelationChanClosed: true,
			ExpectedIsWorkersChanClosed:     true,
		},
		{
			TestAlias:                       "20 work items for 16 workers",
			InitWorkerNumber:                16,
			StartWorkerNumber:               20,
			ExpectedWorkerChanLen:           0,
			ExpectedIsCancelationChanClosed: true,
			ExpectedIsWorkersChanClosed:     true,
		},
		{
			TestAlias:                       "200 work items for 16 workers",
			InitWorkerNumber:                16,
			StartWorkerNumber:               200,
			ExpectedWorkerChanLen:           0,
			ExpectedIsCancelationChanClosed: true,
			ExpectedIsWorkersChanClosed:     true,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedWorkerChanLen := testCase.ExpectedWorkerChanLen
		expectedIsCancelationChanClosed := testCase.ExpectedIsCancelationChanClosed
		expectedIsWorkersChanClosed := testCase.ExpectedIsWorkersChanClosed

		fn := func(t *testing.T) {

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			for i := 0; i < startWorkerNumber; i++ {
				reportStart := make(chan struct{})
				go func() { close(reportStart); dep.Do(func() { <-block }, 10) }()
				<-reportStart
			}

			// chanel workers are waiting for close
			closed := make(chan struct{})

			go func() {
				dep.Close()
				close(closed)
			}()

			close(block)

			<-closed

			wp := dep.(*workerPool)

			actualWorkerChanLen := len(wp.workersChan)

			var actualIsCancelationChanClosed bool

			if _, more := <-wp.cancellationChan; !more {
				actualIsCancelationChanClosed = true
			}

			var actualIsWorkersChanClosed bool

			if _, more := <-wp.workersChan; !more {
				actualIsWorkersChanClosed = true
			}

			if actualWorkerChanLen != expectedWorkerChanLen {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immidiate release of workers \r\n has WorkerChanLen length %v workers \r\n while expected %v \r\n", testAlias, actualWorkerChanLen, expectedWorkerChanLen)
			}

			if actualIsCancelationChanClosed != expectedIsCancelationChanClosed {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immidiate release of workers \r\n has CancelationChan Closed ( %v )  \r\n while expected %v \r\n", testAlias, actualIsCancelationChanClosed, expectedIsCancelationChanClosed)
			}

			if actualIsWorkersChanClosed != expectedIsWorkersChanClosed {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immidiate release of workers \r\n has WorkersChan Closed ( %v )  \r\n while expected %v \r\n", testAlias, actualIsWorkersChanClosed, expectedIsWorkersChanClosed)
			}

		}
		t.Run(testAlias, fn)
	}
}

func Test_WorkerPoolCloseRunningWorkersComplete(t *testing.T) {

	testCases := []struct {
		TestAlias           string
		InitWorkerNumber    int
		StartWorkerNumber   int
		ExpectedDoneWorkers int
	}{
		{
			TestAlias:           "20 work items for 0 workers",
			InitWorkerNumber:    0,
			StartWorkerNumber:   20,
			ExpectedDoneWorkers: 0,
		},
		{
			TestAlias:           "1 work item for 1 worker",
			InitWorkerNumber:    1,
			StartWorkerNumber:   1,
			ExpectedDoneWorkers: 1,
		},
		{
			TestAlias:           "1 work item for 10 workers",
			InitWorkerNumber:    10,
			StartWorkerNumber:   1,
			ExpectedDoneWorkers: 1,
		},
		{
			TestAlias:           "20 work items for 16 workers",
			InitWorkerNumber:    16,
			StartWorkerNumber:   20,
			ExpectedDoneWorkers: 16,
		},
		{
			TestAlias:           "200 work items for 16 workers",
			InitWorkerNumber:    16,
			StartWorkerNumber:   200,
			ExpectedDoneWorkers: 16,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedDoneWorkers := testCase.ExpectedDoneWorkers

		fn := func(t *testing.T) {

			// the chanel where workers report about start
			started := make(chan struct{})

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerNumber := 0

			// Aggregate number of started workers
			go func() {
				for {
					_, more := <-started
					if more {
						actualWorkerNumber++
					} else {
						break
					}
				}
			}()

			for i := 0; i < startWorkerNumber; i++ {
				reportStart := make(chan struct{})
				go func() { reportStart <- struct{}{}; dep.Do(func() { started <- struct{}{}; <-block }, 10) }()
				<-reportStart
				close(reportStart)
			}

			// to be sure all workers reported their start
			time.Sleep(time.Duration(10 * startWorkerNumber))

			go dep.Close()

			// stop blocking workers
			close(block)

			actualDoneWorkers := actualWorkerNumber

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() with immidiate WorkerPool.Close() \r\n done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_ErrorOnTimeOut(t *testing.T) {

	blockingFuncTimeout := time.Duration(1000)

	testCases := []struct {
		TestAlias          string
		InitWorkerNumber   int
		StartBlockedWorker int
		TimeOut            time.Duration
		ExpectedError      error
	}{
		{
			TestAlias:          "Generic One",
			InitWorkerNumber:   10,
			StartBlockedWorker: 10,
			TimeOut:            10,
			ExpectedError:      errors.New("The request for a free execution slot has been timed out"),
		},
		{
			TestAlias:          "Generic One",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            100,
			ExpectedError:      nil,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startBlockedWorker := testCase.StartBlockedWorker
		timeOut := testCase.TimeOut
		expectedError := testCase.ExpectedError

		fn := func(t *testing.T) {

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			for i := 0; i < startBlockedWorker; i++ {
				reportStart := make(chan struct{})
				go func() { reportStart <- struct{}{}; dep.Do(func() { <-block }, blockingFuncTimeout) }()
				<-reportStart
				close(reportStart)
			}

			// to be sure all workers managed to start
			time.Sleep(time.Duration(10 * startBlockedWorker))

			actualError := dep.Do(func() { <-block }, timeOut)

			// stop blocking workers
			close(block)

			if !((actualError == nil && expectedError == nil) ||
				(actualError != nil && expectedError != nil &&
					actualError.Error() == expectedError.Error())) {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() with Timeout \r\n returned %v Error \r\n while expected %v \r\n", testAlias, actualError, expectedError)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_ErrorOnClosingWhileObtainingSlot(t *testing.T) {

	testCases := []struct {
		TestAlias          string
		InitWorkerNumber   int
		StartBlockedWorker int
		TimeOut            time.Duration // Keep TimeOut duration with enough room
		ExpectedError      error
	}{
		{
			TestAlias:          "Generic 10/10/100",
			InitWorkerNumber:   10,
			StartBlockedWorker: 10,
			TimeOut:            1000000,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
		{
			TestAlias:          "Generic 10/9/100",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            1000000,
			ExpectedError:      nil,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startBlockedWorker := testCase.StartBlockedWorker
		timeOut := testCase.TimeOut
		expectedError := testCase.ExpectedError

		fn := func(t *testing.T) {

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			for i := 0; i < startBlockedWorker; i++ {
				reportStart := make(chan struct{})
				go func() { reportStart <- struct{}{}; dep.Do(func() { <-block }, 10) }()
				<-reportStart
				close(reportStart)
			}

			// to be sure all workers managed to start
			time.Sleep(time.Duration(10 * startBlockedWorker))

			actualError := error(nil)

			gocha := make(chan struct{})

			reportStart := make(chan struct{})
			go func() {
				close(reportStart)
				actualError = dep.Do(func() { <-block }, timeOut)
				close(gocha)
			}()

			<-reportStart

			closed := make(chan struct{})
			go func() { dep.Close(); close(closed) }()

			<-gocha

			// stop blocking workers
			go close(block)

			<-closed

			if !((actualError == nil && expectedError == nil) ||
				(actualError != nil && expectedError != nil &&
					actualError.Error() == expectedError.Error())) {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() waiting for a slot while WP is closing \r\n returned %v Error \r\n while expected %v \r\n", testAlias, actualError, expectedError)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_ErrorOnObtainingSlotAfterClosing(t *testing.T) {

	testCases := []struct {
		TestAlias          string
		InitWorkerNumber   int
		StartBlockedWorker int
		TimeOut            time.Duration
		ExpectedError      error
	}{
		{
			TestAlias:          "Generic 10/10/10",
			InitWorkerNumber:   10,
			StartBlockedWorker: 10,
			TimeOut:            10,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
		{
			TestAlias:          "Generic 10/9/10",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            10,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startBlockedWorker := testCase.StartBlockedWorker
		timeOut := testCase.TimeOut
		expectedError := testCase.ExpectedError

		fn := func(t *testing.T) {

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			for i := 0; i < startBlockedWorker; i++ {
				reportStart := make(chan struct{})
				go func() { close(reportStart); dep.Do(func() { <-block }, 10) }()
				<-reportStart
			}

			// to be sure all workers managed to start
			time.Sleep(time.Duration(10 * startBlockedWorker))

			reportStart := make(chan struct{})
			closed := make(chan struct{})
			go func() { close(reportStart); dep.Close(); close(closed) }()

			<-reportStart

			actualError := dep.Do(func() { <-block }, timeOut)

			// stop blocking workers
			go close(block)

			<-closed

			if !((actualError == nil && expectedError == nil) ||
				(actualError != nil && expectedError != nil &&
					actualError.Error() == expectedError.Error())) {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() right after WP is closing \r\n returned %v Error \r\n while expected %v \r\n", testAlias, actualError, expectedError)
			}
		}
		t.Run(testAlias, fn)
	}

}
