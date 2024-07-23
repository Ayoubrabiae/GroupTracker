const bodyElement = document.body

// Deal with local storage
const localStorageToggle = (key) => {
    if (localStorage.getItem(key)) {
        localStorage.setItem(key, "")
    } else {
        localStorage.setItem(key, "true")
    }
}

// Header functionality
/* const headerElement = document.querySelector("header")

if (headerElement) {
    const headerFunc = () => {
        if (window.scrollY > 100) {
            headerElement.style.top = "0"
        } else {
            headerElement.style.top = "-100%"
        }
    }
    
    headerFunc()
    
    document.addEventListener("scroll", () => {
        headerFunc()
    })
} */

// Dark mode functionality
const darkModeBtn = document.querySelector(".dark-mode-btn")
if (darkModeBtn) {
    if (localStorage.getItem("dark")) {
        bodyElement.classList.add("dark")
        darkModeBtn.children[0].classList.add("bxs-sun")
    }
    
    darkModeBtn.addEventListener("click", () => {
        localStorageToggle("dark")
        console.log(localStorage.getItem("dark"))
        bodyElement.classList.toggle("dark")
        darkModeBtn.children[0].classList.toggle("bxs-sun")
    })
}

// Hero functionality
const heroImageElements = document.querySelectorAll(".hero .images img")

if (heroImageElements.length) {
    const artistImages = ["queen.jpeg", "soja.jpeg", "pinkfloyd.jpeg", "scorpions.jpeg",
        "xxxtentacion.jpeg", "macmiller.jpeg", "joynerlucas.jpeg", "kencricklamar.jpeg",
        "acdc.jpeg", "pearljam.jpeg", "katyperry.jpeg", "rihanna.jpeg", "genesis.jpeg",
        "philcollins.jpeg", "ledzeppelin.jpeg", "jimihendrix.jpeg", "beegees.jpeg", "deeppurple.jpeg",
        "aerosmith.jpeg", "direstraits.jpeg", "30.jpeg", "imagineDragons.jpeg", "juiceWrld.jpeg", "logic.jpeg",
        "alecBeenjamin.jpeg", "bobbyMcFerrin.jpeg", "R3HAB.jpeg", "postmalone.jpeg", "travisScott.jpeg", "jCole.jpeg",
        "nickelback.jpeg", "mobbDeep.jpeg", "gunsNRoses.jpeg", "nwa.jpeg", "u2.jpeg", "arcticmonkeys.jpeg", "fallOutBoy.jpeg",
        "gorillaz.jpeg", "eagles.jpeg", "linkinpark.jpeg", "redhotchilipeppers.jpeg", "eminem.jpeg", "greenday.jpeg", "metallica.jpeg",
        "coldplay.jpeg", "maroon5.jpeg", "twentyonepilots.jpeg", "therollingstones.jpeg", "muse.jpeg", "foofighters.jpeg", "thechainsmokers.jpeg"]

    setInterval(() => {
        const randomElementNum = Math.floor(Math.random() * heroImageElements.length)
        const randomImageNum = Math.floor(Math.random() * artistImages.length)
        const imageSrc = `https://groupietrackers.herokuapp.com/api/images/${artistImages[randomImageNum]}`
        heroImageElements[randomElementNum].src = imageSrc
    }, 10000)
}
