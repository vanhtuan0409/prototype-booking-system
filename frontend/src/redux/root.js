import { combineReducers } from "redux";

function defaultReducer(state = 1, action) {
  return state;
}

export default combineReducers({
  sample: defaultReducer
});
