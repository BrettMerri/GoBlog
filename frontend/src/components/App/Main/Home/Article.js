import React from 'react';
import Comment from './Comment';

const Article = (props) => {
    var commentArray = [];
    if (props.articleData.hasOwnProperty('comments')) {
        for (var i = 0; i < props.articleData.comments.length; i++) {
            commentArray.push(<Comment key={i} commentData={props.articleData.comments[i]} />);
        }
    }
    return (
        <div className="card article" data-id={props.articleData._id}>
            <p className="title">{props.articleData.title}</p>
            <p data-id={props.articleData._id}>Posted by: {props.articleData.user.username}</p>
            <p>{props.articleData.body}</p>
            <div className="comment-container">
                {commentArray}
            </div>
            
        </div>
    )

}

export default Article;