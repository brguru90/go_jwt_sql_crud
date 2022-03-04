import React, { useRef } from "react"
import { useEffect } from "react"
import { Link } from "react-router-dom"
import { useNavigate } from "react-router-dom"
import {exeFetch} from "../../modules"
import "./style.scss"

export default function signup() {
    let navigate = useNavigate()

    let email = useRef(null)
    let name = useRef(null)
    let description = useRef(null)

    const createNewUser = () => {
        const newUserData = {
            email: email.current.value,
            name: name.current.value,
            description: description.current.value,
        }

        exeFetch("/api/sign_up", {
            method: "post",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newUserData),
        })
            .then(({ body }) => {
                navigate("/user_profile", {
                    state: body.data,
                })
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    useEffect(() => {
        console.log("-----------sign up----------")

    }, [])
    

    return (
        <div className="sign_up">
            <center>
                <h1>Signup</h1>
            </center>
            <br />

            <fieldset>
                <legend>User detail</legend>
                <table>
                    <tbody>
                        <tr>
                            <td>Username</td>
                            <td>
                                <input type="text" ref={name} />
                            </td>
                        </tr>
                        <tr>
                            <td>Email</td>
                            <td>
                                <input type="text" ref={email} />
                            </td>
                        </tr>
                        <tr>
                            <td>Detail</td>
                            <td>
                                <textarea
                                    className="ver_resizable"
                                    ref={description}
                                ></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td colSpan={2}>
                                <input
                                    type="button"
                                    value="submit"
                                    onClick={createNewUser}
                                />
                            </td>
                        </tr>
                    </tbody>
                </table>
            </fieldset>

            <label>
                Already have account? <Link to="/">Login</Link>
            </label>
        </div>
    )
}
