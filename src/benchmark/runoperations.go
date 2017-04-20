package main

import (
	"helpoptimal"
	"math/rand"
	"time"
)
// type RunOperations struct {
// 	testSanity bool
// 	threadId int64
// 	addPercent int64
// 	removePercent int64
// 	keyRange int64
// 	numberOfOps int64
// 	hoLFList helpoptimal.HelpOptimalLFList
// 	randOp int64
// 	randKey float64
// 	results int[]
// }

func sanityRun(bm *benchmark, tid int) {
	chooseOperation := random(0, 2)
	key := random(0, *bm.keySpaceSize)
	numberOfAdd := make([]int, *bm.keySpaceSize)
	numberOfRemove := make([]int, *bm.keySpaceSize)

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
	for i := 0; i < *bm.keySpaceSize; i++ {
		bm.sanityAdds[tid][i] += numberOfAdd[i]
		bm.sanityRemoves[tid][i] += numberOfRemove[i]
	}
}
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
