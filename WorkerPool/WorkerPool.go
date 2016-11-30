package WorkerPool

import (
	"errors"
	"time"
)

const (
	ERR_DEPARTMENTSHUTDOWN = "The department is shutting down and wont take new assignments"
	ERR_TIMEDOUTREQUSTSLOT = "The request for a free execution slot has been timed out"
)

// Represents simple assignment
type WorkItem func()

// Represents simple workers pool operating on Projects and WorkITems
type WorkerPool interface {

	// Synchronously requests worker slot and end exectes/does WorkItem in parallel routine as soon as slot is obtained
	// If no slot aquired upon timeOut exceeds - returns ERR_TIMEDOUTREQUSTSLOT
	Do(wi WorkItem, timeOut time.Duration) error

	// Closes the department
	Close()
}

// Private custom  implementation of department
type workerPool struct {
	isShuttingDown bool
	workersChan    chan struct{}
}

// A new Deapertment Factory
func NewDepartment(initWorkerNumber int) WorkerPool {
	// instantiate  pool
	workersChan := make(chan struct{}, initWorkerNumber)

	// fill up pool
	// for each initially empty slot we shoul put one value
	for i := 0; i < initWorkerNumber; i++ {
		workersChan <- struct{}{}
	}

	return &workerPool{
		isShuttingDown: false,
		workersChan:    workersChan,
	}

}

// Implements Department.Do(wi WorkItem) method
func (this *workerPool) Do(wi WorkItem, timeOut time.Duration) error {

	if this.isShuttingDown {
		return errors.New(ERR_DEPARTMENTSHUTDOWN)
	}

	t := time.NewTimer(timeOut)

	select {
	case _ = <-this.workersChan:
		if !t.Stop() {
			<-t.C
		}
	case _ = <-t.C:
		return errors.New(ERR_TIMEDOUTREQUSTSLOT)
	}

	if this.isShuttingDown {
		return errors.New(ERR_DEPARTMENTSHUTDOWN)
	}

	go func() {
		defer this.releaseSlot()
		wi()
	}()

	return nil
}

func (this *workerPool) Close() {
	this.isShuttingDown = true

	// wait while all left assignments are done
	for i := 0; i < cap(this.workersChan); i++ {
		<-this.workersChan
	}

	close(this.workersChan)

}

func (this *workerPool) releaseSlot() {
	this.workersChan <- struct{}{}
}
