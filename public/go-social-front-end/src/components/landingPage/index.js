import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Form, TextArea } from 'semantic-ui-react'
import SegmentReact from '../segment';

//component
import Header from '../header';
import { homepage, postData } from '../../redux/actions/posts';
import { getItem } from '../../redux/helpers/storage';

class landingPage extends Component {
    constructor() {
        super();
        this.state = {
            posts: [],
            inputs: {
                post: ''
            },
            error: {
                post: ''
            }
        }
    }

    getPosts() {
        const { dispatch } = this.props
        const authToken = getItem('AuthToken')
        if (authToken) {
            dispatch(homepage(authToken))
        }
    }

    componentWillMount() {
        this.getPosts()
    }

    componentWillReceiveProps(props) {
        if (!props || !props.homepageData) {
            return
        }
        if (props.homepageData.error) {
            alert(props.homepageData.error)
            return
        }
        if (props.homepageData.post) {
            let posts = this.state.posts
            posts.unshift(props.homepageData.data)
            this.setState({ posts })
            return
        }
        this.setState({ posts: props.homepageData.data })
    }

    setInputFieldEmpty(field) {
        let temp  = this.state.inputs
        temp[field] = ''
        this.setState({inputs: temp})
    }

    fieldChange(data, type) {
        let obj = this.state.inputs
        obj[type] = data
        this.setState({ inputs: obj })
    }

    submitForm(e) {
        e.preventDefault()
        const { dispatch } = this.props
        const { post } = this.state.inputs
        if (!post) {
            return
        }
        const authToken = getItem('AuthToken')
        if (authToken) {
            dispatch(postData(post, authToken))
            this.setInputFieldEmpty('post')
        }
    }

    render() {
        const { posts, inputs } = this.state
        return (
            <div>
                <Header></Header>
                <div className="main-body row container-fluid margin-top">
                    <div className="col-md-3">
                        <div className='row'>
                            <div className="col-md-offset-7 col-md-3 border">
                                <ul>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                    <li><a href="/">Test link</a></li>
                                </ul>
                            </div>
                        </div>
                    </div>
                    <div className="col-md-6">
                        <div className="card">
                            <div className="card-body">
                                <h4 className="card-title">Go-Social</h4>
                                <span className="card-text">
                                    <Form onSubmit={(e) => { this.submitForm(e) }}>
                                        <TextArea onChange={(e) => this.fieldChange(e.target.value, 'post')} value={inputs.post} placeholder='Try sharing your thoughts' className="input-post" style={{ minHeight: 100 }} />
                                        <span className="span_error postError">{this.state.error.post}</span>
                                        <button type='submit' className='btn btn-success homepage-post'>POST</button>
                                        {posts.map(post => {
                                            return (<span key={post.ID}>
                                                <SegmentReact question={post.Owner.Username} answer={post.Text} ></SegmentReact>
                                            </span>)
                                        })}
                                    </Form>
                                </span>
                            </div>
                        </div>
                    </div>
                    <div className="col-md-3">
                        {/* not sure */}
                    </div>
                </div>
            </div>
        )
    }
}


const mapStateToProps = (state) => {
    // console.log('sending to componentWillReceiveProps');
    return { homepageData: state.homepage.homepageData }
}

export default connect(mapStateToProps)(landingPage)