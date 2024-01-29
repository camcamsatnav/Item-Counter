const express = require('express')
const app = express()
app.use(express.json())
const port = 3696
app.listen(port, () => console.log(`Listening on port ${port}`))

// for (let i = 0; i < file.nbt.length; i++) {
//     if (file.nbt[2].name === 'minecraft:chest') {
//         for (let j = 0; j < file.nbt[4].value.list.length; j++) {
//             console.log(file.nbt[4].value.list)
//         }
//     }
// }
// for (let j = 0; j < file.nbt[4].value.list.length; j++) {
//     if (file.nbt[4].value.list[j][1].value === 'minecraft:string') {
//         console.log(file.nbt[4].value.list[j])
//     }
// }
app.get('/chest/:search', (req, res) => {
    let file = require('./final.json')
    let count = 0;
    for (let i = 0; i < file.nbt.length; i++) {
        if (file.nbt[2].name === 'minecraft:chest') {
            console.log("yes chest")
            for (let j = 0; j < file.nbt[4].value.list.length; j++) {
                if (file.nbt[4].value.list[j][1].value === req.params.search) {
                    count += file.nbt[4].value.list[j][2].value
                }
            }
        }
    }
    res.send({count: count})
})