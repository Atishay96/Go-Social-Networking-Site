import React, { Component } from 'react';
import "../../css/login.css";
class loginPage extends Component {
    constructor() {
        super();
        this.state = {

        }
    }
    render() {
        return (
            <div className="login_form container">
                <form className="form-horizontal">
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="email">Email:</label>
                        <div className="col-sm-10">
                            <input type="email" className="form-control" id="email" placeholder="Enter email" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label className="control-label col-sm-2" htmlFor="pwd">Password:</label>
                        <div className="col-sm-10">
                            <input type="password" className="form-control" id="pwd" placeholder="Enter password" />
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-4 left">
                            <div className="checkbox">
                                <label><input type="checkbox" /> Remember me</label>
                            </div>
                        </div>
                        <div className="signup col-sm-offset-2 col-sm-4">
                            <label><a href="#">Signup Instead?</a></label>
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-offset-2 col-sm-10">
                            <button type="submit" className="btn btn-default">Submit</button>
                        </div>
                    </div>
                </form>
            </div>
        )
    }
}
export default loginPage;