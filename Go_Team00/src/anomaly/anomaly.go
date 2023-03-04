package anomaly

import (
	// "fmt"
	"fmt"
	"math"
	// "math/rand"
	// "time"
)

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	det := Init()
// 	sd := rand.Float64() * 1.2 + 0.3
// 	mean := rand.Float64() * 20 - 10
// 	for k := 0; k < 20000; k++ {
// 		d := rand.NormFloat64() * sd + mean
// 		det.Do(float64(d))
// 	}
// 	fmt.Println(mean, sd)
// }

// type statistic struct {
// }

type AnomalyDetector struct {
	// stat statistic
	sD float64
	mean float64
	len int64
	k float64
	anomalyCount int
	Do func(float64) bool
}

func Init(kinit float64) *AnomalyDetector {
	var ad AnomalyDetector
	ad.Do = ad.firstValue
	ad.k = kinit
	return &ad
}

func (ad *AnomalyDetector) firstValue(value float64) bool {
	ad.len = 1
	ad.mean = float64(value)
	ad.sD = 0
	ad.Do = ad.notEnoughValues
	return false
}

func (ad *AnomalyDetector) notEnoughValues(value float64) bool {
	ad.statisticCount(value)
	if ad.len > 100 {
		ad.Do = ad.enoughValues
	}
	return false
}

func (ad *AnomalyDetector) enoughValues(value float64) bool {
	sigma := ad.sD * ad.k
	if math.Abs(value - ad.mean) > sigma {
		fmt.Println("!!! ANOMALY !!! ", ad.anomalyCount, value, ad.len, ad.mean, ad.sD)
		ad.anomalyCount++
		return true
	}
	ad.statisticCount(value)
	return false
}

func (ad *AnomalyDetector) statisticCount(value float64) {
	ad.len++
	ad.sD = math.Sqrt((float64(ad.len - 1) / float64(ad.len)) * (ad.sD * ad.sD + (math.Pow((ad.mean - value), 2) / float64(ad.len))))
	ad.mean = (ad.mean * float64(ad.len - 1) / float64(ad.len)) + value / float64(ad.len)
}

// func main() {
// 	t := Create()
// 	// t.init()
// 	// t.foo = t.one
// 	t.foo()
// 	t.foo()
// 	t.foo()
// }


// type test struct {
// 	foo func()
// }

// func Create() (*test) {
// 	var t test
// 	t.foo = t.one
// 	return &t
// }

// func (t *test) one() {
// 	fmt.Println("one")
// 	t.foo = t.two
// }

// func (t *test) two() {
// 	fmt.Println("two")
// }
