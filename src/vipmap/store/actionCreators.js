import axios from "axios"
import * as constants from './constants';

export const changeCurrentCity = (city, zoom) => ({
  type: constants.CHANGE_CURRENT_CITY,
  currentCity: city,
  currentZoom: zoom,
})

export const InitCityList = (data) => {

  // 数据处理 原数据是data.json
  var result = {};
  for ( var i=0; i<data.length; i++){
    var city = data[i].city;
    if (city in result){
      result[city].push(data[i])
    }else {
      result[city] = [data[i]]
    }
  }

  return {
    type: constants.INIT_CITY_LIST,
    cityList: result,
    srcCityList: data,
  }};

export const changeCurrentZoom = (value) => ({
  type: constants.CHANGE_CURRENT_ZOOM,
  value,
})

// 使用了 redux-thunk 之后 返回可以是一个函数
export const getCityList = () => {
  return (dispatch) => { // dispatch: 如果action是函数的话会自动接收到dispatch方法
    // ajax request
    var url = window.location.href
    axios.get(url + "/data.json").then((res) => {
      dispatch(InitCityList(res.data.data)) // action change store
      }).catch(() => { // ajax request error
        console.log("error")
    })
  }
}