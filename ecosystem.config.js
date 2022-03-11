module.exports = {
    apps: [
        {
            name: "go_jwt_sql_demo",
            "script": "yarn start_prod",
            "exec_interpreter": "none",
            watch: false,   
            exec_mode: "fork_mode",
            env_pm2: {
                "NODE_ENV": "production"
            }
        }
    ]
}