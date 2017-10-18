export function articlesAreLoading(state = false, action) {
    switch (action.type) {
        case 'ARTICLES_ARE_LOADING':
            return action.isLoading;

        default:
            return state;
    }
}

export function articleData(state = {}, action) {
    switch (action.type) {
        case 'ARTICLE_FETCH_DATA_SUCCESS':
            return action.articleData;

        default:
            return state;
    }
}