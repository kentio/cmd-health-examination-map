import React, {Component} from "react";
import {connect} from "react-redux"
// import { useRef, useEffect, useState } from 'react';
import { Map, APILoader } from '@uiw/react-baidu-map';
import {getCityList, changeCurrentCity,} from "./store/actionCreators"

class VipMap extends Component {

  render() {
    const {list,srcList, handleClick, currentCity, currentZoom} = this.props;

    return (
      <div >
        <div style={{width: '80%', height: 'auto', margin: '0 auto', border: '1px solid #000', marginTop: '50px'}}>
          {
            Object.keys(list).map((item, index) => {
              // var city = item.city;
              return <button
                onClick={() => handleClick({item},13)}
                key={index}>{item}
              </button>
            })

          }
          <button onClick={() => handleClick("北京", 5)}>全国</button>
          <div>
            {currentZoom}
            {currentCity}
          </div>
        </div>

        <div style={{width: '80%', height: '700px', margin: '0 auto', border: '1px solid #000', marginTop: '2px',}}>
          <APILoader akay="X3dS7gyDF4AFBDzF9wWcn3CY49Di4sYQ">
            <Map
              zoom={currentZoom}
              currentCity={currentCity}
              center={currentCity}
              enableScrollWheelZoom={true}
            />
          </APILoader>
        </div>
      </div>

    )
  }

  componentDidMount() {
    this.props.changeCityList();
    // this.bindEvents();
  }
}

// link 规则（方式）映射关系
const mapState = (state) => {
  return {
    list: state.getIn(['vipmap','cityList']),
    srcList: state.getIn(['vipmap','srcCityList']),
    currentCity: state.getIn(['vipmap','currentCity']),
    currentZoom: state.getIn(['vipmap','currentZoom']),
  }
};

// redux 数据修改逻辑映射 store.dispatch, props
const mapDispatch = (dispatch) => {
  return {
    // load city list
    changeCityList(){
      dispatch(getCityList())
    },

    // city btn click
    handleClick(city, zoom) {

      if (typeof city === 'object'){
        city = city.item
      }

      const action = changeCurrentCity(city, zoom);
      dispatch(action)
    },

  }
};

// 组件是通过connect获取到state的数据
export default connect(mapState, mapDispatch)(VipMap)
