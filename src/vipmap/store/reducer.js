import {INIT_CITY_LIST, CHANGE_CURRENT_CITY} from './constants'

const defaultState = {
  cityList: [],
  currentCity: "上海",
};


// state    整个DOM的数据库
// action
// reducer 可以接收state，但是不可以在修改stacurrentCityte
export default (state = defaultState, action) => {

  // redux ajax
  if (action.type === INIT_CITY_LIST){
    const newState = JSON.parse(JSON.stringify(state)); // create new state
    newState.cityList = action.data;
    return newState;
  }
  // change current city
  if (action.type === CHANGE_CURRENT_CITY) {
    const newState = JSON.parse(JSON.stringify(state)); // create new state
    newState.currentCity = action.value;
    return newState
  }

  return state;
}