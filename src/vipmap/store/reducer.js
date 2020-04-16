import * as constants from './constants';
import { fromJS} from 'immutable';

const defaultState = fromJS({
  cityList: [],
  srcCityList: [],
  // currentCity: "北京",
  currentCity: "上海",
  currentZoom: 13,
});


const changeCurrentCity = (state, action) => {
  var result = {
    "currentCity": action.currentCity,
    "currentZoom": action.currentZoom,
  }
  return state.merge(result)
}

const initCityList = (state, action) => {
  return state.merge({
    cityList: action.cityList,
    srcCityList: action.srcCityList
  })
}

// state    整个DOM的数据库
// action
// reducer 可以接收state，但是不可以在修改stacurrentCityte
export default (state = defaultState, action) => {
  switch(action.type) {
    case constants.INIT_CITY_LIST:
      return initCityList(state, action)
    case constants.CHANGE_CURRENT_CITY:
      return changeCurrentCity(state, action)
    default:
      return state;
  }
}