import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from "react-redux"
import VipMap from './vipmap';
import store from "./vipmap/store";

const App = (
  // 子组件可以使用store，都可以通过connect做连接
  <Provider store={store}>
    <VipMap/>
  </Provider>
)

ReactDOM.render(App, document.getElementById('root'));

