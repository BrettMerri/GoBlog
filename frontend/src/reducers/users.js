export function usersAreLoading(state = false, action) {
    switch (action.type) {
        case 'USERS_ARE_LOADING':
            return action.isLoading;

        default:
            return state;
    }
}

export function userData(state = {}, action) {
    switch (action.type) {
        case 'USER_FETCH_DATA_SUCCESS':
            return action.userData;

        default:
            return state;
    }
}

export function userSelected(state = "", action) {
    switch (action.type) {
        case 'USER_SELECTED':
            return action.userId;
            
        default:
            return state;
    }
}