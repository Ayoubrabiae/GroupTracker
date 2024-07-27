/* const bodyElement = document.body
const headerElement = document.querySelector("header")
const darkModeBtn = document.querySelector(".dark-mode-btn")

const localStorageToggle = (key) => {
    if (localStorage.getItem(key)) {
        localStorage.setItem(key, "")
    } else {
        localStorage.setItem(key, "true")
    }
}

// Dark mode functionality
if (localStorage.getItem("dark")) {
    bodyElement.classList.add("dark")
}

darkModeBtn.addEventListener("click", () => {
    localStorageToggle("dark")
    console.log(localStorage.getItem("dark"))
    bodyElement.classList.toggle("dark")
    darkModeBtn.children[0].classList.toggle("bxs-sun")
})

// Hero functionality
const artistImages = ["queen.jpeg", "soja.jpeg", "pinkfloyd.jpeg", "scorpions.jpeg",
                    "xxxtentacion.jpeg", "macmiller.jpeg", "joynerlucas.jpeg", "kencricklamar.jpeg",
                    "acdc.jpeg", "pearljam.jpeg", "katyperry.jpeg", "rihanna.jpeg", "genesis.jpeg",
                    "philcollins.jpeg", "ledzeppelin.jpeg", "jimihendrix.jpeg", "beegees.jpeg", "deeppurple.jpeg",
                    "aerosmith.jpeg", "direstraits.jpeg", "30.jpeg", "imagineDragons.jpeg", "juiceWrld.jpeg", "logic.jpeg",
                    "alecBeenjamin.jpeg", "bobbyMcFerrin.jpeg", "R3HAB.jpeg", "postmalone.jpeg", "travisScott.jpeg", "jCole.jpeg",
                    "nickelback.jpeg", "mobbDeep.jpeg", "gunsNRoses.jpeg", "nwa.jpeg", "u2.jpeg", "arcticmonkeys.jpeg", "fallOutBoy.jpeg",
                    "gorillaz.jpeg", "eagles.jpeg", "linkinpark.jpeg", "redhotchilipeppers.jpeg", "eminem.jpeg", "greenday.jpeg", "metallica.jpeg",
                    "coldplay.jpeg", "maroon5.jpeg", "twentyonepilots.jpeg", "therollingstones.jpeg", "muse.jpeg", "foofighters.jpeg", "thechainsmokers.jpeg"]

const heroImageElements = document.querySelectorAll(".hero .images img")

setInterval(() => {
    const randomElementNum = Math.floor(Math.random() * heroImageElements.length) 
    const randomImageNum = Math.floor(Math.random() * artistImages.length)
    const imageSrc = `https://groupietrackers.herokuapp.com/api/images/${artistImages[randomImageNum]}`
    heroImageElements[randomElementNum].src = imageSrc
}, 10000) */

/*Artist Details Slide Cards*/
let items = document.querySelectorAll('.info .info-child');
let prev = document.getElementById('prev');
let next = document.getElementById('next');

let active = 3;
function loadShow(){
    let setting = 0;
    items[active].style.transform = `none`;
    items[active].style.zIndex = 1;
    items[active].style.filter = 'none';
    items[active].style.opacity = 1;
    for(var i = active + 1; i < items.length; i++){
        setting++;
        items[i].style.transform = `translateX(${120*setting}px) scale(${1 - 0.2*setting}) perspective(16px) rotateY(-1deg)`;
        items[i].style.zIndex = -setting;
        items[i].style.filter = 'blur(5px)';
        items[i].style.opacity = setting > 2 ? 0 : 0.6;
    }
    setting = 0;
      for(var i = active - 1; i >= 0; i--) {
        setting++;
        items[i].style.transform = `translateX(${-120*setting}px) scale(${1 - 0.2*setting}) perspective(16px) rotateY(1deg)`;
        items[i].style.zIndex = -setting;
        items[i].style.filter= 'blur(5px)';
        items[i].style.opacity = setting > 2 ? 0 : 0.6;
    }
}
loadShow();
next.onclick = function() {
    active = active + 1 < items.length ? active + 1 : active;
    loadShow();
}

prev.onclick = function() {
    active = active - 1 >= 0 ? active - 1 : active;
    loadShow();
}