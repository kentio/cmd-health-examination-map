import { combineReducers } from 'redux-immutable';
import { reducer as vipmapReducer } from '../vipmap/store';

const reducer = combineReducers({
	vipmap: vipmapReducer,
});

export default reducer;
