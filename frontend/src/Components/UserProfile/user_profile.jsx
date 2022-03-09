import React, { useState, useRef, useEffect } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import { exeFetch } from "../../modules"
import "./style.scss"


export default function user_profile() {
    let { state } = useLocation()
    let navigate = useNavigate()

    const [userData, setUserData] = useState({ uuid: state?.uuid })
    const [activeSessions, setUserActiveSessions] = useState([])
    let email = useRef(null)
    let name = useRef(null)
    let description = useRef(null)

    const gerUserData = () => {
        exeFetch("/api/user/")
            .then(({ body }) => {
                setUserData(body.data)
            }, () => navigate("/"))
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    const gerUserActiveSessions = () => {
        exeFetch("/api/user/active_sessions/")
            .then(({ body }) => {
                setUserActiveSessions(body.data)
            }, () => navigate("/"))
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    const updateUserData = () => {
        const newUserData = {
            email: email.current.value,
            name: name.current.value,
            description: description.current.value,
        }

        exeFetch("/api/user/", {
            method: "put",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                newUserData,
            }),
        }, () => navigate("/"))
            .then(() => {
                newUserData.uuid = userData.uuid
                setUserData(newUserData)
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    const removeAccount = () => {
        exeFetch("/api/user/", {
            method: "delete",
        })
            .then(({ body }) => {
                alert(JSON.stringify(body))
                navigate("/signup")
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }


    const blockToken = (token_id, exp) => {
        exeFetch("/api/user/block_token/", {
            method: "post",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                token_id,
                exp
            }),
        }, () => navigate("/"))
            .then(() => {
                gerUserActiveSessions()
                // alert(JSON.stringify(body))
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }

    const Logout = () => {
        exeFetch("/api/user/logout/", {}, () => navigate("/"))
            .then(() => {
                navigate("/")
            })
            .catch(e => alert("Error\n" + JSON.stringify(e)))
    }


    let _interval = null

    const login_check = () => {
        exeFetch("/api/login_status/", {}, () =>  navigate("/"))
    }

    useEffect(() => {
        console.log("-----------user profile----------")
        gerUserData()
        gerUserActiveSessions()

        // if (!_interval) {
        //     _interval = setInterval(() => {
        //         login_check()
        //     }, 2000);
        // }

        // return () => {
        //     clearInterval(_interval)
        // }


    }, [])


    return (
        <div className="user_profile">
            <center>
                <h1>User profile</h1>
            </center>
            <br />

            <div className="form_container">
                <form className="user_view">
                    <fieldset>
                        <legend>User detail</legend>
                        <table>
                            <tbody>
                                <tr>
                                    <td>UUID</td>
                                    <td>
                                        <input
                                            type="text"
                                            disabled
                                            value={userData?.uuid || ""}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Username</td>
                                    <td>
                                        <input
                                            type="text"
                                            disabled
                                            value={userData?.name || ""}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Email</td>
                                    <td>
                                        <input
                                            type="text"
                                            disabled
                                            value={userData?.email || ""}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Detail</td>
                                    <td>
                                        <pre className="ver_resizable">
                                            {userData?.description || ""}
                                        </pre>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </fieldset>
                </form>

                <form className="user_update">
                    <fieldset>
                        <legend>Update User detail</legend>
                        <table>
                            <tbody>
                                <tr>
                                    <td>UUID</td>
                                    <td>
                                        <input
                                            type="text"
                                            disabled
                                            value={userData?.uuid || ""}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Username</td>
                                    <td>
                                        <input
                                            type="text"
                                            defaultValue={userData?.name || ""}
                                            ref={name}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Email</td>
                                    <td>
                                        <input
                                            type="text"
                                            defaultValue={userData?.email || ""}
                                            ref={email}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Detail</td>
                                    <td>
                                        <textarea
                                            className="ver_resizable"
                                            ref={description}
                                            defaultValue={userData?.description || ""}
                                        ></textarea>
                                    </td>
                                </tr>
                                <tr>
                                    <td colSpan={2}>
                                        <input
                                            type="button"
                                            value="submit"
                                            onClick={updateUserData}
                                        />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </fieldset>
                </form>

                <form className="user_logout">
                    <fieldset>
                        <legend>Delete User detail</legend>
                        <table>
                            <tbody>
                                <tr>
                                    <td>UUID</td>
                                    <td>
                                        <input
                                            type="text"
                                            disabled
                                            value={userData?.uuid || ""}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td colSpan={2}>&nbsp;</td>
                                </tr>
                                <tr>
                                    <td colSpan={2}>
                                        <input
                                            type="button"
                                            value="Logout"
                                            onClick={Logout}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td colSpan={2}>
                                        <input
                                            type="button"
                                            value="Delete my account"
                                            onClick={removeAccount}
                                        />
                                    </td>
                                </tr>
                                <tr>
                                    <td><input /></td>
                                </tr>
                                <tr>
                                    <th>Cookie:</th>
                                    <td>
                                        {decodeURIComponent(document.cookie)}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </fieldset>
                </form>


                <form className="active_sessions">
                    <fieldset>
                        <legend>Active Sessions</legend>
                        <table>
                            <thead>
                                <tr>
                                    <th>Remove</th>
                                    <th>IP</th>
                                    <th>Token</th>
                                    <th>Status</th>
                                    <th>Expire</th>
                                    <th>Use rAgent</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    activeSessions.map(({ ip, ua, token_id, exp, status }) => {
                                        return (
                                            <tr key={token_id}>
                                                <td >
                                                   {status=="active" && <input type="button" value="delete" onClick={() => blockToken(token_id, exp)} />} 
                                                </td>
                                                <td className="ip">{ip}</td>
                                                <td>{token_id}</td>
                                                <td className="ip">{status}</td>
                                                <td>{new Date(exp).toLocaleString()}</td>
                                                <td>
                                                    <div className="sub_tbl_sect">
                                                        {
                                                            typeof(ua)=="object"?
                                                            Object.entries(JSON.parse(ua)).map(([key, val]) => {
                                                                return <div key={token_id + "_" + key}>
                                                                    <b>{key}</b>: <span>{JSON.stringify(val)}</span>
                                                                </div>
                                                            })
                                                            :ua
                                                            // JSON.stringify(Object.entries(JSON.parse(ua)))
                                                        }
                                                    </div>
                                                </td>
                                            </tr>
                                        )
                                    })
                                }



                            </tbody>
                        </table>
                    </fieldset>
                </form>



            </div>
        </div>
    )
}
