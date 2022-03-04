let csrf_token;


const exeFetch = (url, options = {}, navigate = () => { }) => {
    console.log(url,csrf_token)
    csrf_token= csrf_token || localStorage.getItem("csrf_token")
    if(csrf_token){
        options.headers=Object.assign( options.headers || {},{csrf_token},)
    }
    return fetch(url, options)
        .then(async (res) => {
            if (res.ok) {
                let _csrf_token=res.headers.get('csrf_token')
                if(_csrf_token){
                    csrf_token=_csrf_token
                    localStorage.setItem("csrf_token",csrf_token)
                }
                console.log("csrf_token",csrf_token)
                return {
                    body: await res.json(),
                }
            }
            else if (res.status == 401) {
                throw { body: await res.json(), ...(navigate() || {}) }
            }
            else {
                throw { err: res.status,body: await res.json(), ...(navigate() || {}) }
            }
        })
}


export {
    exeFetch
}