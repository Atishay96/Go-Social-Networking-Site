import React, {Component} from 'react';
import { Form, TextArea } from 'semantic-ui-react'
import SegmentReact from '../segment';

//component
import Header from '../header';

class landingPage extends Component {
    constructor() {
        super()
    }
    render() {
        return (
        <div>
            Activation mail sent to your email. Please check
            <div className="signup">
                <a href="/login">Looking for Login?</a>
            </div>
        </div>
        )
    }
}
export default landingPage;