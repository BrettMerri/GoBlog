import React, { Component } from 'react';
import { connect } from 'react-redux';
import { fetchArticleData, showAddArticleForm, postNewArticle } from '../../../../actions/article';
import { fetchUserData, userSelected } from '../../../../actions/user';
import ArticleContainer from './ArticleContainer';
import UserSelector from './UserSelector';
import AddArticleButton from './AddArticleButton';
import AddArticleForm from './AddArticleForm';
import './Home.css';

class Home extends Component {
  componentDidMount() {
      this.props.fetchArticleData();
      this.props.fetchUserData();
  }

  handleUserChange = (values) => {
    if (values.userId === undefined)
      values.userId = "";
    this.props.updateUserData(values.userId);
  }

  addArticleClick(userSelected) {
    this.props.showAddArticleForm(true);
  }

  addArticleSubmit = (values) => {
    this.props.postNewArticle(values, this.props.userSelected)
  }

  render() {
    return (
      <div id="HomeContainer">
        <UserSelector
          userData={this.props.userData}
          isLoading={this.props.usersAreLoading}
          onChange={this.handleUserChange}
        />
        <AddArticleButton
          userSelected={this.props.userSelected}
          onClick={() => this.addArticleClick(this.props.userSelected)}
        />
        <AddArticleForm
          display={this.props.displayAddArticleForm}
          onSubmit={this.addArticleSubmit}
          isPosting={this.props.isPosting}
        />
        <ArticleContainer
          articleData={this.props.articleData}
          isLoading={this.props.articlesAreLoading}
        />
      </div>
    )
  }
}

const mapStateToProps = (state) => {
  return {
      articleData: state.articleData,
      articlesAreLoading: state.articlesAreLoading,
      displayAddArticleForm: state.showAddArticleForm,
      articleIsPosting: state.articleIsPosting,
      userData: state.userData,
      usersAreLoading: state.usersAreLoading,
      userSelected: state.userSelected,
      isPosting: state.articleIsPosting
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
      fetchArticleData: () => dispatch(fetchArticleData()),
      fetchUserData: () => dispatch(fetchUserData()),
      updateUserData: (userId) => dispatch(userSelected(userId)),
      showAddArticleForm: (bool) => dispatch(showAddArticleForm(bool)),
      postNewArticle: (values, userId) => dispatch(postNewArticle(values, userId))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Home);