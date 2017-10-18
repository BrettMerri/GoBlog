import React, { Component } from 'react';
import { connect } from 'react-redux';
import { fetchArticleData } from '../../../../actions/chart';
import ArticleContainer from './ArticleContainer/ArticleContainer';

class Home extends Component {
  componentDidMount() {
      this.props.fetchArticleData();
  }

  render() {
    return (
      <div id="HomeContainer">
        <ArticleContainer
          articleData={this.props.articleData}
          isLoading={this.props.isLoading}
        />
      </div>
    )
  }
}

const mapStateToProps = (state) => {
  return {
      articleData: state.articleData,
      isLoading: state.articlesAreLoading
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
      fetchArticleData: () => dispatch(fetchArticleData())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Home);