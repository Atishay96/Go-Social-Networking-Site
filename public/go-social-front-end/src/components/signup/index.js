import React, { Component } from 'react';
import { connect } from 'react-redux'

import "../../css/login.css";
import { notEmpty } from '../../helpers/arrayHelper'
import { signUp } from '../../redux/actions/login';

class signUpPage extends Component {
    constructor(props) {
        super();
        this.state = {
            user: {
                firstName: '',
                lastName: '',
                userName: '',
                email: '',
                password: '',
                accept: false
            },
            error: {
                firstName: false,
                lastName: false,
                userName: false,
                email: false,
                password: false,
                accept: false
            },
            touched: {
                firstName: '',
                lastName: '',
                userName: '',
                email: '',
                password: '',
                accept: false
            },
            signUpError: ''
        }
    }

    componentWillReceiveProps(props) {
        console.log(props)
        this.setState({ signUpError: props.signUp.error })
    }

    onSubmit(e) {
        e.preventDefault()
        const { user } = this.state
        const { dispatch } = this.props;
        this.setState({
            error: {
                firstName: false,
                lastName: false,
                userName: false,
                email: false,
                password: false,
                accept: false
            }
        })
        this.setState({
            touched: {
                firstName: true,
                lastName: true,
                userName: true,
                email: true,
                password: true,
                accept: true,
            }
        })
        let missingParams = notEmpty(user)
        let error = this.state.error
        missingParams.map((param) => {
            error[param] = true
        })
        this.setState({ error })
        if (missingParams.length !== 0) {
            console.log(missingParams)
            return
        }
        if (user.password.trim().length <= 5) {
            this.setState({ signUpError: 'Password should be greater than 5' })
            return
        }
        dispatch(signUp(user))
    }

    fieldChange(e, param) {
        const { user } = this.state
        console.log(user)
        user[param] = e.target.value
        this.setState({ user })
        if (!this.state.touched[param]) {
            let touched = this.state.touched
            touched[param] = true
            this.setState({ touched })
        } else {
            if (!this.state.user[param]) {
                let error = this.state.error
                error[param] = true
                this.setState({ error })
            } else {
                let error = this.state.error
                error[param] = false
                this.setState({ error })
            }
        }
    }

    render() {
        return (
            <div className="signup_form container">
                <div className='sign_up_heading'>Sign up</div>
                <form className="form-horizontal" onSubmit={(e) => { this.onSubmit(e) }}>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="fname">FirstName: </label>
                        <div className="col-sm-10">
                            <input type="text" className="form-control righttwo" id="fname" onChange={(e) => { this.fieldChange(e, 'firstName') }} placeholder="Enter your first name" />
                        </div>
                        <span className="span_error firstNameError">{this.state.error.firstName ? 'First name is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="lname">LastName: </label>
                        <div className="col-sm-10">
                            <input type="text" className="form-control righttwo" id="lname" onChange={(e) => { this.fieldChange(e, 'lastName') }} placeholder="Enter your last name" />
                        </div>
                        <span className="span_error lastNameError">{this.state.error.lastName ? 'Last name is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="username">Username: </label>
                        <div className="col-sm-10">
                            <input type="text" className="form-control righttwo" id="username" onChange={(e) => { this.fieldChange(e, 'userName') }} placeholder="Enter Username" />
                        </div>
                        <span className="span_error userNameError">{this.state.error.userName ? 'User name is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="email">Email: </label>
                        <div className="col-sm-10">
                            <input type="email" className="form-control righttwo" id="email" onChange={(e) => { this.fieldChange(e, 'email') }} placeholder="Enter email" />
                        </div>
                        <span className="span_error emailError">{this.state.error.email ? 'Email is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="pwd">Password: </label>
                        <div className="col-sm-10">
                            <input type="password" className="form-control righttwo" id="pwd" onChange={(e) => { this.fieldChange(e, 'password') }} placeholder="Enter password" />
                        </div>
                        <span className="span_error passwordError">{this.state.error.password ? 'Password name is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-4 left">
                            <div className="checkbox">
                                <label><input type="checkbox" onChange={(e) => { this.fieldChange(e, 'accept') }} /> I Accept all the conditions</label>
                            </div>
                            <span className="span_error acceptError">{this.state.error.accept ? 'Please accept the condtions' : ''}</span>
                        </div>
                        <div className="signup col-sm-offset-2 col-sm-4">
                            <label><a href="/login">Looking for Login?</a></label>
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-10">
                            <button type="submit" className="btn btn-default">Create My account</button>
                        </div>
                    </div>
                    <span className="span_error signUpError">{this.state.signUpError}</span>
                </form>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    console.log('sending to componentWillReceiveProps');
    return { signUp: state.login.signUp }
}

export default connect(mapStateToProps)(signUpPage)