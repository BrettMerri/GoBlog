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

export function showAddArticleForm(bool) {
    return {
        type: 'SHOW_ADD_ARTICLE_FORM',
        display: bool
    }
}

export function articleIsPosting(bool) {
    return {
        type: 'ARTICLE_IS_POSTING',
        isPosting: bool
    }
}

export function articlePostNewArticleSuccess(articleData) {
    return {
        type: 'ARTICLE_POST_NEW_ARTICLE_SUCCESS',
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

export function postNewArticle(values, userId) {
    return (dispatch) => {
        dispatch(articleIsPosting(true));
        let sendJson = {
            "title": values.title,
            "body": values.body,
            "user": {
                "_id": userId
            }
        };
        fetch('/api/article/add', {
            method: "POST",
            body: JSON.stringify(sendJson)
        })
        .then(response => {
            if (!response.ok) {
                throw Error(response.statusText);
            }
            dispatch(articleIsPosting(false));
            return response;
        })
        .then(response => response.json())
        .then(articleData => {
            articleData.result && dispatch(articlePostNewArticleSuccess(articleData.article));
        })
        .catch((err) => console.log(err));
    }
}