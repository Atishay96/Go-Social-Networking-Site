import React, {Component} from 'react';

class Header extends Component {
    render() {
        return (
        <div>
            <header>
                <nav className="navigation">
                    <div className="container">
                        <div className="navigation-header margin-top">
                            <h2>HOLA</h2>
                        </div>
                        <div className="navigation-body" id="myNavbar">
                            <input type="text" className="search form-control" placeholder="How you doin' today?"/>
                        </div>
                        <span>
                            <div className="navigation-body-right">              
                                <ul>
                                    <li className="dropdown">
                                        <span><img id="image" src="https://avatars1.githubusercontent.com/u/20838008?s=460&v=4" alt="profile" /></span>
                                        <div className="dropdown-profile">
                                            <div className="arrow-up"></div>
                                            <div id="profile-dropdown" className="dropdown-content">
                                                <ul>                      
                                                    <li>
                                                        <h4>Profile</h4>
                                                        <hr/>
                                                    </li>
                                                    <li>
                                                        <h4>Settings</h4>
                                                        <hr/>
                                                    </li>
                                                    <li>
                                                        <h4>Logout</h4>
                                                    </li>
                                                </ul>
                                            </div>
                                        </div>
                                    </li>
                                    <li>
                                        <span> <i className="fas fa-bell"></i> Notifications</span>
                                    </li>
                                    <li>
                                        <span> <i className="fas fa-question-circle"></i> Questions</span>
                                    </li>
                                    <li className="active">
                                        <span> <i className="fas fa-home"></i> Home</span>
                                    </li>
                                </ul>
                            </div>
                        </span>
                    </div>
                </nav>
            </header>
        </div>            
        )
    }
}
export default Header;