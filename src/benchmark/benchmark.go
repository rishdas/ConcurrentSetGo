package main

import (
	"flag"
	"fmt"
	"os"
	"helpoptimal"
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
func main() {
	bm := newBenchmark()
	bm.initFlags()
	if *bm.testSanity == true {
		bm.sanityTest()
	}
}

func newBenchmark() *benchmark {
	return new(benchmark)
}
func (bm *benchmark)initFlags() {
	bm.algo = flag.String("a", "HelpOptimalLFList", "Available Algorithms  (default=HelpOptimalLFList)")
	bm.testSanity = flag.Bool("t", false, "Sanity check (default=false)")
	bm.duration = flag.Int("d", 2, "Test duration in seconds (0=infinite, default=2s)")
	bm.numOfThreads = flag.Int("n", 2, "Number of threads (default=2)")
	bm.searchFraction = flag.Int("r", 0, "Fraction of search operations (default=0%)")
	bm.insertUpdateFraction = flag.Int("i", 50, "Fraction of insert/add operations (default=50%)")
	bm.deleteFraction = flag.Int("x", 50, "Fraction of delete operations (default=50%)")
	bm.warmUpTime = flag.Int("w", 2, "Go Runtime warm up time in seconds(default=2s)")
	bm.keySpaceSize = flag.Int("k", 100, "Number of possible keys (default=100)")

	flag.Parse()
	fmt.Println(*bm.insertUpdateFraction + *bm.deleteFraction + *bm.searchFraction)
	if (*bm.insertUpdateFraction + *bm.deleteFraction + *bm.searchFraction) > 100 {
		fmt.Println("(addPercent+removePercent+searchPercent) > 100")
		os.Exit(1)
	}
	bm.results = make([]int, *(bm.numOfThreads))
	if *bm.testSanity {
		bm.presentKeys = make([]int, *(bm.numOfThreads))
		bm.sanityAdds = make([][]int, *(bm.numOfThreads))
		for i := range bm.sanityAdds {
			bm.sanityAdds[i] = make([]int, *(bm.keySpaceSize))
		}
		bm.sanityRemoves = make([][]int, *(bm.numOfThreads))
		for i := range bm.sanityRemoves {
			bm.sanityRemoves[i] = make([]int, *(bm.keySpaceSize))
		}
	}
	// Print arguments
	// fmt.Println("benchmark:", *(bm.algo))
	// fmt.Println("benchmark:", *(bm.testSanity))
	// fmt.Println("benchmark:", *(bm.duration))
	// fmt.Println("benchmark:", *(bm.numOfThreads))
	bm.defineSet()
	
}
func (bm *benchmark) sanityTest() {
	fmt.Println("Entering Test Sanity")
	for i := 0; i < *bm.numOfThreads; i++ {
		go sanityRun(bm, i)
	}
}
func (bm *benchmark) defineSet() {
	fmt.Println("Define Set")
	bm.hoLFList = helpoptimal.NewHelpOptimalLFList()
}
