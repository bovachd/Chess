import React, { Component } from 'react';
import './styles/FrontPage.css';
import UserForm from './UserForm';
import Intro from './Intro';
import RankDir from './RankDir';
import About from './About';
import axios from 'axios';
import {Redirect} from 'react-router-dom';
import CryptoJS from "crypto-js";
import config from '../../../config';
// import history from '../../../history'

export default class FrontPage extends Component {
    constructor(props){
        super(props)

        this.state = {
            isLoginForm: false,
            token: localStorage.getItem("token"),
            loginStatus: localStorage.getItem("loginStatus"),
            isRedirect: false,
            top: []
        }
    }

    UNSAFE_componentWillMount() {
        
        let token = localStorage.getItem("token")
        if (token !== null)
            this.loginWithToken(token)
    }

    componentDidMount() {
        this.getTopPlayer()
    }

    loginOnClick = () => {
        if (this.state.isLoginForm === false){
            this.setState({
                isLoginForm: !this.state.isLoginForm
            })
        }
    }
    loginOnClick2 = () => {
        // this.props.getConfirm(this.state.loginStatus, "HomePage")
        // this.setState({
        //     isRedirect: true
        // })
        if (this.state.loginStatus !== undefined || this.state.loginStatus !== null){
            this.props.history.push("/Home")
        }
    }
    mainOnClick = () => {
        if (this.state.isLoginForm === true){
            this.setState({
                isLoginForm: !this.state.isLoginForm
            })
        }
    }

    getTopPlayer = () => {
        let url = config.API_URL + "/api/v1/be/account/accounts/top10";
        axios.get(url)
            .then((response) => {
                if (response.data.success === true)
                    this.setState({
                        top: response.data.data
                    })
            })
            .catch((error) => {
            });
    }

    login = (userName, password) => {
        let url = config.API_URL + "/api/v1/be/access/login";
        axios.post(url, {
            name: userName, // Dữ liệu được gửi lên endpoint '/user'
            password: password
        })
            .then((response) => {
                // console.log(response)
                // this.setState({
                    
                // })
                if (response.data.success === true) {
                    this.setState({           
                        isLoginForm: false,
                        token: response.data.token,
                        loginStatus: response.data.success,  
                    })
                    let encryptName = CryptoJS.AES.encrypt(userName, "secret")
                    let encryptPass = CryptoJS.AES.encrypt(password, "secret")
                    localStorage.setItem("userName", encryptName)
                    localStorage.setItem("userPass", encryptPass)
                    localStorage.setItem("token", response.data.token)
                    localStorage.setItem("loginStatus", response.data.success)
                }
                
            })
            .catch((error) => {
            });

        // this.props.passDataToHP(userName, password)
    }

    loginWithToken = (token) => {
        var headerconfig = {
            header: {
                "Authorization": token
            }
        }
        let url = config.API_URL + "/api/v1/be/access/login/token";
        axios.post(url,
            {}, headerconfig
        )
            .then((response) => {
                if (response.data.success === true)
                localStorage.setItem("loginStatus", response.data.success)
            })
            .catch((error) => {
            });
    }


    register = (userName, password) => {
        let url = config.API_URL + "/api/v1/be/account/create";
        axios.post(url, {
            name: userName, // Dữ liệu được gửi lên endpoint '/user'
            password: password
        })
            .then((response) => {
                if (response.data.success === true) {
                    this.setState({           
                        isLoginForm: false,
                        token: response.data.token,
                        loginStatus: response.data.success,  
                    })
                    localStorage.setItem("token", response.data.token)
                    localStorage.setItem("loginStatus", response.data.success)
                    this.login(userName, password)
                }
            })
            .catch((error) => {
            });
    }
    render() {
        if (this.state.isRedirect === true){
            return <Redirect to="/Home"></Redirect>
        }
        return (
            <div className="fp-container" >
            <div className="spinner-border" role="status">
                <span className="sr-only">Loading...</span>
            </div>
                <div className="fp-bg">

                </div>
                <div className="fp-bg-box">
                    <div className="fp-bg-img">

                    </div>
                    <div className="fp-logo">
                        <p>
                            <i className='fas fa-chess-queen'></i>
                            Chess Online
                        </p>            
                    </div>
                </div>     
                <UserForm 
                    isLoginForm={this.state.isLoginForm} 
                    onRegisterSubmit={this.register} 
                    login={this.login}
                ></UserForm> 
                <div className="fp-main" 
                    onClick={this.mainOnClick}
                > 
                    <div className="fp-main-content">
                        <Intro 
                            isLoginForm={this.state.isLoginForm}
                        ></Intro>
                        <RankDir 
                            isLoginForm={this.state.isLoginForm} 
                            loginOnClick={this.state.loginStatus===null?this.loginOnClick:this.loginOnClick2} 
                            token={this.state.token}
                            getTop={this.state.top}
                        ></RankDir>
                    </div>                        
                </div>    
                <About 
                    isLoginForm={this.state.isLoginForm}
                ></About>
            </div>
        )
    }
}
