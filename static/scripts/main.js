const bodyElement = document.body
const headerElement = document.querySelector("header")
const darkModeBtn = document.querySelector(".dark-mode-btn")

// Dark mode functionality
darkModeBtn.addEventListener("click", () => {
    console.log("Welcome")
    bodyElement.classList.toggle("dark")
    darkModeBtn.children[0].classList.toggle("bxs-sun")
})

// Hero functionality
const artistNames = ["queen.jpeg", "soja.jpeg", "pinkfloyd.jpeg", "scorpions.jpeg", "xxxtentacion.jpeg", "macmiller.jpeg", "joynerlucas.jpeg", "kencricklamar.jpeg", "acdc.jpeg", "pearljam.jpeg", "katyperry.jpeg", "rihanna.jpeg", "genesis.jpeg", "philcollins.jpeg", "ledzeppelin.jpeg", "jimihendrix.jpeg", "beegees.jpeg", "deeppurple.jpeg", "aerosmith.jpeg", "direstraits.jpeg", "mamonasassassinas.jpeg", "30.jpeg", "imagineDragons.jpeg", "juiceWrld.jpeg", "logic.jpeg", "alecBeenjamin.jpeg", "bobbyMcFerrin.jpeg", "R3HAB.jpeg", "postmalone.jpeg", "travisScott.jpeg", "jCole.jpeg", "nickelback.jpeg", "mobbDeep.jpeg", "gunsNRoses.jpeg", "nwa.jpeg", "u2.jpeg", "arcticmonkeys.jpeg", "fallOutBoy.jpeg", "gorillaz.jpeg", "eagles.jpeg", "linkinpark.jpeg", "redhotchilipeppers.jpeg", "eminem.jpeg", "greenday.jpeg", "metallica.jpeg", "coldplay.jpeg", "maroon5.jpeg", "twentyonepilots.jpeg", "therollingstones.jpeg", "muse.jpeg", "foofighters.jpeg", "thechainsmokers.jpeg"]
const heroImages = document.querySelectorAll(".hero .images img")

setInterval(() => {
    for (let i = 0; i < heroImages.length; i++) {
        const randomNum = Math.floor(Math.random() * artistNames.length)
        const imageSrc = `https://groupietrackers.herokuapp.com/api/images/${artistNames[randomNum]}`
        heroImages[i].src = imageSrc
    }
}, 10000)