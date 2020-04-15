import {createStore, applyMiddleware, compose} from 'redux';
import thunk from 'redux-thunk';
import reducer from "./reducer";

/*
  redux-sagas 中间件把异步放到单独的文件进行管理
  redux-thunk 中间件把异步放到action中
 */

// https://github.com/zalmoxisus/redux-devtools-extension/#12-advanced-store-setup
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ? window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({}) : compose;
const enhancer = composeEnhancers(
  applyMiddleware(thunk),
  // other store enhancers if any
);

const store = createStore(
  reducer,
  enhancer // redux 中间件
);

export default store;