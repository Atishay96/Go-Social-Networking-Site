import React, { Component } from 'react';
import { connect } from 'react-redux'

import { notEmpty } from '../../helpers/arrayHelper'
import { login } from '../../redux/actions/login';

import "../../css/login.css";

class loginPage extends Component {
    constructor() {
        super();
        this.state = {
            user: {
                email: '',
                password: '',
                rememberMe: false
            },
            error: {
                email: '',
                password: '',
                rememberMe: false
            },
            touched: {
                email: '',
                password: '',
                rememberMe: false
            },
            loginError: ''
        }
    }
    
    componentWillReceiveProps(props) {
        console.log(props)
        this.setState({ loginError: props.signIn.error })
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

    onSubmit(e) {
        e.preventDefault()
        const { user } = this.state
        const { dispatch } = this.props;
        this.setState({
            error: {
                email: false,
                password: false,
                rememberMe: false
            }
        })
        this.setState({
            touched: {
                email: true,
                password: true,
                rememberMe: true
            }
        })
        let userCopy = Object.assign({}, user)
        delete userCopy['rememberMe']
        let missingParams = notEmpty(userCopy)
        let error = this.state.error
        missingParams.map((param) => {
            error[param] = true
        })
        this.setState({ error })
        if (missingParams.length !== 0) {
            console.log(missingParams)
            return
        }
        dispatch(login(user))
    }

    render() {
        return (
            <div className="login_form container">
                <form className="form-horizontal" onSubmit={(e) => { this.onSubmit(e) }}>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="email">Email:</label>
                        <div className="col-sm-10">
                            <input type="email" className="form-control" id="email" onChange={(e) => { this.fieldChange(e, 'email') }} placeholder="Enter email" />
                        </div>
                        <span className="span_error emailError">{this.state.error.email ? 'Email is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="pwd">Password:</label>
                        <div className="col-sm-10">
                            <input type="password" className="form-control" id="pwd" onChange={(e) => { this.fieldChange(e, 'password') }} placeholder="Enter password" />
                        </div>
                        <span className="span_error passwordError">{this.state.error.password ? 'Password name is required' : ''}</span>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-4 left">
                            <div className="checkbox">
                                <label><input type="checkbox" onChange={(e) => { this.fieldChange(e, 'rememberMe') }} /> Remember me</label>
                            </div>
                        </div>
                        <div className="signup col-sm-offset-2 col-sm-4">
                            <label><a href="/signup">Signup Instead?</a></label>
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-10">
                            <button type="submit" className="btn btn-default">Submit</button>
                        </div>
                    </div>
                    <span className="span_error signUpError">{this.state.loginError}</span>
                </form>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    console.log('sending to componentWillReceiveProps');
    return { signIn: state.login.signIn }
}

export default connect(mapStateToProps)(loginPage)