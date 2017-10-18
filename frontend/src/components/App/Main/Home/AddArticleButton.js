import React from 'react';

const AddArticleButton = (props) => {
    if (props.userSelected === "") {
        return null;
    }
    return (
        <button onClick={props.onClick} className="btn btn-default add-article-button">Add Article</button>
    )
}

export default AddArticleButton;