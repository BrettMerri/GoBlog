export function articlesAreLoading(state = false, action) {
    switch (action.type) {
        case 'ARTICLES_ARE_LOADING':
            return action.isLoading;
        default:
            return state;
    }
}

export function articleData(state = [], action) {
    switch (action.type) {
        case 'ARTICLE_FETCH_DATA_SUCCESS':
            return action.articleData; 
        case 'ARTICLE_POST_NEW_ARTICLE_SUCCESS':
            return [action.articleData, ...state]
        default:
            return state;
    }
}

export function showAddArticleForm(state = false, action) {
    switch (action.type) {
        case 'SHOW_ADD_ARTICLE_FORM':
            return action.display;
        default:
            return state;
    }
}

export function articleIsPosting(state = false, action) {
    switch (action.type) {
        case 'ARTICLE_IS_POSTING':
            return action.isPosting;
        default:
            return state;
    }
}