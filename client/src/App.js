// import React from 'react'
// 
// import RightSide from './components/pages/HomePage/RightSide'
// import LeftSide from './components/pages/HomePage/LeftSide'
// import ChatBox from './components/pages/HomePage/ChatBox'
// import Play_Mode from './components/pages/HomePage/PlayMode'



// export default function App() {
//   let frontPage = <FrontPage token={this.getToken}></FrontPage>
//   return (
//     <Router>
//     <Route path="/" exact component={FrontPage}></Route>
//     <Route path="/ChessPage" component={ChessPage}></Route>
//     <Route path="/HomePage" component={HomePage}></Route>
//     </Router>
//   )
// }

import React, { Component } from 'react'
import {BrowserRouter as Router, Route} from 'react-router-dom'

import FrontPage from './components/pages/FrontPage/FrontPage'
import ChessPage from './components/pages/chessPage/ChessPage'
import HomePage from './components/pages/HomePage/HomePage'



export default class App extends Component {
  constructor(props) {
    super(props)

    this.state = {
         token:"",
         isRedirect: false,
         nextPage: "",
         curPage: "/"
    }
    
  }
  UNSAFE_componentWillUnmount(){
    this.setState({
      token:localStorage.getItem('token'),
      isRedirect: false
    })
  }
  getToken = (token) => {
    this.setState({
      token: token,
    })
  }

  // getConfirm = (confirm, name) =>{
  //     this.setState({
  //       isRedirect: confirm,
  //       nextPage: name
  //     })
  // }

  // redirect = () =>{
  //   console.log(this.state.nextPage)
  //   if (this.state.isRedirect)
  //   return <Redirect from="/" to={this.state.nextPage}></Redirect>
  //   else 
  //   return <Redirect to="/" ></Redirect>
  // }

  render() { 
    return (
      <div>
      <Router>
      {/* {this.redirect()} */}
      <Route path="/" exact>
        <FrontPage name="FrontPage"></FrontPage>
      </Route>
      <Route path={"/ChessPage"}>
        <ChessPage></ChessPage>
      </Route>
      <Route path={"/HomePage"}>
        {localStorage.getItem("loginStatus") !== null?<HomePage></HomePage>:""}
      </Route>
      </Router>   
      </div>
   )
  }
}
