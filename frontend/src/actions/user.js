export function usersAreLoading(bool) {
    return {
        type: 'USERS_ARE_LOADING',
        isLoading: bool
    }
}

export function userFetchDataSuccess(userData) {
    return {
        type: 'USER_FETCH_DATA_SUCCESS',
        userData
    }
}

export function userSelected(userId) {
    return {
        type: 'USER_SELECTED',
        userId: userId
    }
}

export function fetchUserData() {
    return (dispatch) => {
        dispatch(usersAreLoading(true));
        fetch('/api/user/read')
            .then(response => {
                if (!response.ok) {
                    throw Error(response.statusText);
                }
                dispatch(usersAreLoading(false));

                return response;
            })
            .then(response => response.json())
            .then(userData => dispatch(userFetchDataSuccess(userData)))
            .catch((err) => console.log(err));
    };
}