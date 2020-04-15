import React, {Component} from "react";
import {connect} from "react-redux"

// import { useRef, useEffect, useState } from 'react';

import {Map, APILoader} from '@uiw/react-baidu-map';

import {getCityList, changeCurrentCity} from "./store/actionCreators"
import store from "./store";


class VipMap extends Component {

  render() {
    const {list, handleClick, currentCity} = this.props;

    return (
      <div >
        <div style={{width: '80%', height: 'auto', margin: '0 auto', border: '1px solid #000', marginTop: '50px',}}>
          {
            Object.keys(list).map((item, index) => {
              // var city = item.city;
              return <button
                onClick={() => handleClick({item})}
                key={index}>{item}
                {/*style={{float:'left',}}*/}
              </button>
            })
          }
        </div>

        <div style={{width: '80%', height: '600px', margin: '0 auto', border: '1px solid #000', marginTop: '2px',}}>

          <APILoader akay="X3dS7gyDF4AFBDzF9wWcn3CY49Di4sYQ">
            <Map enableScrollWheelZoom={this.enableScrollWheelZoom} zoom={12} center={currentCity}/>
          </APILoader>
        </div>
      </div>

    )
  }

  componentDidMount() {
    const action = getCityList()
    store.dispatch(action)
  }
}

// link 规则（方式）映射关系
const mapStateToProps = (state) => {
  return {
    list: state.cityList,
    currentCity: state.currentCity,
  }
};

// redux 数据修改逻辑映射 store.dispatch, props
const mapDispatchToProps = (dispatch) => {
  return {
    // city btn click
    handleClick(e) {
      const action = changeCurrentCity(e.item);
      dispatch(action)
    }

  }
};

// 组件是通过connect获取到state的数据
export default connect(mapStateToProps, mapDispatchToProps)(VipMap)
