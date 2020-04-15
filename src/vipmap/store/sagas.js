import {takeEvery, put} from 'redux-saga/effects'
import axios from 'axios'


import {initListAction} from './actionCreators'
import {GET_INIT_CITY_LIST} from "./constants";

function* getInitList() {
  try {
    const res = yield axios.get('/data.json');
    console.log(res, "----")
    console.log(res.data)
    const action = initListAction(res.data.data);
    yield put(action);
  } catch (e) {
    console.log("getInitList error.")
  }
}

// generator 函数
function* mySaga() {
  yield takeEvery(GET_INIT_CITY_LIST, getInitList)
}

export default mySaga;