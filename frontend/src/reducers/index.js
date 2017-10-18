import { combineReducers } from 'redux';
import { articleData, articlesAreLoading } from './articles';
import { reducer as reduxFormReducer } from 'redux-form';

export default combineReducers({
    articleData,
    articlesAreLoading,
    form: reduxFormReducer
});