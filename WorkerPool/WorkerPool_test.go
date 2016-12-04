package WorkerPool

import (
	"errors"
	"sync"
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

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerNumber := 0
			mtx := sync.RWMutex{}

			timeOut := time.Duration(time.Millisecond)

			for i := 0; i < startWorkerNumber; i++ {
				_ = dep.Do(func() {
					mtx.Lock()
					actualWorkerNumber++
					mtx.Unlock()
					<-block
				}, timeOut)
			}

			_ = <-time.After(timeOut)

			mtx.RLock()
			actualStartedWorkers := actualWorkerNumber
			mtx.RUnlock()

			// release workers
			defer close(block)

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

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerDone := 0
			mtx := sync.RWMutex{}

			timeOut := time.Duration(time.Millisecond)

			for i := 0; i < startWorkerNumber; i++ {
				_ = dep.Do(func() {
					mtx.Lock()
					actualWorkerDone++
					mtx.Unlock()
				}, timeOut)
			}

			_ = <-time.After(timeOut)

			// wait while all left assignments are done
			mtx.RLock()
			actualDoneWorkers := actualWorkerDone
			mtx.RUnlock()

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

			_ = <-time.After(1)

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

			// chanel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerNumber := 0
			mtx := sync.RWMutex{}

			timeOut := time.Duration(time.Microsecond)

			for i := 0; i < startWorkerNumber; i++ {
				reportStart := make(chan struct{})
				go func() {
					close(reportStart)
					_ = dep.Do(func() {
						mtx.Lock()
						actualWorkerNumber++
						mtx.Unlock()
						<-block
					}, timeOut)
				}()
				<-reportStart
			}

			_ = <-time.After(time.Millisecond)

			go dep.Close()

			// stop blocking workers
			close(block)

			mtx.RLock()
			actualDoneWorkers := actualWorkerNumber
			mtx.RUnlock()

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() with immidiate WorkerPool.Close() \r\n done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_ErrorOnTimeOut(t *testing.T) {

	blockingFuncTimeout := time.Duration(time.Millisecond)

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
			TimeOut:            time.Millisecond,
			ExpectedError:      errors.New("The request for a free execution slot has been timed out"),
		},
		{
			TestAlias:          "Generic Two",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            time.Millisecond,
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
				go func() { close(reportStart); _ = dep.Do(func() { <-block }, blockingFuncTimeout) }()
				<-reportStart
			}

			_ = <-time.After(time.Microsecond)

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
			TimeOut:            time.Millisecond,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
		{
			TestAlias:          "Generic 10/9/100",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            time.Millisecond,
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

			//timeOut := time.Millisecond

			for i := 0; i < startBlockedWorker; i++ {
				reportStart := make(chan struct{})
				go func() { close(reportStart); _ = dep.Do(func() { <-block }, timeOut) }()
				_ = <-reportStart

			}

			_ = <-time.After(time.Millisecond)

			actualError := error(nil)

			gocha := make(chan struct{})

			reportStart := make(chan struct{})
			go func() {
				close(reportStart)
				actualError = dep.Do(func() { <-block }, timeOut)
				close(gocha)
			}()

			_ = <-reportStart

			closed := make(chan struct{})
			go func() { dep.Close(); close(closed) }()

			_ = <-gocha

			// stop blocking workers
			go close(block)

			_ = <-closed

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
			TimeOut:            time.Millisecond,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
		{
			TestAlias:          "Generic 10/9/10",
			InitWorkerNumber:   10,
			StartBlockedWorker: 9,
			TimeOut:            time.Millisecond,
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
				_ = <-reportStart
			}

			_ = <-time.After(1)

			reportStart := make(chan struct{})
			closed := make(chan struct{})
			go func() { close(reportStart); dep.Close(); close(closed) }()

			_ = <-reportStart

			actualError := dep.Do(func() { <-block }, timeOut)

			// stop blocking workers
			go close(block)

			_ = <-closed

			if !((actualError == nil && expectedError == nil) ||
				(actualError != nil && expectedError != nil &&
					actualError.Error() == expectedError.Error())) {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() right after WP is closing \r\n returned %v Error \r\n while expected %v \r\n", testAlias, actualError, expectedError)
			}
		}
		t.Run(testAlias, fn)
	}

}
