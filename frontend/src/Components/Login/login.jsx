import React from "react"
import { useEffect } from "react"
import { useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import {exeFetch} from "../../modules"
import "./style.scss"


export default function login() {
    let navigate = useNavigate()
    const [email, setEmail] = useState(null)

    const Login = () => {
        exeFetch("/api/login", {
            method: "post",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email: email,
            }),
        })
            .then(({ body }) => {
                navigate("/user_profile", {
                    state: body.data,
                })
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    const login_check = () => {
        exeFetch("/api/login_status")
            .then(({ body }) => {
                console.log("/user_profile")
                navigate("/user_profile", {
                    state: body.data,
                })
            })
    }

    useEffect(() => {
        console.log("-----------Login----------")
        login_check()
    }, [])


    return (
        <div className="login">
            <center>
                <h1>Login</h1>
            </center>
            <br />

            <form>
                <fieldset>
                    <legend>User detail</legend>
                    <table>
                        <tbody>
                            <tr>
                                <td>Email</td>
                                <td>
                                    <input
                                        type="email"
                                        onChange={(e) => setEmail(e.target.value)}
                                    />
                                </td>
                            </tr>
                            <tr>
                                <td colSpan={2}>
                                    <input type="button" value="Login" onClick={Login} />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </fieldset>
            </form>

            <label>
                New user? <Link to="/signup">Sign up</Link>
            </label>
        </div>
    )
}
