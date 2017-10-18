import React from 'react';
import { Field, reduxForm } from 'redux-form';

const validate = values => {
    const errors = {}
    if (!values.title) {
        errors.title = 'Required';
    } else if (values.title.length < 3 || values.title.length > 50) {
        errors.title = 'Must be 3 to 50 characters';
    }
    if (!values.body) {
        errors.body = 'Required';
    } else if (values.body.length < 3 || values.title.length > 200) {
        errors.body = 'Must be 3 to 200 characters';
    }
    return errors
}


const renderField = ({ input, label, type, meta: { touched, error, warning } }) => (
    <div>
        <div className="form-group mx-sm-3">
            <span>{label}</span>
            {type === "textarea" ?
            <textarea {...input} className="form-control" placeholder={label}/> :
            <input {...input} className="form-control" placeholder={label} type={type}/>}
        </div>
        {touched && (error && <span className="error">{error}</span>)}
    </div>
  )
  
const AddArticleForm = (props) => {
    if (!props.display)
        return null;
    return (
    <form className="form" onSubmit={props.handleSubmit}>
        <Field
            name="title"
            label="Title"
            component={renderField}
            type="text"
            placeholder="Title"
        />
        <Field
            name="body"
            label="Body"
            component={renderField}
            type="textarea"
            placeholder="Body"
        />
        <div className="form-group mx-sm-3">
            <button type="submit" className="btn btn-default" disabled={props.submitting}>Submit</button>
        </div>
    </form>
    );
}

export default reduxForm({
    form: 'addArticle',
    validate
})(AddArticleForm);