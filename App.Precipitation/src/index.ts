require('dotenv').config()
import express from 'express'

const PORT = process.env.PORT || 4000

const main = async () => {

    const app = express()

    app.get('/', (_, res) => {
        res.send("Server is working fine for App.Precipitation!")
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