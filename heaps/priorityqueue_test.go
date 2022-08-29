package heaps

import (
	"log"
	"sort"
	"testing"
)

const (
	pass = "\u2713"
	fail = "\u2717"
)


func lessFunc(x, y *int) bool {
	return *x < *y
}

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int](2, lessFunc)

	if pq.Count() != 0 {
		t.Fatalf("\t%s\t\ta newly instantiated queue should have a count of 0", fail)
	} else {
		t.Logf("\t%s\t\ta newly instantiated queue should have a count of 0", pass)
	}
}

func TestNewPriorityQueueFromList(t *testing.T) {
	list := []int{1,3,7,8,-20,56}
	pq := NewPriorityQueueFromList(list, lessFunc)
	n := len(list)
	if pq.Count() != 6 {
		t.Fatalf("\t%s\t\ta newly instantiated queue should have a count of %d", fail, n)
	} else {
		t.Logf("\t%s\t\ta newly instantiated queue should have a count of %d", pass, n)
	}
	sort.Ints(list)
	ix := 0
	for  ; ix < n && pq.Count() > 0 ; ix++{
		min := pq.Delete()
		if min == nil {
			log.Fatalf("\t%s\tcurrent min element should be:\t%d, but nil was returned\n", fail,list[ix])
		}
		if *min != list[ix] {
			log.Fatalf("\t%s\tcurrent min element should be:\t%d, but it was:\t%d\n", fail,list[ix], *min)
		} else {
			log.Printf("\t%s\tcurrent min element should be:\t%d", pass,*min)
		}
	}

	if pq.Len() != 0 {
		t.Fatalf("\t%s\t\ta newly instantiated queue should have a count of %d", fail, 0)
	} else {
		t.Logf("\t%s\t\ta newly instantiated queue should have a count of %d", pass, 0)
	}
}

func TestNewPriorityQueueFromWithRandomReadWrites(t *testing.T) {
	list := []int{1,3,7,8,-20,56}
	sort.Ints(list)
	pq := NewPriorityQueue(0, lessFunc)
	min := pq.Delete()
	if min != nil {
		t.Fatalf("\t%s\t\ta newly instantiated queue should a min or max of nil\n", fail)
	} else {
		t.Logf("\t%s\t\ta newly instantiated queue should a min or max value of%v\n", pass, min)
	}
	ix := 0
	n := len(list)
	for  ; ix < n  ; ix++{
		pq.Insert(list[ix])

		if pq.Count() != 1 {
			log.Fatalf("\t%s\tcurrent count should be: %d\n", fail,1)
		}

		log.Printf("inserting %v to the queue\n",list[ix] )

		min := pq.MinOrMax()
		minDelete := pq.Delete()
		if min == nil || minDelete == nil {
			log.Fatalf("\t%s\tcurrent min element should be:\t%d, but nil was returned\n", fail,list[ix])
		}
		if *min != list[0] && *minDelete != list[ix] {
			log.Fatalf("\t%s\tcurrent min element should be:\t%d, but it was:\t%d\n", fail,list[ix], *min)
		} else {
			log.Printf("\t%s\tcurrent min element should be:\t%d", pass,*min)
		}
		if pq.Count() != 0 {
			log.Fatalf("\t%s\tcurrent count should be: %d\n", fail,0)
		}
	}

	if pq.Len() != 0 {
		t.Fatalf("\t%s\t\ta newly instantiated queue should have a count of %d", fail, 0)
	} else {
		t.Logf("\t%s\t\ta newly instantiated queue should have a count of %d", pass, 0)
	}
}
