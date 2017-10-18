import React from 'react';
import { Field, reduxForm } from 'redux-form';

const UserSelector = (props) => {
    if (props.isLoading) {
        return <p>Loading...</p>
    }
    var userArray = [];
    for (var i = 0; i < props.userData.length; i++) {
        userArray.push(<option key={i} value={props.userData[i]._id}>{props.userData[i].username}</option>);
    }
    return (
        <div className="userSelectorContainer">
            <label>Select user</label>
            <Field name="userId" component="select">
                <option></option>
                {userArray}
            </Field>
        </div>
    )
}

export default reduxForm({
    form: 'user'
})(UserSelector);
