import axios from "axios"
import {INIT_CITY_LIST, CHANGE_CURRENT_CITY} from './constants'

export const defaultListData = [
  '明天11点开小组会议.',
];

export const changeCurrentCity = (value) => ({
  type: CHANGE_CURRENT_CITY,
  value,
})

export const InitCityList = (data) => ({
  type: INIT_CITY_LIST,
  data
});


// 使用了 redux-thunk 之后 返回可以是一个函数
export const getCityList = () => {
  return (dispatch) => { // dispatch: 如果action是函数的话会自动接收到dispatch方法
    // ajax request
    axios.get("/data.json").then((res) => {
      const data = res.data.data

      // 数据处理 原数据是data.json
      var tmp = {};
      for ( var i=0; i<data.length; i++){
        var city = data[i].city;
        if (city in tmp){
          tmp[city].push(data[i])
        }else {
          tmp[city] = [data[i]]
        }
      }

      const action = InitCityList(tmp) // action change store
      dispatch(action)

      }).catch(() => { // ajax request error
        console.log("error")
    })
  }
}