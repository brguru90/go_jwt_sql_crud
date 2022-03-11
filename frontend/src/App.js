import React, {Component} from "react"
import {HashRouter, Routes, Route} from "react-router-dom"
import SignUp from "./pages/signup"
import Login from "./pages/login"
import UserProfile from "./pages/user_profile"

export default class App extends Component {
    render() {
        return (
            <HashRouter>
                <Routes>
                    <Route path="/" exact element={<Login />} />
                    <Route path="signup" exact element={<SignUp />} />
                    <Route path="user_profile" exact element={<UserProfile />} />
                </Routes>
            </HashRouter>
        )
    }
}
