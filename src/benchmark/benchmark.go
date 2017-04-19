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
	duration *int64
	numOfThreads *int64
	searchFraction *int64
	insertUpdateFraction *int64
	deleteFraction *int64
	warmUpTime *int64
	keySpaceSize *int64
	results []int64
	presentKeys []int64
	sanityAdds [][]int64
	sanityRemoves [][]int64
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
	bm.duration = flag.Int64("d", 2, "Test duration in seconds (0=infinite, default=2s)")
	bm.numOfThreads = flag.Int64("n", 2, "Number of threads (default=2)")
	bm.searchFraction = flag.Int64("r", 0, "Fraction of search operations (default=0%)")
	bm.insertUpdateFraction = flag.Int64("i", 50, "Fraction of insert/add operations (default=50%)")
	bm.deleteFraction = flag.Int64("x", 50, "Fraction of delete operations (default=50%)")
	bm.warmUpTime = flag.Int64("w", 2, "Go Runtime warm up time in seconds(default=2s)")
	bm.keySpaceSize = flag.Int64("k", 100, "Number of possible keys (default=100)")

	flag.Parse()
	fmt.Println(*bm.insertUpdateFraction + *bm.deleteFraction + *bm.searchFraction)
	if (*bm.insertUpdateFraction + *bm.deleteFraction + *bm.searchFraction) > 100 {
		fmt.Println("(addPercent+removePercent+searchPercent) > 100")
		os.Exit(1)
	}
	bm.results = make([]int64, *(bm.numOfThreads))
	if *bm.testSanity {
		bm.presentKeys = make([]int64, *(bm.numOfThreads))
		bm.sanityAdds = make([][]int64, *(bm.numOfThreads))
		for i := range bm.sanityAdds {
			bm.sanityAdds[i] = make([]int64, *(bm.keySpaceSize))
		}
		bm.sanityRemoves = make([][]int64, *(bm.numOfThreads))
		for i := range bm.sanityRemoves {
			bm.sanityRemoves[i] = make([]int64, *(bm.keySpaceSize))
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
}
func (bm *benchmark) defineSet() {
	fmt.Println("Define Set")
	bm.hoLFList = helpoptimal.NewHelpOptimalLFList()
}
