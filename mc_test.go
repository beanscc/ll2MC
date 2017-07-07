package ll2MC

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// go test -v -run TestConvertLL2MC
func TestConvertLL2MC(t *testing.T) {
	lng := 109.60355
	lat := 35.17318

	lngMc, latMc := Convert(lng, lat)

	t.Logf("lngMc=%f, latMc=%f \n", lngMc, latMc)
}

// go test -v -run TestGetLocationByLL
func TestGetLocationByLL(t *testing.T) {
	//http://api.map.baidu.com/?qt=rgc&x=12953722.17&y=4837205.83&dis_poi=100&poi_num=10&ie=utf-8&oue=1&fromproduct=jsapi&res=api&callback=BMap._rd._cbk35101&ak=E4805d16520de693a3fe707cdc962045
	lng := 109.60355
	lat := 35.17318

	lngMc, latMc := Convert(lng, lat)
	// baiduMapURL := fmt.Sprintf("http://api.map.baidu.com/?qt=rgc&x=%.2f&y=%.2f&dis_poi=100&poi_num=10&ie=utf-8&oue=1&fromproduct=jsapi&res=api&callback=BMap._rd._cbk35101", lngMc, latMc)

	baiduMapURL := fmt.Sprintf("http://api.map.baidu.com/?qt=rgc&x=%.2f&y=%.2f&dis_poi=100&poi_num=10&ie=utf-8&oue=1&fromproduct=jsapi&res=api", lngMc, latMc)

	t.Logf("request url=%v", baiduMapURL)
	resp, err := http.Get(baiduMapURL)
	if err != nil {
		t.Fatalf("http get err. err=%v", err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ret: %s", respBody)
}
