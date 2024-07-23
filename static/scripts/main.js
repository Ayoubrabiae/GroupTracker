const bodyElement = document.body

// Deal with local storage
const localStorageToggle = (key) => {
    if (localStorage.getItem(key)) {
        localStorage.setItem(key, "")
    } else {
        localStorage.setItem(key, "true")
    }
}

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

// Top Button Functionality
const topBtn = document.querySelector(".top-arrow")

if (topBtn) {
    topBtn.addEventListener("click", () => {
        window.scrollTo(0, 0)
    })

    window.addEventListener("scroll", () => {
        if (window.scrollY > 400) {
            if (!topBtn.classList.contains("show")) {
                topBtn.classList.add("show")
            }
        } else {
            if (topBtn.classList.contains("show")) {
                topBtn.classList.remove("show")
            }
        }
    })

}