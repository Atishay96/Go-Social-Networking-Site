import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Form, TextArea } from 'semantic-ui-react'
import SegmentReact from '../segment';

//component
import Header from '../header';
import { getUser } from '../../redux/actions/profile';
import { getItem } from '../../redux/helpers/storage';

//css
import '../../css/profile.css'
class landingPage extends Component {
    constructor() {
        super();
        this.state = {
            user: {}
        }
    }

    componentWillMount() {
        const { dispatch } = this.props
        const authToken = getItem('AuthToken')
        let url = window.location.href
        url = url.split('/')
        let userId = url[url.length - 1]
        if (userId[userId.length - 1] === '#') {
            userId.substring(0, userId.length - 1)
        }
        if (authToken) {
            dispatch(getUser(authToken, userId))
        }
    }

    componentWillReceiveProps(props) {
        
    }

    render() {
        return (
            <div>
                <Header></Header>
                <div className="main-body container-fluid rows marginTop">
                    <div className="col-md-3">
                    </div>
                    <div className="col-md-6">
                        <div className="rows">
                            <div className="profilePicture col-md-6">
                                Profile Picture
                            </div>
                            <div className="profileName col-md-6">
                                Name
                            </div>
                            <div className="firstName marginTop col-md-6">
                                FirstName
                            </div>
                            <div className="lastName marginTop col-md-6">
                                LastName
                            </div>
                            <div className="userName marginTop col-md-12">
                                UserName
                            </div>
                        </div>
                    </div>
                    <div className="col-md-3">
                    </div>
                </div>
                >            </div>
        )
    }
}


const mapStateToProps = (state) => {
    console.log('sending to componentWillReceiveProps')
    console.log(state)
    return {}
}

export default connect(mapStateToProps)(landingPage)