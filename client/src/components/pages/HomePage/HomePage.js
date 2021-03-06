import React, { Component } from 'react';
import './styles/HomePage.css';
import PlayModes from './PlayModes';
import Chat from './Chat';
import Profile from './Profile';
import Friend from './Friend';
import RoomComponent from './RoomComponent';
import axios from 'axios';
import CryptoJS from 'crypto-js';
import config from '../../../config'
// import {Switch, Route} from 'react-router-dom';
// import ChessPage from '../chessPage/ChessPage';


export default class HomePage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            isVsBot: true,
            userData: {},
            friends: [],
            redirectToVsBot: false,
        }

        this.userName = this.props.userName
        this.pass = this.props.userPass
    }

    componentDidMount(){
        let loginStatus = localStorage.getItem("loginStatus")
        // console.log(loginStatus)
        if ((loginStatus === true && this.props.userName === "") || (loginStatus === undefined)){
            this.props.history.push("/")
        }
    }


    async componentWillMount() {
        let token = localStorage.getItem("token");
        if (token !== null){
            try {
                let response = {}
                let userNameDeCrypt = CryptoJS.AES.decrypt(localStorage.getItem("userName"), "secret")
                let userPassDeCrypt = CryptoJS.AES.decrypt(localStorage.getItem("userPass"), "secret")
                let name = userNameDeCrypt.toString(CryptoJS.enc.Utf8)
                let pass = userPassDeCrypt.toString(CryptoJS.enc.Utf8)
                response = await this.login(name, pass);
                this.getFriends(token, response.user.id)
                this.setState({ 
                    userData: response.user,
                });
            } catch (error) {
                console.log(error);
            }
        }
    }

    getFriends = (token, userId) => {
        var headerconfig = {
            headers: {
                'Authorization': token
            }
        }
        var url = config.API_URL + '/api/v1/be/friend/friends/all/' + String(userId);
        axios.get(url, headerconfig).then((response) => {
            if (response.data.success === true){
                this.setState({
                    friends: response.data.data 
                })
            }
        })
    }

    addFriend = async (token, userID, friendID) => {
        var headerconfig = {
            headers: {
                'Authorization': token
            }
        }
        let url = config.API_URL + "/api/v1/be/friend/friends/new" ;

        axios.post(url, {
            friendId: friendID,
            userId: userID
            }, headerconfig).then((response) => {
                if (response.data.success){
                    // alert("12331");
                    this.getFriends(token, userID)
                }
            }).catch((error) =>{
                console.log(error)
            })
            
    }

    addFriendOnClick =  (id) => {
        let userID = this.state.userData.id
        let frId = Number(id)
        // console.log(typeof(userID), typeof(frId))
        if (userID !== frId){
            this.addFriend(localStorage.getItem("token"), userID, frId)
        }
        
    }

    removeFriend = async (token, userId, frId) => {
        var headerconfig = {
            headers: {
                'Authorization': token
            }
        }
        var url = config.API_URL +  '/api/v1/be/friend/friends/' + String(userId) + "/"  +String(frId)

        const response = await axios.delete(url, headerconfig)
        if (response.data.success === true){
            var newurl = config.API_URL +  '/api/v1/be/friend/friends/all/' + String(userId)
            const response = await axios.get(newurl, headerconfig)
            if (response.data.success === false) {
                this.setState({
                    friends: []
                })
            } else {
                this.setState({
                    friends: response.data.data
                })
            }
            // this.getFriends(token, userId)
        }
        
    }

    removeFriendOnClick = (frId) => {
        let userID = this.state.userData.id
        let _frId = Number(frId)
        this.removeFriend(localStorage.getItem("token"), userID, _frId)
        // this.getFriends(localStorage.getItem("token"), this.state.userData.id)
    }
    report = (token ,message, reportedId, reporterID) => {
        var headerconfig = {
            headers: {
                'Authorization': token
            }
        }
        let url = config.API_URL + "/api/v1/be/report/reports"
        axios.post(url, {
            message: message,
            reportedAccountId: reportedId,
            reporterId: reporterID
        }, headerconfig).then((response) => {
            if(response.data.success === true) {
            }
        })
    }

    reportOnClick = (message, reportid, reportMsg) => {
        // console.log(reportid, reportMsg)
        let _repedID = Number(reportid)
        let msg = message + ", user_message : " + reportMsg
        let repID = this.state.userData.id
        
        this.report(localStorage.getItem("token") ,msg , _repedID, repID)
    }

    login = async (userName, password) => {
        let url = config.API_URL + "/api/v1/be/access/login"
        const response = await axios.post(url, {
            name: userName, // Dữ liệu được gửi lên endpoint '/user'
            password: password
        })
        // console.log(response)
        return response.data.data
    }

    redirectToVsBot = () => {
        this.props.history.push("/Home/Bot")
    
    }

    vsMan = (isVsBot) => {
        this.setState({isVsBot: isVsBot})
    }

    render() {
        const {userData, friends} = this.state;
        var friendData = friends;
        return (
            <div className="hp-container">
                <div className="hp-bg-box">
                    <div className="hp-bg-img">

                    </div>
                    <div className="hp-logo">
                        <p>
                            <i className='fas fa-chess-queen'></i>
                            Chess Online
                        </p>            
                    </div>
                </div>
                <div className="hp-main">
                    <div className="hp-main-content">
                        <div className="hp-col-3 mr-15">
                            <div className="hp-row-1">
                                <PlayModes vsMan={this.vsMan} redirectToVsBot={this.redirectToVsBot}></PlayModes>
                            </div>
                            <div className="hp-row-2">
                                <RoomComponent isVsBot={this.state.isVsBot} history={this.props.history} passRoomIdToVsMan={this.props.passRoomIdToVsMan}></RoomComponent>
                                <Chat isVsBot={this.state.isVsBot} addFriend={this.addFriendOnClick} userID={this.state.userData} report={this.reportOnClick}></Chat>
                            </div>                            
                        </div>
                        <div className="hp-col-1 sr">
                            <div className="hp-row-3">
                                <Profile name={userData.name} rank={userData.Rank} point={userData.Point}/>
                            </div>
                            <div className="hp-row-4">
                                {<Friend friends={friendData} remove={this.removeFriendOnClick}/>}
                            </div>
                        </div>
                    </div>     
                </div>    
            </div>     
        )
    }
}