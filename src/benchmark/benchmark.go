package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"sync"
	"helpoptimal"
	"runtime"
)
type benchmark struct {
	algo *string
	testSanity *bool
	duration *int
	numOfThreads *int
	searchFraction *int
	insertUpdateFraction *int
	deleteFraction *int
	warmUpTime *int
	keySpaceSize *int
	results []int
	presentKeys []int
	sanityAdds [][]int
	sanityRemoves [][]int
	hoLFList *helpoptimal.HelpOptimalLFList
}

func newBenchmark() *benchmark {
	return new(benchmark)
}
func (bm *benchmark)initFlags() {
	bm.algo = flag.String("a", "HelpOptimalLFList",
		"Available Algorithms  (default=HelpOptimalLFList)")
	bm.testSanity = flag.Bool("t", false, "Sanity check (default=false)")
	bm.duration = flag.Int("d", 2,
		"Test duration in seconds (0=infinite, default=2s)")
	bm.numOfThreads = flag.Int("n", 4, "Number of threads (default=2)")
	bm.searchFraction = flag.Int("r", 0,
		"Fraction of search operations (default=0%)")
	bm.insertUpdateFraction = flag.Int("i", 50,
		"Fraction of insert/add operations (default=50%)")
	bm.deleteFraction = flag.Int("x", 50,
		"Fraction of delete operations (default=50%)")
	bm.warmUpTime = flag.Int("w", 2,
		"Go Runtime warm up time in seconds(default=2s)")
	bm.keySpaceSize = flag.Int("k", 100,
		"Number of possible keys (default=100)")

	flag.Parse()
	if (*bm.insertUpdateFraction + *bm.deleteFraction + *bm.searchFraction) > 100 {
		fmt.Println("(addPercent+removePercent+searchPercent) > 100")
		os.Exit(1)
	}
	bm.results = make([]int, *(bm.numOfThreads))
	if *bm.testSanity {
		bm.presentKeys = make([]int, *(bm.keySpaceSize))
		bm.sanityAdds = make([][]int, *(bm.numOfThreads))
		for i := range bm.sanityAdds {
			bm.sanityAdds[i] = make([]int, *(bm.keySpaceSize))
		}
		bm.sanityRemoves = make([][]int, *(bm.numOfThreads))
		for i := range bm.sanityRemoves {
			bm.sanityRemoves[i] = make([]int, *(bm.keySpaceSize))
		}
	}
	bm.defineSet()
	
}
func (bm *benchmark) sanityTest() {
	fmt.Println("Entering Test Sanity")
	var keyAdded int
	var keyRemoved int
	var wg sync.WaitGroup
	stopFlag := make(chan bool)
	startFlag := make(chan bool)
	for i := 0; i < *bm.numOfThreads; i++ {
		wg.Add(1)
		go func(tid int) {
			fmt.Printf("Entering thread %v\n", tid)
			chooseOperation := random(0, 2)
			key := random(0, *bm.keySpaceSize)
			numberOfAdd := make([]int, *bm.keySpaceSize)
			numberOfRemove := make([]int, *bm.keySpaceSize)
			for {
				select {
				case <- startFlag:
					break
				default:
					continue
				}
				break
			}
			for {
				select {
				case <- stopFlag:
					 break
				default:

					if chooseOperation == 1 {
						if bm.hoLFList.Add(helpoptimal.NewKeyValue(float64(key))) {
							numberOfAdd[key]++
						} else if bm.hoLFList.Remove(helpoptimal.NewKeyValue(float64(key))) {
							numberOfRemove[key]++
						}
					} else {
						if bm.hoLFList.Remove(helpoptimal.NewKeyValue(float64(key))) {
							numberOfRemove[key]++
						} else if bm.hoLFList.Add(helpoptimal.NewKeyValue(float64(key))) {
							numberOfAdd[key]++
						}
					}
					continue
				}
				break
			}
			for i := 0; i < *bm.keySpaceSize; i++ {
				bm.sanityAdds[tid][i] += numberOfAdd[i]
				bm.sanityRemoves[tid][i] += numberOfRemove[i]
			}
			fmt.Printf("Exiting thread %v\n", tid)
			wg.Done()			
		}(i)
	}
	for i := 0; i < *bm.numOfThreads; i++ {
		startFlag <- true
	}
	time.Sleep(time.Second * 100)
	for i := 0; i < *bm.numOfThreads; i++ {
		stopFlag <- true
	}
	wg.Wait()
	failedSanity := false
	for k := 0; k < *bm.keySpaceSize; k++ {
		keyAdded = bm.presentKeys[k]
		keyRemoved = 0
		for tid := 0; tid < *bm.numOfThreads; tid++ {
			keyAdded += bm.sanityAdds[tid][k]
			keyRemoved += bm.sanityRemoves[tid][k]
		}

		if bm.hoLFList.Contains(helpoptimal.NewKeyValue(float64(k))) == true {
			if keyAdded != keyRemoved + 1 {
				fmt.Println("First Sanity passed")
				failedSanity = true
			}
		} else if (keyAdded != keyRemoved) {
			fmt.Println("Second Sanity passed")
			failedSanity = true
		}
			
	}
	if failedSanity == false {
		fmt.Println("Sanity Test Complete")
	}
	fmt.Println("Traversal Test :")
	fmt.Println(bm.hoLFList.TraversalTest());
}
func (bm *benchmark) defineSet() {
	fmt.Println("Define Set")
	bm.hoLFList = helpoptimal.NewHelpOptimalLFList()
}

