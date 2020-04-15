import * as constants from './constants';
import { fromJS} from 'immutable';

const defaultState = fromJS({
  cityList: [],
  currentCity: "北京",
  currentZoom: 5,
});


const changeCurrentCity = (state, action) => {
  var result = {"currentCity": action.currentCity,}

  if (action.currentZoom !== state.get('currentZoom')){
    result['currentZoom'] = action.currentZoom
  }
  return state.merge(result)
}

// state    整个DOM的数据库
// action
// reducer 可以接收state，但是不可以在修改stacurrentCityte
export default (state = defaultState, action) => {
  switch(action.type) {
    case constants.INIT_CITY_LIST:
      return state.set("cityList", action.cityList)
    case constants.CHANGE_CURRENT_CITY:
      return changeCurrentCity(state, action)
    default:
      return state;
  }
}