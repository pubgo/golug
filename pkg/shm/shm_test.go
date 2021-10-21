package shm

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestAtomic(t *testing.T) {
	span, err := Alloc("TestAtomic", 256)
	if err != nil {
		t.Error(err)
	}
	block, err := span.Alloc(128)

	counter := (*uint32)(unsafe.Pointer(block))
	expected := 10000
	cpu := runtime.NumCPU()
	wg := sync.WaitGroup{}

	wg.Add(cpu)
	for i := 0; i < cpu; i++ {
		go func() {
			for j := 0; j < expected/cpu; j++ {
				atomic.AddUint32(counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if *counter != uint32(expected) {
		t.Errorf("counter error, expected %d, actual %d", 10000, *counter)
	}

	if err := Free(span); nil != err {
		log.Fatalln(err)
	}

}

func TestConsitency(t *testing.T) {
	span, err := Alloc("TestConsitency", 256)
	if err != nil {
		t.Error(err)
	}

	_, err = Alloc("TestConsitency", 512)
	if err == nil && err.Error() == "mmap target path ./mosn_shm_TestConsitency exists and its size 256 mismatch 512" {
		t.Error()
	}

	if err := Free(span); nil != err {
		log.Fatalln(err)
	}

}

func BenchmarkPointerCast_Raw(b *testing.B) {
	var counter uint32 = 0
	ptr := &counter
	for i := 0; i < b.N; i++ {
		atomic.AddUint32(ptr, 1)
	}
}

func BenchmarkPointerCast_Cast(b *testing.B) {
	var counter uint32 = 0
	ptr := uintptr(unsafe.Pointer(&counter))
	for i := 0; i < b.N; i++ {
		atomic.AddUint32((*uint32)(unsafe.Pointer(ptr)), 1)
	}
}
