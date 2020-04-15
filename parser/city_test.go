package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents)
	//fmt.Println(result)

	const resultSize = 7
	//	expectedCity := "北京"
	//	expectedHospitals := []model.Hospital{
	//		model.Hospital{
	//			City: expectedCity,
	//			Name: "爱康国宾北京日坛体检分院",
	//			Content: `
	//营业时间：周二到周日 8:00-11:00 采血截止时间 10:30
	//门店地址：北京朝阳区日坛东路7号
	//联系电话：010-85619296`},
	//		model.Hospital{
	//			City: expectedCity,
	//			Name: "爱康国宾北京白云路体检分院",
	//			Content: `
	//门店地址：北京市西城区莲花池东路甲5号院1号楼二层
	//联系电话：010-63377277`},
	//		model.Hospital{
	//			City: expectedCity,
	//			Name: "爱康国宾北京宣武门体检分院VIP部",
	//			Content: `
	//营业时间：周二到周日采血时间7:40-10:30
	//门店地址：北京宣武门外大街甲1号，环球财讯中心D座2层
	//联系电话：010-83163355转1063`},
	//	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}

}
