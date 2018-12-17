import React, { Component } from 'react';
import "../../css/login.css";
import { notEmpty } from '../../helpers/arrayHelper'

class signUpPage extends Component {
    constructor() {
        super();
        this.state = {
            user: {
                firstName: '',
                lastName: '',
                userName: '',
                email: '',
                password: '',
                accept: false
            }
        }
    }

    onSubmit(e) {
        e.preventDefault()
        const { user } = this.state
        let missingParams = notEmpty(user)
        if (missingParams.length != 0) {
            console.log(missingParams)
            return alert("Parameters missing")
        }
        alert("READY TO SUBMIT")
    }

    fieldChange(e, param) {
        const { user } = this.state
        console.log(user)
        user[param] = e.target.value
        this.setState(user)
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
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="lname">LastName: </label>
                        <div className="col-sm-10">
                            <input type="text" className="form-control righttwo" id="lname" onChange={(e) => { this.fieldChange(e, 'lastName') }} placeholder="Enter your last name" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="username">Username: </label>
                        <div className="col-sm-10">
                            <input type="text" className="form-control righttwo" id="username" onChange={(e) => { this.fieldChange(e, 'userName') }} placeholder="Enter Username" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="email">Email: </label>
                        <div className="col-sm-10">
                            <input type="email" className="form-control righttwo" id="email" onChange={(e) => { this.fieldChange(e, 'email') }} placeholder="Enter email" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="pwd">Password: </label>
                        <div className="col-sm-10">
                            <input type="password" className="form-control righttwo" id="pwd" onChange={(e) => { this.fieldChange(e, 'password') }} placeholder="Enter password" />
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-4 left">
                            <div className="checkbox">
                                <label><input type="checkbox" onChange={(e) => { this.fieldChange(e, 'accept') }} /> I Accept all the conditions</label>
                            </div>
                        </div>
                        <div className="signup col-sm-offset-2 col-sm-4">
                            <label><a href="#">Looking for Login?</a></label>
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-10">
                            <button type="submit" className="btn btn-default">Create My account</button>
                        </div>
                    </div>
                </form>
            </div>
        )
    }
}
export default signUpPage