func (bm *benchmark) initializeSet() {
	var key int
	var added bool
	fmt.Println("Intialize Set")
	for i := 0; i < *bm.keySpaceSize/2; {
		key = random(0, *bm.keySpaceSize);
		added = bm.hoLFList.Add(helpoptimal.NewKeyValue(float64(key)))
		if added == true {
			i++
		}
		if added == true && *bm.testSanity {
			bm.presentKeys[key] = bm.presentKeys[key] + 1
		}
	}
}
func (bm *benchmark) warmupVM() {
	//First Round
	var wg sync.WaitGroup
	stopFlag := make(chan bool)
	startFlag := make(chan bool)
	for i := 0; i < *bm.numOfThreads; i++ {
		wg.Add(1)
		go func(tid int) {
			fmt.Printf("Entering thread %v\n", tid)
			chooseOperation := random(0, 100)
			key := random(0, *bm.keySpaceSize)
			numberOfOps := 0
			for {
				select {
				case <- startFlag:
					break
				default:
					continue
				}
				break
			}
			for {
				select {
				case <- stopFlag:
					break
				default:

					if chooseOperation < *bm.insertUpdateFraction {
						bm.hoLFList.Add(helpoptimal.NewKeyValue(float64(key)))
					} else if (chooseOperation < *bm.deleteFraction){
						bm.hoLFList.Remove(helpoptimal.NewKeyValue(float64(key)))
					} else {
						bm.hoLFList.Contains(helpoptimal.NewKeyValue(float64(key)))
					}
					numberOfOps++
					continue
				}
				break
			}
			bm.results[tid] = numberOfOps
			fmt.Printf("Exiting thread %v\n", tid)
			wg.Done()			
		}(i)
	}
	for i := 0; i < *bm.numOfThreads; i++ {
		startFlag <- true
	}
	time.Sleep(time.Second * 100)
	for i := 0; i < *bm.numOfThreads; i++ {
		stopFlag <- true
	}
	wg.Wait()
	
}

func (bm *benchmark) doBenchmark() {
}

func main() {
	bm := newBenchmark()
	bm.initFlags()
	bm.initializeSet()
	if *bm.testSanity == true {
		bm.sanityTest()
	} else {
		// memory cleanup 
		runtime.GC()
		bm.warmupVM()
		runtime.GC()
		bm.defineSet()
		runtime.GC()
		bm.initializeSet()
		bm.doBenchmark()
	}
}
