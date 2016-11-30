package WorkerPool

import (
	"testing"
	"time"
)

func Test_NewDepartment(t *testing.T) {
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
			id := NewDepartment(initWorkerNumber)

			d, ok := id.(*workerPool)

			if !ok {
				t.Skipf("\r\n For TestAlias: '%s' the NewDepartment(%v)  returned unknown implementation of Department inteface\r\n", initWorkerNumber, testAlias)
			}

			actualWorkerNumber := len(d.workersChan)

			if actualWorkerNumber != expectedWorkerNumber {
				t.Errorf("\r\n For TestAlias: '%s' the NewDepartment(%v)  returned department{} \r\n with numberOfWorkers = %v \r\n while expected %v \r\n", testAlias, initWorkerNumber, actualWorkerNumber, expectedWorkerNumber)
			}

			actualChanCapacity := cap(d.workersChan)

			if actualChanCapacity != expectedChanCapacity {
				t.Errorf("\r\n For TestAlias: '%s' the NewDepartment(%v)  returned department{} \r\n with workersPool cap = %v \r\n while expected %v \r\n", testAlias, initWorkerNumber, actualChanCapacity, expectedChanCapacity)
			}

		}
		t.Run(testAlias, fn)

	}
}

func Test_DepartmentDoWorkersLimit(t *testing.T) {

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

			// chanel workes are waiting for close
			block := make(chan struct{})

			dep := NewDepartment(initWorkerNumber)

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
			time.Sleep(1000)

			actualStartedWorkers := actualWorkerNumber

			// cancell workers
			close(block)

			if actualStartedWorkers != expectedStartedWorkers {
				t.Errorf("For TestAlias '%s' Department.Do()  \r\n started %v workers \r\n while expected %v \r\n", testAlias, actualStartedWorkers, expectedStartedWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_DepartmentDoProcessAllWorkers(t *testing.T) {

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

			dep := NewDepartment(initWorkerNumber)

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

			for i := 0; i < startWorkerNumber; i++ {
				go func() {
					dep.Do(func() {
						done <- struct{}{}
					}, 100000000)

				}()
			}

			// Give a bit time to start workers
			time.Sleep(100000000)

			// wait while all left assignments are done
			actualDoneWorkers := actualWorkerDone

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' Department.Do()  \r\n has done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}

func Test_DepartmentClose(t *testing.T) {

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

			// chanel workes are waiting for close
			block := make(chan struct{})

			dep := NewDepartment(initWorkerNumber)

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
			time.Sleep(1000)

			go dep.Close()

			// stop blocking workers
			close(block)

			actualDoneWorkers := actualWorkerNumber

			if actualDoneWorkers != expectedDoneWorkers {
				t.Errorf("For TestAlias '%s' Department.Do() with immidiate Department.Close() \r\n done %v workers \r\n while expected %v \r\n", testAlias, actualDoneWorkers, expectedDoneWorkers)
			}
		}
		t.Run(testAlias, fn)
	}

}
