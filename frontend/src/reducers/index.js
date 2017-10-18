import { combineReducers } from 'redux';
import { articleData, articlesAreLoading, showAddArticleForm, articleIsPosting, postedArticleData } from './articles';
import { userData, usersAreLoading, userSelected } from './users';
import { reducer as reduxFormReducer } from 'redux-form';

export default combineReducers({
    articleData,
    articlesAreLoading,
    showAddArticleForm,
    articleIsPosting,
    postedArticleData,
    userData,
    usersAreLoading,
    userSelected,
    form: reduxFormReducer
});