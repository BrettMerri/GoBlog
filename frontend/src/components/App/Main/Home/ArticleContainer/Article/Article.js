import React from 'react';
import Comment from './Comment/Comment';

const Article = (props) => {
    var commentArray = [];
    if (props.articleData.hasOwnProperty('comments')) {
        for (var i = 0; i < props.articleData.comments.length; i++) {
            commentArray.push(<Comment key={i} commentData={props.articleData.comments[i]} />);
        }
    }
    return (
        <div className="article" data-id={props.articleData._id}>
            <h1>{props.articleData.title}</h1>
            <p>{props.articleData.body}</p>
            <p data-id={props.articleData._id}>{props.articleData.user.username}</p>
            {commentArray}
        </div>
    )

}

export default Article;