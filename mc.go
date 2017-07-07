// Package ll2MC
// ll2MC包提供了将将纬度转换为MC坐标的方法
package ll2MC

import (
	"fmt"
	"math"
	"strconv"
)

func getRange(cf, ce, t float64) float64 {
	if 0 != cf {
		cf = math.Max(cf, ce)
	}

	if 0 != t {
		cf = math.Min(cf, t)
	}

	return cf
}

func getLoop(cf, ce, t float64) float64 {
	for cf < ce {
		cf += t - ce
	}

	for cf > t {
		cf -= t - ce
	}

	return cf
}

func converter(longitude, latitude float64, cg []float64) (float64, float64) {
	if nil == cg {
		return 0.0, 0.0
	}

	t := cg[0] + cg[1]*math.Abs(longitude)
	ce := math.Abs(latitude) / cg[9]
	ch := cg[2] + cg[3]*ce + cg[4]*ce*ce + cg[5]*ce*ce*ce + cg[6]*ce*ce*ce*ce + cg[7]*ce*ce*ce*ce*ce + cg[8]*ce*ce*ce*ce*ce*ce
	if longitude < 0 {
		t *= -1
	} else {
		t *= 1
	}

	if latitude < 0 {
		ch *= -1
	} else {
		ch *= 1
	}

	longitudeRet, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", t), 64)
	latitudeRet, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", ch), 64)

	return longitudeRet, latitudeRet
}

// 经纬度转MC坐标
func Convert(longitude, latitude float64) (float64, float64) {
	LLBAND := []float64{75, 60, 45, 30, 15, 0}
	var LL2MC [][]float64
	LL2MC = [][]float64{
		[]float64{-0.0015702102444, 111320.7020616939, 1704480524535203, -10338987376042340, 26112667856603880, -35149669176653700, 26595700718403920, -10725012454188240, 1800819912950474, 82.5},
		[]float64{0.0008277824516172526, 111320.7020463578, 647795574.6671607, -4082003173.641316, 10774905663.51142, -15171875531.51559, 12053065338.62167, -5124939663.577472, 913311935.9512032, 67.5},
		[]float64{0.00337398766765, 111320.7020202162, 4481351.045890365, -23393751.19931662, 79682215.47186455, -115964993.2797253, 97236711.15602145, -43661946.33752821, 8477230.501135234, 52.5},
		[]float64{0.00220636496208, 111320.7020209128, 51751.86112841131, 3796837.749470245, 992013.7397791013, -1221952.21711287, 1340652.697009075, -620943.6990984312, 144416.9293806241, 37.5},
		[]float64{-0.0003441963504368392, 111320.7020576856, 278.2353980772752, 2485758.690035394, 6070.750963243378, 54821.18345352118, 9540.606633304236, -2710.55326746645, 1405.483844121726, 22.5},
		[]float64{-0.0003218135878613132, 111320.7020701615, 0.00369383431289, 823725.6402795718, 0.46104986909093, 2351.343141331292, 1.58060784298199, 8.77738589078284, 0.37238884252424, 7.45},
	}

	longitude = getLoop(longitude, -180, 180)
	latitude = getRange(latitude, -74, 74)

	var cg []float64
	for cf := 0; cf < len(LLBAND); cf++ {
		if latitude >= LLBAND[cf] {
			cg = LL2MC[cf]
			break
		}
	}

	if nil != cg {
		for cf := len(LLBAND) - 1; cf > 0; cf-- {
			if longitude <= LLBAND[cf] {
				cg = LL2MC[cf]
				break
			}
		}
	}

	longitudeMc, latitudeMc := converter(longitude, latitude, cg)

	return longitudeMc, latitudeMc
}
