import React from 'react';
import Article from './Article';
const ArticleContainer = (props) => {
    if (props.isLoading) {
        return <p>Loading...</p>
    }
    var articleArray = [];
    for (var i = 0; i < props.articleData.length; i++) {
        articleArray.push(<Article key={i} number={i + 1} articleData={props.articleData[i]} />);
    }
    return (
        <div className="articleContainer">
            {articleArray}
        </div>
    )
}

export default ArticleContainer;