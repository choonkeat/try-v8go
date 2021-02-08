/* globals Elm */
console.log = println
const app = Elm.Main.init({
    flags: { initial : 7 }
})

console.log(app)
sleep(5000)

console.log