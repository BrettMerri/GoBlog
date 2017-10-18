import React from 'react';

const Comment = (props) => (
    <div data-id={props.commentData.id}>
        <p><span data-id={props.commentData.user._id}>{props.commentData.user.username}</span>: {props.commentData.body}</p>
    </div>
)

export default Comment;