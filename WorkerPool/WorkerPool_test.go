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
			TestAlias:              "2 work items for 0 workers",
			InitWorkerNumber:       0,
			StartWorkerNumber:      2,
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

			// channel workers are waiting for close
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
		TestAlias           string
		InitWorkerNumber    int
		StartWorkerNumber   int
		ExpectedDoneWorkers int
	}{
		{
			TestAlias:           "2 work items for 0 workers",
			InitWorkerNumber:    0,
			StartWorkerNumber:   2,
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
			ExpectedDoneWorkers: 20,
		},
		{
			TestAlias:           "1000 work items for 16 workers",
			InitWorkerNumber:    16,
			StartWorkerNumber:   1000,
			ExpectedDoneWorkers: 1000,
		},
		{
			TestAlias:           "1000 work items for 1 workers",
			InitWorkerNumber:    1,
			StartWorkerNumber:   1000,
			ExpectedDoneWorkers: 1000,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedDoneWorkers := testCase.ExpectedDoneWorkers

		fn := func(t *testing.T) {

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerDone := 0
			mtx := sync.RWMutex{}

			// timeout should be long enough not to happen
			timeOut := time.Duration(time.Millisecond)

			wg := sync.WaitGroup{}

			for i := 0; i < startWorkerNumber; i++ {
				wg.Add(1)
				err := dep.Do(func() {
					mtx.Lock()
					actualWorkerDone++
					mtx.Unlock()
					wg.Done()
				}, timeOut)
				if err != nil {
					wg.Done()
				}
			}

			wg.Wait()

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
		TestAlias                        string
		InitWorkerNumber                 int
		StartWorkerNumber                int
		ExpectedWorkerChanLen            int
		ExpectedIsCancellationChanClosed bool
		ExpectedIsWorkersChanClosed      bool
	}{
		{
			TestAlias:                        "20 work items for 0 workers",
			InitWorkerNumber:                 0,
			StartWorkerNumber:                20,
			ExpectedWorkerChanLen:            0,
			ExpectedIsCancellationChanClosed: true,
			ExpectedIsWorkersChanClosed:      true,
		},
		{
			TestAlias:                        "1 work item for 1 worker",
			InitWorkerNumber:                 1,
			StartWorkerNumber:                1,
			ExpectedWorkerChanLen:            0,
			ExpectedIsCancellationChanClosed: true,
			ExpectedIsWorkersChanClosed:      true,
		},
		{
			TestAlias:                        "1 work item for 10 workers",
			InitWorkerNumber:                 10,
			StartWorkerNumber:                1,
			ExpectedWorkerChanLen:            0,
			ExpectedIsCancellationChanClosed: true,
			ExpectedIsWorkersChanClosed:      true,
		},
		{
			TestAlias:                        "20 work items for 16 workers",
			InitWorkerNumber:                 16,
			StartWorkerNumber:                20,
			ExpectedWorkerChanLen:            0,
			ExpectedIsCancellationChanClosed: true,
			ExpectedIsWorkersChanClosed:      true,
		},
		{
			TestAlias:                        "200 work items for 16 workers",
			InitWorkerNumber:                 16,
			StartWorkerNumber:                200,
			ExpectedWorkerChanLen:            0,
			ExpectedIsCancellationChanClosed: true,
			ExpectedIsWorkersChanClosed:      true,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		initWorkerNumber := testCase.InitWorkerNumber
		startWorkerNumber := testCase.StartWorkerNumber
		expectedWorkerChanLen := testCase.ExpectedWorkerChanLen
		expectedIsCancellationChanClosed := testCase.ExpectedIsCancellationChanClosed
		expectedIsWorkersChanClosed := testCase.ExpectedIsWorkersChanClosed

		fn := func(t *testing.T) {

			// channel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			wg := sync.WaitGroup{}

			for i := 0; i < startWorkerNumber; i++ {
				wg.Add(1)
				go func() { wg.Done(); dep.Do(func() { <-block }, 10) }()
			}

			wg.Wait()

			// channel closed on Close() complete
			closed := make(chan struct{})

			go func() {
				dep.Close()
				close(closed)
			}()

			close(block)

			<-closed

			wp := dep.(*workerPool)

			actualWorkerChanLen := len(wp.workersChan)

			var actualIsCancellationChanClosed bool

			if _, more := <-wp.cancellationChan; !more {
				actualIsCancellationChanClosed = true
			}

			var actualIsWorkersChanClosed bool

			if _, more := <-wp.workersChan; !more {
				actualIsWorkersChanClosed = true
			}

			if actualWorkerChanLen != expectedWorkerChanLen {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immediate release of workers \r\n has WorkerChanLen length %v workers \r\n while expected %v \r\n", testAlias, actualWorkerChanLen, expectedWorkerChanLen)
			}

			if actualIsCancellationChanClosed != expectedIsCancellationChanClosed {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immediate release of workers \r\n has CancellationChan Closed ( %v )  \r\n while expected %v \r\n", testAlias, actualIsCancellationChanClosed, expectedIsCancellationChanClosed)
			}

			if actualIsWorkersChanClosed != expectedIsWorkersChanClosed {
				t.Errorf("For TestAlias '%s' WorkerPool.Close() with immediate release of workers \r\n has WorkersChan Closed ( %v )  \r\n while expected %v \r\n", testAlias, actualIsWorkersChanClosed, expectedIsWorkersChanClosed)
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

			// channel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			actualWorkerNumber := 0
			mtx := sync.RWMutex{}

			timeOut := time.Duration(time.Microsecond)

			wg := sync.WaitGroup{}

			for i := 0; i < startWorkerNumber; i++ {
				wg.Add(1)

				go func() {
					err := dep.Do(func() {
						mtx.Lock()
						actualWorkerNumber++
						mtx.Unlock()
						wg.Done()
						<-block
					}, timeOut)
					if err != nil {
						wg.Done()
					}
				}()
			}

			wg.Wait()

			go dep.Close()

			// release workers
			close(block)

			mtx.RLock()
			actualDoneWorkers := actualWorkerNumber
			mtx.RUnlock()

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' WorkerPool.Do() with immediate WorkerPool.Close() \r\n done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
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

			// channel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			wg := sync.WaitGroup{}

			for i := 0; i < startBlockedWorker; i++ {
				wg.Add(1)
				go func() { _ = dep.Do(func() { <-block }, blockingFuncTimeout); wg.Done() }()
			}

			wg.Wait()

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
			TestAlias:          "Generic 10/10/Millisecond",
			InitWorkerNumber:   10,
			StartBlockedWorker: 10,
			TimeOut:            time.Millisecond,
			ExpectedError:      errors.New("The worker pool is shutting down and wont take new assignments"),
		},
		{
			TestAlias:          "Generic 10/9/Millisecond",
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

			// channel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			wg := sync.WaitGroup{}

			for i := 0; i < startBlockedWorker; i++ {
				wg.Add(1)
				go func() { _ = dep.Do(func() { <-block }, timeOut); wg.Done() }()
			}

			wg.Wait()

			actualError := error(nil)
			mtx := sync.Mutex{}

			wg.Add(1)
			go func() {
				wg.Done()
				mtx.Lock()
				actualError = dep.Do(func() { <-block }, timeOut)
				mtx.Unlock()
			}()

			wg.Wait()

			_ = <-time.After(1)

			wg.Add(1)
			closed := make(chan struct{})
			go func() { wg.Done(); dep.Close(); close(closed) }()

			wg.Wait()

			_ = <-time.After(1)

			// stop blocking workers
			go close(block)

			_ = <-closed

			mtx.Lock()

			// The defer in TD tests worked badly in 1.7.0...4
			defer mtx.Unlock()
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

			// channel workers are waiting for close
			block := make(chan struct{})

			dep := NewWorkerPool(initWorkerNumber)

			wg := sync.WaitGroup{}

			for i := 0; i < startBlockedWorker; i++ {
				wg.Add(1)
				go func() { wg.Done(); dep.Do(func() { <-block }, 10) }()
			}

			wg.Wait()

			closed := make(chan struct{})
			wg.Add(1)
			go func() { wg.Done(); dep.Close(); close(closed) }()

			wg.Wait()

			actualError := dep.Do(func() { <-block }, timeOut)

			// release workers
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
