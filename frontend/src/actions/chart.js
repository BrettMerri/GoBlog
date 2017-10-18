export function articlesAreLoading(bool) {
    return {
        type: 'ARTICLES_ARE_LOADING',
        isLoading: bool
    }
}

export function articleFetchDataSuccess(articleData) {
    return {
        type: 'ARTICLE_FETCH_DATA_SUCCESS',
        articleData
    }
}

export function fetchArticleData() {
    return (dispatch) => {
        dispatch(articlesAreLoading(true));

        fetch('/api/article/read')
            .then(response => {
                if (!response.ok) {
                    throw Error(response.statusText);
                }
                dispatch(articlesAreLoading(false));

                return response;
            })
            .then(response => response.json())
            .then(articleData => dispatch(articleFetchDataSuccess(articleData)))
            .catch((err) => console.log(err));
    };
}