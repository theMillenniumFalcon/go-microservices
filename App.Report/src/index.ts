import "reflect-metadata"
require('dotenv').config()
import express from 'express'
import { createConnection } from 'typeorm'
import { COOKIE, __prod__ } from "./constants/constants"
import session from 'express-session'
import path from "path"
import cors from "cors"
import { Report } from "./entities/Report"

const PORT = process.env.PORT || 4003

const main = async () => {

    const connection = await createConnection({
        type: "postgres",
        username: "postgres",
        password: "postgres",
        database: "report",
        logging: true,
        synchronize: false,
        migrations: [path.join(__dirname, "./migrations/*")],
        entities: [Report]
    })

    await connection.runMigrations()

    const app = express()

    app.use(cors({
        origin: 'http://localhost:3000',
        credentials: true
    }))

    app.use(session({
        name: COOKIE,
        cookie: {
            maxAge: 1000 * 60 * 60 * 24 * 365 * 10,
            httpOnly: true,
            sameSite: 'lax', // * csrf
            secure: __prod__ // * cookie only works in https
        },
        saveUninitialized: false,
        secret: "asdfghjklqwertyuiopasdfghjkl",
        resave: false,
    }))

    app.get('/', (_, res) => {
        res.send("Server is working fine!")
    })

    const server = app.listen(PORT, () => {
        console.log(`Server listening on port ${PORT}`)
    })
    process.on('unhandledRejection', (err, _) => {
        console.log(`Logged Error: ${err}`)
        server.close(() => process.exit(1))
    })

}
main().catch((error) => {
    console.error(error)
})